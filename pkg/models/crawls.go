// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Crawl is an object representing the database table.
type Crawl struct {
	ID              int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	State           string    `boil:"state" json:"state" toml:"state" yaml:"state"`
	StartedAt       time.Time `boil:"started_at" json:"started_at" toml:"started_at" yaml:"started_at"`
	FinishedAt      null.Time `boil:"finished_at" json:"finished_at,omitempty" toml:"finished_at" yaml:"finished_at,omitempty"`
	UpdatedAt       time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	CreatedAt       time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	CrawledPeers    null.Int  `boil:"crawled_peers" json:"crawled_peers,omitempty" toml:"crawled_peers" yaml:"crawled_peers,omitempty"`
	DialablePeers   null.Int  `boil:"dialable_peers" json:"dialable_peers,omitempty" toml:"dialable_peers" yaml:"dialable_peers,omitempty"`
	UndialablePeers null.Int  `boil:"undialable_peers" json:"undialable_peers,omitempty" toml:"undialable_peers" yaml:"undialable_peers,omitempty"`

	R *crawlR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L crawlL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CrawlColumns = struct {
	ID              string
	State           string
	StartedAt       string
	FinishedAt      string
	UpdatedAt       string
	CreatedAt       string
	CrawledPeers    string
	DialablePeers   string
	UndialablePeers string
}{
	ID:              "id",
	State:           "state",
	StartedAt:       "started_at",
	FinishedAt:      "finished_at",
	UpdatedAt:       "updated_at",
	CreatedAt:       "created_at",
	CrawledPeers:    "crawled_peers",
	DialablePeers:   "dialable_peers",
	UndialablePeers: "undialable_peers",
}

var CrawlTableColumns = struct {
	ID              string
	State           string
	StartedAt       string
	FinishedAt      string
	UpdatedAt       string
	CreatedAt       string
	CrawledPeers    string
	DialablePeers   string
	UndialablePeers string
}{
	ID:              "crawls.id",
	State:           "crawls.state",
	StartedAt:       "crawls.started_at",
	FinishedAt:      "crawls.finished_at",
	UpdatedAt:       "crawls.updated_at",
	CreatedAt:       "crawls.created_at",
	CrawledPeers:    "crawls.crawled_peers",
	DialablePeers:   "crawls.dialable_peers",
	UndialablePeers: "crawls.undialable_peers",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Int struct{ field string }

func (w whereHelpernull_Int) EQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int) NEQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Int) LT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int) LTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int) GT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int) GTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var CrawlWhere = struct {
	ID              whereHelperint
	State           whereHelperstring
	StartedAt       whereHelpertime_Time
	FinishedAt      whereHelpernull_Time
	UpdatedAt       whereHelpertime_Time
	CreatedAt       whereHelpertime_Time
	CrawledPeers    whereHelpernull_Int
	DialablePeers   whereHelpernull_Int
	UndialablePeers whereHelpernull_Int
}{
	ID:              whereHelperint{field: "\"crawls\".\"id\""},
	State:           whereHelperstring{field: "\"crawls\".\"state\""},
	StartedAt:       whereHelpertime_Time{field: "\"crawls\".\"started_at\""},
	FinishedAt:      whereHelpernull_Time{field: "\"crawls\".\"finished_at\""},
	UpdatedAt:       whereHelpertime_Time{field: "\"crawls\".\"updated_at\""},
	CreatedAt:       whereHelpertime_Time{field: "\"crawls\".\"created_at\""},
	CrawledPeers:    whereHelpernull_Int{field: "\"crawls\".\"crawled_peers\""},
	DialablePeers:   whereHelpernull_Int{field: "\"crawls\".\"dialable_peers\""},
	UndialablePeers: whereHelpernull_Int{field: "\"crawls\".\"undialable_peers\""},
}

// CrawlRels is where relationship names are stored.
var CrawlRels = struct {
	CrawlProperties string
}{
	CrawlProperties: "CrawlProperties",
}

// crawlR is where relationships are stored.
type crawlR struct {
	CrawlProperties CrawlPropertySlice `boil:"CrawlProperties" json:"CrawlProperties" toml:"CrawlProperties" yaml:"CrawlProperties"`
}

// NewStruct creates a new relationship struct
func (*crawlR) NewStruct() *crawlR {
	return &crawlR{}
}

// crawlL is where Load methods for each relationship are stored.
type crawlL struct{}

