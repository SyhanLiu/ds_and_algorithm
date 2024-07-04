package zset

import (
	"cmp"

	"github.com/Senhnn/ds_and_algorithm/skiplist"
)

var (
	maxLevel = 32
)

type ZSet[M, S cmp.Ordered, V any] struct {
	kvMap map[M]*skiplist.Node[M, S, V]
	sl    *skiplist.SkipList[M, S, V]
}

func NewZSet[M, S cmp.Ordered, V any]() *ZSet[M, S, V] {
	this := new(ZSet[M, S, V])
	this.kvMap = make(map[M]*skiplist.Node[M, S, V])
	this.sl = skiplist.NewSkipList[M, S, V]()
	return this
}
