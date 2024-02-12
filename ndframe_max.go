package pandas

import (
	"gitee.com/quant1x/num"
)

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
			if num.StringIsNaN(iv) {
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
			return num.StringNaN
		} else if i > 0 {
			return maxValue
		}
		return num.StringNaN
	case []int64:
		maxValue := num.MinInt64
		//i := 0
		for idx, iv := range rows {
			if num.Float64IsNaN(float64(iv)) {
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
		maxValue := num.MinFloat32
		hasNan := false
		i := 0
		for idx, iv := range rows {
			if num.Float32IsNaN(iv) {
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
			return num.Nil2Float32
		} else if i > 0 {
			return maxValue
		}
		return num.Nil2Float32
	//case []float32:
	//	if this.Len() == 0 {
	//		return Nil2Float32
	//	}
	//	return stat.Max(rows)
	case []float64:
		maxValue := num.MinFloat64
		hasNaN := false
		i := 0
		for idx, iv := range rows {
			if num.Float64IsNaN(iv) {
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
			return num.Nil2Float64
		} else if i > 0 {
			return maxValue
		}
		return num.Nil2Float64
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
	return num.ArgMax2(this.DTypes())
}
