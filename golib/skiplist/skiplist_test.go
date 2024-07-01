package skiplist

var _ = NewSkipList[string, int64, string]()

var _ = NewSkipList[uintptr, int64, any]()

var _ = NewSkipList[*Node[string, int64, any], int64, any]()
