package cartesian

import (
	"testing"
)

func TestTreeSimple(t *testing.T) {
	tree := NewTree()
	if tree.Size() != 0 {
		t.Fatal()
	}
	if !tree.Insert("kek", "mem") {
		t.Fatal()
	}
	if *tree.Get("kek") != "mem" {
		t.Fatal()
	}
	if tree.Insert("kek", "mem") {
		t.Fatal()
	}
	if tree.Size() != 1 {
		t.Fatal()
	}

	keys := []string{"puk", "krya", "meow", "z1", "z2", "z15", "y230"}
	vals := []string{"kek", "mem", "lel", "kek", "kek", "kek", "zupa"}
	if len(keys) != len(vals) {
		t.Fatal()
	}

	for i := 0; i < len(keys); i++ {
		if !tree.Insert(keys[i], vals[i]) {
			t.Fatal()
		}
		if *tree.Get(keys[i]) != vals[i] {
			t.Fatal()
		}
		if tree.Size() != i+2 {
			t.Fatal()
		}
	}

	for i := 0; i < len(keys); i++ {
		if tree.Insert(keys[i], vals[i]) {
			t.Fatal()
		}
		if *tree.Get(keys[i]) != vals[i] {
			t.Fatal()
		}
		if tree.Size() != len(vals)+1 {
			t.Fatal()
		}
	}
}
