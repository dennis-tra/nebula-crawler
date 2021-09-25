// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

func TestUpsert(t *testing.T) {
	t.Run("CrawlProperties", testCrawlPropertiesUpsert)

	t.Run("Crawls", testCrawlsUpsert)

	t.Run("IPAddresses", testIPAddressesUpsert)

	t.Run("Latencies", testLatenciesUpsert)

	t.Run("MultiAddresses", testMultiAddressesUpsert)

	t.Run("MultiAddressesXIPAddresses", testMultiAddressesXIPAddressesUpsert)

	t.Run("Neighbors", testNeighborsUpsert)

	t.Run("Peers", testPeersUpsert)

	t.Run("PegasysConnections", testPegasysConnectionsUpsert)

	t.Run("PegasysNeighbours", testPegasysNeighboursUpsert)

	t.Run("Properties", testPropertiesUpsert)

	t.Run("RawVisits", testRawVisitsUpsert)

	t.Run("Sessions", testSessionsUpsert)

	t.Run("Visits", testVisitsUpsert)
}
