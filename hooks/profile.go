package hooks

import (
	"context"
	"errors"

	"github.com/deicod/oidcmw/viewer"
	"github.com/dlukt/graphql-backend-starter/ent"
	"github.com/dlukt/graphql-backend-starter/ent/hook"
)

func ProfileCreateHook(next ent.Mutator) ent.Mutator {
	return hook.ProfileFunc(func(ctx context.Context, m *ent.ProfileMutation) (ent.Value, error) {
		// Prefer viewer if present
		sub := ""
		if v, _ := viewer.FromContext(ctx); v != nil && viewer.IsAuthenticated(ctx) {
			sub = v.Subject
		}
		if sub == "" {
			return nil, errors.New("no subject in token")
		}
		m.SetSub(sub)
		return next.Mutate(ctx, m)
	})
}
