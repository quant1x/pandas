package pandas

import (
	"gitee.com/quant1x/num"
)

// FillNa 填充NaN的元素为v
// inplace为真是修改series元素的值
// 如果v和Values()返回值的slice类型不一致就会panic
func FillNa[T num.GenericType](s *NDFrame, v T, inplace bool) *NDFrame {
	values := s.Values()
	switch rows := values.(type) {
	case []string:
		for idx, iv := range rows {
			if num.StringIsNaN(iv) && inplace {
				rows[idx] = num.AnyToString(v)
			}
		}
	case []float64:
		for idx, iv := range rows {
			if num.Float64IsNaN(iv) && inplace {
				rows[idx] = num.AnyToFloat64(v)
			}
		}
	}
	return s
}
