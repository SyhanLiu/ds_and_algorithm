package linkedlist

import (
	golib "github.com/Senhnn/ds_and_algorithm"
)

type LinkedList[T any] struct {
	head, tail *Node[T]
	length     uint64
}

type Node[T any] struct {
	prev, next *Node[T]
	owner      *LinkedList[T]
	value      T
}

func (n *Node[T]) GetValue() T {
	return n.value
}

func (n *Node[T]) SetValue(value T) {
	n.value = value
}

func NewLinkedList[T any]() *LinkedList[T] {
	this := new(LinkedList[T])
	this.head = new(Node[T])
	this.tail = new(Node[T])
	this.head.next = this.tail
	this.tail.prev = this.head
	return this
}

func (l *LinkedList[T]) Front() (*Node[T], bool) {
	if l.length == 0 {
		return nil, false
	}
	return l.head.next, true
}

func (l *LinkedList[T]) Back() (*Node[T], bool) {
	if l.length == 0 {
		return nil, false
	}
	return l.tail.prev, true
}

func (l *LinkedList[T]) PushFront(value T) {
	node := &Node[T]{
		owner: l,
		value: value,
	}

	next := l.head.next
	l.head.next = node
	node.prev = l.head
	node.next = next
	next.prev = node
	l.length++
}

func (l *LinkedList[T]) PushBack(value T) {
	node := &Node[T]{
		owner: l,
		value: value,
	}

	prev := l.tail.prev
	l.tail.prev = node
	node.prev = prev
	node.next = l.tail
	prev.next = node
	l.length++
}

func (l *LinkedList[T]) InsertFrontNode(node *Node[T]) {
	node.owner = l
	next := l.head.next
	next.prev = node
	node.next = next
	node.prev = l.head
	l.head.next = node
}

func (l *LinkedList[T]) PopFront() (T, bool) {
	if l.length == 0 {
		return golib.Zero[T](), false
	}

	node := l.head.next
	l.head.next = node.next
	node.next.prev = l.head
	l.length--

	return node.value, true
}

func (l *LinkedList[T]) PopBack() (T, bool) {
	if l.length == 0 {
		return golib.Zero[T](), false
	}

	node := l.tail.prev
	l.tail.prev = node.prev
	node.prev.next = l.tail
	l.length--

	return node.value, true
}

func (l *LinkedList[T]) Iterator(callback func(node *Node[T]) bool) {
	end := l.tail
	for curr := l.head.next; curr != end; curr = curr.next {
		if !callback(curr) {
			return
		}
	}
}

func (l *LinkedList[T]) Len() uint64 {
	return l.length
}

func (l *LinkedList[T]) Remove(node *Node[T]) (*Node[T], bool) {
	if node.owner != l {
		return nil, false
	}

	node.next.prev = node.prev
	node.prev.next = node.next
	node.next = nil
	node.prev = nil
	node.owner = nil
	l.length--

	return node, true
}
