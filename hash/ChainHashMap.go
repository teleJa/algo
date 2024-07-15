package hash

import (
	"container/list"
	"hash/crc32"
	"log"
)

type ChainHashMap struct {
	tables []*list.List
}

type Node struct {
	key string
	val int
}

func newNode(key string, val int) *Node {
	return &Node{
		key: key,
		val: val,
	}
}

func NewChainHashMap(length int) *ChainHashMap {
	return &ChainHashMap{
		tables: make([]*list.List, length),
	}
}

func (h *ChainHashMap) Put(key string, val int) {

	index := h.hash(key)
	l := h.tables[index]
	node := newNode(key, val)

	if l == nil {

		l = list.New()
		l.PushBack(node)
		h.tables[index] = l
		return
	}

	cur := l.Front()
	for cur != nil {
		n, ok := cur.Value.(*Node)
		if ok {
			if n.key == key {
				n.val = val
				return
			}
		}
		cur = cur.Next()
	}
	if cur == nil {
		l.PushBack(node)
	}

}

func (h *ChainHashMap) Get(key string) (bool, int) {
	idx := h.hash(key)
	l := h.tables[idx]
	if l == nil {
		return false, -1
	}

	cur := l.Front()
	for cur != nil {

		n, ok := cur.Value.(*Node)
		if ok {
			if n.key == key {
				return true, n.val
			}
		}
		cur = cur.Next()
	}
	return false, -1
}

func (h *ChainHashMap) Remove(key string) {
	idx := h.hash(key)
	l := h.tables[idx]
	if l == nil {
		return
	}

	cur := l.Front()
	for cur != nil {
		n, ok := cur.Value.(*Node)
		if ok {
			if n.key == key {
				l.Remove(cur)
				return
			}
		}
		cur = cur.Next()
	}
}

func (h *ChainHashMap) print() {
	for _, l := range h.tables {
		if l == nil {
			continue
		}
		cur := l.Front()
		for cur != nil {
			n := cur.Value.(*Node)
			log.Printf("%s-%d\n", n.key, n.val)
			cur = cur.Next()
		}
	}
}

func (h *ChainHashMap) hash(key string) int {
	return int(crc32.ChecksumIEEE([]byte(key))) % len(h.tables)
}
