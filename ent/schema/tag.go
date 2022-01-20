package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Tag struct {
	ent.Schema
}

// Fields of the Group.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Unique(),
	}
}

// Edges of the Group.
func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tag_items", Item.Type),
	}
}
