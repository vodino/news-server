// Code generated by entc, DO NOT EDIT.

package store

import (
	"context"
	"errors"
	"fmt"
	"news/internal/store/newsarticlesource"
	"news/internal/store/predicate"
	"news/internal/store/schema"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NewsArticleSourceUpdate is the builder for updating NewsArticleSource entities.
type NewsArticleSourceUpdate struct {
	config
	hooks    []Hook
	mutation *NewsArticleSourceMutation
}

// Where appends a list predicates to the NewsArticleSourceUpdate builder.
func (nasu *NewsArticleSourceUpdate) Where(ps ...predicate.NewsArticleSource) *NewsArticleSourceUpdate {
	nasu.mutation.Where(ps...)
	return nasu
}

// SetLatestPostURL sets the "latest_post_url" field.
func (nasu *NewsArticleSourceUpdate) SetLatestPostURL(s string) *NewsArticleSourceUpdate {
	nasu.mutation.SetLatestPostURL(s)
	return nasu
}

// SetNillableLatestPostURL sets the "latest_post_url" field if the given value is not nil.
func (nasu *NewsArticleSourceUpdate) SetNillableLatestPostURL(s *string) *NewsArticleSourceUpdate {
	if s != nil {
		nasu.SetLatestPostURL(*s)
	}
	return nasu
}

// ClearLatestPostURL clears the value of the "latest_post_url" field.
func (nasu *NewsArticleSourceUpdate) ClearLatestPostURL() *NewsArticleSourceUpdate {
	nasu.mutation.ClearLatestPostURL()
	return nasu
}

// SetLatestPostSelector sets the "latest_post_selector" field.
func (nasu *NewsArticleSourceUpdate) SetLatestPostSelector(sps *schema.NewsPostSelector) *NewsArticleSourceUpdate {
	nasu.mutation.SetLatestPostSelector(sps)
	return nasu
}

// ClearLatestPostSelector clears the value of the "latest_post_selector" field.
func (nasu *NewsArticleSourceUpdate) ClearLatestPostSelector() *NewsArticleSourceUpdate {
	nasu.mutation.ClearLatestPostSelector()
	return nasu
}

// SetCategoryPostURL sets the "category_post_url" field.
func (nasu *NewsArticleSourceUpdate) SetCategoryPostURL(s string) *NewsArticleSourceUpdate {
	nasu.mutation.SetCategoryPostURL(s)
	return nasu
}

// SetNillableCategoryPostURL sets the "category_post_url" field if the given value is not nil.
func (nasu *NewsArticleSourceUpdate) SetNillableCategoryPostURL(s *string) *NewsArticleSourceUpdate {
	if s != nil {
		nasu.SetCategoryPostURL(*s)
	}
	return nasu
}

// ClearCategoryPostURL clears the value of the "category_post_url" field.
func (nasu *NewsArticleSourceUpdate) ClearCategoryPostURL() *NewsArticleSourceUpdate {
	nasu.mutation.ClearCategoryPostURL()
	return nasu
}

// SetCategoryPostSelector sets the "category_post_selector" field.
func (nasu *NewsArticleSourceUpdate) SetCategoryPostSelector(sps *schema.NewsPostSelector) *NewsArticleSourceUpdate {
	nasu.mutation.SetCategoryPostSelector(sps)
	return nasu
}

// ClearCategoryPostSelector clears the value of the "category_post_selector" field.
func (nasu *NewsArticleSourceUpdate) ClearCategoryPostSelector() *NewsArticleSourceUpdate {
	nasu.mutation.ClearCategoryPostSelector()
	return nasu
}

// SetArticleSelector sets the "article_selector" field.
func (nasu *NewsArticleSourceUpdate) SetArticleSelector(sas *schema.NewsArticleSelector) *NewsArticleSourceUpdate {
	nasu.mutation.SetArticleSelector(sas)
	return nasu
}

// ClearArticleSelector clears the value of the "article_selector" field.
func (nasu *NewsArticleSourceUpdate) ClearArticleSelector() *NewsArticleSourceUpdate {
	nasu.mutation.ClearArticleSelector()
	return nasu
}

// SetCategories sets the "categories" field.
func (nasu *NewsArticleSourceUpdate) SetCategories(m map[string]string) *NewsArticleSourceUpdate {
	nasu.mutation.SetCategories(m)
	return nasu
}

// SetLanguage sets the "language" field.
func (nasu *NewsArticleSourceUpdate) SetLanguage(s string) *NewsArticleSourceUpdate {
	nasu.mutation.SetLanguage(s)
	return nasu
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (nasu *NewsArticleSourceUpdate) SetNillableLanguage(s *string) *NewsArticleSourceUpdate {
	if s != nil {
		nasu.SetLanguage(*s)
	}
	return nasu
}

// SetCountry sets the "country" field.
func (nasu *NewsArticleSourceUpdate) SetCountry(s string) *NewsArticleSourceUpdate {
	nasu.mutation.SetCountry(s)
	return nasu
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (nasu *NewsArticleSourceUpdate) SetNillableCountry(s *string) *NewsArticleSourceUpdate {
	if s != nil {
		nasu.SetCountry(*s)
	}
	return nasu
}

// SetStatus sets the "status" field.
func (nasu *NewsArticleSourceUpdate) SetStatus(b bool) *NewsArticleSourceUpdate {
	nasu.mutation.SetStatus(b)
	return nasu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (nasu *NewsArticleSourceUpdate) SetNillableStatus(b *bool) *NewsArticleSourceUpdate {
	if b != nil {
		nasu.SetStatus(*b)
	}
	return nasu
}

// SetLogo sets the "logo" field.
func (nasu *NewsArticleSourceUpdate) SetLogo(s string) *NewsArticleSourceUpdate {
	nasu.mutation.SetLogo(s)
	return nasu
}

// SetName sets the "name" field.
func (nasu *NewsArticleSourceUpdate) SetName(s string) *NewsArticleSourceUpdate {
	nasu.mutation.SetName(s)
	return nasu
}

// SetURL sets the "url" field.
func (nasu *NewsArticleSourceUpdate) SetURL(s string) *NewsArticleSourceUpdate {
	nasu.mutation.SetURL(s)
	return nasu
}

// Mutation returns the NewsArticleSourceMutation object of the builder.
func (nasu *NewsArticleSourceUpdate) Mutation() *NewsArticleSourceMutation {
	return nasu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nasu *NewsArticleSourceUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(nasu.hooks) == 0 {
		if err = nasu.check(); err != nil {
			return 0, err
		}
		affected, err = nasu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NewsArticleSourceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nasu.check(); err != nil {
				return 0, err
			}
			nasu.mutation = mutation
			affected, err = nasu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nasu.hooks) - 1; i >= 0; i-- {
			if nasu.hooks[i] == nil {
				return 0, fmt.Errorf("store: uninitialized hook (forgotten import store/runtime?)")
			}
			mut = nasu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nasu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (nasu *NewsArticleSourceUpdate) SaveX(ctx context.Context) int {
	affected, err := nasu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nasu *NewsArticleSourceUpdate) Exec(ctx context.Context) error {
	_, err := nasu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nasu *NewsArticleSourceUpdate) ExecX(ctx context.Context) {
	if err := nasu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nasu *NewsArticleSourceUpdate) check() error {
	if v, ok := nasu.mutation.Logo(); ok {
		if err := newsarticlesource.LogoValidator(v); err != nil {
			return &ValidationError{Name: "logo", err: fmt.Errorf(`store: validator failed for field "NewsArticleSource.logo": %w`, err)}
		}
	}
	return nil
}

func (nasu *NewsArticleSourceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   newsarticlesource.Table,
			Columns: newsarticlesource.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: newsarticlesource.FieldID,
			},
		},
	}
	if ps := nasu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nasu.mutation.LatestPostURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: newsarticlesource.FieldLatestPostURL,
		})
	}
	if nasu.mutation.LatestPostURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: newsarticlesource.FieldLatestPostURL,
		})
	}
	if value, ok := nasu.mutation.LatestPostSelector(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: newsarticlesource.FieldLatestPostSelector,
		})
	}
	if nasu.mutation.LatestPostSelectorCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: newsarticlesource.FieldLatestPostSelector,
		})
	}
	if value, ok := nasu.mutation.CategoryPostURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: newsarticlesource.FieldCategoryPostURL,
		})
	}
	if nasu.mutation.CategoryPostURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: newsarticlesource.FieldCategoryPostURL,
		})
	}
	if value, ok := nasu.mutation.CategoryPostSelector(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: newsarticlesource.FieldCategoryPostSelector,
		})
	}
	if nasu.mutation.CategoryPostSelectorCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: newsarticlesource.FieldCategoryPostSelector,
		})
	}
	if value, ok := nasu.mutation.ArticleSelector(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: newsarticlesource.FieldArticleSelector,
		})
	}
	if nasu.mutation.ArticleSelectorCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: newsarticlesource.FieldArticleSelector,
		})
	}
	if value, ok := nasu.mutation.Categories(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: newsarticlesource.FieldCategories,
		})
	}
	if value, ok := nasu.mutation.Language(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: newsarticlesource.FieldLanguage,
		})
	}
	if value, ok := nasu.mutation.Country(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: newsarticlesource.FieldCountry,
		})
	}
	if value, ok := nasu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: newsarticlesource.FieldStatus,
		})
	}
	if value, ok := nasu.mutation.Logo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: newsarticlesource.FieldLogo,
		})
	}
	if value, ok := nasu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: newsarticlesource.FieldName,
		})
	}
	if value, ok := nasu.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: newsarticlesource.FieldURL,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nasu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{newsarticlesource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// NewsArticleSourceUpdateOne is the builder for updating a single NewsArticleSource entity.
type NewsArticleSourceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NewsArticleSourceMutation
}

