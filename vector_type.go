package pandas

import "gitee.com/quant1x/num"

func (this Vector[T]) DTypes() []num.DType {
	return num.Slice2DType([]T(this))
}

func (this Vector[T]) Float32s() []float32 {
	return num.SliceToFloat32([]T(this))
}

func (this Vector[T]) Float64s() []float64 {
	return num.SliceToFloat64([]T(this))
}

func (this Vector[T]) Ints() []int {
	return num.AnyToSlice[int](this, this.Len())
}

func (this Vector[T]) Int32s() []int32 {
	d := make([]int32, this.Len())
	for i, v := range this {
		d[i] = num.AnyToInt32(v)
	}
	return d
}

func (this Vector[T]) Int64s() []int64 {
	return num.AnyToSlice[int64](this, this.Len())
}

func (this Vector[T]) Strings() []string {
	return num.SliceToString(this.Values())
}

func (this Vector[T]) Bools() []bool {
	return ToBool(this)
}
