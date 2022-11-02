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

func testNetworkInformations(t *testing.T) {
	t.Parallel()

	query := NetworkInformations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testNetworkInformationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkInformation{}
	if err = randomize.Struct(seed, o, networkInformationDBTypes, true, networkInformationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkInformation struct: %s", err)
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

	count, err := NetworkInformations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNetworkInformationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkInformation{}
	if err = randomize.Struct(seed, o, networkInformationDBTypes, true, networkInformationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkInformation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := NetworkInformations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := NetworkInformations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNetworkInformationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkInformation{}
	if err = randomize.Struct(seed, o, networkInformationDBTypes, true, networkInformationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkInformation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NetworkInformationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := NetworkInformations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNetworkInformationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkInformation{}
	if err = randomize.Struct(seed, o, networkInformationDBTypes, true, networkInformationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkInformation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := NetworkInformationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if NetworkInformation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected NetworkInformationExists to return true, but got false.")
	}
}

func testNetworkInformationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkInformation{}
	if err = randomize.Struct(seed, o, networkInformationDBTypes, true, networkInformationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkInformation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	networkInformationFound, err := FindNetworkInformation(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if networkInformationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testNetworkInformationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkInformation{}
	if err = randomize.Struct(seed, o, networkInformationDBTypes, true, networkInformationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkInformation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = NetworkInformations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testNetworkInformationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkInformation{}
	if err = randomize.Struct(seed, o, networkInformationDBTypes, true, networkInformationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkInformation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := NetworkInformations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testNetworkInformationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	networkInformationOne := &NetworkInformation{}
	networkInformationTwo := &NetworkInformation{}
	if err = randomize.Struct(seed, networkInformationOne, networkInformationDBTypes, false, networkInformationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkInformation struct: %s", err)
	}
	if err = randomize.Struct(seed, networkInformationTwo, networkInformationDBTypes, false, networkInformationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkInformation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = networkInformationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = networkInformationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := NetworkInformations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testNetworkInformationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	networkInformationOne := &NetworkInformation{}
	networkInformationTwo := &NetworkInformation{}
	if err = randomize.Struct(seed, networkInformationOne, networkInformationDBTypes, false, networkInformationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkInformation struct: %s", err)
	}
	if err = randomize.Struct(seed, networkInformationTwo, networkInformationDBTypes, false, networkInformationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkInformation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = networkInformationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = networkInformationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NetworkInformations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func networkInformationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *NetworkInformation) error {
	*o = NetworkInformation{}
	return nil
}

func networkInformationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *NetworkInformation) error {
	*o = NetworkInformation{}
	return nil
}

func networkInformationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *NetworkInformation) error {
	*o = NetworkInformation{}
	return nil
}

func networkInformationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *NetworkInformation) error {
	*o = NetworkInformation{}
	return nil
}

func networkInformationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *NetworkInformation) error {
	*o = NetworkInformation{}
	return nil
}

func networkInformationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *NetworkInformation) error {
	*o = NetworkInformation{}
	return nil
}

func networkInformationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *NetworkInformation) error {
	*o = NetworkInformation{}
	return nil
}

func networkInformationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *NetworkInformation) error {
	*o = NetworkInformation{}
	return nil
}

func networkInformationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *NetworkInformation) error {
	*o = NetworkInformation{}
	return nil
}

func testNetworkInformationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &NetworkInformation{}
	o := &NetworkInformation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, networkInformationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize NetworkInformation object: %s", err)
	}

	AddNetworkInformationHook(boil.BeforeInsertHook, networkInformationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	networkInformationBeforeInsertHooks = []NetworkInformationHook{}

	AddNetworkInformationHook(boil.AfterInsertHook, networkInformationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	networkInformationAfterInsertHooks = []NetworkInformationHook{}

	AddNetworkInformationHook(boil.AfterSelectHook, networkInformationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	networkInformationAfterSelectHooks = []NetworkInformationHook{}

	AddNetworkInformationHook(boil.BeforeUpdateHook, networkInformationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	networkInformationBeforeUpdateHooks = []NetworkInformationHook{}

	AddNetworkInformationHook(boil.AfterUpdateHook, networkInformationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	networkInformationAfterUpdateHooks = []NetworkInformationHook{}

	AddNetworkInformationHook(boil.BeforeDeleteHook, networkInformationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	networkInformationBeforeDeleteHooks = []NetworkInformationHook{}

	AddNetworkInformationHook(boil.AfterDeleteHook, networkInformationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	networkInformationAfterDeleteHooks = []NetworkInformationHook{}

	AddNetworkInformationHook(boil.BeforeUpsertHook, networkInformationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	networkInformationBeforeUpsertHooks = []NetworkInformationHook{}

	AddNetworkInformationHook(boil.AfterUpsertHook, networkInformationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	networkInformationAfterUpsertHooks = []NetworkInformationHook{}
}

func testNetworkInformationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkInformation{}
	if err = randomize.Struct(seed, o, networkInformationDBTypes, true, networkInformationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkInformation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NetworkInformations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNetworkInformationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkInformation{}
	if err = randomize.Struct(seed, o, networkInformationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize NetworkInformation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(networkInformationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := NetworkInformations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNetworkInformationToOnePeerUsingPeer(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local NetworkInformation
	var foreign Peer

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, networkInformationDBTypes, false, networkInformationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkInformation struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, peerDBTypes, false, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.PeerID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Peer().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := NetworkInformationSlice{&local}
	if err = local.L.LoadPeer(ctx, tx, false, (*[]*NetworkInformation)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Peer == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Peer = nil
	if err = local.L.LoadPeer(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Peer == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testNetworkInformationToOneSetOpPeerUsingPeer(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a NetworkInformation
	var b, c Peer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, networkInformationDBTypes, false, strmangle.SetComplement(networkInformationPrimaryKeyColumns, networkInformationColumnsWithoutDefault)...); err != nil {
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
		err = a.SetPeer(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Peer != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.NetworkInformations[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PeerID != x.ID {
			t.Error("foreign key was wrong value", a.PeerID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PeerID))
		reflect.Indirect(reflect.ValueOf(&a.PeerID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PeerID != x.ID {
			t.Error("foreign key was wrong value", a.PeerID, x.ID)
		}
	}
}

func testNetworkInformationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkInformation{}
	if err = randomize.Struct(seed, o, networkInformationDBTypes, true, networkInformationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkInformation struct: %s", err)
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

func testNetworkInformationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkInformation{}
	if err = randomize.Struct(seed, o, networkInformationDBTypes, true, networkInformationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkInformation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NetworkInformationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testNetworkInformationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkInformation{}
	if err = randomize.Struct(seed, o, networkInformationDBTypes, true, networkInformationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkInformation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := NetworkInformations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	networkInformationDBTypes = map[string]string{`ID`: `integer`, `PeerID`: `bigint`, `SupportsIpv6`: `boolean`, `SupportsIpv6Error`: `text`, `RouterHTML`: `text`, `RouterHTMLError`: `text`, `CreatedAt`: `timestamp with time zone`}
	_                         = bytes.MinRead
)

func testNetworkInformationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(networkInformationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(networkInformationAllColumns) == len(networkInformationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &NetworkInformation{}
	if err = randomize.Struct(seed, o, networkInformationDBTypes, true, networkInformationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkInformation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NetworkInformations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, networkInformationDBTypes, true, networkInformationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NetworkInformation struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testNetworkInformationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(networkInformationAllColumns) == len(networkInformationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &NetworkInformation{}
	if err = randomize.Struct(seed, o, networkInformationDBTypes, true, networkInformationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkInformation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NetworkInformations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, networkInformationDBTypes, true, networkInformationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NetworkInformation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(networkInformationAllColumns, networkInformationPrimaryKeyColumns) {
		fields = networkInformationAllColumns
	} else {
		fields = strmangle.SetComplement(
			networkInformationAllColumns,
			networkInformationPrimaryKeyColumns,
		)
		fields = strmangle.SetComplement(fields, networkInformationGeneratedColumns)
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

	slice := NetworkInformationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testNetworkInformationsUpsert(t *testing.T) {
	t.Parallel()

	if len(networkInformationAllColumns) == len(networkInformationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := NetworkInformation{}
	if err = randomize.Struct(seed, &o, networkInformationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize NetworkInformation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert NetworkInformation: %s", err)
	}

	count, err := NetworkInformations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, networkInformationDBTypes, false, networkInformationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NetworkInformation struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert NetworkInformation: %s", err)
	}

	count, err = NetworkInformations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
