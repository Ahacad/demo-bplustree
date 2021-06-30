package bplustree

type kv struct {
	key   int
	value string
}

type kvs [MAXKV]kv

func (a *kvs) Len() int           { return len(a) }
func (a *kvs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a *kvs) Less(i, j int) bool { return a[i].key < a[j].key }

type leafNode struct {
	kvs   kvs
	count int
	next  *leafNode
	p     *interiorNode
}

func newLeafNode(p *interiorNode) *leafNode {
	return &leafNode{
		p: p,
	}
}
