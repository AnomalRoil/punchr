// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// NetworkInformation is an object representing the database table.
type NetworkInformation struct {
	ID                int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	PeerID            int64       `boil:"peer_id" json:"peer_id" toml:"peer_id" yaml:"peer_id"`
	SupportsIpv6      null.Bool   `boil:"supports_ipv6" json:"supports_ipv6,omitempty" toml:"supports_ipv6" yaml:"supports_ipv6,omitempty"`
	SupportsIpv6Error null.String `boil:"supports_ipv6_error" json:"supports_ipv6_error,omitempty" toml:"supports_ipv6_error" yaml:"supports_ipv6_error,omitempty"`
	RouterHTML        null.String `boil:"router_html" json:"router_html,omitempty" toml:"router_html" yaml:"router_html,omitempty"`
	RouterHTMLError   null.String `boil:"router_html_error" json:"router_html_error,omitempty" toml:"router_html_error" yaml:"router_html_error,omitempty"`
	CreatedAt         time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *networkInformationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L networkInformationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var NetworkInformationColumns = struct {
	ID                string
	PeerID            string
	SupportsIpv6      string
	SupportsIpv6Error string
	RouterHTML        string
	RouterHTMLError   string
	CreatedAt         string
}{
	ID:                "id",
	PeerID:            "peer_id",
	SupportsIpv6:      "supports_ipv6",
	SupportsIpv6Error: "supports_ipv6_error",
	RouterHTML:        "router_html",
	RouterHTMLError:   "router_html_error",
	CreatedAt:         "created_at",
}

var NetworkInformationTableColumns = struct {
	ID                string
	PeerID            string
	SupportsIpv6      string
	SupportsIpv6Error string
	RouterHTML        string
	RouterHTMLError   string
	CreatedAt         string
}{
	ID:                "network_information.id",
	PeerID:            "network_information.peer_id",
	SupportsIpv6:      "network_information.supports_ipv6",
	SupportsIpv6Error: "network_information.supports_ipv6_error",
	RouterHTML:        "network_information.router_html",
	RouterHTMLError:   "network_information.router_html_error",
	CreatedAt:         "network_information.created_at",
}

// Generated where

var NetworkInformationWhere = struct {
	ID                whereHelperint
	PeerID            whereHelperint64
	SupportsIpv6      whereHelpernull_Bool
	SupportsIpv6Error whereHelpernull_String
	RouterHTML        whereHelpernull_String
	RouterHTMLError   whereHelpernull_String
	CreatedAt         whereHelpertime_Time
}{
	ID:                whereHelperint{field: "\"network_information\".\"id\""},
	PeerID:            whereHelperint64{field: "\"network_information\".\"peer_id\""},
	SupportsIpv6:      whereHelpernull_Bool{field: "\"network_information\".\"supports_ipv6\""},
	SupportsIpv6Error: whereHelpernull_String{field: "\"network_information\".\"supports_ipv6_error\""},
	RouterHTML:        whereHelpernull_String{field: "\"network_information\".\"router_html\""},
	RouterHTMLError:   whereHelpernull_String{field: "\"network_information\".\"router_html_error\""},
	CreatedAt:         whereHelpertime_Time{field: "\"network_information\".\"created_at\""},
}

// NetworkInformationRels is where relationship names are stored.
var NetworkInformationRels = struct {
	Peer string
}{
	Peer: "Peer",
}

// networkInformationR is where relationships are stored.
type networkInformationR struct {
	Peer *Peer `boil:"Peer" json:"Peer" toml:"Peer" yaml:"Peer"`
}

// NewStruct creates a new relationship struct
func (*networkInformationR) NewStruct() *networkInformationR {
	return &networkInformationR{}
}

func (r *networkInformationR) GetPeer() *Peer {
	if r == nil {
		return nil
	}
	return r.Peer
}

// networkInformationL is where Load methods for each relationship are stored.
type networkInformationL struct{}

var (
	networkInformationAllColumns            = []string{"id", "peer_id", "supports_ipv6", "supports_ipv6_error", "router_html", "router_html_error", "created_at"}
	networkInformationColumnsWithoutDefault = []string{"peer_id", "created_at"}
	networkInformationColumnsWithDefault    = []string{"id", "supports_ipv6", "supports_ipv6_error", "router_html", "router_html_error"}
	networkInformationPrimaryKeyColumns     = []string{"id"}
	networkInformationGeneratedColumns      = []string{"id"}
)

type (
	// NetworkInformationSlice is an alias for a slice of pointers to NetworkInformation.
	// This should almost always be used instead of []NetworkInformation.
	NetworkInformationSlice []*NetworkInformation
	// NetworkInformationHook is the signature for custom NetworkInformation hook methods
	NetworkInformationHook func(context.Context, boil.ContextExecutor, *NetworkInformation) error

	networkInformationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	networkInformationType                 = reflect.TypeOf(&NetworkInformation{})
	networkInformationMapping              = queries.MakeStructMapping(networkInformationType)
	networkInformationPrimaryKeyMapping, _ = queries.BindMapping(networkInformationType, networkInformationMapping, networkInformationPrimaryKeyColumns)
	networkInformationInsertCacheMut       sync.RWMutex
	networkInformationInsertCache          = make(map[string]insertCache)
	networkInformationUpdateCacheMut       sync.RWMutex
	networkInformationUpdateCache          = make(map[string]updateCache)
	networkInformationUpsertCacheMut       sync.RWMutex
	networkInformationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var networkInformationAfterSelectHooks []NetworkInformationHook

var networkInformationBeforeInsertHooks []NetworkInformationHook
var networkInformationAfterInsertHooks []NetworkInformationHook

var networkInformationBeforeUpdateHooks []NetworkInformationHook
var networkInformationAfterUpdateHooks []NetworkInformationHook

var networkInformationBeforeDeleteHooks []NetworkInformationHook
var networkInformationAfterDeleteHooks []NetworkInformationHook

var networkInformationBeforeUpsertHooks []NetworkInformationHook
var networkInformationAfterUpsertHooks []NetworkInformationHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *NetworkInformation) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range networkInformationAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *NetworkInformation) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range networkInformationBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *NetworkInformation) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range networkInformationAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *NetworkInformation) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range networkInformationBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *NetworkInformation) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range networkInformationAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *NetworkInformation) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range networkInformationBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *NetworkInformation) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range networkInformationAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *NetworkInformation) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range networkInformationBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *NetworkInformation) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range networkInformationAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddNetworkInformationHook registers your hook function for all future operations.
func AddNetworkInformationHook(hookPoint boil.HookPoint, networkInformationHook NetworkInformationHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		networkInformationAfterSelectHooks = append(networkInformationAfterSelectHooks, networkInformationHook)
	case boil.BeforeInsertHook:
		networkInformationBeforeInsertHooks = append(networkInformationBeforeInsertHooks, networkInformationHook)
	case boil.AfterInsertHook:
		networkInformationAfterInsertHooks = append(networkInformationAfterInsertHooks, networkInformationHook)
	case boil.BeforeUpdateHook:
		networkInformationBeforeUpdateHooks = append(networkInformationBeforeUpdateHooks, networkInformationHook)
	case boil.AfterUpdateHook:
		networkInformationAfterUpdateHooks = append(networkInformationAfterUpdateHooks, networkInformationHook)
	case boil.BeforeDeleteHook:
		networkInformationBeforeDeleteHooks = append(networkInformationBeforeDeleteHooks, networkInformationHook)
	case boil.AfterDeleteHook:
		networkInformationAfterDeleteHooks = append(networkInformationAfterDeleteHooks, networkInformationHook)
	case boil.BeforeUpsertHook:
		networkInformationBeforeUpsertHooks = append(networkInformationBeforeUpsertHooks, networkInformationHook)
	case boil.AfterUpsertHook:
		networkInformationAfterUpsertHooks = append(networkInformationAfterUpsertHooks, networkInformationHook)
	}
}

