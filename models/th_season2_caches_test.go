// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testTHSeason2Caches(t *testing.T) {
	t.Parallel()

	query := THSeason2Caches()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTHSeason2CachesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSeason2Cach{}
	if err = randomize.Struct(seed, o, thSeason2CachDBTypes, true, thSeason2CachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeason2Cach struct: %s", err)
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

	count, err := THSeason2Caches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTHSeason2CachesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSeason2Cach{}
	if err = randomize.Struct(seed, o, thSeason2CachDBTypes, true, thSeason2CachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeason2Cach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := THSeason2Caches().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := THSeason2Caches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTHSeason2CachesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSeason2Cach{}
	if err = randomize.Struct(seed, o, thSeason2CachDBTypes, true, thSeason2CachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeason2Cach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := THSeason2CachSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := THSeason2Caches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTHSeason2CachesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSeason2Cach{}
	if err = randomize.Struct(seed, o, thSeason2CachDBTypes, true, thSeason2CachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeason2Cach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := THSeason2CachExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if THSeason2Cach exists: %s", err)
	}
	if !e {
		t.Errorf("Expected THSeason2CachExists to return true, but got false.")
	}
}

func testTHSeason2CachesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSeason2Cach{}
	if err = randomize.Struct(seed, o, thSeason2CachDBTypes, true, thSeason2CachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeason2Cach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	thSeason2CachFound, err := FindTHSeason2Cach(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if thSeason2CachFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTHSeason2CachesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSeason2Cach{}
	if err = randomize.Struct(seed, o, thSeason2CachDBTypes, true, thSeason2CachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeason2Cach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = THSeason2Caches().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTHSeason2CachesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSeason2Cach{}
	if err = randomize.Struct(seed, o, thSeason2CachDBTypes, true, thSeason2CachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeason2Cach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := THSeason2Caches().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTHSeason2CachesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	thSeason2CachOne := &THSeason2Cach{}
	thSeason2CachTwo := &THSeason2Cach{}
	if err = randomize.Struct(seed, thSeason2CachOne, thSeason2CachDBTypes, false, thSeason2CachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeason2Cach struct: %s", err)
	}
	if err = randomize.Struct(seed, thSeason2CachTwo, thSeason2CachDBTypes, false, thSeason2CachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeason2Cach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = thSeason2CachOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = thSeason2CachTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := THSeason2Caches().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTHSeason2CachesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	thSeason2CachOne := &THSeason2Cach{}
	thSeason2CachTwo := &THSeason2Cach{}
	if err = randomize.Struct(seed, thSeason2CachOne, thSeason2CachDBTypes, false, thSeason2CachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeason2Cach struct: %s", err)
	}
	if err = randomize.Struct(seed, thSeason2CachTwo, thSeason2CachDBTypes, false, thSeason2CachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeason2Cach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = thSeason2CachOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = thSeason2CachTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := THSeason2Caches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func thSeason2CachBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *THSeason2Cach) error {
	*o = THSeason2Cach{}
	return nil
}

func thSeason2CachAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *THSeason2Cach) error {
	*o = THSeason2Cach{}
	return nil
}

func thSeason2CachAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *THSeason2Cach) error {
	*o = THSeason2Cach{}
	return nil
}

func thSeason2CachBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *THSeason2Cach) error {
	*o = THSeason2Cach{}
	return nil
}

func thSeason2CachAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *THSeason2Cach) error {
	*o = THSeason2Cach{}
	return nil
}

func thSeason2CachBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *THSeason2Cach) error {
	*o = THSeason2Cach{}
	return nil
}

func thSeason2CachAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *THSeason2Cach) error {
	*o = THSeason2Cach{}
	return nil
}

func thSeason2CachBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *THSeason2Cach) error {
	*o = THSeason2Cach{}
	return nil
}

func thSeason2CachAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *THSeason2Cach) error {
	*o = THSeason2Cach{}
	return nil
}

