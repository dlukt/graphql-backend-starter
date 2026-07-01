package cmd

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/MadAppGang/httplog"
	oidcCfg "github.com/deicod/oidcmw/config"
	oidcMW "github.com/deicod/oidcmw/middleware"
	"github.com/deicod/oidcmw/viewer"
	"github.com/dlukt/graphql-backend-starter/config"
	"github.com/dlukt/graphql-backend-starter/ent"
	"github.com/dlukt/graphql-backend-starter/graph"
	coderws "github.com/coder/websocket"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"github.com/vektah/gqlparser/v2/ast"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/mattn/go-sqlite3"
)

var (
	graphqlDebug = true
)

// graphqlCmd represents the graphql command
var graphqlCmd = &cobra.Command{
	Use:   "graphql",
	Short: "run the graphql backend",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("graphql called")
		setDatabaseURI()

		var client *ent.Client
		if useSQLite {
			fmt.Println("Running with SQLite")
			var e error
			client, e = ent.Open(dialect.SQLite, config.SqliteDSN)
			if e != nil {
				return e
			}
		} else {
			fmt.Println("Running with PostgreSQL")
			client = openDB(config.DatabaseURI)
		}
		if graphqlDebug {
			client = client.Debug()
		}
		defer client.Close()
		if e := client.Schema.Create(
			context.Background(),
		); e != nil {
			log.Fatal("opening ent client", e)
		}

		var oidcConfig oidcCfg.Config

		if graphqlDebug {
			oidcConfig = oidcCfg.Config{
				Issuer: config.OidcConfigDev.Issuer,
				Audiences: []string{
					"account",
					config.OidcConfigDev.Audience,
				},
				AuthorizedParties: []string{
					config.OidcConfigDev.AuthorizedParty,
				},
				AllowAnonymousRequests: true,
			}
		} else {
			oidcConfig = oidcCfg.Config{
				Issuer: config.OidcConfigProd.Issuer,
				Audiences: []string{
					"account",
					config.OidcConfigProd.Audience,
				},
				AuthorizedParties: []string{
					config.OidcConfigProd.AuthorizedParty,
				},
				AllowAnonymousRequests: true,
			}
		}
		mw, e := oidcMW.NewMiddleware(oidcConfig)
		if e != nil {
			return e
		}
		srv := NewDefaultServer(graph.NewSchema(client))
		srv.Use(entgql.Transactioner{TxOpener: client})

		corsHandler := cors.AllowAll()
		fmt.Println("debug:", graphqlDebug)
		if !graphqlDebug {
			http.Handle("/query", corsHandler.Handler(
				httplog.HandlerWithFormatter(
					httplog.DefaultLogFormatter,
					mw(srv),
				)))
		} else {
			http.Handle("/", playground.Handler("graphql", "/query"))
			http.Handle("/query", corsHandler.Handler(
				httplog.HandlerWithFormatter(
					httplog.DefaultLogFormatterWithRequestHeader,
					mw(srv),
				)))
		}

		fmt.Printf("listening on %s", config.ListenAddress)
		return http.ListenAndServe(config.ListenAddress, nil)
	},
}

func init() {
	rootCmd.AddCommand(graphqlCmd)

	graphqlCmd.Flags().BoolVar(
		&graphqlDebug,
		"debug",
		true,
		"debug enabled?",
	)
	graphqlCmd.Flags().StringVarP(
		&config.ListenAddress,
		"addr",
		"a",
		":8081",
		"listen address",
	)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// graphqlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// graphqlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func setDatabaseURI() {
	if config.DatabaseURI == "" {
		config.DatabaseURI = fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s",
			config.DatabaseUser,
			url.PathEscape(config.DatabasePassword),
			config.DatabaseHost,
			config.DatabasePort,
			config.DatabaseName,
		)
	}
}

func openDB(databaseURL string) *ent.Client {
	db, e := sql.Open("pgx", databaseURL)
	if e != nil {
		log.Fatalln(e.Error())
	}
	driver := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(driver))
}

func NewDefaultServer(es graphql.ExecutableSchema) *handler.Server {
	srv := handler.New(es)

	// Default to accepting all origins; restrict to WebsocketAllowedOrigins
	// when set.
	acceptOptions := coderws.AcceptOptions{InsecureSkipVerify: true}
	if len(config.WebsocketAllowedOrigins) > 0 {
		acceptOptions = coderws.AcceptOptions{OriginPatterns: config.WebsocketAllowedOrigins}
	}

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		InitFunc: func(ctx context.Context, p transport.InitPayload) (context.Context, *transport.InitPayload, error) {
			var auth string
			if v := p.GetString("Authorization"); v != "" {
				auth = v
			}
			if auth == "" {
				if h, ok := p["headers"].(map[string]any); ok {
					if s, ok2 := h["Authorization"].(string); ok2 {
						auth = s
					} else {
						for k, val := range h {
							if strings.ToLower(k) == "authorization" {
								if s, ok3 := val.(string); ok3 {
									auth = s
								}
								break
							}
						}
					}
				}
			}
			if auth == "" {
				return ctx, nil, nil
			}
			token := strings.TrimPrefix(auth, "Bearer ")
			if m := decodeJWTClaims(token); m != nil {
				ctx = context.WithValue(ctx, "claims", m)
				v := viewer.FromClaims(m)
				ctx = viewer.WithViewer(ctx, v)
			}
			return ctx, nil, nil
		},
		Implementation: transport.CoderWebsocketImplementation{
			AcceptOptions: acceptOptions,
		},
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	return srv
}

func decodeJWTClaims(token string) map[string]any {
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return nil
	}
	payload := parts[1]
	// Base64url decode
	// add padding if needed
	if l := len(payload) % 4; l != 0 {
		payload += strings.Repeat("=", 4-l)
	}
	b, err := base64.URLEncoding.DecodeString(payload)
	if err != nil {
		return nil
	}
	var out map[string]any
	if err := json.Unmarshal(b, &out); err != nil {
		return nil
	}
	return out
}
