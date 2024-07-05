package lrucache

import "github.com/Senhnn/ds_and_algorithm/linkedlist"

type LruCache[K comparable, V any] struct {
	cache map[K]*linkedlist.Node[V]
	list  *linkedlist.LinkedList[V]
}
