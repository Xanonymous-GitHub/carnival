package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/utils"
	"github.com/google/uuid"
	"time"
)

type TicketStatus string

const (
	TkCreated   TicketStatus = "created"
	TkReplied   TicketStatus = "replied"
	TkCompleted TicketStatus = "completed"
)

var ticketStatus = []TicketStatus{TkCreated, TkReplied, TkCompleted}

// Values get all enum values in ApplicationStatus.
// TODO(TU): extract enum values function
func (TicketStatus) Values() []string {
	return utils.ToValuesFromEnum(ticketStatus)
}

func (TicketStatus) ToMap() map[string]int32 {
	return utils.ToMapFromEnum(ticketStatus)
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
			Values(new(TicketStatus).Values()...).
			Annotations(
				entproto.Field(3),
				entproto.Enum(new(TicketStatus).ToMap()),
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
			Field("application_id").
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
