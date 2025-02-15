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

func testSeasonAreaCaches(t *testing.T) {
	t.Parallel()

	query := SeasonAreaCaches()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSeasonAreaCachesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SeasonAreaCach{}
	if err = randomize.Struct(seed, o, seasonAreaCachDBTypes, true, seasonAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SeasonAreaCach struct: %s", err)
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

	count, err := SeasonAreaCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSeasonAreaCachesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SeasonAreaCach{}
	if err = randomize.Struct(seed, o, seasonAreaCachDBTypes, true, seasonAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SeasonAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := SeasonAreaCaches().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SeasonAreaCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSeasonAreaCachesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SeasonAreaCach{}
	if err = randomize.Struct(seed, o, seasonAreaCachDBTypes, true, seasonAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SeasonAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SeasonAreaCachSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SeasonAreaCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSeasonAreaCachesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SeasonAreaCach{}
	if err = randomize.Struct(seed, o, seasonAreaCachDBTypes, true, seasonAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SeasonAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SeasonAreaCachExists(ctx, tx, o.SeasonID)
	if err != nil {
		t.Errorf("Unable to check if SeasonAreaCach exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SeasonAreaCachExists to return true, but got false.")
	}
}

func testSeasonAreaCachesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SeasonAreaCach{}
	if err = randomize.Struct(seed, o, seasonAreaCachDBTypes, true, seasonAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SeasonAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	seasonAreaCachFound, err := FindSeasonAreaCach(ctx, tx, o.SeasonID)
	if err != nil {
		t.Error(err)
	}

	if seasonAreaCachFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSeasonAreaCachesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SeasonAreaCach{}
	if err = randomize.Struct(seed, o, seasonAreaCachDBTypes, true, seasonAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SeasonAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = SeasonAreaCaches().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSeasonAreaCachesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SeasonAreaCach{}
	if err = randomize.Struct(seed, o, seasonAreaCachDBTypes, true, seasonAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SeasonAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := SeasonAreaCaches().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSeasonAreaCachesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	seasonAreaCachOne := &SeasonAreaCach{}
	seasonAreaCachTwo := &SeasonAreaCach{}
	if err = randomize.Struct(seed, seasonAreaCachOne, seasonAreaCachDBTypes, false, seasonAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SeasonAreaCach struct: %s", err)
	}
	if err = randomize.Struct(seed, seasonAreaCachTwo, seasonAreaCachDBTypes, false, seasonAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SeasonAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = seasonAreaCachOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = seasonAreaCachTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SeasonAreaCaches().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSeasonAreaCachesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	seasonAreaCachOne := &SeasonAreaCach{}
	seasonAreaCachTwo := &SeasonAreaCach{}
	if err = randomize.Struct(seed, seasonAreaCachOne, seasonAreaCachDBTypes, false, seasonAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SeasonAreaCach struct: %s", err)
	}
	if err = randomize.Struct(seed, seasonAreaCachTwo, seasonAreaCachDBTypes, false, seasonAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SeasonAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = seasonAreaCachOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = seasonAreaCachTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SeasonAreaCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func seasonAreaCachBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *SeasonAreaCach) error {
	*o = SeasonAreaCach{}
	return nil
}

func seasonAreaCachAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *SeasonAreaCach) error {
	*o = SeasonAreaCach{}
	return nil
}

func seasonAreaCachAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *SeasonAreaCach) error {
	*o = SeasonAreaCach{}
	return nil
}

func seasonAreaCachBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *SeasonAreaCach) error {
	*o = SeasonAreaCach{}
	return nil
}

func seasonAreaCachAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *SeasonAreaCach) error {
	*o = SeasonAreaCach{}
	return nil
}

func seasonAreaCachBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *SeasonAreaCach) error {
	*o = SeasonAreaCach{}
	return nil
}

func seasonAreaCachAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *SeasonAreaCach) error {
	*o = SeasonAreaCach{}
	return nil
}

func seasonAreaCachBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *SeasonAreaCach) error {
	*o = SeasonAreaCach{}
	return nil
}

func seasonAreaCachAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *SeasonAreaCach) error {
	*o = SeasonAreaCach{}
	return nil
}

func testSeasonAreaCachesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &SeasonAreaCach{}
	o := &SeasonAreaCach{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, seasonAreaCachDBTypes, false); err != nil {
		t.Errorf("Unable to randomize SeasonAreaCach object: %s", err)
	}

	AddSeasonAreaCachHook(boil.BeforeInsertHook, seasonAreaCachBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	seasonAreaCachBeforeInsertHooks = []SeasonAreaCachHook{}

	AddSeasonAreaCachHook(boil.AfterInsertHook, seasonAreaCachAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	seasonAreaCachAfterInsertHooks = []SeasonAreaCachHook{}

	AddSeasonAreaCachHook(boil.AfterSelectHook, seasonAreaCachAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	seasonAreaCachAfterSelectHooks = []SeasonAreaCachHook{}

	AddSeasonAreaCachHook(boil.BeforeUpdateHook, seasonAreaCachBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	seasonAreaCachBeforeUpdateHooks = []SeasonAreaCachHook{}

	AddSeasonAreaCachHook(boil.AfterUpdateHook, seasonAreaCachAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	seasonAreaCachAfterUpdateHooks = []SeasonAreaCachHook{}

	AddSeasonAreaCachHook(boil.BeforeDeleteHook, seasonAreaCachBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	seasonAreaCachBeforeDeleteHooks = []SeasonAreaCachHook{}

	AddSeasonAreaCachHook(boil.AfterDeleteHook, seasonAreaCachAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	seasonAreaCachAfterDeleteHooks = []SeasonAreaCachHook{}

	AddSeasonAreaCachHook(boil.BeforeUpsertHook, seasonAreaCachBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	seasonAreaCachBeforeUpsertHooks = []SeasonAreaCachHook{}

	AddSeasonAreaCachHook(boil.AfterUpsertHook, seasonAreaCachAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	seasonAreaCachAfterUpsertHooks = []SeasonAreaCachHook{}
}

func testSeasonAreaCachesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SeasonAreaCach{}
	if err = randomize.Struct(seed, o, seasonAreaCachDBTypes, true, seasonAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SeasonAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SeasonAreaCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSeasonAreaCachesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SeasonAreaCach{}
	if err = randomize.Struct(seed, o, seasonAreaCachDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SeasonAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(seasonAreaCachColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := SeasonAreaCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSeasonAreaCachesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SeasonAreaCach{}
	if err = randomize.Struct(seed, o, seasonAreaCachDBTypes, true, seasonAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SeasonAreaCach struct: %s", err)
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

func testSeasonAreaCachesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SeasonAreaCach{}
	if err = randomize.Struct(seed, o, seasonAreaCachDBTypes, true, seasonAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SeasonAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SeasonAreaCachSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSeasonAreaCachesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SeasonAreaCach{}
	if err = randomize.Struct(seed, o, seasonAreaCachDBTypes, true, seasonAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SeasonAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SeasonAreaCaches().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	seasonAreaCachDBTypes = map[string]string{`SeasonID`: `bigint`, `CN`: `boolean`, `HK`: `boolean`, `TW`: `boolean`, `TH`: `boolean`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_                     = bytes.MinRead
)

func testSeasonAreaCachesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(seasonAreaCachPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(seasonAreaCachAllColumns) == len(seasonAreaCachPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SeasonAreaCach{}
	if err = randomize.Struct(seed, o, seasonAreaCachDBTypes, true, seasonAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SeasonAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SeasonAreaCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, seasonAreaCachDBTypes, true, seasonAreaCachPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SeasonAreaCach struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSeasonAreaCachesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(seasonAreaCachAllColumns) == len(seasonAreaCachPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SeasonAreaCach{}
	if err = randomize.Struct(seed, o, seasonAreaCachDBTypes, true, seasonAreaCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SeasonAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SeasonAreaCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, seasonAreaCachDBTypes, true, seasonAreaCachPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SeasonAreaCach struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(seasonAreaCachAllColumns, seasonAreaCachPrimaryKeyColumns) {
		fields = seasonAreaCachAllColumns
	} else {
		fields = strmangle.SetComplement(
			seasonAreaCachAllColumns,
			seasonAreaCachPrimaryKeyColumns,
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

	slice := SeasonAreaCachSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSeasonAreaCachesUpsert(t *testing.T) {
	t.Parallel()

	if len(seasonAreaCachAllColumns) == len(seasonAreaCachPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := SeasonAreaCach{}
	if err = randomize.Struct(seed, &o, seasonAreaCachDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SeasonAreaCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SeasonAreaCach: %s", err)
	}

	count, err := SeasonAreaCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, seasonAreaCachDBTypes, false, seasonAreaCachPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SeasonAreaCach struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SeasonAreaCach: %s", err)
	}

	count, err = SeasonAreaCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
