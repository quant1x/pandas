package pandas

import (
	"gitee.com/quant1x/gox/api"
)

// Subset returns a subset of the rows of the original DataFrame based on the
// Series subsetting indexes.
func (self DataFrame) Subset(start, end int) DataFrame {
	if self.Err != nil {
		return self
	}
	columns := make([]Series, self.ncols)
	for i, column := range self.columns {
		s := column.Subset(start, end)
		columns[i] = s
	}
	nrows, ncols, err := checkColumnsDimensions(columns...)
	if err != nil {
		return DataFrame{Err: err}
	}
	return DataFrame{
		columns: columns,
		ncols:   ncols,
		nrows:   nrows,
	}
}

// Sub 选择一个子集, start end 支持从后到前选择
func (self DataFrame) Sub(start, end int) DataFrame {
	sl := api.RangeFinite(start, end)
	return self.SelectRows(sl)
}

// SelectRows 选择一段记录
func (self DataFrame) SelectRows(p api.ScopeLimit) DataFrame {
	columns := []Series{}
	for i := range self.columns {
		columns = append(columns, self.columns[i].Select(p))
	}
	nrows, ncols, err := checkColumnsDimensions(columns...)
	if err != nil {
		return DataFrame{Err: err}
	}
	newDF := DataFrame{
		columns: columns,
		ncols:   ncols,
		nrows:   nrows,
	}
	return newDF
}

func (self DataFrame) Concat(dfb DataFrame) DataFrame {
	if self.Err != nil {
		return self
	}
	if dfb.Err != nil {
		return dfb
	}

	uniques := make(map[string]struct{})
	cols := []string{}
	for _, t := range []DataFrame{self, dfb} {
		for _, u := range t.Names() {
			if _, ok := uniques[u]; !ok {
				uniques[u] = struct{}{}
				cols = append(cols, u)
			}
		}
	}

	expandedSeries := make([]Series, len(cols))
	for k, v := range cols {
		aidx := findInStringSlice(v, self.Names())
		bidx := findInStringSlice(v, dfb.Names())

		// aidx and bidx must not be -1 at the same time.
		var a, b Series
		if aidx != -1 {
			a = self.columns[aidx]
		} else {
			bb := dfb.columns[bidx]
			a = NewSeries(bb.Type(), bb.Name(), make([]struct{}, self.nrows))

		}
		if bidx != -1 {
			b = dfb.columns[bidx]
		} else {
			b = NewSeries(a.Type(), a.Name(), make([]struct{}, dfb.nrows))
		}
		newSeries := a.Concat(b)
		expandedSeries[k] = newSeries
	}
	return NewDataFrame(expandedSeries...)
}

// IndexOf 取一条记录
//
//	idx 为负值时从后往前取
func (self DataFrame) IndexOf(idx int, opt ...any) map[string]any {
	one := map[string]any{}
	if idx < 0 {
		idx = self.Nrow() + idx
	} else if idx >= self.Nrow() {
		idx = self.Nrow() - 1
	}
	var __optInplace = false
	if len(opt) > 0 {
		// 第一个参数为是否copy
		if _opt, ok := opt[0].(bool); ok {
			__optInplace = _opt
		}
	}
	for _, series := range self.columns {
		key := series.Name()
		value := series.IndexOf(idx, __optInplace)
		one[key] = value
	}
	return one
}
