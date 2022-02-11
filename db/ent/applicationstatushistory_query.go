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
	"github.com/Xanonymous-GitHub/carnival/db/ent/application"
	"github.com/Xanonymous-GitHub/carnival/db/ent/applicationstatushistory"
	"github.com/Xanonymous-GitHub/carnival/db/ent/predicate"
	"github.com/google/uuid"
)

// ApplicationStatusHistoryQuery is the builder for querying ApplicationStatusHistory entities.
type ApplicationStatusHistoryQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ApplicationStatusHistory
	// eager-loading edges.
	withApplications *ApplicationQuery
	withFKs          bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ApplicationStatusHistoryQuery builder.
func (ashq *ApplicationStatusHistoryQuery) Where(ps ...predicate.ApplicationStatusHistory) *ApplicationStatusHistoryQuery {
	ashq.predicates = append(ashq.predicates, ps...)
	return ashq
}

// Limit adds a limit step to the query.
func (ashq *ApplicationStatusHistoryQuery) Limit(limit int) *ApplicationStatusHistoryQuery {
	ashq.limit = &limit
	return ashq
}

// Offset adds an offset step to the query.
func (ashq *ApplicationStatusHistoryQuery) Offset(offset int) *ApplicationStatusHistoryQuery {
	ashq.offset = &offset
	return ashq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ashq *ApplicationStatusHistoryQuery) Unique(unique bool) *ApplicationStatusHistoryQuery {
	ashq.unique = &unique
	return ashq
}

// Order adds an order step to the query.
func (ashq *ApplicationStatusHistoryQuery) Order(o ...OrderFunc) *ApplicationStatusHistoryQuery {
	ashq.order = append(ashq.order, o...)
	return ashq
}

