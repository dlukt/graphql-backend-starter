package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type PriceTier struct {
	ent.Schema
}

func (PriceTier) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
	}
}

func (PriceTier) Fields() []ent.Field {
	return []ent.Field{
		field.Int("start_month"),
		field.Int("end_month"),
		field.Int("monthly_fee_cents"),
	}
}

func (PriceTier) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("plan", Plan.Type).Ref("price_tiers").Unique().Required(),
	}
}

func (PriceTier) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
