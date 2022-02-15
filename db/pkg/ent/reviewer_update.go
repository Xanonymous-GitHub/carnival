// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/predicate"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/reviewer"
)

// ReviewerUpdate is the builder for updating Reviewer entities.
type ReviewerUpdate struct {
	config
	hooks    []Hook
	mutation *ReviewerMutation
}

// Where appends a list predicates to the ReviewerUpdate builder.
func (ru *ReviewerUpdate) Where(ps ...predicate.Reviewer) *ReviewerUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetReviewerID sets the "reviewer_id" field.
func (ru *ReviewerUpdate) SetReviewerID(s string) *ReviewerUpdate {
	ru.mutation.SetReviewerID(s)
	return ru
}

// SetReviewerName sets the "reviewer_name" field.
func (ru *ReviewerUpdate) SetReviewerName(s string) *ReviewerUpdate {
	ru.mutation.SetReviewerName(s)
	return ru
}

// SetIimsRole sets the "iims_role" field.
func (ru *ReviewerUpdate) SetIimsRole(rr reviewer.IimsRole) *ReviewerUpdate {
	ru.mutation.SetIimsRole(rr)
	return ru
}

// Mutation returns the ReviewerMutation object of the builder.
func (ru *ReviewerUpdate) Mutation() *ReviewerMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ReviewerUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReviewerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ReviewerUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ReviewerUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ReviewerUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *ReviewerUpdate) check() error {
	if v, ok := ru.mutation.ReviewerID(); ok {
		if err := reviewer.ReviewerIDValidator(v); err != nil {
			return &ValidationError{Name: "reviewer_id", err: fmt.Errorf(`ent: validator failed for field "Reviewer.reviewer_id": %w`, err)}
		}
	}
	if v, ok := ru.mutation.ReviewerName(); ok {
		if err := reviewer.ReviewerNameValidator(v); err != nil {
			return &ValidationError{Name: "reviewer_name", err: fmt.Errorf(`ent: validator failed for field "Reviewer.reviewer_name": %w`, err)}
		}
	}
	if v, ok := ru.mutation.IimsRole(); ok {
		if err := reviewer.IimsRoleValidator(v); err != nil {
			return &ValidationError{Name: "iims_role", err: fmt.Errorf(`ent: validator failed for field "Reviewer.iims_role": %w`, err)}
		}
	}
	return nil
}

func (ru *ReviewerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   reviewer.Table,
			Columns: reviewer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: reviewer.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.ReviewerID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: reviewer.FieldReviewerID,
		})
	}
	if value, ok := ru.mutation.ReviewerName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: reviewer.FieldReviewerName,
		})
	}
	if value, ok := ru.mutation.IimsRole(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: reviewer.FieldIimsRole,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{reviewer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ReviewerUpdateOne is the builder for updating a single Reviewer entity.
type ReviewerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ReviewerMutation
}

// SetReviewerID sets the "reviewer_id" field.
func (ruo *ReviewerUpdateOne) SetReviewerID(s string) *ReviewerUpdateOne {
	ruo.mutation.SetReviewerID(s)
	return ruo
}

// SetReviewerName sets the "reviewer_name" field.
func (ruo *ReviewerUpdateOne) SetReviewerName(s string) *ReviewerUpdateOne {
	ruo.mutation.SetReviewerName(s)
	return ruo
}

// SetIimsRole sets the "iims_role" field.
func (ruo *ReviewerUpdateOne) SetIimsRole(rr reviewer.IimsRole) *ReviewerUpdateOne {
	ruo.mutation.SetIimsRole(rr)
	return ruo
}

// Mutation returns the ReviewerMutation object of the builder.
func (ruo *ReviewerUpdateOne) Mutation() *ReviewerMutation {
	return ruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ReviewerUpdateOne) Select(field string, fields ...string) *ReviewerUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Reviewer entity.
func (ruo *ReviewerUpdateOne) Save(ctx context.Context) (*Reviewer, error) {
	var (
		err  error
		node *Reviewer
	)
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReviewerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ReviewerUpdateOne) SaveX(ctx context.Context) *Reviewer {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ReviewerUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ReviewerUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *ReviewerUpdateOne) check() error {
	if v, ok := ruo.mutation.ReviewerID(); ok {
		if err := reviewer.ReviewerIDValidator(v); err != nil {
			return &ValidationError{Name: "reviewer_id", err: fmt.Errorf(`ent: validator failed for field "Reviewer.reviewer_id": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.ReviewerName(); ok {
		if err := reviewer.ReviewerNameValidator(v); err != nil {
			return &ValidationError{Name: "reviewer_name", err: fmt.Errorf(`ent: validator failed for field "Reviewer.reviewer_name": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.IimsRole(); ok {
		if err := reviewer.IimsRoleValidator(v); err != nil {
			return &ValidationError{Name: "iims_role", err: fmt.Errorf(`ent: validator failed for field "Reviewer.iims_role": %w`, err)}
		}
	}
	return nil
}

func (ruo *ReviewerUpdateOne) sqlSave(ctx context.Context) (_node *Reviewer, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   reviewer.Table,
			Columns: reviewer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: reviewer.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Reviewer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, reviewer.FieldID)
		for _, f := range fields {
			if !reviewer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != reviewer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.ReviewerID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: reviewer.FieldReviewerID,
		})
	}
	if value, ok := ruo.mutation.ReviewerName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: reviewer.FieldReviewerName,
		})
	}
	if value, ok := ruo.mutation.IimsRole(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: reviewer.FieldIimsRole,
		})
	}
	_node = &Reviewer{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{reviewer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
