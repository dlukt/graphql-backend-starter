package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Plan struct {
	ent.Schema
}

func (Plan) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
	}
}

func (Plan) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("technology").
			Values("FTTH", "DSL", "CABLE", "MOBILE"),
		field.String("name"),
		field.String("description").Optional(),
		field.Int("min_term_months").Default(24),
		field.Int("cancel_notice_days").Optional(),
		field.Time("valid_from"),
		field.Time("valid_to").Optional(),
		field.String("source_url"),
		field.String("version_tag"),
	}
}

func (Plan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("provider", Provider.Type).Ref("plans").Unique().Required(),
		edge.To("bandwidth", Bandwidth.Type).Unique(),
		edge.To("price_tiers", PriceTier.Type),
		edge.To("one_time_fees", OneTimeFee.Type),
		edge.To("promos", Promo.Type),
		edge.From("addon", Addon.Type).Ref("plans"),
	}
}

func (Plan) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
