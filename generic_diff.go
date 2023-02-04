package pandas

import (
	"gitee.com/quant1x/pandas/stat"
	"reflect"
)

//func (self *NDFrame) Diff_() float64 {
//	if self.Len() < 1 {
//		return NaN()
//	}
//	fs := make([]float64, 0)
//	self.Apply(func(idx int, v any) {
//		f := AnyToFloat64(v)
//		fs = append(fs, f)
//	})
//	stdDev := avx2.Mean(fs)
//	return stdDev
//}

// Diff 元素的第一个离散差
// First discrete difference of element.
// Calculates the difference of a {klass} element compared with another
// element in the {klass} (default is element in previous row).
func (self *NDFrame) Diff(param any) (s Series) {
	if !(self.type_ == SERIES_TYPE_INT || self.type_ == SERIES_TYPE_FLOAT) {
		return NewSeries(SERIES_TYPE_INVAILD, "", "")
	}
	var N []float32
	switch v := param.(type) {
	case int:
		N = stat.Repeat[float32](float32(v), self.Len())
	case Series:
		vs := v.Values()
		N = sliceToFloat32(vs)
		N = stat.Align(N, Nil2Float32, self.Len())
	default:
		//periods = 1
		N = stat.Repeat[float32](float32(1), self.Len())
	}
	r := RollingAndExpandingMixin{
		window: N,
		series: self,
	}
	var d []float64
	var front = Nil2Float64
	for _, block := range r.getBlocks() {
		vs := reflect.ValueOf(block.Values())
		vl := vs.Len()
		if vl == 0 {
			d = append(d, Nil2Float64)
			continue
		}
		vf := vs.Index(0).Interface()
		vc := vs.Index(vl - 1).Interface()
		cu := AnyToFloat64(vc)
		cf := AnyToFloat64(vf)
		if Float64IsNaN(cu) || Float64IsNaN(front) {
			front = cf
			d = append(d, Nil2Float64)
			continue
		}
		diff := cu - front
		d = append(d, diff)
		front = cf
	}
	s = NewSeries(SERIES_TYPE_FLOAT, r.series.Name(), d)
	return
}