// QueryApplications chains the current query on the "applications" edge.
func (ashq *ApplicationStatusHistoryQuery) QueryApplications() *ApplicationQuery {
	query := &ApplicationQuery{config: ashq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ashq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ashq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(applicationstatushistory.Table, applicationstatushistory.FieldID, selector),
			sqlgraph.To(application.Table, application.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, applicationstatushistory.ApplicationsTable, applicationstatushistory.ApplicationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(ashq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ApplicationStatusHistory entity from the query.
// Returns a *NotFoundError when no ApplicationStatusHistory was found.
func (ashq *ApplicationStatusHistoryQuery) First(ctx context.Context) (*ApplicationStatusHistory, error) {
	nodes, err := ashq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{applicationstatushistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ashq *ApplicationStatusHistoryQuery) FirstX(ctx context.Context) *ApplicationStatusHistory {
	node, err := ashq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ApplicationStatusHistory ID from the query.
// Returns a *NotFoundError when no ApplicationStatusHistory ID was found.
func (ashq *ApplicationStatusHistoryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ashq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{applicationstatushistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ashq *ApplicationStatusHistoryQuery) FirstIDX(ctx context.Context) int {
	id, err := ashq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ApplicationStatusHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ApplicationStatusHistory entity is not found.
// Returns a *NotFoundError when no ApplicationStatusHistory entities are found.
func (ashq *ApplicationStatusHistoryQuery) Only(ctx context.Context) (*ApplicationStatusHistory, error) {
	nodes, err := ashq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{applicationstatushistory.Label}
	default:
		return nil, &NotSingularError{applicationstatushistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ashq *ApplicationStatusHistoryQuery) OnlyX(ctx context.Context) *ApplicationStatusHistory {
	node, err := ashq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ApplicationStatusHistory ID in the query.
// Returns a *NotSingularError when exactly one ApplicationStatusHistory ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ashq *ApplicationStatusHistoryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ashq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{applicationstatushistory.Label}
	default:
		err = &NotSingularError{applicationstatushistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ashq *ApplicationStatusHistoryQuery) OnlyIDX(ctx context.Context) int {
	id, err := ashq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ApplicationStatusHistories.
func (ashq *ApplicationStatusHistoryQuery) All(ctx context.Context) ([]*ApplicationStatusHistory, error) {
	if err := ashq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ashq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ashq *ApplicationStatusHistoryQuery) AllX(ctx context.Context) []*ApplicationStatusHistory {
	nodes, err := ashq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ApplicationStatusHistory IDs.
func (ashq *ApplicationStatusHistoryQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ashq.Select(applicationstatushistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ashq *ApplicationStatusHistoryQuery) IDsX(ctx context.Context) []int {
	ids, err := ashq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ashq *ApplicationStatusHistoryQuery) Count(ctx context.Context) (int, error) {
	if err := ashq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ashq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ashq *ApplicationStatusHistoryQuery) CountX(ctx context.Context) int {
	count, err := ashq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ashq *ApplicationStatusHistoryQuery) Exist(ctx context.Context) (bool, error) {
	if err := ashq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ashq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ashq *ApplicationStatusHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := ashq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ApplicationStatusHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ashq *ApplicationStatusHistoryQuery) Clone() *ApplicationStatusHistoryQuery {
	if ashq == nil {
		return nil
	}
	return &ApplicationStatusHistoryQuery{
		config:           ashq.config,
		limit:            ashq.limit,
		offset:           ashq.offset,
		order:            append([]OrderFunc{}, ashq.order...),
		predicates:       append([]predicate.ApplicationStatusHistory{}, ashq.predicates...),
		withApplications: ashq.withApplications.Clone(),
		// clone intermediate query.
		sql:  ashq.sql.Clone(),
		path: ashq.path,
	}
}

// WithApplications tells the query-builder to eager-load the nodes that are connected to
// the "applications" edge. The optional arguments are used to configure the query builder of the edge.
func (ashq *ApplicationStatusHistoryQuery) WithApplications(opts ...func(*ApplicationQuery)) *ApplicationStatusHistoryQuery {
	query := &ApplicationQuery{config: ashq.config}
	for _, opt := range opts {
		opt(query)
	}
	ashq.withApplications = query
	return ashq
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
//	client.ApplicationStatusHistory.Query().
//		GroupBy(applicationstatushistory.FieldApplicationID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ashq *ApplicationStatusHistoryQuery) GroupBy(field string, fields ...string) *ApplicationStatusHistoryGroupBy {
	group := &ApplicationStatusHistoryGroupBy{config: ashq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ashq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ashq.sqlQuery(ctx), nil
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
//	client.ApplicationStatusHistory.Query().
//		Select(applicationstatushistory.FieldApplicationID).
//		Scan(ctx, &v)
//
func (ashq *ApplicationStatusHistoryQuery) Select(fields ...string) *ApplicationStatusHistorySelect {
	ashq.fields = append(ashq.fields, fields...)
	return &ApplicationStatusHistorySelect{ApplicationStatusHistoryQuery: ashq}
}

func (ashq *ApplicationStatusHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ashq.fields {
		if !applicationstatushistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ashq.path != nil {
		prev, err := ashq.path(ctx)
		if err != nil {
			return err
		}
		ashq.sql = prev
	}
	return nil
}

func (ashq *ApplicationStatusHistoryQuery) sqlAll(ctx context.Context) ([]*ApplicationStatusHistory, error) {
	var (
		nodes       = []*ApplicationStatusHistory{}
		withFKs     = ashq.withFKs
		_spec       = ashq.querySpec()
		loadedTypes = [1]bool{
			ashq.withApplications != nil,
		}
	)
	if ashq.withApplications != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, applicationstatushistory.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ApplicationStatusHistory{config: ashq.config}
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
	if err := sqlgraph.QueryNodes(ctx, ashq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ashq.withApplications; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*ApplicationStatusHistory)
		for i := range nodes {
			if nodes[i].application_status_histories == nil {
				continue
			}
			fk := *nodes[i].application_status_histories
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
				return nil, fmt.Errorf(`unexpected foreign-key "application_status_histories" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Applications = n
			}
		}
	}

	return nodes, nil
}

func (ashq *ApplicationStatusHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ashq.querySpec()
	_spec.Node.Columns = ashq.fields
	if len(ashq.fields) > 0 {
		_spec.Unique = ashq.unique != nil && *ashq.unique
	}
	return sqlgraph.CountNodes(ctx, ashq.driver, _spec)
}

func (ashq *ApplicationStatusHistoryQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ashq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ashq *ApplicationStatusHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   applicationstatushistory.Table,
			Columns: applicationstatushistory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: applicationstatushistory.FieldID,
			},
		},
		From:   ashq.sql,
		Unique: true,
	}
	if unique := ashq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ashq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, applicationstatushistory.FieldID)
		for i := range fields {
			if fields[i] != applicationstatushistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ashq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ashq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ashq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ashq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ashq *ApplicationStatusHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ashq.driver.Dialect())
	t1 := builder.Table(applicationstatushistory.Table)
	columns := ashq.fields
	if len(columns) == 0 {
		columns = applicationstatushistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ashq.sql != nil {
		selector = ashq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ashq.unique != nil && *ashq.unique {
		selector.Distinct()
	}
	for _, p := range ashq.predicates {
		p(selector)
	}
	for _, p := range ashq.order {
		p(selector)
	}
	if offset := ashq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ashq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ApplicationStatusHistoryGroupBy is the group-by builder for ApplicationStatusHistory entities.
type ApplicationStatusHistoryGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ashgb *ApplicationStatusHistoryGroupBy) Aggregate(fns ...AggregateFunc) *ApplicationStatusHistoryGroupBy {
	ashgb.fns = append(ashgb.fns, fns...)
	return ashgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ashgb *ApplicationStatusHistoryGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ashgb.path(ctx)
	if err != nil {
		return err
	}
	ashgb.sql = query
	return ashgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ashgb *ApplicationStatusHistoryGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ashgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ashgb *ApplicationStatusHistoryGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ashgb.fields) > 1 {
		return nil, errors.New("ent: ApplicationStatusHistoryGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ashgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ashgb *ApplicationStatusHistoryGroupBy) StringsX(ctx context.Context) []string {
	v, err := ashgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ashgb *ApplicationStatusHistoryGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ashgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{applicationstatushistory.Label}
	default:
		err = fmt.Errorf("ent: ApplicationStatusHistoryGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ashgb *ApplicationStatusHistoryGroupBy) StringX(ctx context.Context) string {
	v, err := ashgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ashgb *ApplicationStatusHistoryGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ashgb.fields) > 1 {
		return nil, errors.New("ent: ApplicationStatusHistoryGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ashgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ashgb *ApplicationStatusHistoryGroupBy) IntsX(ctx context.Context) []int {
	v, err := ashgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ashgb *ApplicationStatusHistoryGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ashgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{applicationstatushistory.Label}
	default:
		err = fmt.Errorf("ent: ApplicationStatusHistoryGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ashgb *ApplicationStatusHistoryGroupBy) IntX(ctx context.Context) int {
	v, err := ashgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ashgb *ApplicationStatusHistoryGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ashgb.fields) > 1 {
		return nil, errors.New("ent: ApplicationStatusHistoryGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ashgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ashgb *ApplicationStatusHistoryGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ashgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ashgb *ApplicationStatusHistoryGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ashgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{applicationstatushistory.Label}
	default:
		err = fmt.Errorf("ent: ApplicationStatusHistoryGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ashgb *ApplicationStatusHistoryGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ashgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ashgb *ApplicationStatusHistoryGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ashgb.fields) > 1 {
		return nil, errors.New("ent: ApplicationStatusHistoryGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ashgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ashgb *ApplicationStatusHistoryGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ashgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ashgb *ApplicationStatusHistoryGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ashgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{applicationstatushistory.Label}
	default:
		err = fmt.Errorf("ent: ApplicationStatusHistoryGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ashgb *ApplicationStatusHistoryGroupBy) BoolX(ctx context.Context) bool {
	v, err := ashgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ashgb *ApplicationStatusHistoryGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ashgb.fields {
		if !applicationstatushistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ashgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ashgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ashgb *ApplicationStatusHistoryGroupBy) sqlQuery() *sql.Selector {
	selector := ashgb.sql.Select()
	aggregation := make([]string, 0, len(ashgb.fns))
	for _, fn := range ashgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ashgb.fields)+len(ashgb.fns))
		for _, f := range ashgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ashgb.fields...)...)
}

// ApplicationStatusHistorySelect is the builder for selecting fields of ApplicationStatusHistory entities.
type ApplicationStatusHistorySelect struct {
	*ApplicationStatusHistoryQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ashs *ApplicationStatusHistorySelect) Scan(ctx context.Context, v interface{}) error {
	if err := ashs.prepareQuery(ctx); err != nil {
		return err
	}
	ashs.sql = ashs.ApplicationStatusHistoryQuery.sqlQuery(ctx)
	return ashs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ashs *ApplicationStatusHistorySelect) ScanX(ctx context.Context, v interface{}) {
	if err := ashs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ashs *ApplicationStatusHistorySelect) Strings(ctx context.Context) ([]string, error) {
	if len(ashs.fields) > 1 {
		return nil, errors.New("ent: ApplicationStatusHistorySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ashs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ashs *ApplicationStatusHistorySelect) StringsX(ctx context.Context) []string {
	v, err := ashs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ashs *ApplicationStatusHistorySelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ashs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{applicationstatushistory.Label}
	default:
		err = fmt.Errorf("ent: ApplicationStatusHistorySelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ashs *ApplicationStatusHistorySelect) StringX(ctx context.Context) string {
	v, err := ashs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ashs *ApplicationStatusHistorySelect) Ints(ctx context.Context) ([]int, error) {
	if len(ashs.fields) > 1 {
		return nil, errors.New("ent: ApplicationStatusHistorySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ashs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ashs *ApplicationStatusHistorySelect) IntsX(ctx context.Context) []int {
	v, err := ashs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ashs *ApplicationStatusHistorySelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ashs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{applicationstatushistory.Label}
	default:
		err = fmt.Errorf("ent: ApplicationStatusHistorySelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ashs *ApplicationStatusHistorySelect) IntX(ctx context.Context) int {
	v, err := ashs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ashs *ApplicationStatusHistorySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ashs.fields) > 1 {
		return nil, errors.New("ent: ApplicationStatusHistorySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ashs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ashs *ApplicationStatusHistorySelect) Float64sX(ctx context.Context) []float64 {
	v, err := ashs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ashs *ApplicationStatusHistorySelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ashs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{applicationstatushistory.Label}
	default:
		err = fmt.Errorf("ent: ApplicationStatusHistorySelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ashs *ApplicationStatusHistorySelect) Float64X(ctx context.Context) float64 {
	v, err := ashs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ashs *ApplicationStatusHistorySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ashs.fields) > 1 {
		return nil, errors.New("ent: ApplicationStatusHistorySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ashs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ashs *ApplicationStatusHistorySelect) BoolsX(ctx context.Context) []bool {
	v, err := ashs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ashs *ApplicationStatusHistorySelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ashs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{applicationstatushistory.Label}
	default:
		err = fmt.Errorf("ent: ApplicationStatusHistorySelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ashs *ApplicationStatusHistorySelect) BoolX(ctx context.Context) bool {
	v, err := ashs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ashs *ApplicationStatusHistorySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ashs.sql.Query()
	if err := ashs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
