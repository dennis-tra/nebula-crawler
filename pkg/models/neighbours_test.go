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

func testNeighbours(t *testing.T) {
	t.Parallel()

	query := Neighbours()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testNeighboursDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Neighbour{}
	if err = randomize.Struct(seed, o, neighbourDBTypes, true, neighbourColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbour struct: %s", err)
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

	count, err := Neighbours().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNeighboursQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Neighbour{}
	if err = randomize.Struct(seed, o, neighbourDBTypes, true, neighbourColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbour struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Neighbours().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Neighbours().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNeighboursSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Neighbour{}
	if err = randomize.Struct(seed, o, neighbourDBTypes, true, neighbourColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbour struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NeighbourSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Neighbours().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNeighboursExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Neighbour{}
	if err = randomize.Struct(seed, o, neighbourDBTypes, true, neighbourColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbour struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := NeighbourExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Neighbour exists: %s", err)
	}
	if !e {
		t.Errorf("Expected NeighbourExists to return true, but got false.")
	}
}

func testNeighboursFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Neighbour{}
	if err = randomize.Struct(seed, o, neighbourDBTypes, true, neighbourColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbour struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	neighbourFound, err := FindNeighbour(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if neighbourFound == nil {
		t.Error("want a record, got nil")
	}
}

func testNeighboursBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Neighbour{}
	if err = randomize.Struct(seed, o, neighbourDBTypes, true, neighbourColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbour struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Neighbours().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testNeighboursOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Neighbour{}
	if err = randomize.Struct(seed, o, neighbourDBTypes, true, neighbourColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbour struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Neighbours().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testNeighboursAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	neighbourOne := &Neighbour{}
	neighbourTwo := &Neighbour{}
	if err = randomize.Struct(seed, neighbourOne, neighbourDBTypes, false, neighbourColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbour struct: %s", err)
	}
	if err = randomize.Struct(seed, neighbourTwo, neighbourDBTypes, false, neighbourColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbour struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = neighbourOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = neighbourTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Neighbours().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testNeighboursCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	neighbourOne := &Neighbour{}
	neighbourTwo := &Neighbour{}
	if err = randomize.Struct(seed, neighbourOne, neighbourDBTypes, false, neighbourColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbour struct: %s", err)
	}
	if err = randomize.Struct(seed, neighbourTwo, neighbourDBTypes, false, neighbourColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbour struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = neighbourOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = neighbourTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Neighbours().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func neighbourBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Neighbour) error {
	*o = Neighbour{}
	return nil
}

func neighbourAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Neighbour) error {
	*o = Neighbour{}
	return nil
}

func neighbourAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Neighbour) error {
	*o = Neighbour{}
	return nil
}

func neighbourBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Neighbour) error {
	*o = Neighbour{}
	return nil
}

func neighbourAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Neighbour) error {
	*o = Neighbour{}
	return nil
}

func neighbourBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Neighbour) error {
	*o = Neighbour{}
	return nil
}

func neighbourAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Neighbour) error {
	*o = Neighbour{}
	return nil
}

func neighbourBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Neighbour) error {
	*o = Neighbour{}
	return nil
}

func neighbourAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Neighbour) error {
	*o = Neighbour{}
	return nil
}

func testNeighboursHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Neighbour{}
	o := &Neighbour{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, neighbourDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Neighbour object: %s", err)
	}

	AddNeighbourHook(boil.BeforeInsertHook, neighbourBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	neighbourBeforeInsertHooks = []NeighbourHook{}

	AddNeighbourHook(boil.AfterInsertHook, neighbourAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	neighbourAfterInsertHooks = []NeighbourHook{}

	AddNeighbourHook(boil.AfterSelectHook, neighbourAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	neighbourAfterSelectHooks = []NeighbourHook{}

	AddNeighbourHook(boil.BeforeUpdateHook, neighbourBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	neighbourBeforeUpdateHooks = []NeighbourHook{}

	AddNeighbourHook(boil.AfterUpdateHook, neighbourAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	neighbourAfterUpdateHooks = []NeighbourHook{}

	AddNeighbourHook(boil.BeforeDeleteHook, neighbourBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	neighbourBeforeDeleteHooks = []NeighbourHook{}

	AddNeighbourHook(boil.AfterDeleteHook, neighbourAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	neighbourAfterDeleteHooks = []NeighbourHook{}

	AddNeighbourHook(boil.BeforeUpsertHook, neighbourBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	neighbourBeforeUpsertHooks = []NeighbourHook{}

	AddNeighbourHook(boil.AfterUpsertHook, neighbourAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	neighbourAfterUpsertHooks = []NeighbourHook{}
}

func testNeighboursInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Neighbour{}
	if err = randomize.Struct(seed, o, neighbourDBTypes, true, neighbourColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbour struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Neighbours().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNeighboursInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Neighbour{}
	if err = randomize.Struct(seed, o, neighbourDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Neighbour struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(neighbourColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Neighbours().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNeighboursReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Neighbour{}
	if err = randomize.Struct(seed, o, neighbourDBTypes, true, neighbourColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbour struct: %s", err)
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

func testNeighboursReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Neighbour{}
	if err = randomize.Struct(seed, o, neighbourDBTypes, true, neighbourColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbour struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NeighbourSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testNeighboursSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Neighbour{}
	if err = randomize.Struct(seed, o, neighbourDBTypes, true, neighbourColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbour struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Neighbours().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	neighbourDBTypes = map[string]string{`ID`: `integer`, `PeerID`: `character varying`, `NeighbourPeerID`: `character varying`, `CreatedAt`: `timestamp with time zone`, `CrawlStartAt`: `timestamp with time zone`}
	_                = bytes.MinRead
)

func testNeighboursUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(neighbourPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(neighbourAllColumns) == len(neighbourPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Neighbour{}
	if err = randomize.Struct(seed, o, neighbourDBTypes, true, neighbourColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbour struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Neighbours().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, neighbourDBTypes, true, neighbourPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Neighbour struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testNeighboursSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(neighbourAllColumns) == len(neighbourPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Neighbour{}
	if err = randomize.Struct(seed, o, neighbourDBTypes, true, neighbourColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Neighbour struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Neighbours().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, neighbourDBTypes, true, neighbourPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Neighbour struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(neighbourAllColumns, neighbourPrimaryKeyColumns) {
		fields = neighbourAllColumns
	} else {
		fields = strmangle.SetComplement(
			neighbourAllColumns,
			neighbourPrimaryKeyColumns,
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

	slice := NeighbourSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testNeighboursUpsert(t *testing.T) {
	t.Parallel()

	if len(neighbourAllColumns) == len(neighbourPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Neighbour{}
	if err = randomize.Struct(seed, &o, neighbourDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Neighbour struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Neighbour: %s", err)
	}

	count, err := Neighbours().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, neighbourDBTypes, false, neighbourPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Neighbour struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Neighbour: %s", err)
	}

	count, err = Neighbours().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
