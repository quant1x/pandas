package pandas

import (
	"gitee.com/quant1x/gox/api"
)

// Group 分组
func (this DataFrame) Group(columnName string, filter func(kind Type, e any) bool) DataFrame {
	return this.Filter(columnName, filter)
}

// Filter 过滤
func (this DataFrame) Filter(columnName string, filter func(kind Type, e any) bool) DataFrame {
	series := this.Col(columnName)
	if series.Len() == 0 {
		return this
	}
	t := series.Type()
	var indexes []int
	series.Apply(func(idx int, v any) {
		ok := filter(t, v)
		if ok {
			indexes = append(indexes, idx)
		}
	})
	ranges := api.IntsToRanges(indexes)
	df := DataFrame{}
	for _, r := range ranges {
		tmp := this.SelectRows(r)
		df = df.Concat(tmp)
	}
	return df
}
