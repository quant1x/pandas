package pandas

import (
	"gitee.com/quant1x/pandas/stat"
	"reflect"
)

func (self *NDFrame) Apply(f func(idx int, v any)) {
	vv := reflect.ValueOf(self.values)
	vk := vv.Kind()
	switch vk {
	case reflect.Invalid: // {interface} nil
		//series.assign(idx, size, Nil2Float64)
	case reflect.Slice: // 切片, 不定长
		for i := 0; i < vv.Len(); i++ {
			tv := vv.Index(i).Interface()
			f(i, tv)
		}
	case reflect.Array: // 数组, 定长
		for i := 0; i < vv.Len(); i++ {
			tv := vv.Index(i).Interface()
			f(i, tv)
		}
	default:
		// 其它类型忽略
	}
}

func (self *NDFrame) Logic(f func(idx int, v any) bool) []bool {
	x := make([]bool, self.Len())
	vv := reflect.ValueOf(self.values)
	vk := vv.Kind()
	switch vk {
	case reflect.Invalid: // {interface} nil
		//series.assign(idx, size, Nil2Float64)
	case reflect.Slice: // 切片, 不定长
		for i := 0; i < vv.Len(); i++ {
			tv := vv.Index(i).Interface()
			x[i] = f(i, tv)
		}
	case reflect.Array: // 数组, 定长
		for i := 0; i < vv.Len(); i++ {
			tv := vv.Index(i).Interface()
			x[i] = f(i, tv)
		}
	default:
		// 其它类型忽略
	}
	return x
}

func (self *NDFrame) Apply2(f func(idx int, v any) any, args ...bool) stat.Series {
	inplace := false
	if len(args) >= 1 {
		inplace = args[0]
	}
	vv := reflect.ValueOf(self.values)
	vk := vv.Kind()
	switch vk {
	case reflect.Invalid: // {interface} nil
		//series.assign(idx, size, Nil2Float64)
	case reflect.Slice, reflect.Array:
		for i := 0; i < vv.Len(); i++ {
			tv := vv.Index(i).Interface()
			r := f(i, tv)
			if inplace {
				vv.Index(i).Set(reflect.ValueOf(r))
			}
		}
	default:
		// 其它类型忽略
	}
	return self
}
