package bplustree

const (
    MAXKV = 255
    MAXKC = 511
)

type node interface {
    find(key int) (int, bool)
    parent() *interiorNode
    setParent(*interiorNode)
    full() bool
}
