package cqzone

import "testing"

func TestZoneLength(t *testing.T) {
	if len(zones) != 40 {
		t.Errorf("Expected 40 zones, got %d", len(zones))
	}
}

func TestIndices(t *testing.T) {
	for i := range zones {
		if zones[i].number != i+1 {
			t.Errorf("Zone %d has number %d, expected %d", i, zones[i].number, i+1)
		}
	}
}
