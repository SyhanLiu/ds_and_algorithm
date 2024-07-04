package skiplist

import (
	"fmt"
	"math"
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
