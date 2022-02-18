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

type AttachmentType string

const (
	BizCert         AttachmentType = "biz_cert"
	StoreAppearance AttachmentType = "store_appearance"
	IdDocument      AttachmentType = "id_document"
	Other           AttachmentType = "other"
	Supplement      AttachmentType = "supplement"
	Reference       AttachmentType = "reference"
)

var attachmentType = []AttachmentType{
	BizCert, StoreAppearance, IdDocument, Other, Supplement, Reference,
}

// Values get all enum values in AttachmentType.
// TODO(TU): extract enum values function
func (AttachmentType) Values() []string {
	return utils.ToValuesFromEnum(attachmentType)
}

func (AttachmentType) ToMap() map[string]int32 {
	return utils.ToMapFromEnum(attachmentType)
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
		field.UUID("application_id", uuid.UUID{}).
			Annotations(entproto.Field(2)),

		field.Int("ticket_id").
			Optional().
			Nillable().
			Annotations(entproto.Field(3)),

		field.Enum("attachment_type").
			Values(new(AttachmentType).Values()...).
			Annotations(
				entproto.Field(4),
				entproto.Enum(new(AttachmentType).ToMap()),
			),

		field.String("obs_oid").
			MaxLen(obsOidSizeLimit).
			Annotations(entproto.Field(5)),

		field.String("obs_hash").
			MaxLen(obsHashSizeLimit).
			Annotations(entproto.Field(6)),

		field.Time("created_dtime").
			Default(time.Now).
			Immutable().
			Annotations(entproto.Field(7)),
	}
}

// Edges of the Attachment.
func (Attachment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("applications", Application.Type).
			Ref("attachments").
			Field("application_id").
			Unique().
			Required().
			Annotations(entproto.Field(8)),

		edge.From("tickets", Ticket.Type).
			Ref("attachments").
			Field("ticket_id").
			Unique().
			Annotations(entproto.Field(9)),
	}
}

func (Attachment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
