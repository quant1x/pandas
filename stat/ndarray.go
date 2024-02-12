package stat

import (
	"gitee.com/quant1x/gox/exception"
	"gitee.com/quant1x/num"
)

type NDArray[T num.BaseType] []T

func (self NDArray[T]) NaN() any {
	switch any(self).(type) {
	case []bool:
		return num.BoolNaN
	case []string:
		return num.StringNaN
	case []int64:
		return num.Nil2Int64
	case []float32:
		return num.Nil2Float32
	case []float64:
		return num.Nil2Float64
	default:
		panic(num.ErrUnsupportedType)
	}
}

func (self NDArray[T]) Floats() []float32 {
	return num.SliceToFloat32([]T(self))
}

func (self NDArray[T]) DTypes() []num.DType {
	return num.SliceToFloat64([]T(self))
}

func (self NDArray[T]) Ints() []num.Int {
	d := make([]num.Int, self.Len())
	for i, v := range self {
		d[i] = num.AnyToInt32(v)
	}
	return d
}

func (self NDArray[T]) Strings() []string {
	return num.SliceToString(self.Values())
}

func (self NDArray[T]) Bools() []bool {
	return ToBool(self)
}

func (self NDArray[T]) Empty(tv ...Type) Series {
	if len(tv) == 0 {
		empty := []T{}
		return NDArray[T](empty)
	}

	__type := tv[0]
	switch __type {
	case SERIES_TYPE_STRING:
		return NewSeries[string]()
	case SERIES_TYPE_BOOL:
		return NewSeries[bool]()
	case SERIES_TYPE_INT32:
		return NewSeries[int32]()
	case SERIES_TYPE_INT64:
		return NewSeries[int64]()
	case SERIES_TYPE_FLOAT32:
		return NewSeries[float32]()
	case SERIES_TYPE_FLOAT64:
		return NewSeries[float64]()
	default:
		panic(num.ErrUnsupportedType)
	}
}

func (self NDArray[T]) Copy() Series {
	vlen := self.Len()
	return self.Subset(0, vlen, true)
}

func (self NDArray[T]) Records(round ...bool) []string {
	ret := make([]string, self.Len())
	needRound := false
	if len(round) > 0 {
		needRound = round[0]
	}
	t := self.Type()
	self.Apply(func(idx int, v any) {
		val := v
		if needRound && (t == SERIES_TYPE_FLOAT32 || t == SERIES_TYPE_FLOAT64) {
			ret[idx] = num.PrintString(val)
		} else {
			ret[idx] = num.AnyToString(val)
		}
	})
	return ret

}

func (self NDArray[T]) Repeat(x any, repeats int) Series {
	var d any
	switch values := self.Values().(type) {
	case []bool:
		_ = values
		d = num.Repeat(num.AnyToBool(x), repeats)
	case []string:
		d = num.Repeat(num.AnyToString(x), repeats)
	case []int64:
		d = num.Repeat(num.AnyToInt64(x), repeats)
	case []float32:
		d = num.Repeat(num.AnyToFloat32(x), repeats)
	default: //case []float64:
		d = num.Repeat(num.AnyToFloat64(x), repeats)
	}
	return NDArray[T](d.([]T))
}

func (self NDArray[T]) FillNa(v any, inplace bool) Series {
	d := num.FillNa(self, v, inplace)
	return NDArray[T](d)
}

func (self NDArray[T]) Shift(periods int) Series {
	values := self.Values().([]T)
	d := num.Shift(values, periods)
	return NDArray[T](d)
}

func (self NDArray[T]) Mean() num.DType {
	if self.Len() < 1 {
		return num.NaN()
	}
	d := num.Mean2(self)
	return num.Any2DType(d)
}

func (self NDArray[T]) StdDev() num.DType {
	if self.Len() < 1 {
		return num.NaN()
	}
	return self.Std()
}

func (self NDArray[T]) Max() any {
	d := num.Max2(self)
	return d
}

func (self NDArray[T]) Min() any {
	d := num.Min2(self)
	return d
}

func (self NDArray[T]) Apply(f func(idx int, v any)) {
	//inplace := true
	for i, v := range self {
		f(i, v)
	}
}

// Apply2 提供可替换功能的apply方法, 默认不替换
func (self NDArray[T]) Apply2(f func(idx int, v any) any, args ...bool) Series {
	inplace := false
	if len(args) >= 1 {
		inplace = args[0]
	}
	for i, v := range self {
		r := f(i, v)
		if inplace {
			self[i] = num.AnyToGeneric[T](r)
		}
	}
	return self
}

func (self NDArray[T]) Diff(n any) Series {
	d := num.Diff2(self, n)
	return NDArray[T](d)
}

func (self NDArray[T]) Ref(n any) Series {
	values := self.Values().([]T)
	d := num.Shift(values, n)
	return NDArray[T](d)
}

func (self NDArray[T]) Std() num.DType {
	if self.Len() < 1 {
		return num.NaN()
	}
	d := num.Std(self)
	return num.Any2DType(d)
}

func (self NDArray[T]) Sum() num.DType {
	if self.Len() < 1 {
		return num.NaN()
	}
	values := num.Slice2DType(self.Values())
	d := num.Sum(values)
	return num.Any2DType(d)
}

func (self NDArray[T]) Rolling(param any) RollingAndExpandingMixin {
	var N []num.DType
	switch v := param.(type) {
	case int:
		N = num.Repeat[num.DType](num.DType(v), self.Len())
	case []num.DType:
		N = num.Align(v, num.DTypeNaN, self.Len())
	case Series:
		vs := v.DTypes()
		N = num.Align(vs, num.DTypeNaN, self.Len())
	default:
		panic(exception.New(1, "error window"))
	}
	w := RollingAndExpandingMixin{
		Window: N,
		Series: self,
	}
	return w
}

func (self NDArray[T]) EWM(alpha EW) ExponentialMovingWindow {
	atype := AlphaAlpha
	param := 0.00
	adjust := alpha.Adjust
	ignoreNA := alpha.IgnoreNA
	if alpha.Com != 0 {
		atype = AlphaCom
		param = alpha.Com
	} else if alpha.Span != 0 {
		atype = AlphaSpan
		param = alpha.Span
	} else if alpha.HalfLife != 0 {
		atype = AlphaHalfLife
		param = alpha.HalfLife
	} else {
		atype = AlphaAlpha
		param = alpha.Alpha
	}

	dest := NewSeries[num.DType]()
	dest = dest.Append(self)
	return ExponentialMovingWindow{
		Data:     dest,
		AType:    atype,
		Param:    param,
		Adjust:   adjust,
		IgnoreNA: ignoreNA,
		Cb:       alpha.Callback,
	}
}
