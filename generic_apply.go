package pandas

import "reflect"

func (self *NDFrame) apply(f func(idx int, v any)) {
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
