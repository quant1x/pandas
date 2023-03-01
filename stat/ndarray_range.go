package stat

import (
	"golang.org/x/exp/slices"
	"reflect"
)

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
	var __optCopy = false
	if len(opt) > 0 {
		// 第一个参数为是否copy
		if _cp, ok := opt[0].(bool); ok {
			__optCopy = _cp
		}
	}
	values := []T(self)
	rows := self.Len()
	vvs := values[start:end]
	if __optCopy && rows > 0 {
		vvs = slices.Clone(vvs)
	}
	var d Series
	d = NDArray[T](vvs)
	return d
}

func (self NDArray[T]) Select(r ScopeLimit) Series {
	start, end, err := r.Limits(self.Len())
	if err != nil {
		return nil
	}
	series := self.Subset(start, end+1)
	return series
}
