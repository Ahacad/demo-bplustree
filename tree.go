package bplustree

type BTree struct {
	root     *interiorNode
	first    *leafNode
	leaf     int
	interior int
	height   int
}

func newBTree() *BTree {
	leaf := newLeafNode(nil)
	r := newInteriorNode(nil, leaf)
	leaf.p = r
	return &BTree{
		root:     r,
		first:    leaf,
		leaf:     1,
		interior: 1,
		height:   2,
	}
}

func (bt *BTree) First() *leafNode {
	return bt.first
}

// insert key-value into bplustree
func (bt *BTree) Insert(key int, value string) {}
