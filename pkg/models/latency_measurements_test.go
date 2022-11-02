// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testLatencyMeasurements(t *testing.T) {
	t.Parallel()

	query := LatencyMeasurements()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testLatencyMeasurementsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LatencyMeasurement{}
	if err = randomize.Struct(seed, o, latencyMeasurementDBTypes, true, latencyMeasurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LatencyMeasurement struct: %s", err)
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

	count, err := LatencyMeasurements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLatencyMeasurementsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LatencyMeasurement{}
	if err = randomize.Struct(seed, o, latencyMeasurementDBTypes, true, latencyMeasurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LatencyMeasurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := LatencyMeasurements().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := LatencyMeasurements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLatencyMeasurementsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LatencyMeasurement{}
	if err = randomize.Struct(seed, o, latencyMeasurementDBTypes, true, latencyMeasurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LatencyMeasurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LatencyMeasurementSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := LatencyMeasurements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLatencyMeasurementsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LatencyMeasurement{}
	if err = randomize.Struct(seed, o, latencyMeasurementDBTypes, true, latencyMeasurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LatencyMeasurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := LatencyMeasurementExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if LatencyMeasurement exists: %s", err)
	}
	if !e {
		t.Errorf("Expected LatencyMeasurementExists to return true, but got false.")
	}
}

func testLatencyMeasurementsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LatencyMeasurement{}
	if err = randomize.Struct(seed, o, latencyMeasurementDBTypes, true, latencyMeasurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LatencyMeasurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	latencyMeasurementFound, err := FindLatencyMeasurement(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if latencyMeasurementFound == nil {
		t.Error("want a record, got nil")
	}
}

func testLatencyMeasurementsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LatencyMeasurement{}
	if err = randomize.Struct(seed, o, latencyMeasurementDBTypes, true, latencyMeasurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LatencyMeasurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = LatencyMeasurements().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testLatencyMeasurementsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LatencyMeasurement{}
	if err = randomize.Struct(seed, o, latencyMeasurementDBTypes, true, latencyMeasurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LatencyMeasurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := LatencyMeasurements().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testLatencyMeasurementsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	latencyMeasurementOne := &LatencyMeasurement{}
	latencyMeasurementTwo := &LatencyMeasurement{}
	if err = randomize.Struct(seed, latencyMeasurementOne, latencyMeasurementDBTypes, false, latencyMeasurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LatencyMeasurement struct: %s", err)
	}
	if err = randomize.Struct(seed, latencyMeasurementTwo, latencyMeasurementDBTypes, false, latencyMeasurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LatencyMeasurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = latencyMeasurementOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = latencyMeasurementTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := LatencyMeasurements().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testLatencyMeasurementsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	latencyMeasurementOne := &LatencyMeasurement{}
	latencyMeasurementTwo := &LatencyMeasurement{}
	if err = randomize.Struct(seed, latencyMeasurementOne, latencyMeasurementDBTypes, false, latencyMeasurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LatencyMeasurement struct: %s", err)
	}
	if err = randomize.Struct(seed, latencyMeasurementTwo, latencyMeasurementDBTypes, false, latencyMeasurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LatencyMeasurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = latencyMeasurementOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = latencyMeasurementTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LatencyMeasurements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func latencyMeasurementBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *LatencyMeasurement) error {
	*o = LatencyMeasurement{}
	return nil
}

func latencyMeasurementAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *LatencyMeasurement) error {
	*o = LatencyMeasurement{}
	return nil
}

func latencyMeasurementAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *LatencyMeasurement) error {
	*o = LatencyMeasurement{}
	return nil
}

func latencyMeasurementBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *LatencyMeasurement) error {
	*o = LatencyMeasurement{}
	return nil
}

func latencyMeasurementAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *LatencyMeasurement) error {
	*o = LatencyMeasurement{}
	return nil
}

func latencyMeasurementBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *LatencyMeasurement) error {
	*o = LatencyMeasurement{}
	return nil
}

func latencyMeasurementAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *LatencyMeasurement) error {
	*o = LatencyMeasurement{}
	return nil
}

func latencyMeasurementBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *LatencyMeasurement) error {
	*o = LatencyMeasurement{}
	return nil
}

func latencyMeasurementAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *LatencyMeasurement) error {
	*o = LatencyMeasurement{}
	return nil
}

func testLatencyMeasurementsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &LatencyMeasurement{}
	o := &LatencyMeasurement{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, latencyMeasurementDBTypes, false); err != nil {
		t.Errorf("Unable to randomize LatencyMeasurement object: %s", err)
	}

	AddLatencyMeasurementHook(boil.BeforeInsertHook, latencyMeasurementBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	latencyMeasurementBeforeInsertHooks = []LatencyMeasurementHook{}

	AddLatencyMeasurementHook(boil.AfterInsertHook, latencyMeasurementAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	latencyMeasurementAfterInsertHooks = []LatencyMeasurementHook{}

	AddLatencyMeasurementHook(boil.AfterSelectHook, latencyMeasurementAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	latencyMeasurementAfterSelectHooks = []LatencyMeasurementHook{}

	AddLatencyMeasurementHook(boil.BeforeUpdateHook, latencyMeasurementBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	latencyMeasurementBeforeUpdateHooks = []LatencyMeasurementHook{}

	AddLatencyMeasurementHook(boil.AfterUpdateHook, latencyMeasurementAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	latencyMeasurementAfterUpdateHooks = []LatencyMeasurementHook{}

	AddLatencyMeasurementHook(boil.BeforeDeleteHook, latencyMeasurementBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	latencyMeasurementBeforeDeleteHooks = []LatencyMeasurementHook{}

	AddLatencyMeasurementHook(boil.AfterDeleteHook, latencyMeasurementAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	latencyMeasurementAfterDeleteHooks = []LatencyMeasurementHook{}

	AddLatencyMeasurementHook(boil.BeforeUpsertHook, latencyMeasurementBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	latencyMeasurementBeforeUpsertHooks = []LatencyMeasurementHook{}

	AddLatencyMeasurementHook(boil.AfterUpsertHook, latencyMeasurementAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	latencyMeasurementAfterUpsertHooks = []LatencyMeasurementHook{}
}

func testLatencyMeasurementsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LatencyMeasurement{}
	if err = randomize.Struct(seed, o, latencyMeasurementDBTypes, true, latencyMeasurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LatencyMeasurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LatencyMeasurements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLatencyMeasurementsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LatencyMeasurement{}
	if err = randomize.Struct(seed, o, latencyMeasurementDBTypes, true); err != nil {
		t.Errorf("Unable to randomize LatencyMeasurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(latencyMeasurementColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := LatencyMeasurements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLatencyMeasurementToOneHolePunchResultUsingHolePunchResult(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local LatencyMeasurement
	var foreign HolePunchResult

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, latencyMeasurementDBTypes, false, latencyMeasurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LatencyMeasurement struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, holePunchResultDBTypes, false, holePunchResultColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize HolePunchResult struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.HolePunchResultID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.HolePunchResult().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := LatencyMeasurementSlice{&local}
	if err = local.L.LoadHolePunchResult(ctx, tx, false, (*[]*LatencyMeasurement)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.HolePunchResult == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.HolePunchResult = nil
	if err = local.L.LoadHolePunchResult(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.HolePunchResult == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testLatencyMeasurementToOneMultiAddressUsingMultiAddress(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local LatencyMeasurement
	var foreign MultiAddress

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, latencyMeasurementDBTypes, false, latencyMeasurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LatencyMeasurement struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, multiAddressDBTypes, false, multiAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MultiAddress struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.MultiAddressID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.MultiAddress().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := LatencyMeasurementSlice{&local}
	if err = local.L.LoadMultiAddress(ctx, tx, false, (*[]*LatencyMeasurement)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.MultiAddress == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.MultiAddress = nil
	if err = local.L.LoadMultiAddress(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.MultiAddress == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testLatencyMeasurementToOnePeerUsingRemote(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local LatencyMeasurement
	var foreign Peer

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, latencyMeasurementDBTypes, false, latencyMeasurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LatencyMeasurement struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, peerDBTypes, false, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.RemoteID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Remote().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := LatencyMeasurementSlice{&local}
	if err = local.L.LoadRemote(ctx, tx, false, (*[]*LatencyMeasurement)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Remote == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Remote = nil
	if err = local.L.LoadRemote(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Remote == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testLatencyMeasurementToOneSetOpHolePunchResultUsingHolePunchResult(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a LatencyMeasurement
	var b, c HolePunchResult

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, latencyMeasurementDBTypes, false, strmangle.SetComplement(latencyMeasurementPrimaryKeyColumns, latencyMeasurementColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, holePunchResultDBTypes, false, strmangle.SetComplement(holePunchResultPrimaryKeyColumns, holePunchResultColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, holePunchResultDBTypes, false, strmangle.SetComplement(holePunchResultPrimaryKeyColumns, holePunchResultColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*HolePunchResult{&b, &c} {
		err = a.SetHolePunchResult(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.HolePunchResult != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.LatencyMeasurements[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.HolePunchResultID != x.ID {
			t.Error("foreign key was wrong value", a.HolePunchResultID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.HolePunchResultID))
		reflect.Indirect(reflect.ValueOf(&a.HolePunchResultID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.HolePunchResultID != x.ID {
			t.Error("foreign key was wrong value", a.HolePunchResultID, x.ID)
		}
	}
}
func testLatencyMeasurementToOneSetOpMultiAddressUsingMultiAddress(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a LatencyMeasurement
	var b, c MultiAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, latencyMeasurementDBTypes, false, strmangle.SetComplement(latencyMeasurementPrimaryKeyColumns, latencyMeasurementColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, multiAddressDBTypes, false, strmangle.SetComplement(multiAddressPrimaryKeyColumns, multiAddressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, multiAddressDBTypes, false, strmangle.SetComplement(multiAddressPrimaryKeyColumns, multiAddressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*MultiAddress{&b, &c} {
		err = a.SetMultiAddress(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.MultiAddress != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.LatencyMeasurements[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.MultiAddressID != x.ID {
			t.Error("foreign key was wrong value", a.MultiAddressID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.MultiAddressID))
		reflect.Indirect(reflect.ValueOf(&a.MultiAddressID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.MultiAddressID != x.ID {
			t.Error("foreign key was wrong value", a.MultiAddressID, x.ID)
		}
	}
}
func testLatencyMeasurementToOneSetOpPeerUsingRemote(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a LatencyMeasurement
	var b, c Peer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, latencyMeasurementDBTypes, false, strmangle.SetComplement(latencyMeasurementPrimaryKeyColumns, latencyMeasurementColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, peerDBTypes, false, strmangle.SetComplement(peerPrimaryKeyColumns, peerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, peerDBTypes, false, strmangle.SetComplement(peerPrimaryKeyColumns, peerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Peer{&b, &c} {
		err = a.SetRemote(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Remote != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RemoteLatencyMeasurements[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.RemoteID != x.ID {
			t.Error("foreign key was wrong value", a.RemoteID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.RemoteID))
		reflect.Indirect(reflect.ValueOf(&a.RemoteID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.RemoteID != x.ID {
			t.Error("foreign key was wrong value", a.RemoteID, x.ID)
		}
	}
}

func testLatencyMeasurementsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LatencyMeasurement{}
	if err = randomize.Struct(seed, o, latencyMeasurementDBTypes, true, latencyMeasurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LatencyMeasurement struct: %s", err)
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

func testLatencyMeasurementsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LatencyMeasurement{}
	if err = randomize.Struct(seed, o, latencyMeasurementDBTypes, true, latencyMeasurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LatencyMeasurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LatencyMeasurementSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testLatencyMeasurementsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LatencyMeasurement{}
	if err = randomize.Struct(seed, o, latencyMeasurementDBTypes, true, latencyMeasurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LatencyMeasurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := LatencyMeasurements().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	latencyMeasurementDBTypes = map[string]string{`ID`: `integer`, `RemoteID`: `bigint`, `HolePunchResultID`: `integer`, `MultiAddressID`: `bigint`, `Mtype`: `enum.latency_measurement_type('TO_RELAY','TO_REMOTE_THROUGH_RELAY','TO_REMOTE_AFTER_HOLEPUNCH')`, `RTTS`: `ARRAYdouble precision`, `RTTAvg`: `double precision`, `RTTMax`: `double precision`, `RTTMin`: `double precision`, `RTTSTD`: `double precision`}
	_                         = bytes.MinRead
)

func testLatencyMeasurementsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(latencyMeasurementPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(latencyMeasurementAllColumns) == len(latencyMeasurementPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &LatencyMeasurement{}
	if err = randomize.Struct(seed, o, latencyMeasurementDBTypes, true, latencyMeasurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LatencyMeasurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LatencyMeasurements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, latencyMeasurementDBTypes, true, latencyMeasurementPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize LatencyMeasurement struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testLatencyMeasurementsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(latencyMeasurementAllColumns) == len(latencyMeasurementPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &LatencyMeasurement{}
	if err = randomize.Struct(seed, o, latencyMeasurementDBTypes, true, latencyMeasurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LatencyMeasurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LatencyMeasurements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, latencyMeasurementDBTypes, true, latencyMeasurementPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize LatencyMeasurement struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(latencyMeasurementAllColumns, latencyMeasurementPrimaryKeyColumns) {
		fields = latencyMeasurementAllColumns
	} else {
		fields = strmangle.SetComplement(
			latencyMeasurementAllColumns,
			latencyMeasurementPrimaryKeyColumns,
		)
		fields = strmangle.SetComplement(fields, latencyMeasurementGeneratedColumns)
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

	slice := LatencyMeasurementSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testLatencyMeasurementsUpsert(t *testing.T) {
	t.Parallel()

	if len(latencyMeasurementAllColumns) == len(latencyMeasurementPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := LatencyMeasurement{}
	if err = randomize.Struct(seed, &o, latencyMeasurementDBTypes, true); err != nil {
		t.Errorf("Unable to randomize LatencyMeasurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert LatencyMeasurement: %s", err)
	}

	count, err := LatencyMeasurements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, latencyMeasurementDBTypes, false, latencyMeasurementPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize LatencyMeasurement struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert LatencyMeasurement: %s", err)
	}

	count, err = LatencyMeasurements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