func testTHSeason2CachesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &THSeason2Cach{}
	o := &THSeason2Cach{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, thSeason2CachDBTypes, false); err != nil {
		t.Errorf("Unable to randomize THSeason2Cach object: %s", err)
	}

	AddTHSeason2CachHook(boil.BeforeInsertHook, thSeason2CachBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	thSeason2CachBeforeInsertHooks = []THSeason2CachHook{}

	AddTHSeason2CachHook(boil.AfterInsertHook, thSeason2CachAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	thSeason2CachAfterInsertHooks = []THSeason2CachHook{}

	AddTHSeason2CachHook(boil.AfterSelectHook, thSeason2CachAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	thSeason2CachAfterSelectHooks = []THSeason2CachHook{}

	AddTHSeason2CachHook(boil.BeforeUpdateHook, thSeason2CachBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	thSeason2CachBeforeUpdateHooks = []THSeason2CachHook{}

	AddTHSeason2CachHook(boil.AfterUpdateHook, thSeason2CachAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	thSeason2CachAfterUpdateHooks = []THSeason2CachHook{}

	AddTHSeason2CachHook(boil.BeforeDeleteHook, thSeason2CachBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	thSeason2CachBeforeDeleteHooks = []THSeason2CachHook{}

	AddTHSeason2CachHook(boil.AfterDeleteHook, thSeason2CachAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	thSeason2CachAfterDeleteHooks = []THSeason2CachHook{}

	AddTHSeason2CachHook(boil.BeforeUpsertHook, thSeason2CachBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	thSeason2CachBeforeUpsertHooks = []THSeason2CachHook{}

	AddTHSeason2CachHook(boil.AfterUpsertHook, thSeason2CachAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	thSeason2CachAfterUpsertHooks = []THSeason2CachHook{}
}

func testTHSeason2CachesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSeason2Cach{}
	if err = randomize.Struct(seed, o, thSeason2CachDBTypes, true, thSeason2CachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeason2Cach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := THSeason2Caches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTHSeason2CachesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSeason2Cach{}
	if err = randomize.Struct(seed, o, thSeason2CachDBTypes, true); err != nil {
		t.Errorf("Unable to randomize THSeason2Cach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(thSeason2CachColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := THSeason2Caches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTHSeason2CachesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSeason2Cach{}
	if err = randomize.Struct(seed, o, thSeason2CachDBTypes, true, thSeason2CachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeason2Cach struct: %s", err)
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

func testTHSeason2CachesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSeason2Cach{}
	if err = randomize.Struct(seed, o, thSeason2CachDBTypes, true, thSeason2CachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeason2Cach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := THSeason2CachSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTHSeason2CachesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSeason2Cach{}
	if err = randomize.Struct(seed, o, thSeason2CachDBTypes, true, thSeason2CachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeason2Cach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := THSeason2Caches().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	thSeason2CachDBTypes = map[string]string{`ID`: `integer`, `SeasonID`: `bigint`, `IsVip`: `boolean`, `Data`: `json`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_                    = bytes.MinRead
)

func testTHSeason2CachesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(thSeason2CachPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(thSeason2CachAllColumns) == len(thSeason2CachPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &THSeason2Cach{}
	if err = randomize.Struct(seed, o, thSeason2CachDBTypes, true, thSeason2CachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeason2Cach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := THSeason2Caches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, thSeason2CachDBTypes, true, thSeason2CachPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize THSeason2Cach struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTHSeason2CachesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(thSeason2CachAllColumns) == len(thSeason2CachPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &THSeason2Cach{}
	if err = randomize.Struct(seed, o, thSeason2CachDBTypes, true, thSeason2CachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSeason2Cach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := THSeason2Caches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, thSeason2CachDBTypes, true, thSeason2CachPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize THSeason2Cach struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(thSeason2CachAllColumns, thSeason2CachPrimaryKeyColumns) {
		fields = thSeason2CachAllColumns
	} else {
		fields = strmangle.SetComplement(
			thSeason2CachAllColumns,
			thSeason2CachPrimaryKeyColumns,
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

	slice := THSeason2CachSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTHSeason2CachesUpsert(t *testing.T) {
	t.Parallel()

	if len(thSeason2CachAllColumns) == len(thSeason2CachPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := THSeason2Cach{}
	if err = randomize.Struct(seed, &o, thSeason2CachDBTypes, true); err != nil {
		t.Errorf("Unable to randomize THSeason2Cach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert THSeason2Cach: %s", err)
	}

	count, err := THSeason2Caches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, thSeason2CachDBTypes, false, thSeason2CachPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize THSeason2Cach struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert THSeason2Cach: %s", err)
	}

	count, err = THSeason2Caches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
