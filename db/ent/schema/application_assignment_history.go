package schema

import (
	"entgo.io/ent"
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
		field.UUID("application_id", uuid.UUID{}),
		field.String("assigner").
			MaxLen(assignerSizeLimit),
		field.String("assignee").
			MaxLen(assigneeSizeLimit),
		field.Time("created_time").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the ApplicationAssignmentHistory.
func (ApplicationAssignmentHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("applications", Application.Type).
			Ref("assignment_histories").
			Unique().
			Required(),
	}
}
