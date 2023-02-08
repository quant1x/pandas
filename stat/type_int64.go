package stat

import "math"

const (
	MaxInt64          = int64(math.MaxInt64)
	MinInt64          = int64(math.MinInt64)
	Nil2Int64         = int64(0) // 空指针转int64
	Int64NaN          = int64(0) // int64 无效值
	True2Int64        = int64(1) // true转int64
	False2Int64       = int64(0) // false 转int64
	StringBad2Int64   = int64(0) // 字符串解析int64异常
	StringTrue2Int64  = int64(1) // 字符串true转int64
	StringFalse2Int64 = int64(0) // 字符串false转int64
)
