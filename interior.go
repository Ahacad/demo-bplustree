package bplustree

import (
	"sort"
)

type kc struct {
	key   int
	child node
}

type kcs [MAXKC + 1]kc

func (a *kcs) Len() int { return len(a) }

func (a *kcs) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a *kcs) Less(i, j int) bool {
	if a[i].key == 0 {
		return false
	}
	if a[j].key == 0 {
		return true
	}
	return a[i].key < a[j].key
}

type interiorNode struct {
	kcs   kcs
	count int
	p     *interiorNode
}

func newInteriorNode(p *interiorNode, largestChild node) *interiorNode {
	i := &interiorNode{
		p:     p,
		count: 1,
	}

	if largestChild != nil {
		i.kcs[0].child = largestChild
	}
	return i
}

func (in *interiorNode) find(key int) (int, bool) {
	c := func(i int) bool { return in.kcs[i].key > key }

	i := sort.Search(in.count-1, c)

	return i, true
}

func (in *interiorNode) full() bool { return in.count == MAXKC }

func (in *interiorNode) parent() *interiorNode { return in.p }

func (in *interiorNode) setParent(p *interiorNode) { in.p = p }
