package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Product struct {
	ent.Schema
}

func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id").Unique().Immutable(),
		field.String("name").NotEmpty().Comment("商品名"),
		field.String("description").NotEmpty().Optional().Comment("商品説明"),
		field.Float("price").Min(0).Comment("商品価格"),
		field.Uint8("stock").Min(0).Comment("在庫数"),
	}
}

func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("orders", Order.Type),
	}
}

func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
