package stat

import "reflect"

func (self NDArray[T]) IndexOf(index int, opt ...any) any {
	if index < 0 {
		index = self.Len() + index
	} else if index >= self.Len() {
		index = self.Len() - 1
	}
	var __optInplace = false
	if len(opt) > 0 {
		// 第一个参数为是否替换
		if _opt, ok := opt[0].(bool); ok {
			__optInplace = _opt
		}
	}
	value := self[index]
	if __optInplace {
		mv := reflect.ValueOf(self.Values())
		return mv.Index(index)
	}
	return value

}

func (self NDArray[T]) Subset(start, end int, opt ...any) Series {
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
			vs = Clone(vs)
			//vs = slices.Clone(vs)
		}
		rows = vvs.Len()
		var d Series
		d = NDArray[T](vs.([]T))
		return d
	default:
		// 其它类型忽略
	}
	return self.Empty()
}

func (self NDArray[T]) Select(r ScopeLimit) Series {
	start, end, err := r.Limits(self.Len())
	if err != nil {
		return nil
	}
	series := self.Subset(start, end+1)
	return series
}
