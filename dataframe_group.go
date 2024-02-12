package pandas

import (
	"gitee.com/quant1x/gox/api"
)

// Group 分组
func (self DataFrame) Group(columnName string, filter func(kind Type, e any) bool) DataFrame {
	return self.Filter(columnName, filter)
}

// Filter 过滤
func (self DataFrame) Filter(columnName string, filter func(kind Type, e any) bool) DataFrame {
	series := self.Col(columnName)
	if series.Len() == 0 {
		return self
	}
	t := series.Type()
	indexes := []int{}
	series.Apply(func(idx int, v any) {
		ok := filter(t, v)
		if ok {
			indexes = append(indexes, idx)
		}
	})
	ranges := api.IntsToRanges(indexes)
	df := DataFrame{}
	for _, r := range ranges {
		tmp := self.SelectRows(r)
		df = df.Concat(tmp)
	}
	return df
}
