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
	for i := 0; i < 100; i++ {
		r := rand.Intn(101)
		n := newNode(0, strconv.Itoa(r), r, golib.Void{})
		resM[strconv.Itoa(r)] = n
		sl.Insert(n.member, n.score, n.value)
	}
	if sl.length != uint64(len(resM)) {
		t.Errorf("length:%d != resM.len:%d", sl.length, len(resM))
	}
}
