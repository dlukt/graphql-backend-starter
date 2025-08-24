package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Promo struct {
	ent.Schema
}

func (Promo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
	}
}

func (Promo) Fields() []ent.Field {
	return []ent.Field{
		field.String("description"),
		field.Int("discount_cents"),
		field.Int("months_applies").Optional(),
		field.Time("starts_at").Optional(),
		field.Time("ends_at").Optional(),
		field.String("conditions").Optional(),
	}
}

func (Promo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("plan", Plan.Type).Ref("promos").Unique().Required(),
	}
}

func (Promo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
