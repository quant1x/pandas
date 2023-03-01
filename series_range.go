package pandas

import (
	"gitee.com/quant1x/pandas/stat"
	"reflect"
)

// Copy 复制一个副本
func (self *NDFrame) Copy() stat.Series {
	vlen := self.Len()
	return self.Subset(0, vlen, true)
}

func (self *NDFrame) Subset(start, end int, opt ...any) stat.Series {
	// 默认不copy
	var __optCopy = false
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
	case []int32:
		subset := values[start:end]
		rows = len(subset)
		if !__optCopy {
			vs = subset
		} else {
			_vs := make([]int32, 0)
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
	case []float32:
		subset := values[start:end]
		rows = len(subset)
		if !__optCopy {
			vs = subset
		} else {
			_vs := make([]float32, 0)
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
	var s stat.Series
	s = &frame
	return s
}

// Select 选取一段记录
func (self *NDFrame) Select(r stat.ScopeLimit) stat.Series {
	start, end, err := r.Limits(self.Len())
	if err != nil {
		return nil
	}
	series := self.Subset(start, end+1)
	return series
}

func (self *NDFrame) IndexOf(index int, opt ...any) any {
	if index < 0 {
		index = self.Len() + index
	} else if index >= self.Len() {
		index = self.Len() - 1
	}
	var __optInplace = false
	if len(opt) > 0 {
		// 第一个参数为是否copy
		if _opt, ok := opt[0].(bool); ok {
			__optInplace = _opt
		}
	}
	if !__optInplace {
		return reflect.ValueOf(self.Values()).Index(index).Interface()
	}
	vv := reflect.ValueOf(self.values)
	tv := vv.Index(index)
	return tv
}