// SetLatestPostURL sets the "latest_post_url" field.
func (nasuo *NewsArticleSourceUpdateOne) SetLatestPostURL(s string) *NewsArticleSourceUpdateOne {
	nasuo.mutation.SetLatestPostURL(s)
	return nasuo
}

// SetNillableLatestPostURL sets the "latest_post_url" field if the given value is not nil.
func (nasuo *NewsArticleSourceUpdateOne) SetNillableLatestPostURL(s *string) *NewsArticleSourceUpdateOne {
	if s != nil {
		nasuo.SetLatestPostURL(*s)
	}
	return nasuo
}

// ClearLatestPostURL clears the value of the "latest_post_url" field.
func (nasuo *NewsArticleSourceUpdateOne) ClearLatestPostURL() *NewsArticleSourceUpdateOne {
	nasuo.mutation.ClearLatestPostURL()
	return nasuo
}

// SetLatestPostSelector sets the "latest_post_selector" field.
func (nasuo *NewsArticleSourceUpdateOne) SetLatestPostSelector(sps *schema.NewsPostSelector) *NewsArticleSourceUpdateOne {
	nasuo.mutation.SetLatestPostSelector(sps)
	return nasuo
}

// ClearLatestPostSelector clears the value of the "latest_post_selector" field.
func (nasuo *NewsArticleSourceUpdateOne) ClearLatestPostSelector() *NewsArticleSourceUpdateOne {
	nasuo.mutation.ClearLatestPostSelector()
	return nasuo
}

