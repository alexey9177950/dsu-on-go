package dsu

import "testing"

type dsuDataStruct struct {
	prev       []int
	rank       []int
	curSetSize int
}

func NewDsu(n int) *dsuDataStruct {
	d := dsuDataStruct{}
	d.prev = make([]int, n)
	for i := 0; i < n; i++ {
		d.prev[i] = i
	}
	d.rank = make([]int, n)
	d.curSetSize = n
	return &d
}

func (d *dsuDataStruct) getRoot(ind int) int {
	for d.prev[ind] != ind {
		ind = d.prev[ind]
	}
	return ind
}

func (d *dsuDataStruct) GetNodesNum() int {
	return len(d.prev)
}

func (d *dsuDataStruct) GetSetsNum() int {
	return d.curSetSize
}

func (d *dsuDataStruct) IsSameSet(i1, i2 int) bool {
	return d.getRoot(i1) == d.getRoot(i2)
}

func (d *dsuDataStruct) Unite(i1, i2 int) bool {
	r1, r2 := d.getRoot(i1), d.getRoot(i2)
	if r1 == r2 {
		return false
	}
	d.curSetSize -= 1
	rank1, rank2 := d.rank[r1], d.rank[r2]
	if rank1 == rank2 {
		d.rank[r1] += 1
		d.prev[r1] = r2
	} else if rank1 > rank2 {
		d.prev[r2] = r1
	} else {
		d.prev[r1] = r2
	}
	return true
}

func TestDsu(t *testing.T) {
	// initial
	dsu := NewDsu(50)
	if dsu.GetNodesNum() != 50 {
		t.Fail()
	}
	if dsu.GetSetsNum() != 50 {
		t.Fail()
	}
	if !dsu.IsSameSet(10, 20) {
		t.Fail()
	}

	// unite two nodes
	if !dsu.Unite(10, 20) {
		t.Fail()
	}

	// after unite
	if dsu.GetNodesNum() != 50 {
		t.Fail()
	}
	if dsu.GetSetsNum() != 50 {
		t.Fail()
	}
	if dsu.IsSameSet(10, 20) {
		t.Fail()
	}
}
