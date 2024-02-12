package pandas

import (
	"gitee.com/quant1x/num"
)

func (this *NDFrame) Min() any {
	values := this.Values()
	switch rows := values.(type) {
	case []bool:
		min := true
		i := 0
		for idx, iv := range rows {
			if !iv && min {
				min = iv
				i += 1
			}
			_ = idx
		}
		if i > 0 {
			return min
		}
		return false
	case []string:
		min := ""
		hasNaN := false
		i := 0
		for idx, iv := range rows {
			if num.StringIsNaN(iv) {
				hasNaN = true
				break
			} else if i < 1 {
				min = iv
				i += 1
			}
			if iv < min {
				min = iv
				i += 1
			}
			_ = idx
		}
		if hasNaN {
			return num.StringNaN
		} else if i > 0 {
			return min
		}
		return num.StringNaN
	case []int64:
		min := num.MaxInt64
		i := 0
		for idx, iv := range rows {
			if num.Float64IsNaN(float64(iv)) {
				continue
			} else if i < 1 {
				min = iv
				i += 1
			}
			if iv < min {
				min = iv
				i += 1
			}
			_ = idx
		}
		return min
	case []float32:
		min := num.MaxFloat32
		hasNan := false
		i := 0
		for idx, iv := range rows {
			if num.Float32IsNaN(iv) {
				hasNan = true
				break
			}
			if iv < min {
				min = iv
				i += 1
			}
			_ = idx
		}
		if hasNan {
			return num.Nil2Float32
		} else if i > 0 {
			return min
		}
		return num.Nil2Float32
	//case []float32:
	//	if this.Len() == 0 {
	//		return Nil2Float32
	//	}
	//	return stat.Min(rows)
	case []float64:
		min := num.MaxFloat64
		hasNaN := false
		i := 0
		for idx, iv := range rows {
			if num.Float64IsNaN(iv) {
				hasNaN = true
				break
			}
			if iv < min {
				min = iv
				i += 1
			}
			_ = idx
		}
		if hasNaN {
			return num.Nil2Float64
		} else if i > 0 {
			return min
		}
		return num.Nil2Float64
	//case []float64:
	//	if this.Len() == 0 {
	//		return Nil2Float64
	//	}
	//	return stat.Min(rows)
	default:
		panic(ErrUnsupportedType)
	}
	return num.Nil2Float64
}

func (this *NDFrame) ArgMin() int {
	return num.ArgMin2(this.DTypes())
}
