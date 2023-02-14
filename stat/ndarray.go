package stat

import (
	gc "github.com/huandu/go-clone"
	"reflect"
)

type NDArray[T BaseType] []T

func (self NDArray[T]) Name() string {
	//TODO implement me
	panic("implement me")
}

func (self NDArray[T]) Rename(name string) {
	//TODO implement me
	panic("implement me")
}

func (self NDArray[T]) Type() Type {
	return checkoutRawType(self)
}

func (self NDArray[T]) Values() any {
	return []T(self)
}

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

func (self NDArray[T]) Empty() Frame {
	var empty []T
	return NDArray[T](empty)
}

func (self NDArray[T]) Copy() Frame {
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

func (self NDArray[T]) Subset(start, end int, opt ...any) Frame {
	// 默认不copy
	var __optCopy bool = false
	if len(opt) > 0 {
		// 第一个参数为是否copy
		if _cp, ok := opt[0].(bool); ok {
			__optCopy = _cp
		}
	}
	var vs any
	var rows int
	vv := reflect.ValueOf(self.Values())
	vk := vv.Kind()
	switch vk {
	case reflect.Slice, reflect.Array: // 切片和数组同样的处理逻辑
		vvs := vv.Slice(start, end)
		vs = vvs.Interface()
		rows = vv.Len()
		if __optCopy && rows > 0 {
			vs = gc.Clone(vs)
		}
		rows = vvs.Len()
		var d Frame
		d = NDArray[T](vs.([]T))
		return d
	default:
		// 其它类型忽略
	}
	return self.Empty()
}

func (self NDArray[T]) Repeat(x any, repeats int) Frame {
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

func (self NDArray[T]) Shift(periods int) Frame {
	values := self.Values().([]T)
	d := Shift(values, periods)
	return NDArray[T](d)
}

func (self NDArray[T]) Mean() DType {
	d := Mean2(self)
	return Any2DType(d)
}

func (self NDArray[T]) StdDev() DType {
	return self.Std()
}

func (self NDArray[T]) FillNa(v any, inplace bool) Frame {
	d := FillNa(self, v, inplace)
	return NDArray[T](d)
}

func (self NDArray[T]) Max() any {
	d := Min2(self)
	return d
}

func (self NDArray[T]) Min() any {
	d := Min2(self)
	return d
}

func (self NDArray[T]) Select(r ScopeLimit) Frame {
	start, end, err := r.Limits(self.Len())
	if err != nil {
		return nil
	}
	series := self.Subset(start, end+1)
	return series
}

func (self NDArray[T]) Apply(f func(idx int, v any)) {
	for i, v := range self {
		f(i, v)
	}
}

func (self NDArray[T]) Logic(f func(idx int, v any) bool) []bool {
	d := make([]bool, self.Len())
	for i, v := range self {
		d[i] = f(i, v)
	}
	return d
}

func (self NDArray[T]) Diff(param any) Frame {
	d := Diff2(self, param)
	return NDArray[T](d)
}

func (self NDArray[T]) Ref(param any) Frame {
	values := self.Values().([]T)
	d := Shift3(values, param)
	return NDArray[T](d)
}

func (self NDArray[T]) Std() DType {
	d := Std(self)
	return Any2DType(d)
}

func (self NDArray[T]) Sum() DType {
	values := Slice2DType(self)
	d := Sum(values)
	return Any2DType(d)
}
