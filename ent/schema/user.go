package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User は、User エンティティのスキーマ定義を保持します。
type User struct {
	ent.Schema
}

// Userのフィールド。
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive(),
		field.String("name").
			Default("unknown"),
	}
}

// Userのエッジ。
func (User) Edges() []ent.Edge {
	return nil
}
