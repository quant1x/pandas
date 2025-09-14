package pandas

import (
	"github.com/quant1x/x/api"
)

// Subset returns a subset of the rows of the original DataFrame based on the
// Series subsetting indexes.
func (this DataFrame) Subset(start, end int) DataFrame {
	if this.Err != nil {
		return this
	}
	columns := make([]Series, this.ncols)
	for i, column := range this.columns {
		s := column.Subset(start, end, false)
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
func (this DataFrame) Sub(start, end int) DataFrame {
	sl := api.RangeFinite(start, end)
	return this.SelectRows(sl)
}

// SelectRows 选择一段记录
func (this DataFrame) SelectRows(p api.ScopeLimit) DataFrame {
	columns := []Series{}
	for i := range this.columns {
		columns = append(columns, this.columns[i].Select(p))
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

func (this DataFrame) Concat(dfb DataFrame) DataFrame {
	if this.Err != nil {
		return this
	}
	if dfb.Err != nil {
		return dfb
	}

	uniques := make(map[string]struct{})
	cols := []string{}
	for _, t := range []DataFrame{this, dfb} {
		for _, u := range t.Names() {
			if _, ok := uniques[u]; !ok {
				uniques[u] = struct{}{}
				cols = append(cols, u)
			}
		}
	}

	expandedSeries := make([]Series, len(cols))
	for k, v := range cols {
		aidx := findInStringSlice(v, this.Names())
		bidx := findInStringSlice(v, dfb.Names())

		// aidx and bidx must not be -1 at the same time.
		var a, b Series
		if aidx != -1 {
			a = this.columns[aidx]
		} else {
			bb := dfb.columns[bidx]
			a = NewSeriesWithType(bb.Type(), bb.Name(), make([]struct{}, this.nrows))

		}
		if bidx != -1 {
			b = dfb.columns[bidx]
		} else {
			b = NewSeriesWithType(a.Type(), a.Name(), make([]struct{}, dfb.nrows))
		}
		newSeries := a.Concat(b)
		expandedSeries[k] = newSeries
	}
	return NewDataFrame(expandedSeries...)
}

// IndexOf 取一条记录
//
//	idx 为负值时从后往前取
func (this DataFrame) IndexOf(idx int, opt ...any) map[string]any {
	one := map[string]any{}
	if idx < 0 {
		idx = this.Nrow() + idx
	} else if idx >= this.Nrow() {
		idx = this.Nrow() - 1
	}
	var __optInplace = false
	if len(opt) > 0 {
		// 第一个参数为是否copy
		if _opt, ok := opt[0].(bool); ok {
			__optInplace = _opt
		}
	}
	for _, series := range this.columns {
		key := series.Name()
		value := series.IndexOf(idx, __optInplace)
		one[key] = value
	}
	return one
}
