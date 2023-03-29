package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Order struct {
	ent.Schema
}

func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id").Unique().Immutable(),
		field.Uint32("product_id").Comment("商品識別子"),
		field.Int("user_id").Comment("ユーザー識別子"),
		field.Uint8("quantity").Comment("注文数"),
		field.Float("total_price").Min(0).Comment("合計金額"),
	}
}

func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("products", Product.Type).Ref("orders").Unique().Required().Field("product_id"),
	}
}

func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
