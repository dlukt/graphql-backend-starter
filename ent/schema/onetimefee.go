package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type OneTimeFee struct {
	ent.Schema
}

func (OneTimeFee) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
	}
}

func (OneTimeFee) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("kind").
			Values("activation", "shipping", "router_purchase", "misc"),
		field.Int("amount_cents"),
	}
}

func (OneTimeFee) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("plan", Plan.Type).Ref("one_time_fees").Unique().Required(),
	}
}

func (OneTimeFee) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
