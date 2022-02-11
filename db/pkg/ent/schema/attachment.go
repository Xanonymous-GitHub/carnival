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

func (AType) ToMap() map[string]int32 {
	r := make(map[string]int32)
	for i, k := range []AType{
		BizCert, StoreAppearance, IdDocument, Other, Supplement, Reference,
	} {
		r[string(k)] = int32(i) + 1
	}
	return r
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

		field.Int32("ticket_id").
			Optional().
			Nillable().
			Annotations(entproto.Field(3)),

		field.Enum("a_type").
			Values(AType("").Values()...).
			Annotations(
				entproto.Field(4),
				entproto.Enum(AType("").ToMap()),
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
			Unique().
			Required().
			Annotations(entproto.Field(8)),

		edge.From("tickets", Ticket.Type).
			Ref("attachments").
			Unique().
			Required().
			Annotations(entproto.Field(9)),
	}
}

func (Attachment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
