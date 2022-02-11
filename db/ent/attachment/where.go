// Code generated by entc, DO NOT EDIT.

package attachment

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Xanonymous-GitHub/carnival/db/ent/predicate"
	"github.com/Xanonymous-GitHub/carnival/db/ent/schema"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ApplicationID applies equality check predicate on the "application_id" field. It's identical to ApplicationIDEQ.
func ApplicationID(v uuid.UUID) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldApplicationID), v))
	})
}

// TicketID applies equality check predicate on the "ticket_id" field. It's identical to TicketIDEQ.
func TicketID(v int32) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTicketID), v))
	})
}

// ObsOid applies equality check predicate on the "obs_oid" field. It's identical to ObsOidEQ.
func ObsOid(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldObsOid), v))
	})
}

// ObsHash applies equality check predicate on the "obs_hash" field. It's identical to ObsHashEQ.
func ObsHash(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldObsHash), v))
	})
}

// CreatedDtime applies equality check predicate on the "created_dtime" field. It's identical to CreatedDtimeEQ.
func CreatedDtime(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedDtime), v))
	})
}

// ApplicationIDEQ applies the EQ predicate on the "application_id" field.
func ApplicationIDEQ(v uuid.UUID) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldApplicationID), v))
	})
}

// ApplicationIDNEQ applies the NEQ predicate on the "application_id" field.
func ApplicationIDNEQ(v uuid.UUID) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldApplicationID), v))
	})
}

// ApplicationIDIn applies the In predicate on the "application_id" field.
func ApplicationIDIn(vs ...uuid.UUID) predicate.Attachment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldApplicationID), v...))
	})
}

// ApplicationIDNotIn applies the NotIn predicate on the "application_id" field.
func ApplicationIDNotIn(vs ...uuid.UUID) predicate.Attachment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldApplicationID), v...))
	})
}

// ApplicationIDGT applies the GT predicate on the "application_id" field.
func ApplicationIDGT(v uuid.UUID) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldApplicationID), v))
	})
}

// ApplicationIDGTE applies the GTE predicate on the "application_id" field.
func ApplicationIDGTE(v uuid.UUID) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldApplicationID), v))
	})
}

// ApplicationIDLT applies the LT predicate on the "application_id" field.
func ApplicationIDLT(v uuid.UUID) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldApplicationID), v))
	})
}

// ApplicationIDLTE applies the LTE predicate on the "application_id" field.
func ApplicationIDLTE(v uuid.UUID) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldApplicationID), v))
	})
}

// TicketIDEQ applies the EQ predicate on the "ticket_id" field.
func TicketIDEQ(v int32) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTicketID), v))
	})
}

// TicketIDNEQ applies the NEQ predicate on the "ticket_id" field.
func TicketIDNEQ(v int32) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTicketID), v))
	})
}

// TicketIDIn applies the In predicate on the "ticket_id" field.
func TicketIDIn(vs ...int32) predicate.Attachment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTicketID), v...))
	})
}

// TicketIDNotIn applies the NotIn predicate on the "ticket_id" field.
func TicketIDNotIn(vs ...int32) predicate.Attachment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTicketID), v...))
	})
}

// TicketIDGT applies the GT predicate on the "ticket_id" field.
func TicketIDGT(v int32) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTicketID), v))
	})
}

// TicketIDGTE applies the GTE predicate on the "ticket_id" field.
func TicketIDGTE(v int32) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTicketID), v))
	})
}

// TicketIDLT applies the LT predicate on the "ticket_id" field.
func TicketIDLT(v int32) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTicketID), v))
	})
}

// TicketIDLTE applies the LTE predicate on the "ticket_id" field.
func TicketIDLTE(v int32) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTicketID), v))
	})
}

// TicketIDIsNil applies the IsNil predicate on the "ticket_id" field.
func TicketIDIsNil() predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTicketID)))
	})
}

// TicketIDNotNil applies the NotNil predicate on the "ticket_id" field.
func TicketIDNotNil() predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTicketID)))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v schema.AType) predicate.Attachment {
	vc := v
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), vc))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v schema.AType) predicate.Attachment {
	vc := v
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), vc))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...schema.AType) predicate.Attachment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...schema.AType) predicate.Attachment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// ObsOidEQ applies the EQ predicate on the "obs_oid" field.
func ObsOidEQ(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldObsOid), v))
	})
}

// ObsOidNEQ applies the NEQ predicate on the "obs_oid" field.
func ObsOidNEQ(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldObsOid), v))
	})
}

// ObsOidIn applies the In predicate on the "obs_oid" field.
func ObsOidIn(vs ...string) predicate.Attachment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldObsOid), v...))
	})
}

// ObsOidNotIn applies the NotIn predicate on the "obs_oid" field.
func ObsOidNotIn(vs ...string) predicate.Attachment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldObsOid), v...))
	})
}

// ObsOidGT applies the GT predicate on the "obs_oid" field.
func ObsOidGT(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldObsOid), v))
	})
}

