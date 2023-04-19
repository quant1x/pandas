package pandas

import (
	"gitee.com/quant1x/pandas/stat"
)

// Group 分组
func (self DataFrame) Group(columnName string, filter func(kind stat.Type, e any) bool) DataFrame {
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
	ranges := stat.IntsToRanges(indexes)
	df := DataFrame{}
	for _, r := range ranges {
		tmp := self.SelectRows(r)
		df = df.Concat(tmp)
	}
	return df
}
