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

func (TicketStatus) ToMap() map[string]int32 {
	r := make(map[string]int32)
	for i, k := range []TicketStatus{TkCreated, TkReplied, TkCompleted} {
		r[string(k)] = int32(i) + 1
	}
	return r
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
		field.Int("id").
			StorageKey("ticket_id").
			Annotations(entproto.Field(1)),

		field.UUID("application_id", uuid.UUID{}).
			Annotations(entproto.Field(2)),

		field.Enum("status").
			Values(TicketStatus("").Values()...).
			Annotations(
				entproto.Field(3),
				entproto.Enum(TicketStatus("").ToMap()),
			),

		field.String("creator").
			MaxLen(creatorSizeLimit).
			Annotations(entproto.Field(4)),

		field.Text("content").
			Annotations(entproto.Field(5)),

		field.Text("reply").
			Annotations(entproto.Field(6)),

		field.String("replier").
			MaxLen(replierSizeLimit).
			Annotations(entproto.Field(7)),

		field.String("reviewer").
			MaxLen(reviewerSizeLimit).
			Annotations(entproto.Field(8)),

		field.Text("review_remark").
			Annotations(entproto.Field(9)),

		field.Time("replied_dtime").
			Annotations(entproto.Field(10)),

		field.Time("reviewed_dtime").
			Annotations(entproto.Field(11)),

		field.Time("created_dtime").
			Default(time.Now).
			Immutable().
			Annotations(entproto.Field(12)),

		field.Time("updated_dtime").
			Annotations(entproto.Field(13)),
	}
}

// Edges of the Ticket.
func (Ticket) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("applications", Application.Type).
			Ref("tickets").
			Unique().
			Required().
			Annotations(entproto.Field(14)),

		edge.To("attachments", Attachment.Type).
			Annotations(entproto.Field(15)),
	}
}

func (Ticket) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
