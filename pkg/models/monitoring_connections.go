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

// MonitoringConnection is an object representing the database table.
type MonitoringConnection struct {
	ID          int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	PeerID      string      `boil:"peer_id" json:"peer_id" toml:"peer_id" yaml:"peer_id"`
	IPAddress   null.String `boil:"ip_address" json:"ip_address,omitempty" toml:"ip_address" yaml:"ip_address,omitempty"`
	DialAttempt null.Time   `boil:"dial_attempt" json:"dial_attempt,omitempty" toml:"dial_attempt" yaml:"dial_attempt,omitempty"`
	Latency     null.String `boil:"latency" json:"latency,omitempty" toml:"latency" yaml:"latency,omitempty"`
	IsSucceed   null.Bool   `boil:"is_succeed" json:"is_succeed,omitempty" toml:"is_succeed" yaml:"is_succeed,omitempty"`

	R *monitoringConnectionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L monitoringConnectionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MonitoringConnectionColumns = struct {
	ID          string
	PeerID      string
	IPAddress   string
	DialAttempt string
	Latency     string
	IsSucceed   string
}{
	ID:          "id",
	PeerID:      "peer_id",
	IPAddress:   "ip_address",
	DialAttempt: "dial_attempt",
	Latency:     "latency",
	IsSucceed:   "is_succeed",
}

var MonitoringConnectionTableColumns = struct {
	ID          string
	PeerID      string
	IPAddress   string
	DialAttempt string
	Latency     string
	IsSucceed   string
}{
	ID:          "monitoring_connections.id",
	PeerID:      "monitoring_connections.peer_id",
	IPAddress:   "monitoring_connections.ip_address",
	DialAttempt: "monitoring_connections.dial_attempt",
	Latency:     "monitoring_connections.latency",
	IsSucceed:   "monitoring_connections.is_succeed",
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

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
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

type whereHelpernull_Bool struct{ field string }

func (w whereHelpernull_Bool) EQ(x null.Bool) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Bool) NEQ(x null.Bool) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Bool) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Bool) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Bool) LT(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Bool) LTE(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Bool) GT(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Bool) GTE(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var MonitoringConnectionWhere = struct {
	ID          whereHelperint
	PeerID      whereHelperstring
	IPAddress   whereHelpernull_String
	DialAttempt whereHelpernull_Time
	Latency     whereHelpernull_String
	IsSucceed   whereHelpernull_Bool
}{
	ID:          whereHelperint{field: "\"monitoring_connections\".\"id\""},
	PeerID:      whereHelperstring{field: "\"monitoring_connections\".\"peer_id\""},
	IPAddress:   whereHelpernull_String{field: "\"monitoring_connections\".\"ip_address\""},
	DialAttempt: whereHelpernull_Time{field: "\"monitoring_connections\".\"dial_attempt\""},
	Latency:     whereHelpernull_String{field: "\"monitoring_connections\".\"latency\""},
	IsSucceed:   whereHelpernull_Bool{field: "\"monitoring_connections\".\"is_succeed\""},
}

// MonitoringConnectionRels is where relationship names are stored.
var MonitoringConnectionRels = struct {
}{}

// monitoringConnectionR is where relationships are stored.
type monitoringConnectionR struct {
}

// NewStruct creates a new relationship struct
func (*monitoringConnectionR) NewStruct() *monitoringConnectionR {
	return &monitoringConnectionR{}
}

// monitoringConnectionL is where Load methods for each relationship are stored.
type monitoringConnectionL struct{}

var (
	monitoringConnectionAllColumns            = []string{"id", "peer_id", "ip_address", "dial_attempt", "latency", "is_succeed"}
	monitoringConnectionColumnsWithoutDefault = []string{"peer_id", "ip_address", "dial_attempt", "latency", "is_succeed"}
	monitoringConnectionColumnsWithDefault    = []string{"id"}
	monitoringConnectionPrimaryKeyColumns     = []string{"id"}
)

type (
	// MonitoringConnectionSlice is an alias for a slice of pointers to MonitoringConnection.
	// This should almost always be used instead of []MonitoringConnection.
	MonitoringConnectionSlice []*MonitoringConnection
	// MonitoringConnectionHook is the signature for custom MonitoringConnection hook methods
	MonitoringConnectionHook func(context.Context, boil.ContextExecutor, *MonitoringConnection) error

	monitoringConnectionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	monitoringConnectionType                 = reflect.TypeOf(&MonitoringConnection{})
	monitoringConnectionMapping              = queries.MakeStructMapping(monitoringConnectionType)
	monitoringConnectionPrimaryKeyMapping, _ = queries.BindMapping(monitoringConnectionType, monitoringConnectionMapping, monitoringConnectionPrimaryKeyColumns)
	monitoringConnectionInsertCacheMut       sync.RWMutex
	monitoringConnectionInsertCache          = make(map[string]insertCache)
	monitoringConnectionUpdateCacheMut       sync.RWMutex
	monitoringConnectionUpdateCache          = make(map[string]updateCache)
	monitoringConnectionUpsertCacheMut       sync.RWMutex
	monitoringConnectionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var monitoringConnectionBeforeInsertHooks []MonitoringConnectionHook
var monitoringConnectionBeforeUpdateHooks []MonitoringConnectionHook
var monitoringConnectionBeforeDeleteHooks []MonitoringConnectionHook
var monitoringConnectionBeforeUpsertHooks []MonitoringConnectionHook

var monitoringConnectionAfterInsertHooks []MonitoringConnectionHook
var monitoringConnectionAfterSelectHooks []MonitoringConnectionHook
var monitoringConnectionAfterUpdateHooks []MonitoringConnectionHook
var monitoringConnectionAfterDeleteHooks []MonitoringConnectionHook
var monitoringConnectionAfterUpsertHooks []MonitoringConnectionHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *MonitoringConnection) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range monitoringConnectionBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *MonitoringConnection) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range monitoringConnectionBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *MonitoringConnection) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range monitoringConnectionBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *MonitoringConnection) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range monitoringConnectionBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *MonitoringConnection) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range monitoringConnectionAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *MonitoringConnection) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range monitoringConnectionAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *MonitoringConnection) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range monitoringConnectionAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *MonitoringConnection) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range monitoringConnectionAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *MonitoringConnection) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range monitoringConnectionAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMonitoringConnectionHook registers your hook function for all future operations.
func AddMonitoringConnectionHook(hookPoint boil.HookPoint, monitoringConnectionHook MonitoringConnectionHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		monitoringConnectionBeforeInsertHooks = append(monitoringConnectionBeforeInsertHooks, monitoringConnectionHook)
	case boil.BeforeUpdateHook:
		monitoringConnectionBeforeUpdateHooks = append(monitoringConnectionBeforeUpdateHooks, monitoringConnectionHook)
	case boil.BeforeDeleteHook:
		monitoringConnectionBeforeDeleteHooks = append(monitoringConnectionBeforeDeleteHooks, monitoringConnectionHook)
	case boil.BeforeUpsertHook:
		monitoringConnectionBeforeUpsertHooks = append(monitoringConnectionBeforeUpsertHooks, monitoringConnectionHook)
	case boil.AfterInsertHook:
		monitoringConnectionAfterInsertHooks = append(monitoringConnectionAfterInsertHooks, monitoringConnectionHook)
	case boil.AfterSelectHook:
		monitoringConnectionAfterSelectHooks = append(monitoringConnectionAfterSelectHooks, monitoringConnectionHook)
	case boil.AfterUpdateHook:
		monitoringConnectionAfterUpdateHooks = append(monitoringConnectionAfterUpdateHooks, monitoringConnectionHook)
	case boil.AfterDeleteHook:
		monitoringConnectionAfterDeleteHooks = append(monitoringConnectionAfterDeleteHooks, monitoringConnectionHook)
	case boil.AfterUpsertHook:
		monitoringConnectionAfterUpsertHooks = append(monitoringConnectionAfterUpsertHooks, monitoringConnectionHook)
	}
}

