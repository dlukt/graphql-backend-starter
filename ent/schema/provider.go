package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Provider struct {
	ent.Schema
}

func (Provider) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
	}
}

func (Provider) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("website"),
	}
}

func (Provider) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("plans", Plan.Type),
	}
}

func (Provider) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
