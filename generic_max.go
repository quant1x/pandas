package pandas

import "gitee.com/quant1x/pandas/stat"

func (self *NDFrame) Max() any {
	values := self.Values()
	switch rows := values.(type) {
	case []bool:
		max := false
		i := 0
		for idx, iv := range rows {
			if iv && !max {
				max = iv
				i += 1
			}
			_ = idx
		}
		if i > 0 {
			return max
		}
		return false
	case []string:
		max := ""
		i := 0
		for idx, iv := range rows {
			if StringIsNaN(iv) {
				continue
			}
			if iv > max {
				max = iv
				i += 1
			}
			_ = idx
		}
		if i > 0 {
			return max
		}
		return StringNaN
	case []int64:
		max := int64(0)
		//i := 0
		for idx, iv := range rows {
			if Float64IsNaN(float64(iv)) {
				continue
			}
			if iv > max {
				max = iv
				//i = idx
			}
			_ = idx
		}
		return max
	//case []float32:
	//	max := float32(0)
	//	i := 0
	//	for idx, iv := range rows {
	//		if Float32IsNaN(iv) {
	//			continue
	//		}
	//		if iv > max {
	//			max = iv
	//			i += 1
	//		}
	//		_ = idx
	//	}
	//	if i > 0 {
	//		return max
	//	}
	//	return Nil2Float32
	case []float32:
		return stat.Max(rows)
	//case []float64:
	//	max := float64(0)
	//	i := 0
	//	for idx, iv := range rows {
	//		if Float64IsNaN(iv) {
	//			continue
	//		}
	//		if iv > max {
	//			max = iv
	//			i += 1
	//		}
	//		_ = idx
	//	}
	//	if i > 0 {
	//		return max
	//	}
	//	return Nil2Float64
	case []float64:
		return stat.Max(rows)
	default:
		panic(ErrUnsupportedType)
	}
	//return Nil2Float64
}