// SetCategoryPostURL sets the "category_post_url" field.
func (nasuo *NewsArticleSourceUpdateOne) SetCategoryPostURL(s string) *NewsArticleSourceUpdateOne {
	nasuo.mutation.SetCategoryPostURL(s)
	return nasuo
}

// SetNillableCategoryPostURL sets the "category_post_url" field if the given value is not nil.
func (nasuo *NewsArticleSourceUpdateOne) SetNillableCategoryPostURL(s *string) *NewsArticleSourceUpdateOne {
	if s != nil {
		nasuo.SetCategoryPostURL(*s)
	}
	return nasuo
}

// ClearCategoryPostURL clears the value of the "category_post_url" field.
func (nasuo *NewsArticleSourceUpdateOne) ClearCategoryPostURL() *NewsArticleSourceUpdateOne {
	nasuo.mutation.ClearCategoryPostURL()
	return nasuo
}

// SetCategoryPostSelector sets the "category_post_selector" field.
func (nasuo *NewsArticleSourceUpdateOne) SetCategoryPostSelector(sps *schema.NewsPostSelector) *NewsArticleSourceUpdateOne {
	nasuo.mutation.SetCategoryPostSelector(sps)
	return nasuo
}

// ClearCategoryPostSelector clears the value of the "category_post_selector" field.
func (nasuo *NewsArticleSourceUpdateOne) ClearCategoryPostSelector() *NewsArticleSourceUpdateOne {
	nasuo.mutation.ClearCategoryPostSelector()
	return nasuo
}

