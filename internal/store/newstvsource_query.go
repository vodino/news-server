// Code generated by entc, DO NOT EDIT.

package store

import (
	"context"
	"errors"
	"fmt"
	"math"
	"news/internal/store/newstvsource"
	"news/internal/store/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NewsTvSourceQuery is the builder for querying NewsTvSource entities.
type NewsTvSourceQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.NewsTvSource
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NewsTvSourceQuery builder.
func (ntsq *NewsTvSourceQuery) Where(ps ...predicate.NewsTvSource) *NewsTvSourceQuery {
	ntsq.predicates = append(ntsq.predicates, ps...)
	return ntsq
}

// Limit adds a limit step to the query.
func (ntsq *NewsTvSourceQuery) Limit(limit int) *NewsTvSourceQuery {
	ntsq.limit = &limit
	return ntsq
}

// Offset adds an offset step to the query.
func (ntsq *NewsTvSourceQuery) Offset(offset int) *NewsTvSourceQuery {
	ntsq.offset = &offset
	return ntsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ntsq *NewsTvSourceQuery) Unique(unique bool) *NewsTvSourceQuery {
	ntsq.unique = &unique
	return ntsq
}

// Order adds an order step to the query.
func (ntsq *NewsTvSourceQuery) Order(o ...OrderFunc) *NewsTvSourceQuery {
	ntsq.order = append(ntsq.order, o...)
	return ntsq
}

// First returns the first NewsTvSource entity from the query.
// Returns a *NotFoundError when no NewsTvSource was found.
func (ntsq *NewsTvSourceQuery) First(ctx context.Context) (*NewsTvSource, error) {
	nodes, err := ntsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{newstvsource.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ntsq *NewsTvSourceQuery) FirstX(ctx context.Context) *NewsTvSource {
	node, err := ntsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NewsTvSource ID from the query.
// Returns a *NotFoundError when no NewsTvSource ID was found.
func (ntsq *NewsTvSourceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ntsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{newstvsource.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ntsq *NewsTvSourceQuery) FirstIDX(ctx context.Context) int {
	id, err := ntsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NewsTvSource entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one NewsTvSource entity is found.
// Returns a *NotFoundError when no NewsTvSource entities are found.
func (ntsq *NewsTvSourceQuery) Only(ctx context.Context) (*NewsTvSource, error) {
	nodes, err := ntsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{newstvsource.Label}
	default:
		return nil, &NotSingularError{newstvsource.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ntsq *NewsTvSourceQuery) OnlyX(ctx context.Context) *NewsTvSource {
	node, err := ntsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NewsTvSource ID in the query.
// Returns a *NotSingularError when more than one NewsTvSource ID is found.
// Returns a *NotFoundError when no entities are found.
func (ntsq *NewsTvSourceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ntsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{newstvsource.Label}
	default:
		err = &NotSingularError{newstvsource.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ntsq *NewsTvSourceQuery) OnlyIDX(ctx context.Context) int {
	id, err := ntsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NewsTvSources.
func (ntsq *NewsTvSourceQuery) All(ctx context.Context) ([]*NewsTvSource, error) {
	if err := ntsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ntsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ntsq *NewsTvSourceQuery) AllX(ctx context.Context) []*NewsTvSource {
	nodes, err := ntsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NewsTvSource IDs.
func (ntsq *NewsTvSourceQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ntsq.Select(newstvsource.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ntsq *NewsTvSourceQuery) IDsX(ctx context.Context) []int {
	ids, err := ntsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ntsq *NewsTvSourceQuery) Count(ctx context.Context) (int, error) {
	if err := ntsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ntsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ntsq *NewsTvSourceQuery) CountX(ctx context.Context) int {
	count, err := ntsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ntsq *NewsTvSourceQuery) Exist(ctx context.Context) (bool, error) {
	if err := ntsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ntsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ntsq *NewsTvSourceQuery) ExistX(ctx context.Context) bool {
	exist, err := ntsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NewsTvSourceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ntsq *NewsTvSourceQuery) Clone() *NewsTvSourceQuery {
	if ntsq == nil {
		return nil
	}
	return &NewsTvSourceQuery{
		config:     ntsq.config,
		limit:      ntsq.limit,
		offset:     ntsq.offset,
		order:      append([]OrderFunc{}, ntsq.order...),
		predicates: append([]predicate.NewsTvSource{}, ntsq.predicates...),
		// clone intermediate query.
		sql:    ntsq.sql.Clone(),
		path:   ntsq.path,
		unique: ntsq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Logo string `json:"logo,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.NewsTvSource.Query().
//		GroupBy(newstvsource.FieldLogo).
//		Aggregate(store.Count()).
//		Scan(ctx, &v)
//
func (ntsq *NewsTvSourceQuery) GroupBy(field string, fields ...string) *NewsTvSourceGroupBy {
	group := &NewsTvSourceGroupBy{config: ntsq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ntsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ntsq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Logo string `json:"logo,omitempty"`
//	}
//
//	client.NewsTvSource.Query().
//		Select(newstvsource.FieldLogo).
//		Scan(ctx, &v)
//
func (ntsq *NewsTvSourceQuery) Select(fields ...string) *NewsTvSourceSelect {
	ntsq.fields = append(ntsq.fields, fields...)
	return &NewsTvSourceSelect{NewsTvSourceQuery: ntsq}
}

func (ntsq *NewsTvSourceQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ntsq.fields {
		if !newstvsource.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("store: invalid field %q for query", f)}
		}
	}
	if ntsq.path != nil {
		prev, err := ntsq.path(ctx)
		if err != nil {
			return err
		}
		ntsq.sql = prev
	}
	return nil
}

func (ntsq *NewsTvSourceQuery) sqlAll(ctx context.Context) ([]*NewsTvSource, error) {
	var (
		nodes = []*NewsTvSource{}
		_spec = ntsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &NewsTvSource{config: ntsq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("store: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, ntsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ntsq *NewsTvSourceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ntsq.querySpec()
	_spec.Node.Columns = ntsq.fields
	if len(ntsq.fields) > 0 {
		_spec.Unique = ntsq.unique != nil && *ntsq.unique
	}
	return sqlgraph.CountNodes(ctx, ntsq.driver, _spec)
}

func (ntsq *NewsTvSourceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ntsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("store: check existence: %w", err)
	}
	return n > 0, nil
}

func (ntsq *NewsTvSourceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   newstvsource.Table,
			Columns: newstvsource.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: newstvsource.FieldID,
			},
		},
		From:   ntsq.sql,
		Unique: true,
	}
	if unique := ntsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ntsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, newstvsource.FieldID)
		for i := range fields {
			if fields[i] != newstvsource.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ntsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ntsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ntsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ntsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ntsq *NewsTvSourceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ntsq.driver.Dialect())
	t1 := builder.Table(newstvsource.Table)
	columns := ntsq.fields
	if len(columns) == 0 {
		columns = newstvsource.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ntsq.sql != nil {
		selector = ntsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ntsq.unique != nil && *ntsq.unique {
		selector.Distinct()
	}
	for _, p := range ntsq.predicates {
		p(selector)
	}
	for _, p := range ntsq.order {
		p(selector)
	}
	if offset := ntsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ntsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NewsTvSourceGroupBy is the group-by builder for NewsTvSource entities.
type NewsTvSourceGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ntsgb *NewsTvSourceGroupBy) Aggregate(fns ...AggregateFunc) *NewsTvSourceGroupBy {
	ntsgb.fns = append(ntsgb.fns, fns...)
	return ntsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ntsgb *NewsTvSourceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ntsgb.path(ctx)
	if err != nil {
		return err
	}
	ntsgb.sql = query
	return ntsgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ntsgb *NewsTvSourceGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ntsgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ntsgb *NewsTvSourceGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ntsgb.fields) > 1 {
		return nil, errors.New("store: NewsTvSourceGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ntsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ntsgb *NewsTvSourceGroupBy) StringsX(ctx context.Context) []string {
	v, err := ntsgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ntsgb *NewsTvSourceGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ntsgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{newstvsource.Label}
	default:
		err = fmt.Errorf("store: NewsTvSourceGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ntsgb *NewsTvSourceGroupBy) StringX(ctx context.Context) string {
	v, err := ntsgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ntsgb *NewsTvSourceGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ntsgb.fields) > 1 {
		return nil, errors.New("store: NewsTvSourceGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ntsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ntsgb *NewsTvSourceGroupBy) IntsX(ctx context.Context) []int {
	v, err := ntsgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ntsgb *NewsTvSourceGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ntsgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{newstvsource.Label}
	default:
		err = fmt.Errorf("store: NewsTvSourceGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ntsgb *NewsTvSourceGroupBy) IntX(ctx context.Context) int {
	v, err := ntsgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ntsgb *NewsTvSourceGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ntsgb.fields) > 1 {
		return nil, errors.New("store: NewsTvSourceGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ntsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ntsgb *NewsTvSourceGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ntsgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ntsgb *NewsTvSourceGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ntsgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{newstvsource.Label}
	default:
		err = fmt.Errorf("store: NewsTvSourceGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ntsgb *NewsTvSourceGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ntsgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ntsgb *NewsTvSourceGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ntsgb.fields) > 1 {
		return nil, errors.New("store: NewsTvSourceGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ntsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ntsgb *NewsTvSourceGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ntsgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ntsgb *NewsTvSourceGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ntsgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{newstvsource.Label}
	default:
		err = fmt.Errorf("store: NewsTvSourceGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ntsgb *NewsTvSourceGroupBy) BoolX(ctx context.Context) bool {
	v, err := ntsgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ntsgb *NewsTvSourceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ntsgb.fields {
		if !newstvsource.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ntsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ntsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ntsgb *NewsTvSourceGroupBy) sqlQuery() *sql.Selector {
	selector := ntsgb.sql.Select()
	aggregation := make([]string, 0, len(ntsgb.fns))
	for _, fn := range ntsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ntsgb.fields)+len(ntsgb.fns))
		for _, f := range ntsgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ntsgb.fields...)...)
}

// NewsTvSourceSelect is the builder for selecting fields of NewsTvSource entities.
type NewsTvSourceSelect struct {
	*NewsTvSourceQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ntss *NewsTvSourceSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ntss.prepareQuery(ctx); err != nil {
		return err
	}
	ntss.sql = ntss.NewsTvSourceQuery.sqlQuery(ctx)
	return ntss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ntss *NewsTvSourceSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ntss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ntss *NewsTvSourceSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ntss.fields) > 1 {
		return nil, errors.New("store: NewsTvSourceSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ntss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ntss *NewsTvSourceSelect) StringsX(ctx context.Context) []string {
	v, err := ntss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ntss *NewsTvSourceSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ntss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{newstvsource.Label}
	default:
		err = fmt.Errorf("store: NewsTvSourceSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ntss *NewsTvSourceSelect) StringX(ctx context.Context) string {
	v, err := ntss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ntss *NewsTvSourceSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ntss.fields) > 1 {
		return nil, errors.New("store: NewsTvSourceSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ntss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ntss *NewsTvSourceSelect) IntsX(ctx context.Context) []int {
	v, err := ntss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ntss *NewsTvSourceSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ntss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{newstvsource.Label}
	default:
		err = fmt.Errorf("store: NewsTvSourceSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ntss *NewsTvSourceSelect) IntX(ctx context.Context) int {
	v, err := ntss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ntss *NewsTvSourceSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ntss.fields) > 1 {
		return nil, errors.New("store: NewsTvSourceSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ntss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ntss *NewsTvSourceSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ntss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ntss *NewsTvSourceSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ntss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{newstvsource.Label}
	default:
		err = fmt.Errorf("store: NewsTvSourceSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ntss *NewsTvSourceSelect) Float64X(ctx context.Context) float64 {
	v, err := ntss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ntss *NewsTvSourceSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ntss.fields) > 1 {
		return nil, errors.New("store: NewsTvSourceSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ntss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ntss *NewsTvSourceSelect) BoolsX(ctx context.Context) []bool {
	v, err := ntss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ntss *NewsTvSourceSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ntss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{newstvsource.Label}
	default:
		err = fmt.Errorf("store: NewsTvSourceSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ntss *NewsTvSourceSelect) BoolX(ctx context.Context) bool {
	v, err := ntss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ntss *NewsTvSourceSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ntss.sql.Query()
	if err := ntss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
