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

func testRawVisits(t *testing.T) {
	t.Parallel()

	query := RawVisits()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRawVisitsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RawVisit{}
	if err = randomize.Struct(seed, o, rawVisitDBTypes, true, rawVisitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawVisit struct: %s", err)
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

	count, err := RawVisits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRawVisitsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RawVisit{}
	if err = randomize.Struct(seed, o, rawVisitDBTypes, true, rawVisitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawVisit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RawVisits().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RawVisits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRawVisitsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RawVisit{}
	if err = randomize.Struct(seed, o, rawVisitDBTypes, true, rawVisitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawVisit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RawVisitSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RawVisits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRawVisitsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RawVisit{}
	if err = randomize.Struct(seed, o, rawVisitDBTypes, true, rawVisitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawVisit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RawVisitExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if RawVisit exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RawVisitExists to return true, but got false.")
	}
}

func testRawVisitsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RawVisit{}
	if err = randomize.Struct(seed, o, rawVisitDBTypes, true, rawVisitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawVisit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	rawVisitFound, err := FindRawVisit(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if rawVisitFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRawVisitsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RawVisit{}
	if err = randomize.Struct(seed, o, rawVisitDBTypes, true, rawVisitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawVisit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RawVisits().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRawVisitsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RawVisit{}
	if err = randomize.Struct(seed, o, rawVisitDBTypes, true, rawVisitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawVisit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RawVisits().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRawVisitsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	rawVisitOne := &RawVisit{}
	rawVisitTwo := &RawVisit{}
	if err = randomize.Struct(seed, rawVisitOne, rawVisitDBTypes, false, rawVisitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawVisit struct: %s", err)
	}
	if err = randomize.Struct(seed, rawVisitTwo, rawVisitDBTypes, false, rawVisitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawVisit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = rawVisitOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = rawVisitTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RawVisits().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRawVisitsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	rawVisitOne := &RawVisit{}
	rawVisitTwo := &RawVisit{}
	if err = randomize.Struct(seed, rawVisitOne, rawVisitDBTypes, false, rawVisitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawVisit struct: %s", err)
	}
	if err = randomize.Struct(seed, rawVisitTwo, rawVisitDBTypes, false, rawVisitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawVisit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = rawVisitOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = rawVisitTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RawVisits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func rawVisitBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *RawVisit) error {
	*o = RawVisit{}
	return nil
}

func rawVisitAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *RawVisit) error {
	*o = RawVisit{}
	return nil
}

func rawVisitAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *RawVisit) error {
	*o = RawVisit{}
	return nil
}

func rawVisitBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RawVisit) error {
	*o = RawVisit{}
	return nil
}

func rawVisitAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RawVisit) error {
	*o = RawVisit{}
	return nil
}

func rawVisitBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RawVisit) error {
	*o = RawVisit{}
	return nil
}

func rawVisitAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RawVisit) error {
	*o = RawVisit{}
	return nil
}

func rawVisitBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RawVisit) error {
	*o = RawVisit{}
	return nil
}

func rawVisitAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RawVisit) error {
	*o = RawVisit{}
	return nil
}

func testRawVisitsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &RawVisit{}
	o := &RawVisit{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, rawVisitDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RawVisit object: %s", err)
	}

	AddRawVisitHook(boil.BeforeInsertHook, rawVisitBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	rawVisitBeforeInsertHooks = []RawVisitHook{}

	AddRawVisitHook(boil.AfterInsertHook, rawVisitAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	rawVisitAfterInsertHooks = []RawVisitHook{}

	AddRawVisitHook(boil.AfterSelectHook, rawVisitAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	rawVisitAfterSelectHooks = []RawVisitHook{}

	AddRawVisitHook(boil.BeforeUpdateHook, rawVisitBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	rawVisitBeforeUpdateHooks = []RawVisitHook{}

	AddRawVisitHook(boil.AfterUpdateHook, rawVisitAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	rawVisitAfterUpdateHooks = []RawVisitHook{}

	AddRawVisitHook(boil.BeforeDeleteHook, rawVisitBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	rawVisitBeforeDeleteHooks = []RawVisitHook{}

	AddRawVisitHook(boil.AfterDeleteHook, rawVisitAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	rawVisitAfterDeleteHooks = []RawVisitHook{}

	AddRawVisitHook(boil.BeforeUpsertHook, rawVisitBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	rawVisitBeforeUpsertHooks = []RawVisitHook{}

	AddRawVisitHook(boil.AfterUpsertHook, rawVisitAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	rawVisitAfterUpsertHooks = []RawVisitHook{}
}

func testRawVisitsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RawVisit{}
	if err = randomize.Struct(seed, o, rawVisitDBTypes, true, rawVisitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawVisit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RawVisits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRawVisitsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RawVisit{}
	if err = randomize.Struct(seed, o, rawVisitDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RawVisit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(rawVisitColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RawVisits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRawVisitsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RawVisit{}
	if err = randomize.Struct(seed, o, rawVisitDBTypes, true, rawVisitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawVisit struct: %s", err)
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

func testRawVisitsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RawVisit{}
	if err = randomize.Struct(seed, o, rawVisitDBTypes, true, rawVisitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawVisit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RawVisitSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRawVisitsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RawVisit{}
	if err = randomize.Struct(seed, o, rawVisitDBTypes, true, rawVisitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawVisit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RawVisits().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	rawVisitDBTypes = map[string]string{`ID`: `integer`, `CrawlID`: `integer`, `CrawlStartedAt`: `timestamp with time zone`, `CrawlEndedAt`: `timestamp with time zone`, `ConnectDuration`: `interval`, `CrawlDuration`: `interval`, `AgentVersion`: `character varying`, `PeerMultiHash`: `character varying`, `Protocols`: `ARRAYcharacter varying`, `MultiAddresses`: `ARRAYcharacter varying`, `DialError`: `enum.dial_error('io_timeout','connection_refused','protocol_not_supported','peer_id_mismatch','no_route_to_host','network_unreachable','no_good_addresses','context_deadline_exceeded','no_public_ip','max_dial_attempts_exceeded','unknown','maddr_reset')`, `Error`: `text`, `CreatedAt`: `timestamp with time zone`}
	_               = bytes.MinRead
)

func testRawVisitsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(rawVisitPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(rawVisitAllColumns) == len(rawVisitPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RawVisit{}
	if err = randomize.Struct(seed, o, rawVisitDBTypes, true, rawVisitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawVisit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RawVisits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, rawVisitDBTypes, true, rawVisitPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RawVisit struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRawVisitsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(rawVisitAllColumns) == len(rawVisitPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RawVisit{}
	if err = randomize.Struct(seed, o, rawVisitDBTypes, true, rawVisitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RawVisit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RawVisits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, rawVisitDBTypes, true, rawVisitPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RawVisit struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(rawVisitAllColumns, rawVisitPrimaryKeyColumns) {
		fields = rawVisitAllColumns
	} else {
		fields = strmangle.SetComplement(
			rawVisitAllColumns,
			rawVisitPrimaryKeyColumns,
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

	slice := RawVisitSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRawVisitsUpsert(t *testing.T) {
	t.Parallel()

	if len(rawVisitAllColumns) == len(rawVisitPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RawVisit{}
	if err = randomize.Struct(seed, &o, rawVisitDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RawVisit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RawVisit: %s", err)
	}

	count, err := RawVisits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, rawVisitDBTypes, false, rawVisitPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RawVisit struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RawVisit: %s", err)
	}

	count, err = RawVisits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
