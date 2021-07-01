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

func verifyNode(n node, parent *interiorNode, t *testing.T) {
	switch nn := n.(type) {
	case *interiorNode:
		if nn.count < MAXKC/2 {
			t.Errorf("interior.min.child: want >= %d, got = %d", MAXKC/2, nn.count)
		}
		if nn.count > MAXKC {
			t.Errorf("interior.max.child: wnat <= %d, got = %d", MAXKC, nn.count)
		}
		if nn.parent() != parent {
			t.Errorf("interior.parent: want = %p, got = %p", parent, nn.parent())
		}

		var last int
		for i := 0; i < nn.count; i++ {
			key := nn.kcs[i].key
			if key != 0 && key < last {
				t.Errorf("interior.sort.key: want > %d, got = %d", last, key)
			}
			last = key

			if i == nn.count-1 && key != 0 {
				t.Errorf("interior.last.key: want = 0, got = %d", key)
			}
			verifyNode(nn.kcs[i].child, nn, t)
		}

	case *leafNode:
		if nn.parent() != parent {
			t.Errorf("leaf.parent: want = %p got = %p", parent, nn.parent())
		}
		if nn.count < MAXKV/2 {
			t.Errorf("leaf.min.child: want = %d, got = %d", MAXKV/2, nn.count)
		}
		if nn.count > MAXKV {
			t.Errorf("leaf.max.child: want <= %d, got = %d", MAXKV, nn.count)
		}
	}
}

func verifyLeaf(leftMost *leafNode, count int, t *testing.T) {
	curr := leftMost
	last := 0
	c := 0

	for curr != nil {
		for i := 0; i < curr.count; i++ {
			key := curr.kvs[i].key

			if key <= last {
				t.Errorf("leaf.sort.key: want > %d, got = %d", last, key)
			}
			last = key
			c++
		}
		curr = curr.next
	}

	if c != count {
		t.Errorf("leaf.count: want = %d, got = %d", count, c)
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
