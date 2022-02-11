// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Xanonymous-GitHub/carnival/db/ent/predicate"
	"github.com/Xanonymous-GitHub/carnival/db/ent/reviewer"
)

// ReviewerDelete is the builder for deleting a Reviewer entity.
type ReviewerDelete struct {
	config
	hooks    []Hook
	mutation *ReviewerMutation
}

// Where appends a list predicates to the ReviewerDelete builder.
func (rd *ReviewerDelete) Where(ps ...predicate.Reviewer) *ReviewerDelete {
	rd.mutation.Where(ps...)
	return rd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rd *ReviewerDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rd.hooks) == 0 {
		affected, err = rd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReviewerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rd.mutation = mutation
			affected, err = rd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rd.hooks) - 1; i >= 0; i-- {
			if rd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (rd *ReviewerDelete) ExecX(ctx context.Context) int {
	n, err := rd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rd *ReviewerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: reviewer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: reviewer.FieldID,
			},
		},
	}
	if ps := rd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, rd.driver, _spec)
}

// ReviewerDeleteOne is the builder for deleting a single Reviewer entity.
type ReviewerDeleteOne struct {
	rd *ReviewerDelete
}

// Exec executes the deletion query.
func (rdo *ReviewerDeleteOne) Exec(ctx context.Context) error {
	n, err := rdo.rd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{reviewer.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rdo *ReviewerDeleteOne) ExecX(ctx context.Context) {
	rdo.rd.ExecX(ctx)
}
