package formula

import (
	"reflect"

	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas"
)

// COUNT 统计S为真的天数
func COUNT(S any, N any) pandas.Series {
	ds := []bool{}
	switch s := S.(type) {
	case pandas.Series:
		ds = s.Bools()
	case []int8, []uint8, []int16, []uint16, []int32, []uint32, []int64, []uint64, []int, []uint, []uintptr, []float32, []float64, []bool, []string:
		//sh := (*reflect.SliceHeader)(unsafe.Pointer(&S))
		//fmt.Println(sh)
		v := reflect.ValueOf(s)
		length := v.Len()
		ds = num.AnyToSlice[bool](s, length)
	}
	s := V1COUNT(ds, N)
	return pandas.SliceToSeries(s)
}

// V1COUNT 统计S为真的天数
func V1COUNT(S []bool, N any) []int {
	xLen := len(S)
	x := num.Rolling(S, N)
	ret := make([]int, xLen)
	for i := 0; i < len(x); i++ {
		n := 0
		for _, v := range x[i] {
			if num.AnyToBool(v) {
				n++
			}
		}
		ret[i] = n
	}
	return ret
}
