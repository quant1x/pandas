package pandas

import (
	"gitee.com/quant1x/pandas/stat"
	"reflect"
)

// Diff 元素的第一个离散差
// First discrete difference of element.
// Calculates the difference of a {klass} element compared with another
// element in the {klass} (default is element in previous row).
func (self *NDFrame) Diff(param any) (s Series) {
	if !(self.type_ == SERIES_TYPE_INT64 || self.type_ == SERIES_TYPE_FLOAT32 || self.type_ == SERIES_TYPE_FLOAT64) {
		return NewSeries(SERIES_TYPE_INVAILD, "", "")
	}
	var N []stat.DType
	switch v := param.(type) {
	case int:
		N = stat.Repeat[stat.DType](stat.DType(v), self.Len())
	case Series:
		vs := v.DTypes()
		N = stat.Align(vs, stat.DTypeNaN, self.Len())
	default:
		//periods = 1
		N = stat.Repeat[stat.DType](stat.DType(1), self.Len())
	}
	r := RollingAndExpandingMixin{
		window: N,
		series: self,
	}
	var d []stat.DType
	var front = stat.DTypeNaN
	for _, block := range r.getBlocks() {
		vs := reflect.ValueOf(block.Values())
		vl := vs.Len()
		if vl == 0 {
			d = append(d, stat.DTypeNaN)
			continue
		}
		vf := vs.Index(0).Interface()
		vc := vs.Index(vl - 1).Interface()
		cu := stat.Any2DType(vc)
		cf := stat.Any2DType(vf)
		if stat.DTypeIsNaN(cu) || stat.DTypeIsNaN(front) {
			front = cf
			d = append(d, stat.DTypeNaN)
			continue
		}
		diff := cu - front
		d = append(d, diff)
		front = cf
	}
	s = NewSeries(SERIES_TYPE_DTYPE, r.series.Name(), d)
	return
}
