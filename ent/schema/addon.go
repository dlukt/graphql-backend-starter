package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Addon struct {
	ent.Schema
}

func (Addon) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
	}
}

func (Addon) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("monthly_fee_cents"),
	}
}

func (Addon) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("plans", Plan.Type),
	}
}

func (Addon) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
