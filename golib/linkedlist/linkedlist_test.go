package linkedlist

import "testing"

func TestLinkedList1(t *testing.T) {
	l := NewLinkedList[int]()
	res := make([]int, 0)
	for i := 0; i < 100; i++ {
		l.PushBack(i)
	}

	l.Iterator(func(node *Node[int]) bool {
		res = append(res, node.value)
		return true
	})

	for i := 0; ; i++ {
		n, ok := l.PopFront()
		if !ok {
			break
		}
		if n != res[i] {
			t.Fatalf("list error n:%d != res[%d]:%d", n, i, res[i])
		}
	}

	if l.Len() != 0 {
		t.Fatalf("error")
	}
}

func TestLinkedList2(t *testing.T) {
	l := NewLinkedList[int]()
	res := make([]int, 0)
	for i := 0; i < 100; i++ {
		l.PushFront(i)
	}

	l.Iterator(func(node *Node[int]) bool {
		res = append([]int{node.value}, res...)
		return true
	})

	for i := 0; ; i++ {
		n, ok := l.PopBack()
		if !ok {
			break
		}
		if n != res[i] {
			t.Fatalf("list error n:%d != res[%d]:%d", n, i, res[i])
		}
	}

	if l.Len() != 0 {
		t.Fatalf("error")
	}
}

func TestLinkedList3(t *testing.T) {
	l := NewLinkedList[int]()
	res := make([]*Node[int], 0)
	m := make(map[int]struct{})
	for i := 0; i < 100; i++ {
		m[i] = struct{}{}
		l.PushFront(i)
		f, ok := l.Front()
		if !ok {
			t.Fatalf("error")
		}
		res = append(res, f)
	}

	for _, n := range res {
		delete(m, n.value)
		l.Remove(n)
	}

	if l.Len() != 0 || len(m) != 0 {
		t.Fatalf("error")
	}
}
