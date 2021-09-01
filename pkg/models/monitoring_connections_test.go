// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testMonitoringConnections(t *testing.T) {
	t.Parallel()

	query := MonitoringConnections()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMonitoringConnectionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringConnection{}
	if err = randomize.Struct(seed, o, monitoringConnectionDBTypes, true, monitoringConnectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringConnection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MonitoringConnections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMonitoringConnectionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringConnection{}
	if err = randomize.Struct(seed, o, monitoringConnectionDBTypes, true, monitoringConnectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringConnection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := MonitoringConnections().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MonitoringConnections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMonitoringConnectionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringConnection{}
	if err = randomize.Struct(seed, o, monitoringConnectionDBTypes, true, monitoringConnectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringConnection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MonitoringConnectionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MonitoringConnections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMonitoringConnectionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringConnection{}
	if err = randomize.Struct(seed, o, monitoringConnectionDBTypes, true, monitoringConnectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringConnection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MonitoringConnectionExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if MonitoringConnection exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MonitoringConnectionExists to return true, but got false.")
	}
}

func testMonitoringConnectionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringConnection{}
	if err = randomize.Struct(seed, o, monitoringConnectionDBTypes, true, monitoringConnectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringConnection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	monitoringConnectionFound, err := FindMonitoringConnection(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if monitoringConnectionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMonitoringConnectionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringConnection{}
	if err = randomize.Struct(seed, o, monitoringConnectionDBTypes, true, monitoringConnectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringConnection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = MonitoringConnections().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMonitoringConnectionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringConnection{}
	if err = randomize.Struct(seed, o, monitoringConnectionDBTypes, true, monitoringConnectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringConnection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := MonitoringConnections().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMonitoringConnectionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	monitoringConnectionOne := &MonitoringConnection{}
	monitoringConnectionTwo := &MonitoringConnection{}
	if err = randomize.Struct(seed, monitoringConnectionOne, monitoringConnectionDBTypes, false, monitoringConnectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringConnection struct: %s", err)
	}
	if err = randomize.Struct(seed, monitoringConnectionTwo, monitoringConnectionDBTypes, false, monitoringConnectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringConnection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = monitoringConnectionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = monitoringConnectionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MonitoringConnections().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMonitoringConnectionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	monitoringConnectionOne := &MonitoringConnection{}
	monitoringConnectionTwo := &MonitoringConnection{}
	if err = randomize.Struct(seed, monitoringConnectionOne, monitoringConnectionDBTypes, false, monitoringConnectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringConnection struct: %s", err)
	}
	if err = randomize.Struct(seed, monitoringConnectionTwo, monitoringConnectionDBTypes, false, monitoringConnectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringConnection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = monitoringConnectionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = monitoringConnectionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MonitoringConnections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func monitoringConnectionBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *MonitoringConnection) error {
	*o = MonitoringConnection{}
	return nil
}

func monitoringConnectionAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *MonitoringConnection) error {
	*o = MonitoringConnection{}
	return nil
}

func monitoringConnectionAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *MonitoringConnection) error {
	*o = MonitoringConnection{}
	return nil
}

func monitoringConnectionBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MonitoringConnection) error {
	*o = MonitoringConnection{}
	return nil
}

func monitoringConnectionAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MonitoringConnection) error {
	*o = MonitoringConnection{}
	return nil
}

func monitoringConnectionBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MonitoringConnection) error {
	*o = MonitoringConnection{}
	return nil
}

func monitoringConnectionAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MonitoringConnection) error {
	*o = MonitoringConnection{}
	return nil
}

func monitoringConnectionBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MonitoringConnection) error {
	*o = MonitoringConnection{}
	return nil
}

func monitoringConnectionAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MonitoringConnection) error {
	*o = MonitoringConnection{}
	return nil
}

func testMonitoringConnectionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &MonitoringConnection{}
	o := &MonitoringConnection{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, monitoringConnectionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize MonitoringConnection object: %s", err)
	}

	AddMonitoringConnectionHook(boil.BeforeInsertHook, monitoringConnectionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	monitoringConnectionBeforeInsertHooks = []MonitoringConnectionHook{}

	AddMonitoringConnectionHook(boil.AfterInsertHook, monitoringConnectionAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	monitoringConnectionAfterInsertHooks = []MonitoringConnectionHook{}

	AddMonitoringConnectionHook(boil.AfterSelectHook, monitoringConnectionAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	monitoringConnectionAfterSelectHooks = []MonitoringConnectionHook{}

	AddMonitoringConnectionHook(boil.BeforeUpdateHook, monitoringConnectionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	monitoringConnectionBeforeUpdateHooks = []MonitoringConnectionHook{}

	AddMonitoringConnectionHook(boil.AfterUpdateHook, monitoringConnectionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	monitoringConnectionAfterUpdateHooks = []MonitoringConnectionHook{}

	AddMonitoringConnectionHook(boil.BeforeDeleteHook, monitoringConnectionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	monitoringConnectionBeforeDeleteHooks = []MonitoringConnectionHook{}

	AddMonitoringConnectionHook(boil.AfterDeleteHook, monitoringConnectionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	monitoringConnectionAfterDeleteHooks = []MonitoringConnectionHook{}

	AddMonitoringConnectionHook(boil.BeforeUpsertHook, monitoringConnectionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	monitoringConnectionBeforeUpsertHooks = []MonitoringConnectionHook{}

	AddMonitoringConnectionHook(boil.AfterUpsertHook, monitoringConnectionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	monitoringConnectionAfterUpsertHooks = []MonitoringConnectionHook{}
}

func testMonitoringConnectionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringConnection{}
	if err = randomize.Struct(seed, o, monitoringConnectionDBTypes, true, monitoringConnectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringConnection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MonitoringConnections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMonitoringConnectionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringConnection{}
	if err = randomize.Struct(seed, o, monitoringConnectionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MonitoringConnection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(monitoringConnectionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := MonitoringConnections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMonitoringConnectionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringConnection{}
	if err = randomize.Struct(seed, o, monitoringConnectionDBTypes, true, monitoringConnectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringConnection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMonitoringConnectionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringConnection{}
	if err = randomize.Struct(seed, o, monitoringConnectionDBTypes, true, monitoringConnectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringConnection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MonitoringConnectionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMonitoringConnectionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringConnection{}
	if err = randomize.Struct(seed, o, monitoringConnectionDBTypes, true, monitoringConnectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringConnection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MonitoringConnections().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	monitoringConnectionDBTypes = map[string]string{`ID`: `integer`, `PeerID`: `character varying`, `IPAddress`: `character varying`, `DialAttempt`: `timestamp with time zone`, `Latency`: `interval`, `IsSucceed`: `boolean`}
	_                           = bytes.MinRead
)

func testMonitoringConnectionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(monitoringConnectionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(monitoringConnectionAllColumns) == len(monitoringConnectionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringConnection{}
	if err = randomize.Struct(seed, o, monitoringConnectionDBTypes, true, monitoringConnectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringConnection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MonitoringConnections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, monitoringConnectionDBTypes, true, monitoringConnectionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MonitoringConnection struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMonitoringConnectionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(monitoringConnectionAllColumns) == len(monitoringConnectionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringConnection{}
	if err = randomize.Struct(seed, o, monitoringConnectionDBTypes, true, monitoringConnectionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringConnection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MonitoringConnections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, monitoringConnectionDBTypes, true, monitoringConnectionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MonitoringConnection struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(monitoringConnectionAllColumns, monitoringConnectionPrimaryKeyColumns) {
		fields = monitoringConnectionAllColumns
	} else {
		fields = strmangle.SetComplement(
			monitoringConnectionAllColumns,
			monitoringConnectionPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := MonitoringConnectionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMonitoringConnectionsUpsert(t *testing.T) {
	t.Parallel()

	if len(monitoringConnectionAllColumns) == len(monitoringConnectionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := MonitoringConnection{}
	if err = randomize.Struct(seed, &o, monitoringConnectionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MonitoringConnection struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MonitoringConnection: %s", err)
	}

	count, err := MonitoringConnections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, monitoringConnectionDBTypes, false, monitoringConnectionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MonitoringConnection struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MonitoringConnection: %s", err)
	}

	count, err = MonitoringConnections().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}