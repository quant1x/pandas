package pandas

import "gitee.com/quant1x/pandas/stat"

func slice_any_to_float32[T Number](s []T) []float32 {
	count := len(s)
	if count == 0 {
		return []float32{}
	}
	d := make([]float32, count)
	for idx, iv := range s {
		// 强制转换
		d[idx] = float32(iv)
	}
	return d
}

// SliceToFloat32 any输入只能是一维slice或者数组
func SliceToFloat32(v any) []float32 {
	var vs []float32
	switch values := v.(type) {
	case []int8:
		return slice_any_to_float32(values)
	case []uint8:
		return slice_any_to_float32(values)
	case []int16:
		return slice_any_to_float32(values)
	case []uint16:
		return slice_any_to_float32(values)
	case []int32:
		return slice_any_to_float32(values)
	case []uint32:
		return slice_any_to_float32(values)
	case []int64:
		return slice_any_to_float32(values)
	case []uint64:
		return slice_any_to_float32(values)
	case []int:
		return slice_any_to_float32(values)
	case []uint:
		return slice_any_to_float32(values)
	case []float32:
		// TODO:直接返回会不会有问题
		return values
	case []float64:
		return slice_any_to_float32(values)
	case []bool:
		count := len(values)
		if count == 0 {
			return []float32{}
		}
		vs = make([]float32, count)
		for idx, iv := range values {
			vs[idx] = boolToFloat32(iv)
		}
	case []string:
		count := len(values)
		if count == 0 {
			return []float32{}
		}
		vs = make([]float32, count)
		for idx, iv := range values {
			vs[idx] = float32(stat.AnyToFloat64(iv))
		}
	}
	return []float32{}
}
