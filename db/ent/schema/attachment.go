package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type AType string

const (
	BizCert         AType = "biz_cert"
	StoreAppearance AType = "store_appearance"
	IdDocument      AType = "id_document"
	Other           AType = "other"
	Supplement      AType = "supplement"
	Reference       AType = "reference"
)

// Values get all enum values in AType.
// TODO(TU): extract enum values function
func (AType) Values() (kinds []string) {
	for _, k := range []AType{
		BizCert, StoreAppearance, IdDocument, Other, Supplement, Reference,
	} {
		kinds = append(kinds, string(k))
	}
	return
}

// Attachment holds the schema definition for the Attachment entity.
type Attachment struct {
	ent.Schema
}

// Fields of the Attachment.
func (Attachment) Fields() []ent.Field {
	const (
		obsOidSizeLimit  = 32
		obsHashSizeLimit = 128
	)

	return []ent.Field{
		field.UUID("application_id", uuid.UUID{}),
		field.Int32("ticket_id").
			Optional().
			Nillable(),
		field.Enum("type").
			GoType(AType("")),
		field.String("obs_oid").
			MaxLen(obsOidSizeLimit),
		field.String("obs_hash").
			MaxLen(obsHashSizeLimit),
		field.Time("created_dtime").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the Attachment.
func (Attachment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("applications", Application.Type).
			Ref("attachments").
			Unique().
			Required(),
		edge.From("tickets", Ticket.Type).
			Ref("attachments").
			Unique().
			Required(),
	}
}
