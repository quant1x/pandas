package pandas

import "gitee.com/quant1x/pandas/stat"

func (this *NDFrame) Max() any {
	values := this.Values()
	switch rows := values.(type) {
	case []bool:
		maxValue := false
		i := 0
		for idx, iv := range rows {
			if iv && !maxValue {
				maxValue = iv
				i += 1
			}
			_ = idx
		}
		if i > 0 {
			return maxValue
		}
		return false
	case []string:
		maxValue := ""
		hasNaN := false
		i := 0
		for idx, iv := range rows {
			if stat.StringIsNaN(iv) {
				hasNaN = true
				break
			}
			if iv > maxValue {
				maxValue = iv
				i += 1
			}
			_ = idx
		}
		if hasNaN {
			return stat.StringNaN
		} else if i > 0 {
			return maxValue
		}
		return stat.StringNaN
	case []int64:
		maxValue := stat.MinInt64
		//i := 0
		for idx, iv := range rows {
			if stat.Float64IsNaN(float64(iv)) {
				continue
			}
			if iv > maxValue {
				maxValue = iv
				//i = idx
			}
			_ = idx
		}
		return maxValue
	case []float32:
		maxValue := stat.MinFloat32
		hasNan := false
		i := 0
		for idx, iv := range rows {
			if stat.Float32IsNaN(iv) {
				hasNan = true
				break
			}
			if iv > maxValue {
				maxValue = iv
				i += 1
			}
			_ = idx
		}
		if hasNan {
			return stat.Nil2Float32
		} else if i > 0 {
			return maxValue
		}
		return stat.Nil2Float32
	//case []float32:
	//	if this.Len() == 0 {
	//		return Nil2Float32
	//	}
	//	return stat.Max(rows)
	case []float64:
		maxValue := stat.MinFloat64
		hasNaN := false
		i := 0
		for idx, iv := range rows {
			if stat.Float64IsNaN(iv) {
				hasNaN = true
				break
			}
			if iv > maxValue {
				maxValue = iv
				i += 1
			}
			_ = idx
		}
		if hasNaN {
			return stat.Nil2Float64
		} else if i > 0 {
			return maxValue
		}
		return stat.Nil2Float64
	//case []float64:
	//	if this.Len() == 0 {
	//		return Nil2Float64
	//	}
	//	return stat.Max(rows)
	default:
		panic(ErrUnsupportedType)
	}
	//return Nil2Float64
}

func (this *NDFrame) ArgMax() int {
	return stat.ArgMax2(this.DTypes())
}
