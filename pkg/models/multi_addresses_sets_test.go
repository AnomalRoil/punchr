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

func testMultiAddressesSets(t *testing.T) {
	t.Parallel()

	query := MultiAddressesSets()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMultiAddressesSetsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MultiAddressesSet{}
	if err = randomize.Struct(seed, o, multiAddressesSetDBTypes, true, multiAddressesSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MultiAddressesSet struct: %s", err)
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

	count, err := MultiAddressesSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMultiAddressesSetsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MultiAddressesSet{}
	if err = randomize.Struct(seed, o, multiAddressesSetDBTypes, true, multiAddressesSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MultiAddressesSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := MultiAddressesSets().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MultiAddressesSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMultiAddressesSetsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MultiAddressesSet{}
	if err = randomize.Struct(seed, o, multiAddressesSetDBTypes, true, multiAddressesSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MultiAddressesSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MultiAddressesSetSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MultiAddressesSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMultiAddressesSetsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MultiAddressesSet{}
	if err = randomize.Struct(seed, o, multiAddressesSetDBTypes, true, multiAddressesSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MultiAddressesSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MultiAddressesSetExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if MultiAddressesSet exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MultiAddressesSetExists to return true, but got false.")
	}
}

func testMultiAddressesSetsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MultiAddressesSet{}
	if err = randomize.Struct(seed, o, multiAddressesSetDBTypes, true, multiAddressesSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MultiAddressesSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	multiAddressesSetFound, err := FindMultiAddressesSet(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if multiAddressesSetFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMultiAddressesSetsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MultiAddressesSet{}
	if err = randomize.Struct(seed, o, multiAddressesSetDBTypes, true, multiAddressesSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MultiAddressesSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = MultiAddressesSets().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMultiAddressesSetsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MultiAddressesSet{}
	if err = randomize.Struct(seed, o, multiAddressesSetDBTypes, true, multiAddressesSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MultiAddressesSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := MultiAddressesSets().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMultiAddressesSetsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	multiAddressesSetOne := &MultiAddressesSet{}
	multiAddressesSetTwo := &MultiAddressesSet{}
	if err = randomize.Struct(seed, multiAddressesSetOne, multiAddressesSetDBTypes, false, multiAddressesSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MultiAddressesSet struct: %s", err)
	}
	if err = randomize.Struct(seed, multiAddressesSetTwo, multiAddressesSetDBTypes, false, multiAddressesSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MultiAddressesSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = multiAddressesSetOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = multiAddressesSetTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MultiAddressesSets().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMultiAddressesSetsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	multiAddressesSetOne := &MultiAddressesSet{}
	multiAddressesSetTwo := &MultiAddressesSet{}
	if err = randomize.Struct(seed, multiAddressesSetOne, multiAddressesSetDBTypes, false, multiAddressesSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MultiAddressesSet struct: %s", err)
	}
	if err = randomize.Struct(seed, multiAddressesSetTwo, multiAddressesSetDBTypes, false, multiAddressesSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MultiAddressesSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = multiAddressesSetOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = multiAddressesSetTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MultiAddressesSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func multiAddressesSetBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *MultiAddressesSet) error {
	*o = MultiAddressesSet{}
	return nil
}

func multiAddressesSetAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *MultiAddressesSet) error {
	*o = MultiAddressesSet{}
	return nil
}

func multiAddressesSetAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *MultiAddressesSet) error {
	*o = MultiAddressesSet{}
	return nil
}

func multiAddressesSetBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MultiAddressesSet) error {
	*o = MultiAddressesSet{}
	return nil
}

func multiAddressesSetAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MultiAddressesSet) error {
	*o = MultiAddressesSet{}
	return nil
}

func multiAddressesSetBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MultiAddressesSet) error {
	*o = MultiAddressesSet{}
	return nil
}

func multiAddressesSetAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MultiAddressesSet) error {
	*o = MultiAddressesSet{}
	return nil
}

func multiAddressesSetBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MultiAddressesSet) error {
	*o = MultiAddressesSet{}
	return nil
}

func multiAddressesSetAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MultiAddressesSet) error {
	*o = MultiAddressesSet{}
	return nil
}

func testMultiAddressesSetsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &MultiAddressesSet{}
	o := &MultiAddressesSet{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, multiAddressesSetDBTypes, false); err != nil {
		t.Errorf("Unable to randomize MultiAddressesSet object: %s", err)
	}

	AddMultiAddressesSetHook(boil.BeforeInsertHook, multiAddressesSetBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	multiAddressesSetBeforeInsertHooks = []MultiAddressesSetHook{}

	AddMultiAddressesSetHook(boil.AfterInsertHook, multiAddressesSetAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	multiAddressesSetAfterInsertHooks = []MultiAddressesSetHook{}

	AddMultiAddressesSetHook(boil.AfterSelectHook, multiAddressesSetAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	multiAddressesSetAfterSelectHooks = []MultiAddressesSetHook{}

	AddMultiAddressesSetHook(boil.BeforeUpdateHook, multiAddressesSetBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	multiAddressesSetBeforeUpdateHooks = []MultiAddressesSetHook{}

	AddMultiAddressesSetHook(boil.AfterUpdateHook, multiAddressesSetAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	multiAddressesSetAfterUpdateHooks = []MultiAddressesSetHook{}

	AddMultiAddressesSetHook(boil.BeforeDeleteHook, multiAddressesSetBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	multiAddressesSetBeforeDeleteHooks = []MultiAddressesSetHook{}

	AddMultiAddressesSetHook(boil.AfterDeleteHook, multiAddressesSetAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	multiAddressesSetAfterDeleteHooks = []MultiAddressesSetHook{}

	AddMultiAddressesSetHook(boil.BeforeUpsertHook, multiAddressesSetBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	multiAddressesSetBeforeUpsertHooks = []MultiAddressesSetHook{}

	AddMultiAddressesSetHook(boil.AfterUpsertHook, multiAddressesSetAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	multiAddressesSetAfterUpsertHooks = []MultiAddressesSetHook{}
}

func testMultiAddressesSetsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MultiAddressesSet{}
	if err = randomize.Struct(seed, o, multiAddressesSetDBTypes, true, multiAddressesSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MultiAddressesSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MultiAddressesSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMultiAddressesSetsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MultiAddressesSet{}
	if err = randomize.Struct(seed, o, multiAddressesSetDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MultiAddressesSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(multiAddressesSetColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := MultiAddressesSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMultiAddressesSetToManyListenMultiAddressesSetHolePunchResults(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MultiAddressesSet
	var b, c HolePunchResult

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, multiAddressesSetDBTypes, true, multiAddressesSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MultiAddressesSet struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, holePunchResultDBTypes, false, holePunchResultColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, holePunchResultDBTypes, false, holePunchResultColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ListenMultiAddressesSetID = a.ID
	c.ListenMultiAddressesSetID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ListenMultiAddressesSetHolePunchResults().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ListenMultiAddressesSetID == b.ListenMultiAddressesSetID {
			bFound = true
		}
		if v.ListenMultiAddressesSetID == c.ListenMultiAddressesSetID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := MultiAddressesSetSlice{&a}
	if err = a.L.LoadListenMultiAddressesSetHolePunchResults(ctx, tx, false, (*[]*MultiAddressesSet)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ListenMultiAddressesSetHolePunchResults); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ListenMultiAddressesSetHolePunchResults = nil
	if err = a.L.LoadListenMultiAddressesSetHolePunchResults(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ListenMultiAddressesSetHolePunchResults); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testMultiAddressesSetToManyAddOpListenMultiAddressesSetHolePunchResults(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MultiAddressesSet
	var b, c, d, e HolePunchResult

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, multiAddressesSetDBTypes, false, strmangle.SetComplement(multiAddressesSetPrimaryKeyColumns, multiAddressesSetColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*HolePunchResult{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, holePunchResultDBTypes, false, strmangle.SetComplement(holePunchResultPrimaryKeyColumns, holePunchResultColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*HolePunchResult{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddListenMultiAddressesSetHolePunchResults(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ListenMultiAddressesSetID {
			t.Error("foreign key was wrong value", a.ID, first.ListenMultiAddressesSetID)
		}
		if a.ID != second.ListenMultiAddressesSetID {
			t.Error("foreign key was wrong value", a.ID, second.ListenMultiAddressesSetID)
		}

		if first.R.ListenMultiAddressesSet != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.ListenMultiAddressesSet != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ListenMultiAddressesSetHolePunchResults[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ListenMultiAddressesSetHolePunchResults[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ListenMultiAddressesSetHolePunchResults().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testMultiAddressesSetsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MultiAddressesSet{}
	if err = randomize.Struct(seed, o, multiAddressesSetDBTypes, true, multiAddressesSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MultiAddressesSet struct: %s", err)
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

func testMultiAddressesSetsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MultiAddressesSet{}
	if err = randomize.Struct(seed, o, multiAddressesSetDBTypes, true, multiAddressesSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MultiAddressesSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MultiAddressesSetSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMultiAddressesSetsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MultiAddressesSet{}
	if err = randomize.Struct(seed, o, multiAddressesSetDBTypes, true, multiAddressesSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MultiAddressesSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MultiAddressesSets().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	multiAddressesSetDBTypes = map[string]string{`ID`: `integer`, `MultiAddressesIds`: `ARRAYinteger`, `Digest`: `bytea`}
	_                        = bytes.MinRead
)

func testMultiAddressesSetsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(multiAddressesSetPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(multiAddressesSetAllColumns) == len(multiAddressesSetPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MultiAddressesSet{}
	if err = randomize.Struct(seed, o, multiAddressesSetDBTypes, true, multiAddressesSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MultiAddressesSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MultiAddressesSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, multiAddressesSetDBTypes, true, multiAddressesSetPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MultiAddressesSet struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMultiAddressesSetsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(multiAddressesSetAllColumns) == len(multiAddressesSetPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MultiAddressesSet{}
	if err = randomize.Struct(seed, o, multiAddressesSetDBTypes, true, multiAddressesSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MultiAddressesSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MultiAddressesSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, multiAddressesSetDBTypes, true, multiAddressesSetPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MultiAddressesSet struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(multiAddressesSetAllColumns, multiAddressesSetPrimaryKeyColumns) {
		fields = multiAddressesSetAllColumns
	} else {
		fields = strmangle.SetComplement(
			multiAddressesSetAllColumns,
			multiAddressesSetPrimaryKeyColumns,
		)
		fields = strmangle.SetComplement(fields, multiAddressesSetGeneratedColumns)
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

	slice := MultiAddressesSetSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMultiAddressesSetsUpsert(t *testing.T) {
	t.Parallel()

	if len(multiAddressesSetAllColumns) == len(multiAddressesSetPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := MultiAddressesSet{}
	if err = randomize.Struct(seed, &o, multiAddressesSetDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MultiAddressesSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MultiAddressesSet: %s", err)
	}

	count, err := MultiAddressesSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, multiAddressesSetDBTypes, false, multiAddressesSetPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MultiAddressesSet struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MultiAddressesSet: %s", err)
	}

	count, err = MultiAddressesSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
