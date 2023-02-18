package stat

import (
	"github.com/mymmsc/gox/exception"
)

// Rolling returns an array with elements that roll beyond the last position are re-introduced at the first.
// 滑动窗口, 数据不足是用空数组占位
func Rolling[T BaseType](S []T, N any) [][]T {
	sLen := len(S)
	// 这样就具备了序列化滑动窗口的特性了
	var window []DType
	switch vn := N.(type) {
	case int:
		window = Repeat(DType(vn), sLen)
	case []int:
		_N := Slice2DType(vn)
		//nd := _N[len(_N) - 1]
		window = Align[DType](_N, DTypeNaN, sLen)
	case []DType:
		window = Align(vn, DTypeNaN, sLen)
	case []T: // 这块到不了, N和S不是同一个泛型类型
		window = Slice2DType(vn)
		window = Align[DType](window, DTypeNaN, sLen)
	default:
		panic(exception.New(1, "error window"))
	}
	blocks := make([][]T, sLen)
	for i := 0; i < sLen; i++ {
		n := window[i]
		shift := int(n)
		if DTypeIsNaN(n) || shift > i+1 {
			blocks[i] = []T{}
			continue
		}
		start := i + 1 - shift
		end := i + 1
		subSet := S[start:end]
		blocks[i] = subSet
	}
	return blocks
}
