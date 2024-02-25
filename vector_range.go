package pandas

import (
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/num"
	"reflect"
	"slices"
)

func (this vector[T]) IndexOf(index int, opt ...any) any {
	if index < 0 {
		index = this.Len() + index
	} else if index >= this.Len() {
		index = this.Len() - 1
	}
	var __optInplace = false
	if len(opt) > 0 {
		// 第一个参数为是否替换
		if _opt, ok := opt[0].(bool); ok {
			__optInplace = _opt
		}
	}
	value := this[index]
	if __optInplace {
		mv := reflect.ValueOf(this.Values())
		return mv.Index(index)
	}
	return value

}

func (this vector[T]) Set(index int, v any) {
	tmp := num.AnyToGeneric[T](v)
	this[index] = tmp
}

func (this vector[T]) Subset(start, end int, opt ...any) Series {
	// 默认不copy
	var __optCopy = false
	if len(opt) > 0 {
		// 第一个参数为是否copy
		if _cp, ok := opt[0].(bool); ok {
			__optCopy = _cp
		}
	}
	values := []T(this)
	rows := this.Len()
	vs := values[start:end]
	if __optCopy && rows > 0 {
		vs = slices.Clone(vs)
	}
	return vector[T](vs)
}

func (this vector[T]) Select(r api.ScopeLimit) Series {
	start, end, err := r.Limits(this.Len())
	if err != nil {
		return nil
	}
	series := this.Subset(start, end+1)
	return series
}
