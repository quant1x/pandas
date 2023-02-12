package pandas

import "gitee.com/quant1x/pandas/stat"

func slice_any_to_float64[T Number](s []T) []float64 {
	count := len(s)
	if count == 0 {
		return []float64{}
	}
	d := make([]float64, count)
	for idx, iv := range s {
		d[idx] = float64(iv)
	}
	return d
}

// any输入只能是一维slice或者数组
func numberToFloat64(v any) []float64 {
	var vs []float64
	switch values := v.(type) {
	case []float64:
		return values
	case []int64:
		return slice_any_to_float64(values)
	case []int32:
		return slice_any_to_float64(values)
	case []int:
		return slice_any_to_float64(values)
	case []bool:
		count := len(values)
		if count == 0 {
			return []float64{}
		}
		vs = make([]float64, count)
		for idx, iv := range values {
			vs[idx] = stat.AnyToFloat64(iv)
		}
	case []string:
		count := len(values)
		if count == 0 {
			return []float64{}
		}
		vs = make([]float64, count)
		for idx, iv := range values {
			vs[idx] = stat.AnyToFloat64(iv)
		}
	}
	return vs
}
