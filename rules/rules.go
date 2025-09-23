package rules

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/deicod/oidcmw/viewer"
	"github.com/dlukt/graphql-backend-starter/ent/privacy"
	"github.com/google/uuid"
)

func PrintQueryToken() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		if v, _ := viewer.FromContext(ctx); v != nil {
			spew.Dump(v)
			return privacy.Skip
		}
		return privacy.Skip
	})
}

func DenyIfNoToken() privacy.MutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		sub := ""
		if v, _ := viewer.FromContext(ctx); v != nil && viewer.IsAuthenticated(ctx) {
			sub = v.Subject
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
