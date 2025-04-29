package cartesian

import (
	"hash/fnv"
)

type node struct {
	left  *node
	right *node
	key   string
	value string
}

func getYVal(ptr *node) uint32 {
	h := fnv.New32a()
	h.Write([]byte(ptr.key))
	return h.Sum32()
}

func merge(left *node, right *node) *node {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	if getYVal(left) < getYVal(right) {
		left.right = merge(left.right, right)
		return left
	} else {
		right.left = merge(left, right.left)
		return right
	}
}

func split(root *node, s string) (*node, *node) {
	if root == nil {
		return nil, nil
	}
	if s < root.key {
		leftSub, rightSub := split(root.left, s)
		root.left = rightSub
		return leftSub, root
	} else {
		leftSub, rightSub := split(root.right, s)
		root.right = leftSub
		return root, rightSub
	}
}

func search(root *node, key string) *string {
	if root == nil {
		return nil
	}
	if root.key == key {
		return &root.value
	} else if key < root.key {
		return search(root.left, key)
	} else {
		return search(root.right, key)
	}
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func max_height(root *node) int {
	if root == nil {
		return 0
	}
	return maxInt(max_height(root.left), max_height(root.right)) + 1
}

type Tree struct {
	root *node
	size int
}

func NewTree() Tree {
	return Tree{root: nil, size: 0}
}

func (t *Tree) Size() int {
	return t.size
}

func (t *Tree) Get(key string) *string {
	return search(t.root, key)
}

func (t *Tree) WildCard(pref string) []string {
	panic("unimplemented")
}

func (t *Tree) Insert(key string, value string) bool {
	// for simplicity. TODO: fix
	if t.Get(key) != nil {
		return false
	}

	t.size += 1
	midSub := &node{left: nil, right: nil, key: key, value: value}
	leftSub, rightSub := split(t.root, key)
	t.root = merge(leftSub, midSub)
	t.root = merge(t.root, rightSub)
	return true
}
