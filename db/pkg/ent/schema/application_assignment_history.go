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

// ApplicationAssignmentHistory holds the schema definition for the ApplicationAssignmentHistory entity.
type ApplicationAssignmentHistory struct {
	ent.Schema
}

// Fields of the ApplicationAssignmentHistory.
func (ApplicationAssignmentHistory) Fields() []ent.Field {
	const (
		assignerSizeLimit = 32
		assigneeSizeLimit = 32
	)

	return []ent.Field{
		field.UUID("application_id", uuid.UUID{}).
			Annotations(entproto.Field(2)),

		field.String("assigner").
			MaxLen(assignerSizeLimit).
			Annotations(entproto.Field(3)),

		field.String("assignee").
			MaxLen(assigneeSizeLimit).
			Annotations(entproto.Field(4)),

		field.Time("created_time").
			Default(time.Now).
			Immutable().
			Annotations(entproto.Field(5)),
	}
}

// Edges of the ApplicationAssignmentHistory.
func (ApplicationAssignmentHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("applications", Application.Type).
			Ref("assignment_histories").
			Field("application_id").
			Unique().
			Required().
			Annotations(entproto.Field(6)),
	}
}

func (ApplicationAssignmentHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
