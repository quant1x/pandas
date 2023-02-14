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
		hasNaN := false
		i := 0
		for idx, iv := range rows {
			if stat.StringIsNaN(iv) {
				hasNaN = true
				break
			}
			if iv > max {
				max = iv
				i += 1
			}
			_ = idx
		}
		if hasNaN {
			return stat.StringNaN
		} else if i > 0 {
			return max
		}
		return stat.StringNaN
	case []int64:
		max := stat.MinInt64
		//i := 0
		for idx, iv := range rows {
			if stat.Float64IsNaN(float64(iv)) {
				continue
			}
			if iv > max {
				max = iv
				//i = idx
			}
			_ = idx
		}
		return max
	case []float32:
		max := stat.MinFloat32
		hasNan := false
		i := 0
		for idx, iv := range rows {
			if stat.Float32IsNaN(iv) {
				hasNan = true
				break
			}
			if iv > max {
				max = iv
				i += 1
			}
			_ = idx
		}
		if hasNan {
			return stat.Nil2Float32
		} else if i > 0 {
			return max
		}
		return stat.Nil2Float32
	//case []float32:
	//	if self.Len() == 0 {
	//		return Nil2Float32
	//	}
	//	return stat.Max(rows)
	case []float64:
		max := stat.MinFloat64
		hasNaN := false
		i := 0
		for idx, iv := range rows {
			if stat.Float64IsNaN(iv) {
				hasNaN = true
				break
			}
			if iv > max {
				max = iv
				i += 1
			}
			_ = idx
		}
		if hasNaN {
			return stat.Nil2Float64
		} else if i > 0 {
			return max
		}
		return stat.Nil2Float64
	//case []float64:
	//	if self.Len() == 0 {
	//		return Nil2Float64
	//	}
	//	return stat.Max(rows)
	default:
		panic(ErrUnsupportedType)
	}
	//return Nil2Float64
}
