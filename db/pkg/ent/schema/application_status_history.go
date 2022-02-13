package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// ApplicationStatusHistory holds the schema definition for the ApplicationStatusHistory entity.
type ApplicationStatusHistory struct {
	ent.Schema
}

// Fields of the ApplicationStatusHistory.
func (ApplicationStatusHistory) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("application_id", uuid.UUID{}).
			Annotations(entproto.Field(2)),

		field.Enum("status").
			Values(new(ApplicationStatus).Values()...).
			Annotations(
				entproto.Field(3),
				entproto.Enum(new(ApplicationStatus).ToMap()),
			),

		field.Time("created_time").
			Default(time.Now).
			Immutable().
			Annotations(entproto.Field(4)),
	}
}

// Edges of the ApplicationStatusHistory.
func (ApplicationStatusHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("applications", Application.Type).
			Ref("status_histories").
			Unique().
			Required().
			Annotations(entproto.Field(5)),
	}
}

func (ApplicationStatusHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