var (
	crawlAllColumns            = []string{"id", "state", "started_at", "finished_at", "updated_at", "created_at", "crawled_peers", "dialable_peers", "undialable_peers"}
	crawlColumnsWithoutDefault = []string{"state", "started_at", "finished_at", "updated_at", "created_at", "crawled_peers", "dialable_peers", "undialable_peers"}
	crawlColumnsWithDefault    = []string{"id"}
	crawlPrimaryKeyColumns     = []string{"id"}
)

type (
	// CrawlSlice is an alias for a slice of pointers to Crawl.
	// This should almost always be used instead of []Crawl.
	CrawlSlice []*Crawl
	// CrawlHook is the signature for custom Crawl hook methods
	CrawlHook func(context.Context, boil.ContextExecutor, *Crawl) error

	crawlQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	crawlType                 = reflect.TypeOf(&Crawl{})
	crawlMapping              = queries.MakeStructMapping(crawlType)
	crawlPrimaryKeyMapping, _ = queries.BindMapping(crawlType, crawlMapping, crawlPrimaryKeyColumns)
	crawlInsertCacheMut       sync.RWMutex
	crawlInsertCache          = make(map[string]insertCache)
	crawlUpdateCacheMut       sync.RWMutex
	crawlUpdateCache          = make(map[string]updateCache)
	crawlUpsertCacheMut       sync.RWMutex
	crawlUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var crawlBeforeInsertHooks []CrawlHook
var crawlBeforeUpdateHooks []CrawlHook
var crawlBeforeDeleteHooks []CrawlHook
var crawlBeforeUpsertHooks []CrawlHook

var crawlAfterInsertHooks []CrawlHook
var crawlAfterSelectHooks []CrawlHook
var crawlAfterUpdateHooks []CrawlHook
var crawlAfterDeleteHooks []CrawlHook
var crawlAfterUpsertHooks []CrawlHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Crawl) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range crawlBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Crawl) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range crawlBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Crawl) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range crawlBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Crawl) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range crawlBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Crawl) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range crawlAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Crawl) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range crawlAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Crawl) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range crawlAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Crawl) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range crawlAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Crawl) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range crawlAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCrawlHook registers your hook function for all future operations.
func AddCrawlHook(hookPoint boil.HookPoint, crawlHook CrawlHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		crawlBeforeInsertHooks = append(crawlBeforeInsertHooks, crawlHook)
	case boil.BeforeUpdateHook:
		crawlBeforeUpdateHooks = append(crawlBeforeUpdateHooks, crawlHook)
	case boil.BeforeDeleteHook:
		crawlBeforeDeleteHooks = append(crawlBeforeDeleteHooks, crawlHook)
	case boil.BeforeUpsertHook:
		crawlBeforeUpsertHooks = append(crawlBeforeUpsertHooks, crawlHook)
	case boil.AfterInsertHook:
		crawlAfterInsertHooks = append(crawlAfterInsertHooks, crawlHook)
	case boil.AfterSelectHook:
		crawlAfterSelectHooks = append(crawlAfterSelectHooks, crawlHook)
	case boil.AfterUpdateHook:
		crawlAfterUpdateHooks = append(crawlAfterUpdateHooks, crawlHook)
	case boil.AfterDeleteHook:
		crawlAfterDeleteHooks = append(crawlAfterDeleteHooks, crawlHook)
	case boil.AfterUpsertHook:
		crawlAfterUpsertHooks = append(crawlAfterUpsertHooks, crawlHook)
	}
}

// One returns a single crawl record from the query.
func (q crawlQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Crawl, error) {
	o := &Crawl{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for crawls")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Crawl records from the query.
func (q crawlQuery) All(ctx context.Context, exec boil.ContextExecutor) (CrawlSlice, error) {
	var o []*Crawl

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Crawl slice")
	}

	if len(crawlAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Crawl records in the query.
func (q crawlQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count crawls rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q crawlQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if crawls exists")
	}

	return count > 0, nil
}

// CrawlProperties retrieves all the crawl_property's CrawlProperties with an executor.
func (o *Crawl) CrawlProperties(mods ...qm.QueryMod) crawlPropertyQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"crawl_properties\".\"crawl_id\"=?", o.ID),
	)

	query := CrawlProperties(queryMods...)
	queries.SetFrom(query.Query, "\"crawl_properties\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"crawl_properties\".*"})
	}

	return query
}

