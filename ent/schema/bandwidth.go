package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Bandwidth struct {
	ent.Schema
}

func (Bandwidth) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
	}
}

func (Bandwidth) Fields() []ent.Field {
	return []ent.Field{
		field.Int("down_mbps"),
		field.Int("up_mbps"),
	}
}

func (Bandwidth) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("plan", Plan.Type).Ref("bandwidth").Unique().Required(),
	}
}

func (Bandwidth) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
