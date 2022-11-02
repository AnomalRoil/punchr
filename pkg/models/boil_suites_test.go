// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Authorizations", testAuthorizations)
	t.Run("Clients", testClients)
	t.Run("ConnectionEvents", testConnectionEvents)
	t.Run("HolePunchAttempts", testHolePunchAttempts)
	t.Run("HolePunchResults", testHolePunchResults)
	t.Run("HolePunchResultsXMultiAddresses", testHolePunchResultsXMultiAddresses)
	t.Run("IPAddresses", testIPAddresses)
	t.Run("LatencyMeasurements", testLatencyMeasurements)
	t.Run("MultiAddresses", testMultiAddresses)
	t.Run("MultiAddressesSets", testMultiAddressesSets)
	t.Run("NetworkInformations", testNetworkInformations)
	t.Run("PeerLogs", testPeerLogs)
	t.Run("Peers", testPeers)
}

func TestDelete(t *testing.T) {
	t.Run("Authorizations", testAuthorizationsDelete)
	t.Run("Clients", testClientsDelete)
	t.Run("ConnectionEvents", testConnectionEventsDelete)
	t.Run("HolePunchAttempts", testHolePunchAttemptsDelete)
	t.Run("HolePunchResults", testHolePunchResultsDelete)
	t.Run("HolePunchResultsXMultiAddresses", testHolePunchResultsXMultiAddressesDelete)
	t.Run("IPAddresses", testIPAddressesDelete)
	t.Run("LatencyMeasurements", testLatencyMeasurementsDelete)
	t.Run("MultiAddresses", testMultiAddressesDelete)
	t.Run("MultiAddressesSets", testMultiAddressesSetsDelete)
	t.Run("NetworkInformations", testNetworkInformationsDelete)
	t.Run("PeerLogs", testPeerLogsDelete)
	t.Run("Peers", testPeersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Authorizations", testAuthorizationsQueryDeleteAll)
	t.Run("Clients", testClientsQueryDeleteAll)
	t.Run("ConnectionEvents", testConnectionEventsQueryDeleteAll)
	t.Run("HolePunchAttempts", testHolePunchAttemptsQueryDeleteAll)
	t.Run("HolePunchResults", testHolePunchResultsQueryDeleteAll)
	t.Run("HolePunchResultsXMultiAddresses", testHolePunchResultsXMultiAddressesQueryDeleteAll)
	t.Run("IPAddresses", testIPAddressesQueryDeleteAll)
	t.Run("LatencyMeasurements", testLatencyMeasurementsQueryDeleteAll)
	t.Run("MultiAddresses", testMultiAddressesQueryDeleteAll)
	t.Run("MultiAddressesSets", testMultiAddressesSetsQueryDeleteAll)
	t.Run("NetworkInformations", testNetworkInformationsQueryDeleteAll)
	t.Run("PeerLogs", testPeerLogsQueryDeleteAll)
	t.Run("Peers", testPeersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Authorizations", testAuthorizationsSliceDeleteAll)
	t.Run("Clients", testClientsSliceDeleteAll)
	t.Run("ConnectionEvents", testConnectionEventsSliceDeleteAll)
	t.Run("HolePunchAttempts", testHolePunchAttemptsSliceDeleteAll)
	t.Run("HolePunchResults", testHolePunchResultsSliceDeleteAll)
	t.Run("HolePunchResultsXMultiAddresses", testHolePunchResultsXMultiAddressesSliceDeleteAll)
	t.Run("IPAddresses", testIPAddressesSliceDeleteAll)
	t.Run("LatencyMeasurements", testLatencyMeasurementsSliceDeleteAll)
	t.Run("MultiAddresses", testMultiAddressesSliceDeleteAll)
	t.Run("MultiAddressesSets", testMultiAddressesSetsSliceDeleteAll)
	t.Run("NetworkInformations", testNetworkInformationsSliceDeleteAll)
	t.Run("PeerLogs", testPeerLogsSliceDeleteAll)
	t.Run("Peers", testPeersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Authorizations", testAuthorizationsExists)
	t.Run("Clients", testClientsExists)
	t.Run("ConnectionEvents", testConnectionEventsExists)
	t.Run("HolePunchAttempts", testHolePunchAttemptsExists)
	t.Run("HolePunchResults", testHolePunchResultsExists)
	t.Run("HolePunchResultsXMultiAddresses", testHolePunchResultsXMultiAddressesExists)
	t.Run("IPAddresses", testIPAddressesExists)
	t.Run("LatencyMeasurements", testLatencyMeasurementsExists)
	t.Run("MultiAddresses", testMultiAddressesExists)
	t.Run("MultiAddressesSets", testMultiAddressesSetsExists)
	t.Run("NetworkInformations", testNetworkInformationsExists)
	t.Run("PeerLogs", testPeerLogsExists)
	t.Run("Peers", testPeersExists)
}

func TestFind(t *testing.T) {
	t.Run("Authorizations", testAuthorizationsFind)
	t.Run("Clients", testClientsFind)
	t.Run("ConnectionEvents", testConnectionEventsFind)
	t.Run("HolePunchAttempts", testHolePunchAttemptsFind)
	t.Run("HolePunchResults", testHolePunchResultsFind)
	t.Run("HolePunchResultsXMultiAddresses", testHolePunchResultsXMultiAddressesFind)
	t.Run("IPAddresses", testIPAddressesFind)
	t.Run("LatencyMeasurements", testLatencyMeasurementsFind)
	t.Run("MultiAddresses", testMultiAddressesFind)
	t.Run("MultiAddressesSets", testMultiAddressesSetsFind)
	t.Run("NetworkInformations", testNetworkInformationsFind)
	t.Run("PeerLogs", testPeerLogsFind)
	t.Run("Peers", testPeersFind)
}

func TestBind(t *testing.T) {
	t.Run("Authorizations", testAuthorizationsBind)
	t.Run("Clients", testClientsBind)
	t.Run("ConnectionEvents", testConnectionEventsBind)
	t.Run("HolePunchAttempts", testHolePunchAttemptsBind)
	t.Run("HolePunchResults", testHolePunchResultsBind)
	t.Run("HolePunchResultsXMultiAddresses", testHolePunchResultsXMultiAddressesBind)
	t.Run("IPAddresses", testIPAddressesBind)
	t.Run("LatencyMeasurements", testLatencyMeasurementsBind)
	t.Run("MultiAddresses", testMultiAddressesBind)
	t.Run("MultiAddressesSets", testMultiAddressesSetsBind)
	t.Run("NetworkInformations", testNetworkInformationsBind)
	t.Run("PeerLogs", testPeerLogsBind)
	t.Run("Peers", testPeersBind)
}

func TestOne(t *testing.T) {
	t.Run("Authorizations", testAuthorizationsOne)
	t.Run("Clients", testClientsOne)
	t.Run("ConnectionEvents", testConnectionEventsOne)
	t.Run("HolePunchAttempts", testHolePunchAttemptsOne)
	t.Run("HolePunchResults", testHolePunchResultsOne)
	t.Run("HolePunchResultsXMultiAddresses", testHolePunchResultsXMultiAddressesOne)
	t.Run("IPAddresses", testIPAddressesOne)
	t.Run("LatencyMeasurements", testLatencyMeasurementsOne)
	t.Run("MultiAddresses", testMultiAddressesOne)
	t.Run("MultiAddressesSets", testMultiAddressesSetsOne)
	t.Run("NetworkInformations", testNetworkInformationsOne)
	t.Run("PeerLogs", testPeerLogsOne)
	t.Run("Peers", testPeersOne)
}

func TestAll(t *testing.T) {
	t.Run("Authorizations", testAuthorizationsAll)
	t.Run("Clients", testClientsAll)
	t.Run("ConnectionEvents", testConnectionEventsAll)
	t.Run("HolePunchAttempts", testHolePunchAttemptsAll)
	t.Run("HolePunchResults", testHolePunchResultsAll)
	t.Run("HolePunchResultsXMultiAddresses", testHolePunchResultsXMultiAddressesAll)
	t.Run("IPAddresses", testIPAddressesAll)
	t.Run("LatencyMeasurements", testLatencyMeasurementsAll)
	t.Run("MultiAddresses", testMultiAddressesAll)
	t.Run("MultiAddressesSets", testMultiAddressesSetsAll)
	t.Run("NetworkInformations", testNetworkInformationsAll)
	t.Run("PeerLogs", testPeerLogsAll)
	t.Run("Peers", testPeersAll)
}

func TestCount(t *testing.T) {
	t.Run("Authorizations", testAuthorizationsCount)
	t.Run("Clients", testClientsCount)
	t.Run("ConnectionEvents", testConnectionEventsCount)
	t.Run("HolePunchAttempts", testHolePunchAttemptsCount)
	t.Run("HolePunchResults", testHolePunchResultsCount)
	t.Run("HolePunchResultsXMultiAddresses", testHolePunchResultsXMultiAddressesCount)
	t.Run("IPAddresses", testIPAddressesCount)
	t.Run("LatencyMeasurements", testLatencyMeasurementsCount)
	t.Run("MultiAddresses", testMultiAddressesCount)
	t.Run("MultiAddressesSets", testMultiAddressesSetsCount)
	t.Run("NetworkInformations", testNetworkInformationsCount)
	t.Run("PeerLogs", testPeerLogsCount)
	t.Run("Peers", testPeersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Authorizations", testAuthorizationsHooks)
	t.Run("Clients", testClientsHooks)
	t.Run("ConnectionEvents", testConnectionEventsHooks)
	t.Run("HolePunchAttempts", testHolePunchAttemptsHooks)
	t.Run("HolePunchResults", testHolePunchResultsHooks)
	t.Run("HolePunchResultsXMultiAddresses", testHolePunchResultsXMultiAddressesHooks)
	t.Run("IPAddresses", testIPAddressesHooks)
	t.Run("LatencyMeasurements", testLatencyMeasurementsHooks)
	t.Run("MultiAddresses", testMultiAddressesHooks)
	t.Run("MultiAddressesSets", testMultiAddressesSetsHooks)
	t.Run("NetworkInformations", testNetworkInformationsHooks)
	t.Run("PeerLogs", testPeerLogsHooks)
	t.Run("Peers", testPeersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Authorizations", testAuthorizationsInsert)
	t.Run("Authorizations", testAuthorizationsInsertWhitelist)
	t.Run("Clients", testClientsInsert)
	t.Run("Clients", testClientsInsertWhitelist)
	t.Run("ConnectionEvents", testConnectionEventsInsert)
	t.Run("ConnectionEvents", testConnectionEventsInsertWhitelist)
	t.Run("HolePunchAttempts", testHolePunchAttemptsInsert)
	t.Run("HolePunchAttempts", testHolePunchAttemptsInsertWhitelist)
	t.Run("HolePunchResults", testHolePunchResultsInsert)
	t.Run("HolePunchResults", testHolePunchResultsInsertWhitelist)
	t.Run("HolePunchResultsXMultiAddresses", testHolePunchResultsXMultiAddressesInsert)
	t.Run("HolePunchResultsXMultiAddresses", testHolePunchResultsXMultiAddressesInsertWhitelist)
	t.Run("IPAddresses", testIPAddressesInsert)
	t.Run("IPAddresses", testIPAddressesInsertWhitelist)
	t.Run("LatencyMeasurements", testLatencyMeasurementsInsert)
	t.Run("LatencyMeasurements", testLatencyMeasurementsInsertWhitelist)
	t.Run("MultiAddresses", testMultiAddressesInsert)
	t.Run("MultiAddresses", testMultiAddressesInsertWhitelist)
	t.Run("MultiAddressesSets", testMultiAddressesSetsInsert)
	t.Run("MultiAddressesSets", testMultiAddressesSetsInsertWhitelist)
	t.Run("NetworkInformations", testNetworkInformationsInsert)
	t.Run("NetworkInformations", testNetworkInformationsInsertWhitelist)
	t.Run("PeerLogs", testPeerLogsInsert)
	t.Run("PeerLogs", testPeerLogsInsertWhitelist)
	t.Run("Peers", testPeersInsert)
	t.Run("Peers", testPeersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("ClientToAuthorizationUsingAuthorization", testClientToOneAuthorizationUsingAuthorization)
	t.Run("ClientToPeerUsingPeer", testClientToOnePeerUsingPeer)
	t.Run("ConnectionEventToPeerUsingLocal", testConnectionEventToOnePeerUsingLocal)
	t.Run("ConnectionEventToMultiAddressUsingConnMultiAddress", testConnectionEventToOneMultiAddressUsingConnMultiAddress)
	t.Run("ConnectionEventToPeerUsingRemote", testConnectionEventToOnePeerUsingRemote)
	t.Run("HolePunchAttemptToHolePunchResultUsingHolePunchResult", testHolePunchAttemptToOneHolePunchResultUsingHolePunchResult)
	t.Run("HolePunchResultToMultiAddressesSetUsingListenMultiAddressesSet", testHolePunchResultToOneMultiAddressesSetUsingListenMultiAddressesSet)
	t.Run("HolePunchResultToPeerUsingLocal", testHolePunchResultToOnePeerUsingLocal)
	t.Run("HolePunchResultToPeerUsingRemote", testHolePunchResultToOnePeerUsingRemote)
	t.Run("HolePunchResultsXMultiAddressToHolePunchResultUsingHolePunchResult", testHolePunchResultsXMultiAddressToOneHolePunchResultUsingHolePunchResult)
	t.Run("HolePunchResultsXMultiAddressToMultiAddressUsingMultiAddress", testHolePunchResultsXMultiAddressToOneMultiAddressUsingMultiAddress)
	t.Run("IPAddressToMultiAddressUsingMultiAddress", testIPAddressToOneMultiAddressUsingMultiAddress)
	t.Run("LatencyMeasurementToHolePunchResultUsingHolePunchResult", testLatencyMeasurementToOneHolePunchResultUsingHolePunchResult)
	t.Run("LatencyMeasurementToMultiAddressUsingMultiAddress", testLatencyMeasurementToOneMultiAddressUsingMultiAddress)
	t.Run("LatencyMeasurementToPeerUsingRemote", testLatencyMeasurementToOnePeerUsingRemote)
	t.Run("NetworkInformationToPeerUsingPeer", testNetworkInformationToOnePeerUsingPeer)
	t.Run("PeerLogToPeerUsingPeer", testPeerLogToOnePeerUsingPeer)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("AuthorizationToClients", testAuthorizationToManyClients)
	t.Run("ConnectionEventToMultiAddresses", testConnectionEventToManyMultiAddresses)
	t.Run("HolePunchAttemptToMultiAddresses", testHolePunchAttemptToManyMultiAddresses)
	t.Run("HolePunchResultToHolePunchAttempts", testHolePunchResultToManyHolePunchAttempts)
	t.Run("HolePunchResultToHolePunchResultsXMultiAddresses", testHolePunchResultToManyHolePunchResultsXMultiAddresses)
	t.Run("HolePunchResultToLatencyMeasurements", testHolePunchResultToManyLatencyMeasurements)
	t.Run("MultiAddressToConnMultiAddressConnectionEvents", testMultiAddressToManyConnMultiAddressConnectionEvents)
	t.Run("MultiAddressToConnectionEvents", testMultiAddressToManyConnectionEvents)
	t.Run("MultiAddressToHolePunchAttempts", testMultiAddressToManyHolePunchAttempts)
	t.Run("MultiAddressToHolePunchResultsXMultiAddresses", testMultiAddressToManyHolePunchResultsXMultiAddresses)
	t.Run("MultiAddressToIPAddresses", testMultiAddressToManyIPAddresses)
	t.Run("MultiAddressToLatencyMeasurements", testMultiAddressToManyLatencyMeasurements)
	t.Run("MultiAddressesSetToListenMultiAddressesSetHolePunchResults", testMultiAddressesSetToManyListenMultiAddressesSetHolePunchResults)
	t.Run("PeerToClients", testPeerToManyClients)
	t.Run("PeerToLocalConnectionEvents", testPeerToManyLocalConnectionEvents)
	t.Run("PeerToRemoteConnectionEvents", testPeerToManyRemoteConnectionEvents)
	t.Run("PeerToLocalHolePunchResults", testPeerToManyLocalHolePunchResults)
	t.Run("PeerToRemoteHolePunchResults", testPeerToManyRemoteHolePunchResults)
	t.Run("PeerToRemoteLatencyMeasurements", testPeerToManyRemoteLatencyMeasurements)
	t.Run("PeerToNetworkInformations", testPeerToManyNetworkInformations)
	t.Run("PeerToPeerLogs", testPeerToManyPeerLogs)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("ClientToAuthorizationUsingClients", testClientToOneSetOpAuthorizationUsingAuthorization)
	t.Run("ClientToPeerUsingClients", testClientToOneSetOpPeerUsingPeer)
	t.Run("ConnectionEventToPeerUsingLocalConnectionEvents", testConnectionEventToOneSetOpPeerUsingLocal)
	t.Run("ConnectionEventToMultiAddressUsingConnMultiAddressConnectionEvents", testConnectionEventToOneSetOpMultiAddressUsingConnMultiAddress)
	t.Run("ConnectionEventToPeerUsingRemoteConnectionEvents", testConnectionEventToOneSetOpPeerUsingRemote)
	t.Run("HolePunchAttemptToHolePunchResultUsingHolePunchAttempts", testHolePunchAttemptToOneSetOpHolePunchResultUsingHolePunchResult)
	t.Run("HolePunchResultToMultiAddressesSetUsingListenMultiAddressesSetHolePunchResults", testHolePunchResultToOneSetOpMultiAddressesSetUsingListenMultiAddressesSet)
	t.Run("HolePunchResultToPeerUsingLocalHolePunchResults", testHolePunchResultToOneSetOpPeerUsingLocal)
	t.Run("HolePunchResultToPeerUsingRemoteHolePunchResults", testHolePunchResultToOneSetOpPeerUsingRemote)
	t.Run("HolePunchResultsXMultiAddressToHolePunchResultUsingHolePunchResultsXMultiAddresses", testHolePunchResultsXMultiAddressToOneSetOpHolePunchResultUsingHolePunchResult)
	t.Run("HolePunchResultsXMultiAddressToMultiAddressUsingHolePunchResultsXMultiAddresses", testHolePunchResultsXMultiAddressToOneSetOpMultiAddressUsingMultiAddress)
	t.Run("IPAddressToMultiAddressUsingIPAddresses", testIPAddressToOneSetOpMultiAddressUsingMultiAddress)
	t.Run("LatencyMeasurementToHolePunchResultUsingLatencyMeasurements", testLatencyMeasurementToOneSetOpHolePunchResultUsingHolePunchResult)
	t.Run("LatencyMeasurementToMultiAddressUsingLatencyMeasurements", testLatencyMeasurementToOneSetOpMultiAddressUsingMultiAddress)
	t.Run("LatencyMeasurementToPeerUsingRemoteLatencyMeasurements", testLatencyMeasurementToOneSetOpPeerUsingRemote)
	t.Run("NetworkInformationToPeerUsingNetworkInformations", testNetworkInformationToOneSetOpPeerUsingPeer)
	t.Run("PeerLogToPeerUsingPeerLogs", testPeerLogToOneSetOpPeerUsingPeer)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("AuthorizationToClients", testAuthorizationToManyAddOpClients)
	t.Run("ConnectionEventToMultiAddresses", testConnectionEventToManyAddOpMultiAddresses)
	t.Run("HolePunchAttemptToMultiAddresses", testHolePunchAttemptToManyAddOpMultiAddresses)
	t.Run("HolePunchResultToHolePunchAttempts", testHolePunchResultToManyAddOpHolePunchAttempts)
	t.Run("HolePunchResultToHolePunchResultsXMultiAddresses", testHolePunchResultToManyAddOpHolePunchResultsXMultiAddresses)
	t.Run("HolePunchResultToLatencyMeasurements", testHolePunchResultToManyAddOpLatencyMeasurements)
	t.Run("MultiAddressToConnMultiAddressConnectionEvents", testMultiAddressToManyAddOpConnMultiAddressConnectionEvents)
	t.Run("MultiAddressToConnectionEvents", testMultiAddressToManyAddOpConnectionEvents)
	t.Run("MultiAddressToHolePunchAttempts", testMultiAddressToManyAddOpHolePunchAttempts)
	t.Run("MultiAddressToHolePunchResultsXMultiAddresses", testMultiAddressToManyAddOpHolePunchResultsXMultiAddresses)
	t.Run("MultiAddressToIPAddresses", testMultiAddressToManyAddOpIPAddresses)
	t.Run("MultiAddressToLatencyMeasurements", testMultiAddressToManyAddOpLatencyMeasurements)
	t.Run("MultiAddressesSetToListenMultiAddressesSetHolePunchResults", testMultiAddressesSetToManyAddOpListenMultiAddressesSetHolePunchResults)
	t.Run("PeerToClients", testPeerToManyAddOpClients)
	t.Run("PeerToLocalConnectionEvents", testPeerToManyAddOpLocalConnectionEvents)
	t.Run("PeerToRemoteConnectionEvents", testPeerToManyAddOpRemoteConnectionEvents)
	t.Run("PeerToLocalHolePunchResults", testPeerToManyAddOpLocalHolePunchResults)
	t.Run("PeerToRemoteHolePunchResults", testPeerToManyAddOpRemoteHolePunchResults)
	t.Run("PeerToRemoteLatencyMeasurements", testPeerToManyAddOpRemoteLatencyMeasurements)
	t.Run("PeerToNetworkInformations", testPeerToManyAddOpNetworkInformations)
	t.Run("PeerToPeerLogs", testPeerToManyAddOpPeerLogs)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("ConnectionEventToMultiAddresses", testConnectionEventToManySetOpMultiAddresses)
	t.Run("HolePunchAttemptToMultiAddresses", testHolePunchAttemptToManySetOpMultiAddresses)
	t.Run("MultiAddressToConnectionEvents", testMultiAddressToManySetOpConnectionEvents)
	t.Run("MultiAddressToHolePunchAttempts", testMultiAddressToManySetOpHolePunchAttempts)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("ConnectionEventToMultiAddresses", testConnectionEventToManyRemoveOpMultiAddresses)
	t.Run("HolePunchAttemptToMultiAddresses", testHolePunchAttemptToManyRemoveOpMultiAddresses)
	t.Run("MultiAddressToConnectionEvents", testMultiAddressToManyRemoveOpConnectionEvents)
	t.Run("MultiAddressToHolePunchAttempts", testMultiAddressToManyRemoveOpHolePunchAttempts)
}

func TestReload(t *testing.T) {
	t.Run("Authorizations", testAuthorizationsReload)
	t.Run("Clients", testClientsReload)
	t.Run("ConnectionEvents", testConnectionEventsReload)
	t.Run("HolePunchAttempts", testHolePunchAttemptsReload)
	t.Run("HolePunchResults", testHolePunchResultsReload)
	t.Run("HolePunchResultsXMultiAddresses", testHolePunchResultsXMultiAddressesReload)
	t.Run("IPAddresses", testIPAddressesReload)
	t.Run("LatencyMeasurements", testLatencyMeasurementsReload)
	t.Run("MultiAddresses", testMultiAddressesReload)
	t.Run("MultiAddressesSets", testMultiAddressesSetsReload)
	t.Run("NetworkInformations", testNetworkInformationsReload)
	t.Run("PeerLogs", testPeerLogsReload)
	t.Run("Peers", testPeersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Authorizations", testAuthorizationsReloadAll)
	t.Run("Clients", testClientsReloadAll)
	t.Run("ConnectionEvents", testConnectionEventsReloadAll)
	t.Run("HolePunchAttempts", testHolePunchAttemptsReloadAll)
	t.Run("HolePunchResults", testHolePunchResultsReloadAll)
	t.Run("HolePunchResultsXMultiAddresses", testHolePunchResultsXMultiAddressesReloadAll)
	t.Run("IPAddresses", testIPAddressesReloadAll)
	t.Run("LatencyMeasurements", testLatencyMeasurementsReloadAll)
	t.Run("MultiAddresses", testMultiAddressesReloadAll)
	t.Run("MultiAddressesSets", testMultiAddressesSetsReloadAll)
	t.Run("NetworkInformations", testNetworkInformationsReloadAll)
	t.Run("PeerLogs", testPeerLogsReloadAll)
	t.Run("Peers", testPeersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Authorizations", testAuthorizationsSelect)
	t.Run("Clients", testClientsSelect)
	t.Run("ConnectionEvents", testConnectionEventsSelect)
	t.Run("HolePunchAttempts", testHolePunchAttemptsSelect)
	t.Run("HolePunchResults", testHolePunchResultsSelect)
	t.Run("HolePunchResultsXMultiAddresses", testHolePunchResultsXMultiAddressesSelect)
	t.Run("IPAddresses", testIPAddressesSelect)
	t.Run("LatencyMeasurements", testLatencyMeasurementsSelect)
	t.Run("MultiAddresses", testMultiAddressesSelect)
	t.Run("MultiAddressesSets", testMultiAddressesSetsSelect)
	t.Run("NetworkInformations", testNetworkInformationsSelect)
	t.Run("PeerLogs", testPeerLogsSelect)
	t.Run("Peers", testPeersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Authorizations", testAuthorizationsUpdate)
	t.Run("Clients", testClientsUpdate)
	t.Run("ConnectionEvents", testConnectionEventsUpdate)
	t.Run("HolePunchAttempts", testHolePunchAttemptsUpdate)
	t.Run("HolePunchResults", testHolePunchResultsUpdate)
	t.Run("HolePunchResultsXMultiAddresses", testHolePunchResultsXMultiAddressesUpdate)
	t.Run("IPAddresses", testIPAddressesUpdate)
	t.Run("LatencyMeasurements", testLatencyMeasurementsUpdate)
	t.Run("MultiAddresses", testMultiAddressesUpdate)
	t.Run("MultiAddressesSets", testMultiAddressesSetsUpdate)
	t.Run("NetworkInformations", testNetworkInformationsUpdate)
	t.Run("PeerLogs", testPeerLogsUpdate)
	t.Run("Peers", testPeersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Authorizations", testAuthorizationsSliceUpdateAll)
	t.Run("Clients", testClientsSliceUpdateAll)
	t.Run("ConnectionEvents", testConnectionEventsSliceUpdateAll)
	t.Run("HolePunchAttempts", testHolePunchAttemptsSliceUpdateAll)
	t.Run("HolePunchResults", testHolePunchResultsSliceUpdateAll)
	t.Run("HolePunchResultsXMultiAddresses", testHolePunchResultsXMultiAddressesSliceUpdateAll)
	t.Run("IPAddresses", testIPAddressesSliceUpdateAll)
	t.Run("LatencyMeasurements", testLatencyMeasurementsSliceUpdateAll)
	t.Run("MultiAddresses", testMultiAddressesSliceUpdateAll)
	t.Run("MultiAddressesSets", testMultiAddressesSetsSliceUpdateAll)
	t.Run("NetworkInformations", testNetworkInformationsSliceUpdateAll)
	t.Run("PeerLogs", testPeerLogsSliceUpdateAll)
	t.Run("Peers", testPeersSliceUpdateAll)
}
