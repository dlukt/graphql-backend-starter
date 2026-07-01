# gqlgen with ent starter

This repository is a gqlgen with ent backend starter template.
Initial configuration is time-consuming and complicated.
This is here to make things easier and help people get started with GraphQL and ent.

## Assumptions

- You use [Keycloak](https://www.keycloak.org/) as the OIDC IDP.
  It can be used with any IDP, but a custom viewer would need to be constructed.
  You will have to add the audience in Keycloak to the token, because Keycloak is dumb like that.
- [xid](https://github.com/rs/xid) is used for globally unique IDs
- The `Profile` schema is the root of your related entities.
- Schema introspection is enabled
- PostgreSQL is used as the database. Optionally SQLite in-memory mode can be used or development purposes.
- The Schema is automigrated on each graphql run.
- Relay is used to provide filtering, ordering and pagination

## How to run the backend

### dev mode with SQLite

```bash
go run main.go graphql --sqlite=true --debug=true
```

### dev mode with PostgreSQL

```bash
go run main.go graphql \
--debug=true \
--db_uri="postgres://root:password@localhost:port/?sslmode=disable"

#or

go run main.go graphql \
--debug=true \
--db_user="username" \
--db_password="password" \
--db_name="database name" \
--db_host="localhost" \
--db_port="5432"
```

### restricting websocket origins

By default the websocket transport accepts connections from any origin. Restrict
this in production with `--ws_allowed_origins` (comma-separated host patterns):

```bash
go run main.go graphql \
--debug=true \
--ws_allowed_origins="app.example.com,*.example.com"
```


## Getting started

Let's assume your project is at github.com/user/repo

### 1. Clone the repository

```bash
mkdir -p ~/go/src/github.com/user
cd ~/go/src/github.com/user
git clone https://github.com/dlukt/graphql-backend-starter.git repo
cd repo
```

### 2. Replace the original repository name with your repository name

```bash
chmod +x ./update-repo.sh
./update-repo.sh github.com/user/repo
```

### 3. remove all "starter" occurences and replace with your repo name

```bash
# remember, repo is your repo name
rm graph/generated/starter.generated.go
mv starter.graphql repo.graphql
mv graph/starter.resolvers.go graph/repo.resolvers.go
```

edit `gqlgen.yml` and replace the `- starter.graphql` with your `- repo.graphql`

```yaml
schema:
  - ent.graphql
  - repo.graphql
```

Regenerate

```bash
go generate ./...
```

Change the project name in `cmd/root.go`, line 22.

### 4. Change the git repo url

```bash
rm -rf .git
git init
git add .
git remote add origin github.com/user/repo
git commit -m 'initial'
git push -u origin master
```

## Adding new entities

```bash
alias ent='go run -mod=mod entgo.io/ent/cmd/ent'
```

`cd` into your project root.

```bash
ent new Entity # capitalization matters
```

add the new entity to `gqlgen.yml`

```yaml
autobind:
  - github.com/user/repo/ent
  - github.com/user/repo/ent/profile
  - github.com/user/repo/ent/entity
```

and edit the `ent/schema/entity.go` file.
Afterwards, regenerate all the things.

```bash
go generate ./...
```

## OIDC

Uses [OIDC Middleware Guard](https://github.com/deicod/oidcmw) and its viewer with helper functions.
It assumes a Keycloak claims structure.
If you need the claims `viewer.RawClaims` which returns a `map[string]any`.

## Further reading

[Ent.io GraphQL Tutorial](https://entgo.io/docs/tutorial-todo-gql)

[React Relay](https://relay.dev)
