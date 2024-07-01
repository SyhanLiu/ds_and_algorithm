package skiplist

import (
	"github.com/Senhnn/ds_and_algorithm"
)

var (
	maxLevel = 32
)

type SkipList[K comparable, S golib.Integer, V any] struct {
	kvMap    map[K]*Node[K, S, V]
	length   uint64
	maxLevel int32
	level    int32
	head     *Node[K, S, V]
	tail     *Node[K, S, V]
}

type Level[K comparable, S golib.Integer, V any] struct {
	span uint64
	prev *Node[K, S, V]
}

type Node[K comparable, S golib.Integer, V any] struct {
	key   K
	value V
	score S
	prev  *Node[K, S, V]
	next  *Node[K, S, V]
	level []Level[K, S, V]
}

func NewNode[K comparable, S golib.Integer, V any](level int, key K, score S, value V) *Node[K, S, V] {
	this := new(Node[K, S, V])
	this.key = key
	this.score = score
	this.value = value
	this.level = make([]Level[K, S, V], level)
	return this
}

func NewSkipList[K comparable, S golib.Integer, V any]() *SkipList[K, S, V] {
	this := new(SkipList[K, S, V])
	this.level = 1
	this.length = 0
	this.head = NewNode[K, S, V](maxLevel, nil, 0, nil)
	this.tail = nil
	return this
}
