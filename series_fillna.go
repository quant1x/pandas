package pandas

import "gitee.com/quant1x/pandas/stat"

// FillNa 填充NaN的元素为v
// inplace为真是修改series元素的值
// 如果v和Values()返回值的slice类型不一致就会panic
func FillNa[T stat.GenericType](s *NDFrame, v T, inplace bool) *NDFrame {
	values := s.Values()
	switch rows := values.(type) {
	case []string:
		for idx, iv := range rows {
			if stat.StringIsNaN(iv) && inplace {
				rows[idx] = stat.AnyToString(v)
			}
		}
	case []float64:
		for idx, iv := range rows {
			if stat.Float64IsNaN(iv) && inplace {
				rows[idx] = stat.AnyToFloat64(v)
			}
		}
	}
	return s
}
