package zset

import golib "github.com/Senhnn/ds_and_algorithm"

var _ = NewZSet[string, int64, string]()

var _ = NewZSet[uintptr, int64, any]()

var _ = NewZSet[string, int64, golib.Void]()
