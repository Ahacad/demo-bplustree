package bplustree

import (
	"fmt"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	testCount := 1000000
	bt := newBTree()

	start := time.Now()
	for i := testCount; i > 0; i-- {
		bt.Insert(i, "")
	}
	fmt.Println(time.Now().Sub(start))
	verifyTree(bt, testCount, t)
}

func verifyTree(b *BTree, count int, t *testing.T) {
	verifyRoot(b, t)

	for i := 0; i < b.root.count; i++ {
		verifyNode(b.root.kcs[i].child, b.root, t)
	}

	leftMost := findLeftMost(b.root)

	if leftMost != b.first {
		t.Errorf("bt.first: want = %p, got = %p", b.first, leftMost)
	}
}

func verifyRoot(b *BTree, t *testing.T) {
	if b.root.parent() != nil {
		t.Errorf("root parent: want = nil, got = %p", b.root.parent())
	}
	if b.root.count < 1 {
		t.Errorf("root.min.child: want >= 1, got = %d", b.root.count)
	}
	if b.root.count > MAXKC {
		t.Errorf("root.max.child: want <= %d, got = %d", MAXKC, b.root.count)
	}
}

func findLeftMost(n node) *leafNode {
	switch nn := n.(type) {
	case *interiorNode:
		return findLeftMost(nn.kcs[0].child)
	case *leafNode:
		return nn
	default:
		panic("")
	}
}
