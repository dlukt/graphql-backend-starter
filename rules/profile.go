package rules

import (
	"context"

	"github.com/dlukt/graphql-backend-starter/ent"
	"github.com/dlukt/graphql-backend-starter/ent/privacy"
	"github.com/dlukt/graphql-backend-starter/ent/profile"
	"github.com/dlukt/graphql-backend-starter/rules/claims"
	"github.com/dlukt/graphql-backend-starter/rules/viewer"
)

func ProfileDefaultMutationRule() privacy.MutationRule {
	return privacy.ProfileMutationRuleFunc(func(ctx context.Context, m *ent.ProfileMutation) error {
		// Prefer viewer if available; fall back to raw claims.
		v := viewer.FromContext(ctx)
		claimSubject := ""
		if v != nil && v.IsAuthenticated() {
			claimSubject = v.Subject()
		}
		if claimSubject == "" {
			claimSubject = claims.SubFromContext(ctx)
		}
		if claimSubject == "" {
			return privacy.Denyf("no sub in context")
		}

		switch m.Op() {
		case ent.OpCreate:
			if claimSubject == "" {
				return privacy.Denyf("unauthenticated")
			}
			return privacy.Allow
		case ent.OpUpdateOne:
			// Accept either an explicit sub on the mutation or the old value from DB.
			if sub, ok := m.Sub(); ok {
				if sub != claimSubject {
					return privacy.Denyf("unauthorized to edit this profile")
				}
				return privacy.Allow
			}
			if oldSub, err := m.OldSub(ctx); err == nil {
				if oldSub != claimSubject {
					return privacy.Denyf("unauthorized to edit this profile")
				}
				return privacy.Allow
			}
			return privacy.Denyf("unable to determine owner for update")
		case ent.OpUpdate:
			// For safety deny bulk updates without explicit owner checks.
			return privacy.Denyf("bulk update not allowed")
		case ent.OpDeleteOne:
			// Determine owner by loading the entity by ID.
			if id, ok := m.ID(); ok {
				client := ent.FromContext(ctx)
				if client == nil {
					return privacy.Denyf("no ent client in context")
				}
				p, err := client.Profile.Get(ctx, id)
				if err != nil {
					return privacy.Denyf(err.Error())
				}
				if p.Sub != claimSubject {
					return privacy.Denyf("unauthorized to delete this profile")
				}
				return privacy.Allow
			}
			return privacy.Denyf("missing id for delete")
		case ent.OpDelete:
			// For safety deny bulk deletes.
			return privacy.Denyf("bulk delete not allowed")
		default:
			return privacy.Skip
		}
	})
}

func ProfileCreateIfNotExists() privacy.QueryRule {
	return privacy.ProfileQueryRuleFunc(func(ctx context.Context, q *ent.ProfileQuery) error {
		v := viewer.FromContext(ctx)
		claimSubject := ""
		if v != nil && v.IsAuthenticated() {
			claimSubject = v.Subject()
		}
		if claimSubject == "" {
			claimSubject = claims.SubFromContext(ctx)
		}
		if claimSubject == "" {
			return privacy.Skip
		}
		client := ent.FromContext(ctx)
		if client == nil {
			return privacy.Skip
		}
		allow := privacy.DecisionContext(ctx, privacy.Allow)
		if cnt := client.Profile.Query().Where(profile.Sub(claimSubject)).CountX(allow); cnt == 0 {
			client.Profile.Create().SetSub(claimSubject).ExecX(allow)
		}
		return privacy.Skip
	})
}
