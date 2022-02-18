package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/utils"
	"time"
)

type IIMSRole string

const (
	OavAdmin    IIMSRole = "oav_admin"
	OavReviewer IIMSRole = "oav_reviewer"
)

var iimsRole = []IIMSRole{OavAdmin, OavReviewer}

// Values get all enum values in IIMSRole.
// TODO(TU): extract enum values function
func (IIMSRole) Values() []string {
	return utils.ToValuesFromEnum(iimsRole)
}

func (IIMSRole) ToMap() map[string]int32 {
	return utils.ToMapFromEnum(iimsRole)
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
