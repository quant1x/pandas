package pandas

import (
	"fmt"
	"gitee.com/quant1x/num"
	"reflect"
	"strings"
)

// vector 泛型向量
//
//	私有结构, 不对外开放
//	没有name, 默认是x
type vector[T num.BaseType] []T

func (this vector[T]) v1String() string {
	var t0 T
	records := this.Records(true)
	records = append(records, fmt.Sprintf("dtype: %T", t0))
	return strings.Join(records, "\n")
}

func (this vector[T]) String() string {
	var t0 T
	records := this.Records(true)
	//records = append(records, fmt.Sprintf("dtype: %T", t0))
	text := strings.Join(records, ",")
	return fmt.Sprintf("dtype[%T]: %s", t0, text)
}

func (this vector[T]) elementDefaultValue() any {
	kind := num.CheckoutRawType(this)
	switch kind {
	case reflect.Bool:
		return num.BoolNaN
	case reflect.String:
		return num.StringNaN
	case reflect.Int32:
		return num.Nil2Int32
	case reflect.Int64:
		return num.Nil2Int64
	case reflect.Float32:
		return num.Nil2Float32
	case reflect.Float64:
		return num.Nil2Float64
	default:
		panic(num.TypeError(this))
	}
}

func (this vector[T]) emptySlice() []T {
	s := make([]T, 0)
	return s
}

func (this vector[T]) NaN() any {
	switch this.Values().(type) {
	case []bool:
		return num.BoolNaN
	case []string:
		return num.StringNaN
	case []int32:
		return num.Nil2Int32
	case []int64:
		return num.Nil2Int64
	case []float32:
		return num.Nil2Float32
	case []float64:
		return num.Nil2Float64
	default:
		// 流程走到这里, 有两种情况
		// 1. 切片类型超出了布尔, 整型, 浮点和字符串
		// 2. this是nil
		panic("error type")
		return this.elementDefaultValue()
	}
}

func (this vector[T]) Empty(tv ...Type) Series {
	if len(tv) == 0 {
		// goland提示 Empty slice declaration using a literal
		// 不可以转成 var empty []T, 这样的结果是empty = nil, 无法断言
		// 当然, 也可以传入var empty []T, 后续所有设计引用的地方都需要对vector[T]=nil和ndarray.data进行判断
		//empty := []T{}
		var empty []T
		return vector[T](empty)
	}

	__type := tv[0]
	switch __type {
	case SERIES_TYPE_STRING:
		return vector[string]([]string{})
	case SERIES_TYPE_BOOL:
		return vector[bool]([]bool{})
	case SERIES_TYPE_INT32:
		return vector[int32]([]int32{})
	case SERIES_TYPE_INT64:
		return vector[int64]([]int64{})
	case SERIES_TYPE_FLOAT32:
		return vector[float32]([]float32{})
	case SERIES_TYPE_FLOAT64:
		return vector[float64]([]float64{})
	default:
		panic(num.ErrUnsupportedType)
	}
}

func (this vector[T]) Copy() Series {
	vlen := this.Len()
	return this.Subset(0, vlen, true)
}

func (this vector[T]) Records(round ...bool) []string {
	ret := make([]string, this.Len())
	needRound := false
	if len(round) > 0 {
		needRound = round[0]
	}
	t := this.Type()
	this.Apply(func(idx int, v any) {
		val := v
		if needRound && (t == SERIES_TYPE_FLOAT32 || t == SERIES_TYPE_FLOAT64) {
			ret[idx] = num.PrintString(val)
		} else {
			ret[idx] = num.AnyToString(val)
		}
	})
	return ret

}

func (this vector[T]) Repeat(x any, repeats int) Series {
	var d any
	switch values := this.Values().(type) {
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
	return vector[T](d.([]T))
}

func (this vector[T]) FillNa(v any, inplace bool) Series {
	d := num.FillNa(this, v, inplace)
	return vector[T](d)
}

func (this vector[T]) Shift(periods int) Series {
	values := this.Values().([]T)
	d := num.Shift(values, periods)
	return vector[T](d)
}

func (this vector[T]) Mean() num.DType {
	if this.Len() < 1 {
		return num.NaN()
	}
	d := num.Mean2(this)
	return num.Any2DType(d)
}

func (this vector[T]) StdDev() num.DType {
	if this.Len() < 1 {
		return num.NaN()
	}
	return this.Std()
}

func (this vector[T]) Max() any {
	d := num.Max2(this)
	return d
}

func (this vector[T]) Min() any {
	d := num.Min2(this)
	return d
}

func (this vector[T]) Apply(f func(idx int, v any)) {
	for i, v := range this {
		f(i, v)
	}
}

// Apply2 提供可替换功能的apply方法, 默认不替换
func (this vector[T]) Apply2(f func(idx int, v any) any, inplace ...bool) Series {
	hasInplace := false
	if len(inplace) >= 1 {
		hasInplace = inplace[0]
	}
	for i, v := range this {
		r := f(i, v)
		if hasInplace {
			this[i] = num.AnyToGeneric[T](r)
		}
	}
	return this
}

func (this vector[T]) Diff(n any) Series {
	d := num.Diff2(this, n)
	return vector[T](d)
}

func (this vector[T]) Ref(n any) Series {
	values := this.Values().([]T)
	d := num.Shift(values, n)
	return vector[T](d)
}

func (this vector[T]) Std() num.DType {
	if this.Len() < 1 {
		return num.NaN()
	}
	d := num.Std(this)
	return num.Any2DType(d)
}

func (this vector[T]) Sum() num.DType {
	if this.Len() < 1 {
		return num.NaN()
	}
	values := num.Slice2DType(this.Values())
	d := num.Sum(values)
	return num.Any2DType(d)
}
