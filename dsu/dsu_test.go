package dsu

import "testing"

func TestDsuSimple(t *testing.T) {
	// initial
	dsu := NewDsu(50)
	if dsu.GetNodesNum() != 50 {
		t.Fatal()
	}
	if dsu.GetSetsNum() != 50 {
		t.Fatal()
	}
	if dsu.IsSameSet(10, 20) {
		t.Fatal()
	}

	// unite two nodes
	if !dsu.Unite(10, 20) {
		t.Fatal()
	}

	// after unite
	if dsu.GetNodesNum() != 50 {
		t.Fatal()
	}
	if dsu.GetSetsNum() != 49 {
		t.Fatal()
	}
	if !dsu.IsSameSet(10, 20) {
		t.Fatal()
	}
}
