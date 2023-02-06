package pandas

import "gitee.com/quant1x/pandas/stat"

func (self *NDFrame) Min() any {
	values := self.Values()
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
		i := 0
		for idx, iv := range rows {
			if StringIsNaN(iv) {
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
		if i > 0 {
			return min
		}
		return StringNaN
	case []int64:
		min := int64(0)
		i := 0
		for idx, iv := range rows {
			if Float64IsNaN(float64(iv)) {
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
		if self.Len() == 0 {
			return Nil2Float32
		}
		return stat.Min(rows)
	case []float64:
		if self.Len() == 0 {
			return Nil2Float64
		}
		return stat.Min(rows)
	}
	return Nil2Float64
}
