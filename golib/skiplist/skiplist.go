package skiplist

import (
	"cmp"
	"math"
	"math/rand"
	"time"

	golib "github.com/Senhnn/ds_and_algorithm"
)

const (
	maxLevel   = 32
	skipList_P = 0.5
)

type SkipList[M, S cmp.Ordered, V any] struct {
	length uint64
	level  int
	head   *Node[M, S, V]
	tail   *Node[M, S, V]
	random *rand.Rand
}

func NewSkipList[M, S cmp.Ordered, V any]() *SkipList[M, S, V] {
	this := new(SkipList[M, S, V])
	this.length = 0
	this.level = 1
	this.head = newNode(maxLevel, golib.Zero[M](), golib.Zero[S](), golib.Zero[V]())
	this.tail = nil
	this.random = rand.New(rand.NewSource(time.Now().UnixNano()))
	return this
}

func (sl *SkipList[M, S, V]) Insert(member M, score S, value V) *Node[M, S, V] {
	update := [maxLevel]*Node[M, S, V]{}
	var curr *Node[M, S, V]
	rank := [maxLevel]uint64{}
	var i, level int

	curr = sl.head
	for i = sl.level - 1; i >= 0; i-- {
		if i == sl.level-1 {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}

		for next := curr.level[i].next; (next != nil) && (score > next.score || (score == next.score && next.member < member)); {
			rank[i] += curr.level[i].span
			curr = next
		}

		// 已经节点已经存在则直接修改
		// TODO need fix
		if next := curr.level[0].next; next != nil && (next.score == score) && (next.member == member) {
			curr.value = value
			return curr
		}

		update[i] = curr
	}

	level = sl.slRandomLevel()
	if level > sl.level {
		for i = sl.level; i < level; i++ {
			rank[i] = 0
			update[i] = sl.head
			update[i].level[i].span = sl.length
		}
		sl.level = level
	}
	node := newNode(level, member, score, value)
	for i = 0; i < level; i++ {
		node.level[i].next = update[i].level[i].next
		update[i].level[i].next = node

		// 修改span
		delta := rank[0] - rank[i]
		node.level[i].span = update[i].level[i].span - delta
		update[i].level[i].span = delta + 1
	}

	for i = level; i < sl.level; i++ {
		update[i].level[i].span++
	}

	if update[0] == sl.head {
		node.prev = nil
	} else {
		node.prev = update[0]
	}

	if node.level[0].next != nil {
		node.level[0].next.prev = node
	} else {
		sl.tail = node
	}

	sl.length++
	return node
}

func (sl *SkipList[M, S, V]) Delete(member M, score S) *Node[M, S, V] {
	var update [maxLevel]*Node[M, S, V]
	var i int

	curr := sl.head
	for i = sl.level; i >= 0; i-- {
		for next := curr.level[i].next; next != nil && (score > next.score || (score == next.score && next.member < member)); {
			curr = next
		}
	}

	curr = curr.level[0].next
	if curr != nil && curr.score == score && curr.member == member {
		sl.deleteNode(curr, &update)
		return curr
	}
	return nil
}

func (sl *SkipList[M, S, V]) deleteNode(target *Node[M, S, V], update *[maxLevel]*Node[M, S, V]) {
	var i int
	for i = 0; i < sl.level; i++ {
		if update[i].level[i].next == target {
			update[i].level[i].span += target.level[i].span - 1
			update[i].level[i].next = target.level[i].next
		} else {
			update[i].level[i].span--
		}
	}

	if sl.tail == target {
		sl.tail = target.prev
	} else {
		target.level[0].next.prev = target.prev
	}
	// 上方代码比redis源码更易懂
	//if target.level[0].next != nil {
	//	target.level[0].next.prev = target.prev
	//} else {
	//	sl.tail = target.prev
	//}

	for sl.level > 1 && sl.head.level[sl.level-1].next != nil {
		sl.level--
	}
	sl.length--
}

func (sl *SkipList[M, S, V]) Get(member M, score S) (*Node[M, S, V], uint64) {
	var curr *Node[M, S, V]
	var rank uint64

	curr = sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for next := curr.level[i].next; (next != nil) && (score > next.score || (score == next.score && next.member >= member)); {
			rank += curr.level[i].span
			curr = next
		}

		if (curr.score == score) && (curr.member == member) {
			return curr, rank
		}
	}

	return nil, 0
}

// 在第几层建立节点
// 1层为1/4，2层为1/4*4，以此类推。
func (sl *SkipList[M, S, V]) slRandomLevel() int {
	threshold := int32(skipList_P * float32(math.MaxInt32))
	level := 1
	for sl.random.Int31() < threshold {
		level += 1
	}
	if level < maxLevel {
		return level
	}
	return maxLevel
}

type Level[M, S cmp.Ordered, V any] struct {
	span uint64 // 两个节点之间的跨度
	next *Node[M, S, V]
}

type Node[M, S cmp.Ordered, V any] struct {
	member M
	score  S
	value  V
	prev   *Node[M, S, V]
	level  []Level[M, S, V]
}

func newNode[M, S cmp.Ordered, V any](level int, member M, score S, value V) *Node[M, S, V] {
	this := new(Node[M, S, V])
	this.member = member
	this.score = score
	this.value = value
	this.level = make([]Level[M, S, V], level)
	return this
}

func (n *Node[M, S, V]) GetMember() M {
	return n.member
}

func (n *Node[M, S, V]) GetScore() S {
	return n.score
}

func (n *Node[M, S, V]) GetValue() V {
	return n.value
}

func (n *Node[M, S, V]) SetValue(v V) {
	n.value = v
}
