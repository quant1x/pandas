package stat

import (
	"github.com/mymmsc/gox/exception"
)

type NDArray[T BaseType] []T

func (self NDArray[T]) NaN() any {
	switch any(self).(type) {
	case []bool:
		return BoolNaN
	case []string:
		return StringNaN
	case []int64:
		return Nil2Int64
	case []float32:
		return Nil2Float32
	case []float64:
		return Nil2Float64
	default:
		panic(ErrUnsupportedType)
	}
}

func (self NDArray[T]) Floats() []float32 {
	return SliceToFloat32([]T(self))
}

func (self NDArray[T]) DTypes() []DType {
	return SliceToFloat64([]T(self))
}

func (self NDArray[T]) Ints() []Int {
	d := make([]Int, self.Len())
	for i, v := range self {
		d[i] = AnyToInt32(v)
	}
	return d
}

func (self NDArray[T]) Empty(tv ...Type) Series {
	empty := []T{}
	return NDArray[T](empty)
}

func (self NDArray[T]) Copy() Series {
	vlen := self.Len()
	return self.Subset(0, vlen, true)
}

func (self NDArray[T]) Records() []string {
	ret := make([]string, self.Len())
	self.Apply(func(idx int, v any) {
		ret[idx] = AnyToString(v)
	})
	return ret

}

func (self NDArray[T]) Repeat(x any, repeats int) Series {
	var d any
	switch values := self.Values().(type) {
	case []bool:
		_ = values
		d = Repeat(AnyToBool(x), repeats)
	case []string:
		d = Repeat(AnyToString(x), repeats)
	case []int64:
		d = Repeat(AnyToInt64(x), repeats)
	case []float32:
		d = Repeat(AnyToFloat32(x), repeats)
	default: //case []float64:
		d = Repeat(AnyToFloat64(x), repeats)
	}
	return NDArray[T](d.([]T))
}

func (self NDArray[T]) FillNa(v any, inplace bool) Series {
	d := FillNa(self, v, inplace)
	return NDArray[T](d)
}

func (self NDArray[T]) Shift(periods int) Series {
	values := self.Values().([]T)
	d := Shift(values, periods)
	return NDArray[T](d)
}

func (self NDArray[T]) Mean() DType {
	if self.Len() < 1 {
		return NaN()
	}
	d := Mean2(self)
	return Any2DType(d)
}

func (self NDArray[T]) StdDev() DType {
	if self.Len() < 1 {
		return NaN()
	}
	return self.Std()
}

func (self NDArray[T]) Max() any {
	d := Max2(self)
	return d
}

func (self NDArray[T]) Min() any {
	d := Min2(self)
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
			self[i] = anyToGeneric[T](r)
		}
	}
	return self
}

func (self NDArray[T]) Logic(f func(idx int, v any) bool) []bool {
	d := make([]bool, self.Len())
	for i, v := range self {
		d[i] = f(i, v)
	}
	return d
}

func (self NDArray[T]) Diff(n any) Series {
	d := Diff2(self, n)
	return NDArray[T](d)
}

func (self NDArray[T]) Ref(n any) Series {
	values := self.Values().([]T)
	d := Shift(values, n)
	return NDArray[T](d)
}

func (self NDArray[T]) Std() DType {
	if self.Len() < 1 {
		return NaN()
	}
	d := Std(self)
	return Any2DType(d)
}

func (self NDArray[T]) Sum() DType {
	if self.Len() < 1 {
		return NaN()
	}
	values := Slice2DType(self.Values())
	d := Sum(values)
	return Any2DType(d)
}

func (self NDArray[T]) Rolling(param any) RollingAndExpandingMixin {
	var N []DType
	switch v := param.(type) {
	case int:
		N = Repeat[DType](DType(v), self.Len())
	case []DType:
		N = Align(v, DTypeNaN, self.Len())
	case Series:
		vs := v.DTypes()
		N = Align(vs, DTypeNaN, self.Len())
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

	dest := NewSeries[DType]()
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
