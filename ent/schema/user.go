package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent/schema"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Annotations(
				entproto.Field(2),
			),
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entproto.Field(3),
			),
		field.Time("updated_at").
			Optional().
			Annotations(
				entproto.Field(4),
			),
		field.Time("deleted_at").
			Optional().
			Annotations(
				entproto.Field(5),
			),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Annotations of User.
// used for GRPC delivery.
// generating proto code with Message & Service.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(entproto.PackageName("delivery")),
		entproto.Service(),
	}
}
