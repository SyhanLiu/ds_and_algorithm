package lockfreequeue

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestNewLockFreeQueue(t *testing.T) {
	gNum := 100
	taskNum := 1000000

	q := NewLockFreeQueue[uintptr]()
	wg := sync.WaitGroup{}
	for i := 0; i < gNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < taskNum; j++ {
				q.Enqueue(uintptr(0))
			}
		}()
	}
	wg.Wait()

	var popNum int64 = 0
	for i := 0; i < gNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				_, ok := q.Dequeue()
				if !ok {
					break
				}
				atomic.AddInt64(&popNum, 1)
			}
		}()
	}
	wg.Wait()

	if !q.IsEmpty() {
		t.Errorf("error queue is not empty")
		return
	}

	if popNum != int64(gNum*taskNum) {
		t.Errorf("error popNum:%d != total:%d", popNum, int64(gNum*taskNum))
	} else {
		t.Logf("success")
	}
}
