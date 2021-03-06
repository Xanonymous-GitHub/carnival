// Code generated by entc, DO NOT EDIT.

package reviewer

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ReviewerID applies equality check predicate on the "reviewer_id" field. It's identical to ReviewerIDEQ.
func ReviewerID(v string) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReviewerID), v))
	})
}

// ReviewerName applies equality check predicate on the "reviewer_name" field. It's identical to ReviewerNameEQ.
func ReviewerName(v string) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReviewerName), v))
	})
}

// CreatedDtime applies equality check predicate on the "created_dtime" field. It's identical to CreatedDtimeEQ.
func CreatedDtime(v time.Time) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedDtime), v))
	})
}

// ReviewerIDEQ applies the EQ predicate on the "reviewer_id" field.
func ReviewerIDEQ(v string) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReviewerID), v))
	})
}

// ReviewerIDNEQ applies the NEQ predicate on the "reviewer_id" field.
func ReviewerIDNEQ(v string) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReviewerID), v))
	})
}

// ReviewerIDIn applies the In predicate on the "reviewer_id" field.
func ReviewerIDIn(vs ...string) predicate.Reviewer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Reviewer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldReviewerID), v...))
	})
}

// ReviewerIDNotIn applies the NotIn predicate on the "reviewer_id" field.
func ReviewerIDNotIn(vs ...string) predicate.Reviewer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Reviewer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldReviewerID), v...))
	})
}

// ReviewerIDGT applies the GT predicate on the "reviewer_id" field.
func ReviewerIDGT(v string) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReviewerID), v))
	})
}

// ReviewerIDGTE applies the GTE predicate on the "reviewer_id" field.
func ReviewerIDGTE(v string) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReviewerID), v))
	})
}

// ReviewerIDLT applies the LT predicate on the "reviewer_id" field.
func ReviewerIDLT(v string) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReviewerID), v))
	})
}

// ReviewerIDLTE applies the LTE predicate on the "reviewer_id" field.
func ReviewerIDLTE(v string) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReviewerID), v))
	})
}

// ReviewerIDContains applies the Contains predicate on the "reviewer_id" field.
func ReviewerIDContains(v string) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldReviewerID), v))
	})
}

// ReviewerIDHasPrefix applies the HasPrefix predicate on the "reviewer_id" field.
func ReviewerIDHasPrefix(v string) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldReviewerID), v))
	})
}

// ReviewerIDHasSuffix applies the HasSuffix predicate on the "reviewer_id" field.
func ReviewerIDHasSuffix(v string) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldReviewerID), v))
	})
}

// ReviewerIDEqualFold applies the EqualFold predicate on the "reviewer_id" field.
func ReviewerIDEqualFold(v string) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldReviewerID), v))
	})
}

// ReviewerIDContainsFold applies the ContainsFold predicate on the "reviewer_id" field.
func ReviewerIDContainsFold(v string) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldReviewerID), v))
	})
}

// ReviewerNameEQ applies the EQ predicate on the "reviewer_name" field.
func ReviewerNameEQ(v string) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReviewerName), v))
	})
}

// ReviewerNameNEQ applies the NEQ predicate on the "reviewer_name" field.
func ReviewerNameNEQ(v string) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReviewerName), v))
	})
}

// ReviewerNameIn applies the In predicate on the "reviewer_name" field.
func ReviewerNameIn(vs ...string) predicate.Reviewer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Reviewer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldReviewerName), v...))
	})
}

// ReviewerNameNotIn applies the NotIn predicate on the "reviewer_name" field.
func ReviewerNameNotIn(vs ...string) predicate.Reviewer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Reviewer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldReviewerName), v...))
	})
}

// ReviewerNameGT applies the GT predicate on the "reviewer_name" field.
func ReviewerNameGT(v string) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReviewerName), v))
	})
}

// ReviewerNameGTE applies the GTE predicate on the "reviewer_name" field.
func ReviewerNameGTE(v string) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReviewerName), v))
	})
}

// ReviewerNameLT applies the LT predicate on the "reviewer_name" field.
func ReviewerNameLT(v string) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReviewerName), v))
	})
}

// ReviewerNameLTE applies the LTE predicate on the "reviewer_name" field.
func ReviewerNameLTE(v string) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReviewerName), v))
	})
}

// ReviewerNameContains applies the Contains predicate on the "reviewer_name" field.
func ReviewerNameContains(v string) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldReviewerName), v))
	})
}

// ReviewerNameHasPrefix applies the HasPrefix predicate on the "reviewer_name" field.
func ReviewerNameHasPrefix(v string) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldReviewerName), v))
	})
}

// ReviewerNameHasSuffix applies the HasSuffix predicate on the "reviewer_name" field.
func ReviewerNameHasSuffix(v string) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldReviewerName), v))
	})
}

// ReviewerNameEqualFold applies the EqualFold predicate on the "reviewer_name" field.
func ReviewerNameEqualFold(v string) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldReviewerName), v))
	})
}

// ReviewerNameContainsFold applies the ContainsFold predicate on the "reviewer_name" field.
func ReviewerNameContainsFold(v string) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldReviewerName), v))
	})
}

// IimsRoleEQ applies the EQ predicate on the "iims_role" field.
func IimsRoleEQ(v IimsRole) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIimsRole), v))
	})
}

// IimsRoleNEQ applies the NEQ predicate on the "iims_role" field.
func IimsRoleNEQ(v IimsRole) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIimsRole), v))
	})
}

// IimsRoleIn applies the In predicate on the "iims_role" field.
func IimsRoleIn(vs ...IimsRole) predicate.Reviewer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Reviewer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIimsRole), v...))
	})
}

// IimsRoleNotIn applies the NotIn predicate on the "iims_role" field.
func IimsRoleNotIn(vs ...IimsRole) predicate.Reviewer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Reviewer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIimsRole), v...))
	})
}

// CreatedDtimeEQ applies the EQ predicate on the "created_dtime" field.
func CreatedDtimeEQ(v time.Time) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedDtime), v))
	})
}

// CreatedDtimeNEQ applies the NEQ predicate on the "created_dtime" field.
func CreatedDtimeNEQ(v time.Time) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedDtime), v))
	})
}

// CreatedDtimeIn applies the In predicate on the "created_dtime" field.
func CreatedDtimeIn(vs ...time.Time) predicate.Reviewer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Reviewer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedDtime), v...))
	})
}

// CreatedDtimeNotIn applies the NotIn predicate on the "created_dtime" field.
func CreatedDtimeNotIn(vs ...time.Time) predicate.Reviewer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Reviewer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedDtime), v...))
	})
}

// CreatedDtimeGT applies the GT predicate on the "created_dtime" field.
func CreatedDtimeGT(v time.Time) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedDtime), v))
	})
}

// CreatedDtimeGTE applies the GTE predicate on the "created_dtime" field.
func CreatedDtimeGTE(v time.Time) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedDtime), v))
	})
}

// CreatedDtimeLT applies the LT predicate on the "created_dtime" field.
func CreatedDtimeLT(v time.Time) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedDtime), v))
	})
}

// CreatedDtimeLTE applies the LTE predicate on the "created_dtime" field.
func CreatedDtimeLTE(v time.Time) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedDtime), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Reviewer) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Reviewer) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Reviewer) predicate.Reviewer {
	return predicate.Reviewer(func(s *sql.Selector) {
		p(s.Not())
	})
}
