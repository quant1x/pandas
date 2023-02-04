package pandas

import (
	"gitee.com/quant1x/pandas/exception"
	"gitee.com/quant1x/pandas/stat"
)

// Rolling 滑动窗口
func (self *NDFrame) Rolling(window int) RollingWindow {
	return RollingWindow{
		window: window,
		series: self,
	}
}

// RollingAndExpandingMixin 滚动和扩展静态横切
type RollingAndExpandingMixin struct {
	window []float32
	series Series
}

// Rolling2 RollingAndExpandingMixin
func (self *NDFrame) Rolling2(param any) RollingAndExpandingMixin {
	var N []float32
	switch v := param.(type) {
	case int:
		N = stat.Repeat[float32](float32(v), self.Len())
	case Series:
		vs := v.Values()
		N = sliceToFloat32(vs)
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
		blocks = append(blocks, r.series.Subset(start, end))
	}

	return
}
