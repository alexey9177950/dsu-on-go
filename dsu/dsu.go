package dsu

type DsuDataStruct struct {
	prev       []int
	rank       []int
	curSetSize int
}

func NewDsu(n int) *DsuDataStruct {
	d := DsuDataStruct{}
	d.prev = make([]int, n)
	for i := 0; i < n; i++ {
		d.prev[i] = i
	}
	d.rank = make([]int, n)
	d.curSetSize = n
	return &d
}

func (d *DsuDataStruct) getRoot(ind int) int {
	for d.prev[ind] != ind {
		ind = d.prev[ind]
	}
	return ind
}

func (d *DsuDataStruct) GetNodesNum() int {
	return len(d.prev)
}

func (d *DsuDataStruct) GetSetsNum() int {
	return d.curSetSize
}

func (d *DsuDataStruct) IsSameSet(i1, i2 int) bool {
	return d.getRoot(i1) == d.getRoot(i2)
}

func (d *DsuDataStruct) Unite(i1, i2 int) bool {
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
