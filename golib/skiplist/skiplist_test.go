package skiplist

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"testing"

	golib "github.com/Senhnn/ds_and_algorithm"
)

var _ = NewSkipList[int, int, struct{}]()

func TestNewSkipList(t *testing.T) {
	n := NewSkipList[string, uint64, golib.Void]()
	fmt.Println(n.Get("1", 500))
}

func TestSkipList_slRandomLevel(t *testing.T) {
	n := NewSkipList[string, uint64, golib.Void]()

	res := [maxLevel]uint64{}
	for i := 0; i < math.MaxUint32; i++ {
		res[n.slRandomLevel()-1] += 1
	}
	fmt.Println(res)
}

func TestSkipList_InsertGetDelete(t *testing.T) {
	sl := NewSkipList[string, int, golib.Void]()
	resM := make(map[string]*Node[string, int, golib.Void], 0)
	for i := 0; i < 10000; i++ {
		r := rand.Intn(10001)
		n := newNode(0, strconv.Itoa(r), r, golib.Void{})
		resM[strconv.Itoa(r)] = n
		sl.Insert(n.member, n.score, n.value)
	}
	if sl.length != uint64(len(resM)) {
		t.Errorf("length:%d != resM.len:%d", sl.length, len(resM))
	}

	// 遍历
	var i = 0
	for n := sl.head.level[0].next; n != nil; n = n.level[0].next {
		next := n.level[0].next
		if next == nil {
			break
		}
		if n.score > next.score || (n.score == next.score && n.member >= next.member) {
			t.Fatalf("[FATAL] skiplist don't increase")
		}
		i++
	}

	for _, node := range resM {
		n, rank := sl.Get(node.GetMember(), node.GetScore())
		var prevNum, nextNum uint64
		for j := sl.head.level[0].next; j != nil; j = j.level[0].next {
			if j == n {
				continue
			}
			if j.score < n.score || (j.score == n.score && j.member < n.member) {
				prevNum++
			} else if j.score > n.score || (j.score == n.score && j.member > n.member) {
				nextNum++
			} else {
				t.Fatalf("[FATAL] unexpected branch")
			}
		}
		if prevNum+1 != rank {
			t.Fatalf("[FATAL] prevCal:%d != rank:%d", prevNum+1, rank)
		}
		if prevNum+nextNum+1 != sl.length {
			t.Fatalf("[FATAL] calNum:%d != sl.length:%d", prevNum+nextNum+1, sl.length)
		}
	}

	// 随机删除10个元素
	deleteS := make([]*Node[string, int, golib.Void], 0)
	i = 0
	for _, n := range resM {
		if i >= 10 {
			break
		}
		deleteS = append(deleteS, n)
		i++
	}
	for _, n := range deleteS {
		sl.Delete(n.GetMember(), n.GetScore())
		delete(resM, n.GetMember())
	}

	// 遍历
	i = 0
	for n := sl.head.level[0].next; n != sl.tail; n = n.level[0].next {
		next := n.level[0].next
		if next == nil {
			break
		}
		if n.score > next.score || (n.score == next.score && n.member >= next.member) {
			t.Fatalf("[FATAL] skiplist don't increase")
		}
		i++
	}

	for _, node := range resM {
		n, rank := sl.Get(node.GetMember(), node.GetScore())
		var prevNum, nextNum uint64
		for j := sl.head.level[0].next; j != nil; j = j.level[0].next {
			if j == n {
				continue
			}
			if j.score < n.score || (j.score == n.score && j.member < n.member) {
				prevNum++
			} else if j.score > n.score || (j.score == n.score && j.member > n.member) {
				nextNum++
			} else {
				t.Fatalf("[FATAL] unexpected branch")
			}
		}
		if prevNum+1 != rank {
			t.Fatalf("[FATAL] prevCal:%d != rank:%d", prevNum+1, rank)
		}
		if prevNum+nextNum+1 != sl.length {
			t.Fatalf("[FATAL] calNum:%d != sl.length:%d", prevNum+nextNum+1, sl.length)
		}
	}
}
