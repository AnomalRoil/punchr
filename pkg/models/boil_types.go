// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"strconv"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/strmangle"
)

// M type is for providing columns and column values to UpdateAll.
type M map[string]interface{}

// ErrSyncFail occurs during insert when the record could not be retrieved in
// order to populate default value information. This usually happens when LastInsertId
// fails or there was a primary key configuration that was not resolvable.
var ErrSyncFail = errors.New("models: failed to synchronize data after insert")

type insertCache struct {
	query        string
	retQuery     string
	valueMapping []uint64
	retMapping   []uint64
}

type updateCache struct {
	query        string
	valueMapping []uint64
}

func makeCacheKey(cols boil.Columns, nzDefaults []string) string {
	buf := strmangle.GetBuffer()

	buf.WriteString(strconv.Itoa(cols.Kind))
	for _, w := range cols.Cols {
		buf.WriteString(w)
	}

	if len(nzDefaults) != 0 {
		buf.WriteByte('.')
	}
	for _, nz := range nzDefaults {
		buf.WriteString(nz)
	}

	str := buf.String()
	strmangle.PutBuffer(buf)
	return str
}

// Enum values for HolePunchAttemptOutcome
const (
	HolePunchAttemptOutcomeUNKNOWN        string = "UNKNOWN"
	HolePunchAttemptOutcomeDIRECT_DIAL    string = "DIRECT_DIAL"
	HolePunchAttemptOutcomePROTOCOL_ERROR string = "PROTOCOL_ERROR"
	HolePunchAttemptOutcomeCANCELLED      string = "CANCELLED"
	HolePunchAttemptOutcomeTIMEOUT        string = "TIMEOUT"
	HolePunchAttemptOutcomeFAILED         string = "FAILED"
	HolePunchAttemptOutcomeSUCCESS        string = "SUCCESS"
)

func AllHolePunchAttemptOutcome() []string {
	return []string{
		HolePunchAttemptOutcomeUNKNOWN,
		HolePunchAttemptOutcomeDIRECT_DIAL,
		HolePunchAttemptOutcomePROTOCOL_ERROR,
		HolePunchAttemptOutcomeCANCELLED,
		HolePunchAttemptOutcomeTIMEOUT,
		HolePunchAttemptOutcomeFAILED,
		HolePunchAttemptOutcomeSUCCESS,
	}
}

// Enum values for HolePunchOutcome
const (
	HolePunchOutcomeUNKNOWN             string = "UNKNOWN"
	HolePunchOutcomeNO_CONNECTION       string = "NO_CONNECTION"
	HolePunchOutcomeNO_STREAM           string = "NO_STREAM"
	HolePunchOutcomeCONNECTION_REVERSED string = "CONNECTION_REVERSED"
	HolePunchOutcomeCANCELLED           string = "CANCELLED"
	HolePunchOutcomeFAILED              string = "FAILED"
	HolePunchOutcomeSUCCESS             string = "SUCCESS"
)

func AllHolePunchOutcome() []string {
	return []string{
		HolePunchOutcomeUNKNOWN,
		HolePunchOutcomeNO_CONNECTION,
		HolePunchOutcomeNO_STREAM,
		HolePunchOutcomeCONNECTION_REVERSED,
		HolePunchOutcomeCANCELLED,
		HolePunchOutcomeFAILED,
		HolePunchOutcomeSUCCESS,
	}
}

// Enum values for HolePunchMultiAddressRelationship
const (
	HolePunchMultiAddressRelationshipINITIAL string = "INITIAL"
	HolePunchMultiAddressRelationshipFINAL   string = "FINAL"
)

func AllHolePunchMultiAddressRelationship() []string {
	return []string{
		HolePunchMultiAddressRelationshipINITIAL,
		HolePunchMultiAddressRelationshipFINAL,
	}
}

// Enum values for LatencyMeasurementType
const (
	LatencyMeasurementTypeTO_RELAY                  string = "TO_RELAY"
	LatencyMeasurementTypeTO_REMOTE_THROUGH_RELAY   string = "TO_REMOTE_THROUGH_RELAY"
	LatencyMeasurementTypeTO_REMOTE_AFTER_HOLEPUNCH string = "TO_REMOTE_AFTER_HOLEPUNCH"
)

func AllLatencyMeasurementType() []string {
	return []string{
		LatencyMeasurementTypeTO_RELAY,
		LatencyMeasurementTypeTO_REMOTE_THROUGH_RELAY,
		LatencyMeasurementTypeTO_REMOTE_AFTER_HOLEPUNCH,
	}
}
