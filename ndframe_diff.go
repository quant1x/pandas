package pandas

import (
	"gitee.com/quant1x/num"
	"reflect"
)

// Diff 元素的第一个离散差
// First discrete difference of element.
// Calculates the difference of a {klass} element compared with another
// element in the {klass} (default is element in previous row).
func (this *NDFrame) Diff(param any) (s Series) {
	if !(this.type_ == SERIES_TYPE_INT64 || this.type_ == SERIES_TYPE_FLOAT32 || this.type_ == SERIES_TYPE_FLOAT64) {
		return NewSeries(SERIES_TYPE_INVAILD, "", "")
	}
	var N []num.DType
	switch v := param.(type) {
	case int:
		N = num.Repeat[num.DType](num.DType(v), this.Len())
	case Series:
		vs := v.DTypes()
		N = num.Align(vs, num.DTypeNaN, this.Len())
	default:
		//periods = 1
		N = num.Repeat[num.DType](num.DType(1), this.Len())
	}
	r := RollingAndExpandingMixin{
		Window: N,
		Series: this,
	}
	var d []num.DType
	var front = num.DTypeNaN
	for _, block := range r.GetBlocks() {
		vs := reflect.ValueOf(block.Values())
		vl := vs.Len()
		if vl == 0 {
			d = append(d, num.DTypeNaN)
			continue
		}
		vf := vs.Index(0).Interface()
		vc := vs.Index(vl - 1).Interface()
		cu := num.Any2DType(vc)
		cf := num.Any2DType(vf)
		if num.DTypeIsNaN(cu) || num.DTypeIsNaN(front) {
			front = cf
			d = append(d, num.DTypeNaN)
			continue
		}
		diff := cu - front
		d = append(d, diff)
		front = cf
	}
	s = NewSeries(SERIES_TYPE_DTYPE, r.Series.Name(), d)
	return
}