// One returns a single monitoringConnection record from the query.
func (q monitoringConnectionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*MonitoringConnection, error) {
	o := &MonitoringConnection{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for monitoring_connections")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all MonitoringConnection records from the query.
func (q monitoringConnectionQuery) All(ctx context.Context, exec boil.ContextExecutor) (MonitoringConnectionSlice, error) {
	var o []*MonitoringConnection

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to MonitoringConnection slice")
	}

	if len(monitoringConnectionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all MonitoringConnection records in the query.
func (q monitoringConnectionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count monitoring_connections rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q monitoringConnectionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if monitoring_connections exists")
	}

	return count > 0, nil
}

// MonitoringConnections retrieves all the records using an executor.
func MonitoringConnections(mods ...qm.QueryMod) monitoringConnectionQuery {
	mods = append(mods, qm.From("\"monitoring_connections\""))
	return monitoringConnectionQuery{NewQuery(mods...)}
}

// FindMonitoringConnection retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMonitoringConnection(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*MonitoringConnection, error) {
	monitoringConnectionObj := &MonitoringConnection{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"monitoring_connections\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, monitoringConnectionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from monitoring_connections")
	}

	if err = monitoringConnectionObj.doAfterSelectHooks(ctx, exec); err != nil {
		return monitoringConnectionObj, err
	}

	return monitoringConnectionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MonitoringConnection) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no monitoring_connections provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(monitoringConnectionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	monitoringConnectionInsertCacheMut.RLock()
	cache, cached := monitoringConnectionInsertCache[key]
	monitoringConnectionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			monitoringConnectionAllColumns,
			monitoringConnectionColumnsWithDefault,
			monitoringConnectionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(monitoringConnectionType, monitoringConnectionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(monitoringConnectionType, monitoringConnectionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"monitoring_connections\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"monitoring_connections\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into monitoring_connections")
	}

	if !cached {
		monitoringConnectionInsertCacheMut.Lock()
		monitoringConnectionInsertCache[key] = cache
		monitoringConnectionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the MonitoringConnection.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MonitoringConnection) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	monitoringConnectionUpdateCacheMut.RLock()
	cache, cached := monitoringConnectionUpdateCache[key]
	monitoringConnectionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			monitoringConnectionAllColumns,
			monitoringConnectionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update monitoring_connections, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"monitoring_connections\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, monitoringConnectionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(monitoringConnectionType, monitoringConnectionMapping, append(wl, monitoringConnectionPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update monitoring_connections row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for monitoring_connections")
	}

	if !cached {
		monitoringConnectionUpdateCacheMut.Lock()
		monitoringConnectionUpdateCache[key] = cache
		monitoringConnectionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q monitoringConnectionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for monitoring_connections")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for monitoring_connections")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MonitoringConnectionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), monitoringConnectionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"monitoring_connections\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, monitoringConnectionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in monitoringConnection slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all monitoringConnection")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MonitoringConnection) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no monitoring_connections provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(monitoringConnectionColumnsWithDefault, o)

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

	monitoringConnectionUpsertCacheMut.RLock()
	cache, cached := monitoringConnectionUpsertCache[key]
	monitoringConnectionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			monitoringConnectionAllColumns,
			monitoringConnectionColumnsWithDefault,
			monitoringConnectionColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			monitoringConnectionAllColumns,
			monitoringConnectionPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert monitoring_connections, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(monitoringConnectionPrimaryKeyColumns))
			copy(conflict, monitoringConnectionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"monitoring_connections\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(monitoringConnectionType, monitoringConnectionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(monitoringConnectionType, monitoringConnectionMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert monitoring_connections")
	}

	if !cached {
		monitoringConnectionUpsertCacheMut.Lock()
		monitoringConnectionUpsertCache[key] = cache
		monitoringConnectionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single MonitoringConnection record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MonitoringConnection) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no MonitoringConnection provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), monitoringConnectionPrimaryKeyMapping)
	sql := "DELETE FROM \"monitoring_connections\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from monitoring_connections")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for monitoring_connections")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q monitoringConnectionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no monitoringConnectionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from monitoring_connections")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for monitoring_connections")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MonitoringConnectionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(monitoringConnectionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), monitoringConnectionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"monitoring_connections\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, monitoringConnectionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from monitoringConnection slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for monitoring_connections")
	}

	if len(monitoringConnectionAfterDeleteHooks) != 0 {
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
func (o *MonitoringConnection) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMonitoringConnection(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MonitoringConnectionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MonitoringConnectionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), monitoringConnectionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"monitoring_connections\".* FROM \"monitoring_connections\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, monitoringConnectionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MonitoringConnectionSlice")
	}

	*o = slice

	return nil
}

// MonitoringConnectionExists checks if the MonitoringConnection row exists.
func MonitoringConnectionExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"monitoring_connections\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if monitoring_connections exists")
	}

	return exists, nil
}
