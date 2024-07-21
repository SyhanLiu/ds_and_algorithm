package lrucache

import (
	"math/rand"
	"testing"
)

var op [2]string = [2]string{"get", "set"}

func TestNewLruCache(t *testing.T) {
	dataMap := make(map[int]int)  // key: value
	timesMap := make(map[int]int) // key: times
	c := NewLruCache[int, int](100)
	for i := 0; i < 100000; i++ {
		key := rand.Int()
		value := rand.Int()
		if op[rand.Int()%2] == "get" {
			c.Get(key)
			if _, ok := dataMap[key]; ok {
				timesMap[key]++
			}
		} else {
			c.Set(key, value)
		}
	}
}
