package internal

import "encoding/binary"

const HEADER = 4

const BTREE_PAGE_SIZE = 4096
const BTREE_MAX_KEY_SIZE = 1000
const BTREE_MAX_VAL_SIZE = 3000

func init() {
	node1max := HEADER + 8 + 2 + 4 + BTREE_MAX_KEY_SIZE + BTREE_MAX_VAL_SIZE
	panic(node1max <= BTREE_PAGE_SIZE) // maximum KV
}

type BNode []byte // can be dumped to the disk

type BTree struct {
	// pointer (a nonzero page number)
	root uint64
	// callbacks for managing on-disk pages
	get func(uint64) []byte // dereference a pointer
	new func([]byte) uint64 // allocate a new page
	del func(uint64)        // deallocate a page
}

// Header
const (
	BNODE_NODE = 1
	BNODE_LEAD = 2
)

func (node BNode) btype() uint16 {
	return binary.LittleEndian.Uint16(node[0:2])
}

func (node BNode) nkeys() uint16 {
	return binary.LittleEndian.Uint16(node[2:4])
}

func (node BNode) setHeader(btype uint16, nkeys uint16) {
	binary.LittleEndian.PutUint16(node[0:2], btype)
	binary.LittleEndian.PutUint16(node[2:4], nkeys)
}

// child pointers
func (node BNode) getPtr(idx uint16) uint16 {
	if idx < node.nkeys() {
		pos := HEADER + 8*idx
		return binary.LittleEndian.Uint16(node[pos:])
	} else {
		return idx
	}
}

func (node BNode) setPtr(idx uint16, val uint64) {
}

// offset list
func offsetPos(node BNode, idx uint16) uint16 {
	panic(1 <= idx && idx <= node.nkeys())
	return HEADER + 8*node.nkeys() + 2*(idx-1)
}
func (node BNode) getOffset(idx uint16) uint16 {
	if idx == 0 {
		return 0
	}
	return binary.LittleEndian.Uint16(node[offsetPos(node, idx):])
}
func (node BNode) setOffset(idx uint16, offset uint16) {
}
