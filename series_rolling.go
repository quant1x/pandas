package pandas

import "gitee.com/quant1x/num"

// RollingAndExpandingMixin 滚动和扩展静态横切
type RollingAndExpandingMixin struct {
	Window []num.DType
	Series Series
}

// GetBlocks series分块
func (r RollingAndExpandingMixin) GetBlocks() (blocks []Series) {
	for i := 0; i < r.Series.Len(); i++ {
		N := r.Window[i]
		if num.DTypeIsNaN(N) || int(N) > i+1 {
			blocks = append(blocks, r.Series.Empty())
			continue
		}
		window := int(N)
		start := i + 1 - window
		end := i + 1
		blocks = append(blocks, r.Series.Subset(start, end, false))
	}

	return
}

// Apply 接受一个回调
func (r RollingAndExpandingMixin) Apply(f func(S Series, N num.DType) num.DType) (s Series) {
	values := make([]num.DType, r.Series.Len())
	for i, block := range r.GetBlocks() {
		if block.Len() == 0 {
			values[i] = num.DTypeNaN
			continue
		}
		v := f(block, r.Window[i])
		values[i] = v
	}
	s = r.Series.Empty(SERIES_TYPE_DTYPE)
	s.Rename(r.Series.Name())
	s = s.Append(values)
	return
}
