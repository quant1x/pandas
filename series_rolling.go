package pandas

import "gitee.com/quant1x/num"

// RollingAndExpandingMixin 滚动和扩展静态横切
type RollingAndExpandingMixin struct {
	//Window []num.DType
	Window num.Window[num.DType]
	Series Series
}

// GetBlocks series分块
func (r RollingAndExpandingMixin) GetBlocks() (blocks []Series) {
	for i := 0; i < r.Series.Len(); i++ {
		//N := r.Window[i]
		N := r.Window.At(i)
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

// Apply 接受一个返回DType计算类回调函数
func (r RollingAndExpandingMixin) Apply(f func(S Series, N num.DType) num.DType) (s Series) {
	values := make([]num.DType, r.Series.Len())
	for i, block := range r.GetBlocks() {
		if block.Len() == 0 {
			values[i] = num.NaN()
			continue
		}
		N := r.Window.At(i)
		v := f(block, N)
		values[i] = v
	}
	s = r.Series.Empty(SERIES_TYPE_DTYPE)
	s.Rename(r.Series.Name())
	s = s.Append(values)
	return
}

func (r RollingAndExpandingMixin) Count() Series {
	//if r.Series.Type() != SERIES_TYPE_BOOL {
	//	panic("不支持非bool序列")
	//}
	//values := make([]num.DType, r.Series.Len())
	//for i, block := range r.GetBlocks() {
	//	if block.Len() == 0 {
	//		values[i] = 0
	//		continue
	//	}
	//	bs := block.Values().([]bool)
	//	values[i] = num.DType(num.Count(bs))
	//}
	//s = r.Series.Empty(SERIES_TYPE_DTYPE)
	//s.Rename(r.Series.Name())
	//s = s.Append(values)
	//return
	s := r.Apply(func(S Series, N num.DType) num.DType {
		bs := S.Bools()
		return num.DType(num.Count(bs))
	})
	return s
}

// Aggregation 接受一个聚合回调
func (r RollingAndExpandingMixin) Aggregation(f func(S Series) any) Series {
	s := r.Series.Empty()
	for _, block := range r.GetBlocks() {
		var value any
		if block.Len() == 0 {
			value = block.NaN()
		} else {
			value = f(block)
		}
		s = s.Append(value)
	}
	return s
}

// Max 最大值
func (r RollingAndExpandingMixin) Max() Series {
	s := r.Aggregation(func(S Series) any {
		return S.Max()
	})
	return s
}

// Min 最小值
func (r RollingAndExpandingMixin) Min() Series {
	return r.Aggregation(func(S Series) any {
		return S.Min()
	})
}

// Mean returns the rolling mean.
func (r RollingAndExpandingMixin) Mean() (s Series) {
	var d []num.DType
	for _, block := range r.GetBlocks() {
		d = append(d, block.Mean())
	}
	s = r.Series.Empty(SERIES_TYPE_DTYPE)
	s.Rename(r.Series.Name())
	s = s.Append(d)
	return
}

func (r RollingAndExpandingMixin) Std() Series {
	//s := r.Series.Empty()
	//for _, block := range r.GetBlocks() {
	//	s = s.Append(block.Std())
	//}
	//return s
	return r.Aggregation(func(S Series) any {
		return S.Std()
	})
}

func (r RollingAndExpandingMixin) Sum() Series {
	//var d []num.DType
	//for _, block := range r.GetBlocks() {
	//	d = append(d, block.Sum())
	//}
	//s := r.Series.Empty(SERIES_TYPE_DTYPE)
	//s.Rename(r.Series.Name())
	//s = s.Append(d)
	//return s
	return r.Apply(func(S Series, N num.DType) num.DType {
		return S.Sum()
	})
}
