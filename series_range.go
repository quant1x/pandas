package pandas

import (
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/pandas/stat"
	"reflect"
)

// Copy 复制一个副本
func (this *NDFrame) Copy() stat.Series {
	vlen := this.Len()
	return this.Subset(0, vlen, true)
}

func (this *NDFrame) Subset(start, end int, opt ...any) stat.Series {
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
	switch values := this.values.(type) {
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
		formatter: this.formatter,
		name:      this.name,
		type_:     this.type_,
		nilCount:  0,
		rows:      rows,
		values:    vs,
	}
	var s stat.Series
	s = &frame
	return s
}

// Select 选取一段记录
func (this *NDFrame) Select(r api.ScopeLimit) stat.Series {
	start, end, err := r.Limits(this.Len())
	if err != nil {
		return nil
	}
	series := this.Subset(start, end+1)
	return series
}

func (this *NDFrame) IndexOf(index int, opt ...any) any {
	if index < 0 {
		index = this.Len() + index
	} else if index >= this.Len() {
		index = this.Len() - 1
	}
	var __optInplace = false
	if len(opt) > 0 {
		// 第一个参数为是否copy
		if _opt, ok := opt[0].(bool); ok {
			__optInplace = _opt
		}
	}
	if !__optInplace {
		return reflect.ValueOf(this.Values()).Index(index).Interface()
	}
	vv := reflect.ValueOf(this.values)
	tv := vv.Index(index)
	return tv
}
