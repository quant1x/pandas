package pandas

import (
	"gitee.com/quant1x/pandas/exception"
	"gitee.com/quant1x/pandas/stat"
)

// RollingAndExpandingMixin 滚动和扩展静态横切
type RollingAndExpandingMixin struct {
	window []stat.DType
	series Series
}

// Rolling RollingAndExpandingMixin
func (self *NDFrame) Rolling(param any) RollingAndExpandingMixin {
	var N []stat.DType
	switch v := param.(type) {
	case int:
		N = stat.Repeat[stat.DType](stat.DType(v), self.Len())
	case []stat.DType:
		N = stat.Align(v, stat.DTypeNaN, self.Len())
	case Series:
		vs := v.DTypes()
		N = stat.Align(vs, stat.DTypeNaN, self.Len())
	default:
		panic(exception.New(1, "error window"))
	}
	w := RollingAndExpandingMixin{
		window: N,
		series: self,
	}
	return w
}

func (r RollingAndExpandingMixin) getBlocks() (blocks []Series) {
	for i := 0; i < r.series.Len(); i++ {
		N := r.window[i]
		if stat.DTypeIsNaN(N) || int(N) > i+1 {
			blocks = append(blocks, r.series.Empty())
			continue
		}
		window := int(N)
		start := i + 1 - window
		end := i + 1
		blocks = append(blocks, r.series.Subset(start, end, true))
	}

	return
}

func (r RollingAndExpandingMixin) Apply_v1(f func(S Series, N stat.DType) stat.DType) (s Series) {
	s = r.series.Empty()
	for i, block := range r.getBlocks() {
		if block.Len() == 0 {
			s.Append(stat.DTypeNaN)
			continue
		}
		v := f(block, r.window[i])
		s.Append(v)
	}
	return
}

// Apply 接受一个回调
func (r RollingAndExpandingMixin) Apply(f func(S Series, N stat.DType) stat.DType) (s Series) {
	values := make([]stat.DType, r.series.Len())
	for i, block := range r.getBlocks() {
		if block.Len() == 0 {
			values[i] = stat.DTypeNaN
			continue
		}
		v := f(block, r.window[i])
		values[i] = v
	}
	s = NewSeries(SERIES_TYPE_DTYPE, r.series.Name(), values)
	return
}
