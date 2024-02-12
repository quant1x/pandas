package pandas

import (
	"gitee.com/quant1x/num"
)

// Rolling returns an array with elements that roll beyond the last position are re-introduced at the first.
// 滑动窗口, 数据不足是用空数组占位
func Rolling[T num.BaseType](S []T, N any) [][]T {
	sLen := len(S)
	// 这样就具备了序列化滑动窗口的特性了
	var window []num.DType
	switch vn := N.(type) {
	case int:
		window = num.Repeat(num.DType(vn), sLen)
	case []int:
		_N := num.Slice2DType(vn)
		//nd := _N[len(_N) - 1]
		window = num.Align[num.DType](_N, num.DTypeNaN, sLen)
	case []num.DType:
		window = num.Align(vn, num.DTypeNaN, sLen)
	case []T: // 这块到不了, N和S不是同一个泛型类型
		window = num.Slice2DType(vn)
		window = num.Align[num.DType](window, num.DTypeNaN, sLen)
	case Series:
		window = vn.DTypes()
	default:
		panic(num.ErrInvalidWindow)
	}
	blocks := make([][]T, sLen)
	for i := 0; i < sLen; i++ {
		n := window[i]
		shift := int(n)
		if num.DTypeIsNaN(n) || shift > i+1 {
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
