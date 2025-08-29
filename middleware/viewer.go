package middleware

import (
    "net/http"

    "github.com/dlukt/graphql-backend-starter/rules/claims"
    "github.com/dlukt/graphql-backend-starter/rules/viewer"
)

// WithViewer attaches a viewer to the request context derived from OIDC claims.
// It assumes an upstream auth layer (e.g., OIDC middleware) has already
// validated the token and placed claims in the context.
func WithViewer(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        c := claims.FromContext(r.Context())
        var v *viewer.Viewer
        if c != nil {
            v = viewer.NewFromClaims(c)
        } else {
            v = viewer.NewUnauthenticated()
        }
        ctx := viewer.NewContext(r.Context(), v)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

