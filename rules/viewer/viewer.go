package viewer

import (
    "context"

    "github.com/dlukt/graphql-backend-starter/rules/claims"
)

// Viewer represents the authenticated caller derived from OIDC claims.
// It flattens common attributes for convenient authorization checks.
type Viewer struct {
    // Claims are the raw OIDC claims (Keycloak-shaped by default).
    Claims *claims.Claims

    // Roles is a flattened union of realm and resource roles.
    Roles []string
}

// NewFromClaims constructs a Viewer from Claims, flattening roles.
func NewFromClaims(c *claims.Claims) *Viewer {
    if c == nil {
        return nil
    }
    v := &Viewer{Claims: c}
    // Collect roles from known places (Keycloak standard locations).
    // Ignore duplicates.
    seen := map[string]struct{}{}
    add := func(rs []string) {
        for _, r := range rs {
            if _, ok := seen[r]; ok {
                continue
            }
            seen[r] = struct{}{}
            v.Roles = append(v.Roles, r)
        }
    }
    add(c.RealmAccess.Roles)
    add(c.ResourceAccess.Account.Roles)
    return v
}

// NewUnauthenticated returns a Viewer without claims.
func NewUnauthenticated() *Viewer {
    return &Viewer{Claims: nil, Roles: nil}
}

// IsAuthenticated indicates whether this viewer has a subject.
func (v *Viewer) IsAuthenticated() bool {
    return v != nil && v.Claims != nil && v.Claims.Sub != ""
}

// Subject returns the subject (user id) or empty string.
func (v *Viewer) Subject() string {
    if v == nil || v.Claims == nil {
        return ""
    }
    return v.Claims.Sub
}

// HasRole checks membership in flattened roles.
func (v *Viewer) HasRole(role string) bool {
    if v == nil {
        return false
    }
    for _, r := range v.Roles {
        if r == role {
            return true
        }
    }
    return false
}

// contextKey is an unexported type for keys defined in this package.
type contextKey struct{}

var viewerKey = contextKey{}

// NewContext stores the viewer in the context.
func NewContext(ctx context.Context, v *Viewer) context.Context {
    return context.WithValue(ctx, viewerKey, v)
}

// FromContext extracts the viewer from context, if any.
func FromContext(ctx context.Context) *Viewer {
    if ctx == nil {
        return nil
    }
    v, _ := ctx.Value(viewerKey).(*Viewer)
    return v
}

