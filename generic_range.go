package pandas

import (
	gc "github.com/huandu/go-clone"
	"reflect"
)

// Copy 复制一个副本
func (self *NDFrame) Copy() Series {
	vlen := self.Len()
	return self.Subset(0, vlen, true)
}

func (self *NDFrame) Subset(start, end int, opt ...any) Series {
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
	vv := reflect.ValueOf(self.values)
	vk := vv.Kind()
	switch vk {
	case reflect.Slice, reflect.Array: // 切片和数组同样的处理逻辑
		vs = vv.Slice(start, end).Interface()
		rows = vv.Len()
		if __optCopy && rows > 0 {
			vs = gc.Clone(vs)
		}
		rows = vv.Len()
		frame := NDFrame{
			formatter: self.formatter,
			name:      self.name,
			type_:     self.type_,
			nilCount:  0,
			rows:      rows,
			values:    vs,
		}
		return &frame
	default:
		// 其它类型忽略
	}
	return self.Empty()
}

func (self *NDFrame) oldSubset(start, end int, opt ...any) Series {
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
	switch values := self.values.(type) {
	case []bool:
		subset := values[start:end]
		rows = len(subset)
		if !__optCopy {
			vs = subset
		} else {
			_vs := make([]bool, 0)
			_vs = append(_vs, subset...)
			vs = _vs
		}
	case []string:
		subset := values[start:end]
		rows = len(subset)
		if !__optCopy {
			vs = subset
		} else {
			_vs := make([]string, 0)
			_vs = append(_vs, subset...)
			vs = _vs
		}
	case []int64:
		subset := values[start:end]
		rows = len(subset)
		if !__optCopy {
			vs = subset
		} else {
			_vs := make([]int64, 0)
			_vs = append(_vs, subset...)
			vs = _vs
		}
	case []float64:
		subset := values[start:end]
		rows = len(subset)
		if !__optCopy {
			vs = subset
		} else {
			_vs := make([]float64, 0)
			_vs = append(_vs, subset...)
			vs = _vs
		}
	}
	frame := NDFrame{
		formatter: self.formatter,
		name:      self.name,
		type_:     self.type_,
		nilCount:  0,
		rows:      rows,
		values:    vs,
	}
	var s Series
	s = &frame
	return s
}

// Select 选取一段记录
func (self *NDFrame) Select(r Range) Series {
	start, end, err := r.Limits(self.Len())
	if err != nil {
		return nil
	}
	series := self.Subset(start, end+1)
	return series
}
