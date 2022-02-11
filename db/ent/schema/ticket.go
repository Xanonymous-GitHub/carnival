package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type TicketStatus string

const (
	TkCreated   TicketStatus = "created"
	TkReplied   TicketStatus = "replied"
	TkCompleted TicketStatus = "completed"
)

// Values get all enum values in ApplicationStatus.
// TODO(TU): extract enum values function
func (TicketStatus) Values() (kinds []string) {
	for _, k := range []TicketStatus{TkCreated, TkReplied, TkCompleted} {
		kinds = append(kinds, string(k))
	}
	return
}

// Ticket holds the schema definition for the Ticket entity.
type Ticket struct {
	ent.Schema
}

// Fields of the Ticket.
func (Ticket) Fields() []ent.Field {
	const (
		creatorSizeLimit  = 32
		replierSizeLimit  = 32
		reviewerSizeLimit = 32
	)

	return []ent.Field{
		field.Int32("id").
			StorageKey("ticket_id"),
		field.UUID("application_id", uuid.UUID{}),
		field.Enum("status").
			GoType(TicketStatus("")),
		field.String("creator").
			MaxLen(creatorSizeLimit),
		field.Text("content"),
		field.Text("reply"),
		field.String("replier").
			MaxLen(replierSizeLimit),
		field.String("reviewer").
			MaxLen(reviewerSizeLimit),
		field.Text("review_remark"),
		field.Time("replied_dtime"),
		field.Time("reviewed_dtime"),
		field.Time("created_dtime").
			Default(time.Now).
			Immutable(),
		field.Time("updated_dtime"),
	}
}

// Edges of the Ticket.
func (Ticket) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("applications", Application.Type).
			Ref("tickets").
			Unique().
			Required(),
		edge.To("attachments", Attachment.Type),
	}
}