// One returns a single networkInformation record from the query.
func (q networkInformationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*NetworkInformation, error) {
	o := &NetworkInformation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for network_information")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all NetworkInformation records from the query.
func (q networkInformationQuery) All(ctx context.Context, exec boil.ContextExecutor) (NetworkInformationSlice, error) {
	var o []*NetworkInformation

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to NetworkInformation slice")
	}

	if len(networkInformationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all NetworkInformation records in the query.
func (q networkInformationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count network_information rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q networkInformationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if network_information exists")
	}

	return count > 0, nil
}

// Peer pointed to by the foreign key.
func (o *NetworkInformation) Peer(mods ...qm.QueryMod) peerQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.PeerID),
	}

	queryMods = append(queryMods, mods...)

	return Peers(queryMods...)
}

// LoadPeer allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (networkInformationL) LoadPeer(ctx context.Context, e boil.ContextExecutor, singular bool, maybeNetworkInformation interface{}, mods queries.Applicator) error {
	var slice []*NetworkInformation
	var object *NetworkInformation

	if singular {
		var ok bool
		object, ok = maybeNetworkInformation.(*NetworkInformation)
		if !ok {
			object = new(NetworkInformation)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeNetworkInformation)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeNetworkInformation))
			}
		}
	} else {
		s, ok := maybeNetworkInformation.(*[]*NetworkInformation)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeNetworkInformation)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeNetworkInformation))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &networkInformationR{}
		}
		args = append(args, object.PeerID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &networkInformationR{}
			}

			for _, a := range args {
				if a == obj.PeerID {
					continue Outer
				}
			}

			args = append(args, obj.PeerID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`peers`),
		qm.WhereIn(`peers.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Peer")
	}

	var resultSlice []*Peer
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Peer")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for peers")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for peers")
	}

	if len(networkInformationAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Peer = foreign
		if foreign.R == nil {
			foreign.R = &peerR{}
		}
		foreign.R.NetworkInformations = append(foreign.R.NetworkInformations, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PeerID == foreign.ID {
				local.R.Peer = foreign
				if foreign.R == nil {
					foreign.R = &peerR{}
				}
				foreign.R.NetworkInformations = append(foreign.R.NetworkInformations, local)
				break
			}
		}
	}

	return nil
}

// SetPeer of the networkInformation to the related item.
// Sets o.R.Peer to related.
// Adds o to related.R.NetworkInformations.
func (o *NetworkInformation) SetPeer(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Peer) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"network_information\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"peer_id"}),
		strmangle.WhereClause("\"", "\"", 2, networkInformationPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PeerID = related.ID
	if o.R == nil {
		o.R = &networkInformationR{
			Peer: related,
		}
	} else {
		o.R.Peer = related
	}

	if related.R == nil {
		related.R = &peerR{
			NetworkInformations: NetworkInformationSlice{o},
		}
	} else {
		related.R.NetworkInformations = append(related.R.NetworkInformations, o)
	}

	return nil
}

// NetworkInformations retrieves all the records using an executor.
func NetworkInformations(mods ...qm.QueryMod) networkInformationQuery {
	mods = append(mods, qm.From("\"network_information\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"network_information\".*"})
	}

	return networkInformationQuery{q}
}

// FindNetworkInformation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindNetworkInformation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*NetworkInformation, error) {
	networkInformationObj := &NetworkInformation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"network_information\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, networkInformationObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from network_information")
	}

	if err = networkInformationObj.doAfterSelectHooks(ctx, exec); err != nil {
		return networkInformationObj, err
	}

	return networkInformationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *NetworkInformation) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no network_information provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(networkInformationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	networkInformationInsertCacheMut.RLock()
	cache, cached := networkInformationInsertCache[key]
	networkInformationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			networkInformationAllColumns,
			networkInformationColumnsWithDefault,
			networkInformationColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, networkInformationGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(networkInformationType, networkInformationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(networkInformationType, networkInformationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"network_information\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"network_information\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into network_information")
	}

	if !cached {
		networkInformationInsertCacheMut.Lock()
		networkInformationInsertCache[key] = cache
		networkInformationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the NetworkInformation.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *NetworkInformation) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	networkInformationUpdateCacheMut.RLock()
	cache, cached := networkInformationUpdateCache[key]
	networkInformationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			networkInformationAllColumns,
			networkInformationPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, networkInformationGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update network_information, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"network_information\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, networkInformationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(networkInformationType, networkInformationMapping, append(wl, networkInformationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update network_information row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for network_information")
	}

	if !cached {
		networkInformationUpdateCacheMut.Lock()
		networkInformationUpdateCache[key] = cache
		networkInformationUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q networkInformationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for network_information")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for network_information")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o NetworkInformationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), networkInformationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"network_information\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, networkInformationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in networkInformation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all networkInformation")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *NetworkInformation) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no network_information provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(networkInformationColumnsWithDefault, o)

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

	networkInformationUpsertCacheMut.RLock()
	cache, cached := networkInformationUpsertCache[key]
	networkInformationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			networkInformationAllColumns,
			networkInformationColumnsWithDefault,
			networkInformationColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			networkInformationAllColumns,
			networkInformationPrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, networkInformationGeneratedColumns)
		update = strmangle.SetComplement(update, networkInformationGeneratedColumns)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert network_information, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(networkInformationPrimaryKeyColumns))
			copy(conflict, networkInformationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"network_information\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(networkInformationType, networkInformationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(networkInformationType, networkInformationMapping, ret)
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
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert network_information")
	}

	if !cached {
		networkInformationUpsertCacheMut.Lock()
		networkInformationUpsertCache[key] = cache
		networkInformationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single NetworkInformation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *NetworkInformation) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no NetworkInformation provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), networkInformationPrimaryKeyMapping)
	sql := "DELETE FROM \"network_information\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from network_information")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for network_information")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q networkInformationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no networkInformationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from network_information")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for network_information")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o NetworkInformationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(networkInformationBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), networkInformationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"network_information\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, networkInformationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from networkInformation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for network_information")
	}

	if len(networkInformationAfterDeleteHooks) != 0 {
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
func (o *NetworkInformation) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindNetworkInformation(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *NetworkInformationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := NetworkInformationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), networkInformationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"network_information\".* FROM \"network_information\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, networkInformationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in NetworkInformationSlice")
	}

	*o = slice

	return nil
}

// NetworkInformationExists checks if the NetworkInformation row exists.
func NetworkInformationExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"network_information\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if network_information exists")
	}

	return exists, nil
}
