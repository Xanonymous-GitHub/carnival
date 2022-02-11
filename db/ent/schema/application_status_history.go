package schema

import (
	"entgo.io/ent"
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
		field.UUID("application_id", uuid.UUID{}),
		field.Enum("status").
			GoType(TicketStatus("")),
		field.Time("created_time").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the ApplicationStatusHistory.
func (ApplicationStatusHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("applications", Application.Type).
			Ref("status_histories").
			Unique().
			Required(),
	}
}
