package pandas

import (
	"gitee.com/quant1x/pandas/exception"
	"gitee.com/quant1x/pandas/stat"
)

// RollingAndExpandingMixin 滚动和扩展静态横切
type RollingAndExpandingMixin struct {
	window []float32
	series Series
}

// Rolling RollingAndExpandingMixin
func (self *NDFrame) Rolling(param any) RollingAndExpandingMixin {
	var N []float32
	switch v := param.(type) {
	case int:
		N = stat.Repeat[float32](float32(v), self.Len())
	case []float32:
		N = stat.Align(v, Nil2Float32, self.Len())
	case Series:
		vs := v.Values()
		N = SliceToFloat32(vs)
		N = stat.Align(N, Nil2Float32, self.Len())
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
		if Float32IsNaN(N) || int(N) > i+1 {
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

// Apply 接受一个回调
func (r RollingAndExpandingMixin) Apply(f func(S Series, N float32) float32) (s Series) {
	s = r.series.Empty()
	for i, block := range r.getBlocks() {
		if block.Len() == 0 {
			s.Append(Nil2Float32)
			continue
		}
		v := f(block, r.window[i])
		s.Append(v)
	}
	return
}
