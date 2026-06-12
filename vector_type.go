package pandas

import "github.com/quant1x/num"

func (this vector[T]) DTypes() []num.DType {
	return num.Slice2DType([]T(this))
}

func (this vector[T]) Float32s() []float32 {
	return num.SliceToFloat32([]T(this))
}

func (this vector[T]) Float64s() []float64 {
	return num.SliceToFloat64([]T(this))
}

func (this vector[T]) Ints() []int {
	return num.AnyToSlice[int](this, this.Len())
}

func (this vector[T]) Int32s() []int32 {
	d := make([]int32, this.Len())
	for i, v := range this {
		d[i] = num.AnyToInt32(v)
	}
	return d
}

func (this vector[T]) Int64s() []int64 {
	return num.AnyToSlice[int64](this, this.Len())
}

func (this vector[T]) Strings() []string {
	return num.SliceToString(this.Values())
}

func (this vector[T]) Bools() []bool {
	return ToBool(this)
}
