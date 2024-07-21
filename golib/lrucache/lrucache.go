package lrucache

import (
	"github.com/Senhnn/ds_and_algorithm/linkedlist"
)

type LruCache[K comparable, V any] struct {
	cache map[K]*linkedlist.Node[V]
	list  *linkedlist.LinkedList[V]
	len   uint64
	cap   uint64
}

func NewLruCache[K comparable, V any](cap uint64) *LruCache[K, V] {
	if cap == 0 {
		panic("LruCache: cap must greater 0")
	}
	this := new(LruCache[K, V])
	this.cache = make(map[K]*linkedlist.Node[V])
	this.list = linkedlist.NewLinkedList[V]()
	this.len = 0
	this.cap = cap
	return this
}

func (l *LruCache[K, V]) Get(key K) *linkedlist.Node[V] {
	node, ok := l.cache[key]
	if !ok {
		return nil
	}
	l.list.Remove(node)
	l.list.InsertFrontNode(node)
	return node
}

func (l *LruCache[K, V]) Set(key K, value V) *linkedlist.Node[V] {
	node, ok := l.cache[key]
	if ok {
		l.list.Remove(node)
		l.list.InsertFrontNode(node)
		node.SetValue(value)
		return nil
	}

	if l.len >= l.cap {
		// delete old key
		node, ok = l.list.Back()
		if !ok || node == nil {
			panic("LruCache: list should not empty")
		}
		delete(l.cache, key)
		l.len--
	}

	l.list.PushFront(value)
	node, ok = l.list.Front()
	if !ok || node == nil {
		panic("LruCache: list should not empty")
	}
	l.cache[key] = node
	l.len++
	return node
}

func (l *LruCache[K, V]) Len() uint64 {
	return l.len
}

func (l *LruCache[K, V]) Cap() uint64 {
	return l.cap
}

func (l *LruCache[K, V]) Clear() {
	l.cache = make(map[K]*linkedlist.Node[V])
	l.list = linkedlist.NewLinkedList[V]()
	l.len = 0
}
