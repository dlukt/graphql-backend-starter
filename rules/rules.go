package rules

import (
    "context"

    "github.com/dlukt/graphql-backend-starter/ent/privacy"
    "github.com/dlukt/graphql-backend-starter/rules/claims"
    "github.com/dlukt/graphql-backend-starter/rules/viewer"
    "github.com/davecgh/go-spew/spew"
    "github.com/google/uuid"
)

func PrintQueryToken() privacy.QueryMutationRule {
    return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
        if v := viewer.FromContext(ctx); v != nil {
            spew.Dump(v)
            return privacy.Skip
        }
        c := claims.FromContext(ctx)
        if c == nil {
            return privacy.Skip
        }
        spew.Dump(c)
        return privacy.Skip
    })
}

func DenyIfNoToken() privacy.MutationRule {
    return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
        sub := ""
        if v := viewer.FromContext(ctx); v != nil && v.IsAuthenticated() {
            sub = v.Subject()
        }
        if sub == "" {
            sub = claims.SubFromContext(ctx)
        }
        if sub == "" {
            return privacy.Denyf("unauthenticated")
        }
        _, e := uuid.Parse(sub)
        if e != nil {
			return privacy.Denyf(e.Error())
		}
		return privacy.Skip
	})
}
