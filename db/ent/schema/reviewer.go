package schema

import (
	"entgo.io/ent"
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
			MaxLen(reviewerIdSizeLimit),
		field.String("reviewer_name").
			MaxLen(reviewerNameSizeLimit),
		field.Enum("iims_role").
			GoType(IIMSRole("")),
		field.Time("created_dtime").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the Reviewer.
func (Reviewer) Edges() []ent.Edge {
	return nil
}