// ObsOidGTE applies the GTE predicate on the "obs_oid" field.
func ObsOidGTE(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldObsOid), v))
	})
}

// ObsOidLT applies the LT predicate on the "obs_oid" field.
func ObsOidLT(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldObsOid), v))
	})
}

// ObsOidLTE applies the LTE predicate on the "obs_oid" field.
func ObsOidLTE(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldObsOid), v))
	})
}

// ObsOidContains applies the Contains predicate on the "obs_oid" field.
func ObsOidContains(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldObsOid), v))
	})
}

// ObsOidHasPrefix applies the HasPrefix predicate on the "obs_oid" field.
func ObsOidHasPrefix(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldObsOid), v))
	})
}

// ObsOidHasSuffix applies the HasSuffix predicate on the "obs_oid" field.
func ObsOidHasSuffix(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldObsOid), v))
	})
}

// ObsOidEqualFold applies the EqualFold predicate on the "obs_oid" field.
func ObsOidEqualFold(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldObsOid), v))
	})
}

// ObsOidContainsFold applies the ContainsFold predicate on the "obs_oid" field.
func ObsOidContainsFold(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldObsOid), v))
	})
}

// ObsHashEQ applies the EQ predicate on the "obs_hash" field.
func ObsHashEQ(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldObsHash), v))
	})
}

// ObsHashNEQ applies the NEQ predicate on the "obs_hash" field.
func ObsHashNEQ(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldObsHash), v))
	})
}

// ObsHashIn applies the In predicate on the "obs_hash" field.
func ObsHashIn(vs ...string) predicate.Attachment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldObsHash), v...))
	})
}

// ObsHashNotIn applies the NotIn predicate on the "obs_hash" field.
func ObsHashNotIn(vs ...string) predicate.Attachment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldObsHash), v...))
	})
}

// ObsHashGT applies the GT predicate on the "obs_hash" field.
func ObsHashGT(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldObsHash), v))
	})
}

// ObsHashGTE applies the GTE predicate on the "obs_hash" field.
func ObsHashGTE(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldObsHash), v))
	})
}

// ObsHashLT applies the LT predicate on the "obs_hash" field.
func ObsHashLT(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldObsHash), v))
	})
}

// ObsHashLTE applies the LTE predicate on the "obs_hash" field.
func ObsHashLTE(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldObsHash), v))
	})
}

// ObsHashContains applies the Contains predicate on the "obs_hash" field.
func ObsHashContains(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldObsHash), v))
	})
}

// ObsHashHasPrefix applies the HasPrefix predicate on the "obs_hash" field.
func ObsHashHasPrefix(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldObsHash), v))
	})
}

// ObsHashHasSuffix applies the HasSuffix predicate on the "obs_hash" field.
func ObsHashHasSuffix(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldObsHash), v))
	})
}

// ObsHashEqualFold applies the EqualFold predicate on the "obs_hash" field.
func ObsHashEqualFold(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldObsHash), v))
	})
}

// ObsHashContainsFold applies the ContainsFold predicate on the "obs_hash" field.
func ObsHashContainsFold(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldObsHash), v))
	})
}

// CreatedDtimeEQ applies the EQ predicate on the "created_dtime" field.
func CreatedDtimeEQ(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedDtime), v))
	})
}

// CreatedDtimeNEQ applies the NEQ predicate on the "created_dtime" field.
func CreatedDtimeNEQ(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedDtime), v))
	})
}

// CreatedDtimeIn applies the In predicate on the "created_dtime" field.
func CreatedDtimeIn(vs ...time.Time) predicate.Attachment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
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
func CreatedDtimeNotIn(vs ...time.Time) predicate.Attachment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
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
func CreatedDtimeGT(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedDtime), v))
	})
}

// CreatedDtimeGTE applies the GTE predicate on the "created_dtime" field.
func CreatedDtimeGTE(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedDtime), v))
	})
}

// CreatedDtimeLT applies the LT predicate on the "created_dtime" field.
func CreatedDtimeLT(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedDtime), v))
	})
}

// CreatedDtimeLTE applies the LTE predicate on the "created_dtime" field.
func CreatedDtimeLTE(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedDtime), v))
	})
}

// HasApplications applies the HasEdge predicate on the "applications" edge.
func HasApplications() predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ApplicationsTable, ApplicationFieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ApplicationsTable, ApplicationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasApplicationsWith applies the HasEdge predicate on the "applications" edge with a given conditions (other predicates).
func HasApplicationsWith(preds ...predicate.Application) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ApplicationsInverseTable, ApplicationFieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ApplicationsTable, ApplicationsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTickets applies the HasEdge predicate on the "tickets" edge.
func HasTickets() predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TicketsTable, TicketFieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TicketsTable, TicketsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTicketsWith applies the HasEdge predicate on the "tickets" edge with a given conditions (other predicates).
func HasTicketsWith(preds ...predicate.Ticket) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TicketsInverseTable, TicketFieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TicketsTable, TicketsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Attachment) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Attachment) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
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
func Not(p predicate.Attachment) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		p(s.Not())
	})
}
