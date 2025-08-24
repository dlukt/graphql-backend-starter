package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Snapshot struct {
	ent.Schema
}

func (Snapshot) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
	}
}

func (Snapshot) Fields() []ent.Field {
	return []ent.Field{
		field.Time("month"),
		field.Time("created_at").Default(time.Now),
	}
}
func (Snapshot) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
