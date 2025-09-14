package pandas

import "github.com/quant1x/num"

// RollingAndExpandingMixin 滚动和扩展静态横切
type RollingAndExpandingMixin struct {
	Window num.Window[num.DType]
	Series Series
}

// GetBlocks series分块
func (r RollingAndExpandingMixin) v1GetBlocks() (blocks []Series) {
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

func (r RollingAndExpandingMixin) GetBlocks() (blocks []Series) {
	blocks = make([]Series, r.Series.Len())
	for i := 0; i < r.Series.Len(); i++ {
		N := r.Window.At(i)
		if num.DTypeIsNaN(N) || int(N) > i+1 {
			blocks[i] = r.Series.Empty()
			continue
		}
		window := int(N)
		start := i + 1 - window
		end := i + 1
		blocks[i] = r.Series.Subset(start, end, false)
	}

	return
}

//go:noinline
func (r RollingAndExpandingMixin) block(index int) Series {
	N := r.Window.At(index)
	if num.DTypeIsNaN(N) || int(N) > index+1 {
		return r.Series.Empty()
	}
	window := int(N)
	start := index + 1 - window
	end := index + 1
	return r.Series.Subset(start, end, false)
}

func (r RollingAndExpandingMixin) v1block(index int) Series {
	N := r.Window.At(index)
	if num.DTypeIsNaN(N) || int(N) > index+1 {
		return r.Series.Empty()
	}
	window := int(N)
	start := index + 1 - window
	end := index + 1
	return r.Series.Subset(start, end, false)
}

// Apply 接受一个返回DType计算类回调函数
func (r RollingAndExpandingMixin) v1Apply(f func(S Series, N num.DType) num.DType) (s Series) {
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

func (r RollingAndExpandingMixin) Apply(f func(S Series, N num.DType) num.DType) (s Series) {
	length := r.Series.Len()
	values := make([]num.DType, length)
	for i := 0; i < length; i++ {
		block := r.block(i)
		if block.Len() == 0 {
			values[i] = num.NaN()
			continue
		}
		N := r.Window.At(i)
		v := f(block, N)
		values[i] = v
	}
	s = SeriesWithName(r.Series.Name(), values)
	return
}

func (r RollingAndExpandingMixin) Count() Series {
	return r.v2Count()
}

func (r RollingAndExpandingMixin) v1Count() Series {
	s := r.Apply(func(S Series, N num.DType) num.DType {
		bs := S.Bools()
		return num.DType(num.Count(bs))
	})
	return s
}

func (r RollingAndExpandingMixin) v2Count() Series {
	x := r.Series.Values()
	switch vs := x.(type) {
	case []int32:
		d := num.RollingV1(vs, r.Window, func(N num.DType, values ...int32) int32 {
			return int32(num.Count(values))
		})
		return SliceToSeries(d)
	case []int64:
		d := num.RollingV1(vs, r.Window, func(N num.DType, values ...int64) int64 {
			return int64(num.Count(values))
		})
		return SliceToSeries(d)
	case []float32:
		d := num.RollingV1(vs, r.Window, func(N num.DType, values ...float32) float32 {
			return float32(num.Count(values))
		})
		return SliceToSeries(d)
	case []float64:
		d := num.RollingV1(vs, r.Window, func(N num.DType, values ...float64) float64 {
			return float64(num.Count(values))
		})
		return SliceToSeries(d)
	case []bool:
		length := len(vs)
		periods := num.Periods{Array: r.Window.V, N: r.Window.C}
		array := make([]num.DType, length)
		defaultValue := num.TypeDefault[num.DType]()
		for i := 0; i < length; i++ {
			n, ok := periods.At(i)
			if !ok {
				array[i] = defaultValue
				continue
			}
			shift := int(n)
			offset := i + 1
			start := offset - shift
			end := offset
			block := vs[start:end]
			result := num.Count(block)
			array[i] = num.DType(result)
		}
		return SliceToSeries(array)
	}

	panic(num.ErrUnsupportedType)
}

// Aggregation 接受一个聚合回调
func (r RollingAndExpandingMixin) Aggregation(f func(S Series) any) Series {
	s := r.Series.Copy()
	length := r.Series.Len()
	for i := 0; i < length; i++ {
		block := r.block(i)
		var value any
		if block.Len() == 0 {
			value = block.NaN()
		} else {
			value = f(block)
		}
		s.Set(i, value)
	}
	return s
}

// Aggregation 接受一个聚合回调
func (r RollingAndExpandingMixin) v1Aggregation(f func(S Series) any) Series {
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

// Aggregation 接受一个聚合回调
func (r RollingAndExpandingMixin) v2Aggregation(f func(S Series) any) Series {
	s := r.Series.Empty()
	for i := 0; i < r.Series.Len(); i++ {
		N := r.Window.At(i)
		if num.DTypeIsNaN(N) || int(N) > i+1 {
			s = s.Append(r.Series.NaN())
			continue
		}
		window := int(N)
		start := i + 1 - window
		end := i + 1
		s = s.Append(f(r.Series.Subset(start, end, false)))
	}
	return s
}

// Max 最大值
func (r RollingAndExpandingMixin) Max() Series {
	return r.v2Max()
}

func (r RollingAndExpandingMixin) v1Max() Series {
	s := r.Aggregation(func(S Series) any {
		return S.Max()
	})
	return s
}

func (r RollingAndExpandingMixin) v2Max() Series {
	x := r.Series.Values()
	switch vs := x.(type) {
	case []int32:
		d := num.RollingV1(vs, r.Window, func(N num.DType, values ...int32) int32 {
			return num.Max2(values)
		})
		return SliceToSeries(d)
	case []int64:
		d := num.RollingV1(vs, r.Window, func(N num.DType, values ...int64) int64 {
			return num.Max2(values)
		})
		return SliceToSeries(d)
	case []float32:
		d := num.RollingV1(vs, r.Window, func(N num.DType, values ...float32) float32 {
			return num.Max2(values)
		})
		return SliceToSeries(d)
	case []float64:
		d := num.RollingV1(vs, r.Window, func(N num.DType, values ...float64) float64 {
			return num.Max2(values)
		})
		return SliceToSeries(d)
	case []string:
		d := num.RollingV1(vs, r.Window, func(N num.DType, values ...string) string {
			return num.Max2(values)
		})
		return SliceToSeries(d)
	}

	panic(num.ErrUnsupportedType)
}

// Min 最小值
func (r RollingAndExpandingMixin) Min() Series {
	return r.v2Min()
}

func (r RollingAndExpandingMixin) v1Min() Series {
	return r.Aggregation(func(S Series) any {
		return S.Min()
	})
}

func (r RollingAndExpandingMixin) v2Min() Series {
	x := r.Series.Values()
	switch vs := x.(type) {
	case []int32:
		d := num.RollingV1(vs, r.Window, func(N num.DType, values ...int32) int32 {
			return num.Min2(values)
		})
		return SliceToSeries(d)
	case []int64:
		d := num.RollingV1(vs, r.Window, func(N num.DType, values ...int64) int64 {
			return num.Min2(values)
		})
		return SliceToSeries(d)
	case []float32:
		d := num.RollingV1(vs, r.Window, func(N num.DType, values ...float32) float32 {
			return num.Min2(values)
		})
		return SliceToSeries(d)
	case []float64:
		d := num.RollingV1(vs, r.Window, func(N num.DType, values ...float64) float64 {
			return num.Min2(values)
		})
		return SliceToSeries(d)
	case []string:
		d := num.RollingV1(vs, r.Window, func(N num.DType, values ...string) string {
			return num.Min2(values)
		})
		return SliceToSeries(d)
	}

	panic(num.ErrUnsupportedType)
}

// Mean returns the rolling mean.
func (r RollingAndExpandingMixin) Mean() (s Series) {
	return r.v2Mean()
}

// Mean returns the rolling mean.
func (r RollingAndExpandingMixin) v1Mean() (s Series) {
	var d []num.DType
	for _, block := range r.GetBlocks() {
		d = append(d, block.Mean())
	}
	s = r.Series.Empty(SERIES_TYPE_DTYPE)
	s.Rename(r.Series.Name())
	s = s.Append(d)
	return
}

func (r RollingAndExpandingMixin) v2Mean() Series {
	x := r.Series.Values()
	switch vs := x.(type) {
	case []int32:
		d := num.RollingV1(vs, r.Window, func(N num.DType, values ...int32) int32 {
			return num.Mean2(values)
		})
		return SliceToSeries(d)
	case []int64:
		d := num.RollingV1(vs, r.Window, func(N num.DType, values ...int64) int64 {
			return num.Mean2(values)
		})
		return SliceToSeries(d)
	case []float32:
		d := num.RollingV1(vs, r.Window, func(N num.DType, values ...float32) float32 {
			return num.Mean2(values)
		})
		return SliceToSeries(d)
	case []float64:
		d := num.RollingV1(vs, r.Window, func(N num.DType, values ...float64) float64 {
			return num.Mean2(values)
		})
		return SliceToSeries(d)
	}
	panic(num.ErrUnsupportedType)
}

func (r RollingAndExpandingMixin) Std() Series {
	return r.v2Std()
}

func (r RollingAndExpandingMixin) v1Std() Series {
	return r.Aggregation(func(S Series) any {
		return S.Std()
	})
}

func (r RollingAndExpandingMixin) v2Std() Series {
	x := r.Series.Values()
	switch vs := x.(type) {
	case []int32:
		d := num.RollingV1(vs, r.Window, func(N num.DType, values ...int32) int32 {
			return num.Std(values)
		})
		return SliceToSeries(d)
	case []int64:
		d := num.RollingV1(vs, r.Window, func(N num.DType, values ...int64) int64 {
			return num.Std(values)
		})
		return SliceToSeries(d)
	case []float32:
		d := num.RollingV1(vs, r.Window, func(N num.DType, values ...float32) float32 {
			return num.Std(values)
		})
		return SliceToSeries(d)
	case []float64:
		d := num.RollingV1(vs, r.Window, func(N num.DType, values ...float64) float64 {
			return num.Std(values)
		})
		return SliceToSeries(d)
	}
	panic(num.ErrUnsupportedType)
}

func (r RollingAndExpandingMixin) Sum() Series {
	return r.v3Sum()
}

func (r RollingAndExpandingMixin) v1Sum() Series {
	return r.Apply(func(S Series, N num.DType) num.DType {
		return S.Sum()
	})
}

func (r RollingAndExpandingMixin) v2Sum() Series {
	length := r.Series.Len()
	values := make([]num.DType, length)
	for i := 0; i < length; i++ {
		block := r.block(i)
		if block.Len() == 0 {
			values[i] = num.NaN()
			continue
		}
		//N := r.Window.At(i)
		v := num.Sum(block.DTypes())
		values[i] = v
	}
	s := SeriesWithName(r.Series.Name(), values)
	return s
}

func (r RollingAndExpandingMixin) v3Sum() Series {
	x := r.Series.Values()
	switch vs := x.(type) {
	case []int32:
		d := num.RollingV1(vs, r.Window, func(N num.DType, values ...int32) int32 {
			return num.Sum(values)
		})
		return SliceToSeries(d)
	case []int64:
		d := num.RollingV1(vs, r.Window, func(N num.DType, values ...int64) int64 {
			return num.Sum(values)
		})
		return SliceToSeries(d)
	case []float32:
		d := num.RollingV1(vs, r.Window, func(N num.DType, values ...float32) float32 {
			return num.Sum(values)
		})
		return SliceToSeries(d)
	case []float64:
		d := num.RollingV1(vs, r.Window, func(N num.DType, values ...float64) float64 {
			return num.Sum(values)
		})
		return SliceToSeries(d)
	}
	panic(num.ErrUnsupportedType)
}