// SetArticleSelector sets the "article_selector" field.
func (nasuo *NewsArticleSourceUpdateOne) SetArticleSelector(sas *schema.NewsArticleSelector) *NewsArticleSourceUpdateOne {
	nasuo.mutation.SetArticleSelector(sas)
	return nasuo
}

// ClearArticleSelector clears the value of the "article_selector" field.
func (nasuo *NewsArticleSourceUpdateOne) ClearArticleSelector() *NewsArticleSourceUpdateOne {
	nasuo.mutation.ClearArticleSelector()
	return nasuo
}

// SetCategories sets the "categories" field.
func (nasuo *NewsArticleSourceUpdateOne) SetCategories(m map[string]string) *NewsArticleSourceUpdateOne {
	nasuo.mutation.SetCategories(m)
	return nasuo
}

// SetLanguage sets the "language" field.
func (nasuo *NewsArticleSourceUpdateOne) SetLanguage(s string) *NewsArticleSourceUpdateOne {
	nasuo.mutation.SetLanguage(s)
	return nasuo
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (nasuo *NewsArticleSourceUpdateOne) SetNillableLanguage(s *string) *NewsArticleSourceUpdateOne {
	if s != nil {
		nasuo.SetLanguage(*s)
	}
	return nasuo
}

// SetCountry sets the "country" field.
func (nasuo *NewsArticleSourceUpdateOne) SetCountry(s string) *NewsArticleSourceUpdateOne {
	nasuo.mutation.SetCountry(s)
	return nasuo
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (nasuo *NewsArticleSourceUpdateOne) SetNillableCountry(s *string) *NewsArticleSourceUpdateOne {
	if s != nil {
		nasuo.SetCountry(*s)
	}
	return nasuo
}

// SetStatus sets the "status" field.
func (nasuo *NewsArticleSourceUpdateOne) SetStatus(b bool) *NewsArticleSourceUpdateOne {
	nasuo.mutation.SetStatus(b)
	return nasuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (nasuo *NewsArticleSourceUpdateOne) SetNillableStatus(b *bool) *NewsArticleSourceUpdateOne {
	if b != nil {
		nasuo.SetStatus(*b)
	}
	return nasuo
}

// SetLogo sets the "logo" field.
func (nasuo *NewsArticleSourceUpdateOne) SetLogo(s string) *NewsArticleSourceUpdateOne {
	nasuo.mutation.SetLogo(s)
	return nasuo
}

// SetName sets the "name" field.
func (nasuo *NewsArticleSourceUpdateOne) SetName(s string) *NewsArticleSourceUpdateOne {
	nasuo.mutation.SetName(s)
	return nasuo
}

// SetURL sets the "url" field.
func (nasuo *NewsArticleSourceUpdateOne) SetURL(s string) *NewsArticleSourceUpdateOne {
	nasuo.mutation.SetURL(s)
	return nasuo
}

// Mutation returns the NewsArticleSourceMutation object of the builder.
func (nasuo *NewsArticleSourceUpdateOne) Mutation() *NewsArticleSourceMutation {
	return nasuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nasuo *NewsArticleSourceUpdateOne) Select(field string, fields ...string) *NewsArticleSourceUpdateOne {
	nasuo.fields = append([]string{field}, fields...)
	return nasuo
}

// Save executes the query and returns the updated NewsArticleSource entity.
func (nasuo *NewsArticleSourceUpdateOne) Save(ctx context.Context) (*NewsArticleSource, error) {
	var (
		err  error
		node *NewsArticleSource
	)
	if len(nasuo.hooks) == 0 {
		if err = nasuo.check(); err != nil {
			return nil, err
		}
		node, err = nasuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NewsArticleSourceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nasuo.check(); err != nil {
				return nil, err
			}
			nasuo.mutation = mutation
			node, err = nasuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nasuo.hooks) - 1; i >= 0; i-- {
			if nasuo.hooks[i] == nil {
				return nil, fmt.Errorf("store: uninitialized hook (forgotten import store/runtime?)")
			}
			mut = nasuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nasuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (nasuo *NewsArticleSourceUpdateOne) SaveX(ctx context.Context) *NewsArticleSource {
	node, err := nasuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nasuo *NewsArticleSourceUpdateOne) Exec(ctx context.Context) error {
	_, err := nasuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nasuo *NewsArticleSourceUpdateOne) ExecX(ctx context.Context) {
	if err := nasuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nasuo *NewsArticleSourceUpdateOne) check() error {
	if v, ok := nasuo.mutation.Logo(); ok {
		if err := newsarticlesource.LogoValidator(v); err != nil {
			return &ValidationError{Name: "logo", err: fmt.Errorf(`store: validator failed for field "NewsArticleSource.logo": %w`, err)}
		}
	}
	return nil
}

func (nasuo *NewsArticleSourceUpdateOne) sqlSave(ctx context.Context) (_node *NewsArticleSource, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   newsarticlesource.Table,
			Columns: newsarticlesource.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: newsarticlesource.FieldID,
			},
		},
	}
	id, ok := nasuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`store: missing "NewsArticleSource.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nasuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, newsarticlesource.FieldID)
		for _, f := range fields {
			if !newsarticlesource.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("store: invalid field %q for query", f)}
			}
			if f != newsarticlesource.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nasuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nasuo.mutation.LatestPostURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: newsarticlesource.FieldLatestPostURL,
		})
	}
	if nasuo.mutation.LatestPostURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: newsarticlesource.FieldLatestPostURL,
		})
	}
	if value, ok := nasuo.mutation.LatestPostSelector(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: newsarticlesource.FieldLatestPostSelector,
		})
	}
	if nasuo.mutation.LatestPostSelectorCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: newsarticlesource.FieldLatestPostSelector,
		})
	}
	if value, ok := nasuo.mutation.CategoryPostURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: newsarticlesource.FieldCategoryPostURL,
		})
	}
	if nasuo.mutation.CategoryPostURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: newsarticlesource.FieldCategoryPostURL,
		})
	}
	if value, ok := nasuo.mutation.CategoryPostSelector(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: newsarticlesource.FieldCategoryPostSelector,
		})
	}
	if nasuo.mutation.CategoryPostSelectorCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: newsarticlesource.FieldCategoryPostSelector,
		})
	}
	if value, ok := nasuo.mutation.ArticleSelector(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: newsarticlesource.FieldArticleSelector,
		})
	}
	if nasuo.mutation.ArticleSelectorCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: newsarticlesource.FieldArticleSelector,
		})
	}
	if value, ok := nasuo.mutation.Categories(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: newsarticlesource.FieldCategories,
		})
	}
	if value, ok := nasuo.mutation.Language(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: newsarticlesource.FieldLanguage,
		})
	}
	if value, ok := nasuo.mutation.Country(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: newsarticlesource.FieldCountry,
		})
	}
	if value, ok := nasuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: newsarticlesource.FieldStatus,
		})
	}
	if value, ok := nasuo.mutation.Logo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: newsarticlesource.FieldLogo,
		})
	}
	if value, ok := nasuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: newsarticlesource.FieldName,
		})
	}
	if value, ok := nasuo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: newsarticlesource.FieldURL,
		})
	}
	_node = &NewsArticleSource{config: nasuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nasuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{newsarticlesource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
