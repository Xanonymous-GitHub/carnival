package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

type IIMSRole string

const (
	OavAdmin    IIMSRole = "oav_admin"
	OavReviewer IIMSRole = "oav_reviewer"
)

// Values get all enum values in IIMSRole.
// TODO(TU): extract enum values function
func (IIMSRole) Values() (kinds []string) {
	for _, k := range []IIMSRole{OavAdmin, OavReviewer} {
		kinds = append(kinds, string(k))
	}
	return
}

func (IIMSRole) ToMap() map[string]int32 {
	r := make(map[string]int32)
	for i, k := range []IIMSRole{OavAdmin, OavReviewer} {
		r[string(k)] = int32(i) + 1
	}
	return r
}

// Reviewer holds the schema definition for the Reviewer entity.
type Reviewer struct {
	ent.Schema
}

// Fields of the Reviewer.
func (Reviewer) Fields() []ent.Field {
	const (
		reviewerIdSizeLimit   = 10
		reviewerNameSizeLimit = 32
	)

	return []ent.Field{
		field.String("reviewer_id").
			MaxLen(reviewerIdSizeLimit).
			Unique().
			Annotations(entproto.Field(2)),

		field.String("reviewer_name").
			MaxLen(reviewerNameSizeLimit).
			Annotations(entproto.Field(3)),

		field.Enum("iims_role").
			Values(new(IIMSRole).Values()...).
			Annotations(
				entproto.Field(4),
				entproto.Enum(new(IIMSRole).ToMap()),
			),

		field.Time("created_dtime").
			Default(time.Now).
			Immutable().
			Annotations(entproto.Field(5)),
	}
}

// Edges of the Reviewer.
func (Reviewer) Edges() []ent.Edge {
	return nil
}

func (Reviewer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
