// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/application"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/applicationassignmenthistory"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/predicate"
	"github.com/google/uuid"
)

// ApplicationAssignmentHistoryQuery is the builder for querying ApplicationAssignmentHistory entities.
type ApplicationAssignmentHistoryQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ApplicationAssignmentHistory
	// eager-loading edges.
	withApplications *ApplicationQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ApplicationAssignmentHistoryQuery builder.
func (aahq *ApplicationAssignmentHistoryQuery) Where(ps ...predicate.ApplicationAssignmentHistory) *ApplicationAssignmentHistoryQuery {
	aahq.predicates = append(aahq.predicates, ps...)
	return aahq
}

// Limit adds a limit step to the query.
func (aahq *ApplicationAssignmentHistoryQuery) Limit(limit int) *ApplicationAssignmentHistoryQuery {
	aahq.limit = &limit
	return aahq
}

// Offset adds an offset step to the query.
func (aahq *ApplicationAssignmentHistoryQuery) Offset(offset int) *ApplicationAssignmentHistoryQuery {
	aahq.offset = &offset
	return aahq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aahq *ApplicationAssignmentHistoryQuery) Unique(unique bool) *ApplicationAssignmentHistoryQuery {
	aahq.unique = &unique
	return aahq
}

// Order adds an order step to the query.
func (aahq *ApplicationAssignmentHistoryQuery) Order(o ...OrderFunc) *ApplicationAssignmentHistoryQuery {
	aahq.order = append(aahq.order, o...)
	return aahq
}