// LoadCrawlProperties allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (crawlL) LoadCrawlProperties(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCrawl interface{}, mods queries.Applicator) error {
	var slice []*Crawl
	var object *Crawl

	if singular {
		object = maybeCrawl.(*Crawl)
	} else {
		slice = *maybeCrawl.(*[]*Crawl)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &crawlR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &crawlR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`crawl_properties`),
		qm.WhereIn(`crawl_properties.crawl_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load crawl_properties")
	}

	var resultSlice []*CrawlProperty
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice crawl_properties")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on crawl_properties")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for crawl_properties")
	}

	if len(crawlPropertyAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.CrawlProperties = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &crawlPropertyR{}
			}
			foreign.R.Crawl = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.CrawlID {
				local.R.CrawlProperties = append(local.R.CrawlProperties, foreign)
				if foreign.R == nil {
					foreign.R = &crawlPropertyR{}
				}
				foreign.R.Crawl = local
				break
			}
		}
	}

	return nil
}

// AddCrawlProperties adds the given related objects to the existing relationships
// of the crawl, optionally inserting them as new records.
// Appends related to o.R.CrawlProperties.
// Sets related.R.Crawl appropriately.
func (o *Crawl) AddCrawlProperties(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*CrawlProperty) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.CrawlID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"crawl_properties\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"crawl_id"}),
				strmangle.WhereClause("\"", "\"", 2, crawlPropertyPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.CrawlID = o.ID
		}
	}

	if o.R == nil {
		o.R = &crawlR{
			CrawlProperties: related,
		}
	} else {
		o.R.CrawlProperties = append(o.R.CrawlProperties, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &crawlPropertyR{
				Crawl: o,
			}
		} else {
			rel.R.Crawl = o
		}
	}
	return nil
}

// Crawls retrieves all the records using an executor.
func Crawls(mods ...qm.QueryMod) crawlQuery {
	mods = append(mods, qm.From("\"crawls\""))
	return crawlQuery{NewQuery(mods...)}
}

// FindCrawl retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCrawl(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Crawl, error) {
	crawlObj := &Crawl{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"crawls\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, crawlObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from crawls")
	}

	if err = crawlObj.doAfterSelectHooks(ctx, exec); err != nil {
		return crawlObj, err
	}

	return crawlObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Crawl) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no crawls provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(crawlColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	crawlInsertCacheMut.RLock()
	cache, cached := crawlInsertCache[key]
	crawlInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			crawlAllColumns,
			crawlColumnsWithDefault,
			crawlColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(crawlType, crawlMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(crawlType, crawlMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"crawls\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"crawls\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into crawls")
	}

	if !cached {
		crawlInsertCacheMut.Lock()
		crawlInsertCache[key] = cache
		crawlInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Crawl.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Crawl) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	crawlUpdateCacheMut.RLock()
	cache, cached := crawlUpdateCache[key]
	crawlUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			crawlAllColumns,
			crawlPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update crawls, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"crawls\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, crawlPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(crawlType, crawlMapping, append(wl, crawlPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update crawls row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for crawls")
	}

	if !cached {
		crawlUpdateCacheMut.Lock()
		crawlUpdateCache[key] = cache
		crawlUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q crawlQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for crawls")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for crawls")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CrawlSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), crawlPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"crawls\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, crawlPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in crawl slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all crawl")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Crawl) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no crawls provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(crawlColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	crawlUpsertCacheMut.RLock()
	cache, cached := crawlUpsertCache[key]
	crawlUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			crawlAllColumns,
			crawlColumnsWithDefault,
			crawlColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			crawlAllColumns,
			crawlPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert crawls, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(crawlPrimaryKeyColumns))
			copy(conflict, crawlPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"crawls\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(crawlType, crawlMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(crawlType, crawlMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert crawls")
	}

	if !cached {
		crawlUpsertCacheMut.Lock()
		crawlUpsertCache[key] = cache
		crawlUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Crawl record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Crawl) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Crawl provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), crawlPrimaryKeyMapping)
	sql := "DELETE FROM \"crawls\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from crawls")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for crawls")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q crawlQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no crawlQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from crawls")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for crawls")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CrawlSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(crawlBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), crawlPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"crawls\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, crawlPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from crawl slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for crawls")
	}

	if len(crawlAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Crawl) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCrawl(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CrawlSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CrawlSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), crawlPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"crawls\".* FROM \"crawls\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, crawlPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CrawlSlice")
	}

	*o = slice

	return nil
}

// CrawlExists checks if the Crawl row exists.
func CrawlExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"crawls\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if crawls exists")
	}

	return exists, nil
}
