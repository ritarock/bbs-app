package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("body").
			Default(""),
		field.Time("created_at").
			Default(time.Now()),
		field.Time("updated_at").
			Default(time.Now()),

		field.Int("topic_id").
			Optional(),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("topic", Topic.Type).
			Ref("comments").
			Field("topic_id").
			Unique(),
	}
}