// QueryApplications chains the current query on the "applications" edge.
func (aahq *ApplicationAssignmentHistoryQuery) QueryApplications() *ApplicationQuery {
	query := &ApplicationQuery{config: aahq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aahq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aahq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(applicationassignmenthistory.Table, applicationassignmenthistory.FieldID, selector),
			sqlgraph.To(application.Table, application.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, applicationassignmenthistory.ApplicationsTable, applicationassignmenthistory.ApplicationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(aahq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ApplicationAssignmentHistory entity from the query.
// Returns a *NotFoundError when no ApplicationAssignmentHistory was found.
func (aahq *ApplicationAssignmentHistoryQuery) First(ctx context.Context) (*ApplicationAssignmentHistory, error) {
	nodes, err := aahq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{applicationassignmenthistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aahq *ApplicationAssignmentHistoryQuery) FirstX(ctx context.Context) *ApplicationAssignmentHistory {
	node, err := aahq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ApplicationAssignmentHistory ID from the query.
// Returns a *NotFoundError when no ApplicationAssignmentHistory ID was found.
func (aahq *ApplicationAssignmentHistoryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aahq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{applicationassignmenthistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aahq *ApplicationAssignmentHistoryQuery) FirstIDX(ctx context.Context) int {
	id, err := aahq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ApplicationAssignmentHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ApplicationAssignmentHistory entity is not found.
// Returns a *NotFoundError when no ApplicationAssignmentHistory entities are found.
func (aahq *ApplicationAssignmentHistoryQuery) Only(ctx context.Context) (*ApplicationAssignmentHistory, error) {
	nodes, err := aahq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{applicationassignmenthistory.Label}
	default:
		return nil, &NotSingularError{applicationassignmenthistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aahq *ApplicationAssignmentHistoryQuery) OnlyX(ctx context.Context) *ApplicationAssignmentHistory {
	node, err := aahq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ApplicationAssignmentHistory ID in the query.
// Returns a *NotSingularError when exactly one ApplicationAssignmentHistory ID is not found.
// Returns a *NotFoundError when no entities are found.
func (aahq *ApplicationAssignmentHistoryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aahq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{applicationassignmenthistory.Label}
	default:
		err = &NotSingularError{applicationassignmenthistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aahq *ApplicationAssignmentHistoryQuery) OnlyIDX(ctx context.Context) int {
	id, err := aahq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ApplicationAssignmentHistories.
func (aahq *ApplicationAssignmentHistoryQuery) All(ctx context.Context) ([]*ApplicationAssignmentHistory, error) {
	if err := aahq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aahq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aahq *ApplicationAssignmentHistoryQuery) AllX(ctx context.Context) []*ApplicationAssignmentHistory {
	nodes, err := aahq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ApplicationAssignmentHistory IDs.
func (aahq *ApplicationAssignmentHistoryQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := aahq.Select(applicationassignmenthistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aahq *ApplicationAssignmentHistoryQuery) IDsX(ctx context.Context) []int {
	ids, err := aahq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aahq *ApplicationAssignmentHistoryQuery) Count(ctx context.Context) (int, error) {
	if err := aahq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aahq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aahq *ApplicationAssignmentHistoryQuery) CountX(ctx context.Context) int {
	count, err := aahq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aahq *ApplicationAssignmentHistoryQuery) Exist(ctx context.Context) (bool, error) {
	if err := aahq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aahq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aahq *ApplicationAssignmentHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := aahq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ApplicationAssignmentHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aahq *ApplicationAssignmentHistoryQuery) Clone() *ApplicationAssignmentHistoryQuery {
	if aahq == nil {
		return nil
	}
	return &ApplicationAssignmentHistoryQuery{
		config:           aahq.config,
		limit:            aahq.limit,
		offset:           aahq.offset,
		order:            append([]OrderFunc{}, aahq.order...),
		predicates:       append([]predicate.ApplicationAssignmentHistory{}, aahq.predicates...),
		withApplications: aahq.withApplications.Clone(),
		// clone intermediate query.
		sql:  aahq.sql.Clone(),
		path: aahq.path,
	}
}

// WithApplications tells the query-builder to eager-load the nodes that are connected to
// the "applications" edge. The optional arguments are used to configure the query builder of the edge.
func (aahq *ApplicationAssignmentHistoryQuery) WithApplications(opts ...func(*ApplicationQuery)) *ApplicationAssignmentHistoryQuery {
	query := &ApplicationQuery{config: aahq.config}
	for _, opt := range opts {
		opt(query)
	}
	aahq.withApplications = query
	return aahq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ApplicationID uuid.UUID `json:"application_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ApplicationAssignmentHistory.Query().
//		GroupBy(applicationassignmenthistory.FieldApplicationID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (aahq *ApplicationAssignmentHistoryQuery) GroupBy(field string, fields ...string) *ApplicationAssignmentHistoryGroupBy {
	group := &ApplicationAssignmentHistoryGroupBy{config: aahq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aahq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aahq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ApplicationID uuid.UUID `json:"application_id,omitempty"`
//	}
//
//	client.ApplicationAssignmentHistory.Query().
//		Select(applicationassignmenthistory.FieldApplicationID).
//		Scan(ctx, &v)
//
func (aahq *ApplicationAssignmentHistoryQuery) Select(fields ...string) *ApplicationAssignmentHistorySelect {
	aahq.fields = append(aahq.fields, fields...)
	return &ApplicationAssignmentHistorySelect{ApplicationAssignmentHistoryQuery: aahq}
}

func (aahq *ApplicationAssignmentHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aahq.fields {
		if !applicationassignmenthistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aahq.path != nil {
		prev, err := aahq.path(ctx)
		if err != nil {
			return err
		}
		aahq.sql = prev
	}
	return nil
}

func (aahq *ApplicationAssignmentHistoryQuery) sqlAll(ctx context.Context) ([]*ApplicationAssignmentHistory, error) {
	var (
		nodes       = []*ApplicationAssignmentHistory{}
		_spec       = aahq.querySpec()
		loadedTypes = [1]bool{
			aahq.withApplications != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ApplicationAssignmentHistory{config: aahq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, aahq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := aahq.withApplications; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*ApplicationAssignmentHistory)
		for i := range nodes {
			fk := nodes[i].ApplicationID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(application.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "application_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Applications = n
			}
		}
	}

	return nodes, nil
}

func (aahq *ApplicationAssignmentHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aahq.querySpec()
	_spec.Node.Columns = aahq.fields
	if len(aahq.fields) > 0 {
		_spec.Unique = aahq.unique != nil && *aahq.unique
	}
	return sqlgraph.CountNodes(ctx, aahq.driver, _spec)
}

func (aahq *ApplicationAssignmentHistoryQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := aahq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (aahq *ApplicationAssignmentHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   applicationassignmenthistory.Table,
			Columns: applicationassignmenthistory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: applicationassignmenthistory.FieldID,
			},
		},
		From:   aahq.sql,
		Unique: true,
	}
	if unique := aahq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := aahq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, applicationassignmenthistory.FieldID)
		for i := range fields {
			if fields[i] != applicationassignmenthistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aahq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aahq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aahq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aahq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aahq *ApplicationAssignmentHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aahq.driver.Dialect())
	t1 := builder.Table(applicationassignmenthistory.Table)
	columns := aahq.fields
	if len(columns) == 0 {
		columns = applicationassignmenthistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aahq.sql != nil {
		selector = aahq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aahq.unique != nil && *aahq.unique {
		selector.Distinct()
	}
	for _, p := range aahq.predicates {
		p(selector)
	}
	for _, p := range aahq.order {
		p(selector)
	}
	if offset := aahq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aahq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ApplicationAssignmentHistoryGroupBy is the group-by builder for ApplicationAssignmentHistory entities.
type ApplicationAssignmentHistoryGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aahgb *ApplicationAssignmentHistoryGroupBy) Aggregate(fns ...AggregateFunc) *ApplicationAssignmentHistoryGroupBy {
	aahgb.fns = append(aahgb.fns, fns...)
	return aahgb
}

// Scan applies the group-by query and scans the result into the given value.
func (aahgb *ApplicationAssignmentHistoryGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := aahgb.path(ctx)
	if err != nil {
		return err
	}
	aahgb.sql = query
	return aahgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (aahgb *ApplicationAssignmentHistoryGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := aahgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (aahgb *ApplicationAssignmentHistoryGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(aahgb.fields) > 1 {
		return nil, errors.New("ent: ApplicationAssignmentHistoryGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := aahgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (aahgb *ApplicationAssignmentHistoryGroupBy) StringsX(ctx context.Context) []string {
	v, err := aahgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aahgb *ApplicationAssignmentHistoryGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = aahgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{applicationassignmenthistory.Label}
	default:
		err = fmt.Errorf("ent: ApplicationAssignmentHistoryGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (aahgb *ApplicationAssignmentHistoryGroupBy) StringX(ctx context.Context) string {
	v, err := aahgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (aahgb *ApplicationAssignmentHistoryGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(aahgb.fields) > 1 {
		return nil, errors.New("ent: ApplicationAssignmentHistoryGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := aahgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (aahgb *ApplicationAssignmentHistoryGroupBy) IntsX(ctx context.Context) []int {
	v, err := aahgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aahgb *ApplicationAssignmentHistoryGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = aahgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{applicationassignmenthistory.Label}
	default:
		err = fmt.Errorf("ent: ApplicationAssignmentHistoryGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (aahgb *ApplicationAssignmentHistoryGroupBy) IntX(ctx context.Context) int {
	v, err := aahgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (aahgb *ApplicationAssignmentHistoryGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(aahgb.fields) > 1 {
		return nil, errors.New("ent: ApplicationAssignmentHistoryGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := aahgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (aahgb *ApplicationAssignmentHistoryGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := aahgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aahgb *ApplicationAssignmentHistoryGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = aahgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{applicationassignmenthistory.Label}
	default:
		err = fmt.Errorf("ent: ApplicationAssignmentHistoryGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (aahgb *ApplicationAssignmentHistoryGroupBy) Float64X(ctx context.Context) float64 {
	v, err := aahgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (aahgb *ApplicationAssignmentHistoryGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(aahgb.fields) > 1 {
		return nil, errors.New("ent: ApplicationAssignmentHistoryGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := aahgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (aahgb *ApplicationAssignmentHistoryGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := aahgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aahgb *ApplicationAssignmentHistoryGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = aahgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{applicationassignmenthistory.Label}
	default:
		err = fmt.Errorf("ent: ApplicationAssignmentHistoryGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (aahgb *ApplicationAssignmentHistoryGroupBy) BoolX(ctx context.Context) bool {
	v, err := aahgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (aahgb *ApplicationAssignmentHistoryGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range aahgb.fields {
		if !applicationassignmenthistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := aahgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aahgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (aahgb *ApplicationAssignmentHistoryGroupBy) sqlQuery() *sql.Selector {
	selector := aahgb.sql.Select()
	aggregation := make([]string, 0, len(aahgb.fns))
	for _, fn := range aahgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(aahgb.fields)+len(aahgb.fns))
		for _, f := range aahgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(aahgb.fields...)...)
}

// ApplicationAssignmentHistorySelect is the builder for selecting fields of ApplicationAssignmentHistory entities.
type ApplicationAssignmentHistorySelect struct {
	*ApplicationAssignmentHistoryQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (aahs *ApplicationAssignmentHistorySelect) Scan(ctx context.Context, v interface{}) error {
	if err := aahs.prepareQuery(ctx); err != nil {
		return err
	}
	aahs.sql = aahs.ApplicationAssignmentHistoryQuery.sqlQuery(ctx)
	return aahs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (aahs *ApplicationAssignmentHistorySelect) ScanX(ctx context.Context, v interface{}) {
	if err := aahs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (aahs *ApplicationAssignmentHistorySelect) Strings(ctx context.Context) ([]string, error) {
	if len(aahs.fields) > 1 {
		return nil, errors.New("ent: ApplicationAssignmentHistorySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := aahs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (aahs *ApplicationAssignmentHistorySelect) StringsX(ctx context.Context) []string {
	v, err := aahs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (aahs *ApplicationAssignmentHistorySelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = aahs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{applicationassignmenthistory.Label}
	default:
		err = fmt.Errorf("ent: ApplicationAssignmentHistorySelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (aahs *ApplicationAssignmentHistorySelect) StringX(ctx context.Context) string {
	v, err := aahs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (aahs *ApplicationAssignmentHistorySelect) Ints(ctx context.Context) ([]int, error) {
	if len(aahs.fields) > 1 {
		return nil, errors.New("ent: ApplicationAssignmentHistorySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := aahs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (aahs *ApplicationAssignmentHistorySelect) IntsX(ctx context.Context) []int {
	v, err := aahs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (aahs *ApplicationAssignmentHistorySelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = aahs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{applicationassignmenthistory.Label}
	default:
		err = fmt.Errorf("ent: ApplicationAssignmentHistorySelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (aahs *ApplicationAssignmentHistorySelect) IntX(ctx context.Context) int {
	v, err := aahs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (aahs *ApplicationAssignmentHistorySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(aahs.fields) > 1 {
		return nil, errors.New("ent: ApplicationAssignmentHistorySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := aahs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (aahs *ApplicationAssignmentHistorySelect) Float64sX(ctx context.Context) []float64 {
	v, err := aahs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (aahs *ApplicationAssignmentHistorySelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = aahs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{applicationassignmenthistory.Label}
	default:
		err = fmt.Errorf("ent: ApplicationAssignmentHistorySelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (aahs *ApplicationAssignmentHistorySelect) Float64X(ctx context.Context) float64 {
	v, err := aahs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (aahs *ApplicationAssignmentHistorySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(aahs.fields) > 1 {
		return nil, errors.New("ent: ApplicationAssignmentHistorySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := aahs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (aahs *ApplicationAssignmentHistorySelect) BoolsX(ctx context.Context) []bool {
	v, err := aahs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (aahs *ApplicationAssignmentHistorySelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = aahs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{applicationassignmenthistory.Label}
	default:
		err = fmt.Errorf("ent: ApplicationAssignmentHistorySelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (aahs *ApplicationAssignmentHistorySelect) BoolX(ctx context.Context) bool {
	v, err := aahs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (aahs *ApplicationAssignmentHistorySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := aahs.sql.Query()
	if err := aahs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
