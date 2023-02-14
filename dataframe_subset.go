package pandas

import "gitee.com/quant1x/pandas/stat"

// Subset returns a subset of the rows of the original DataFrame based on the
// Series subsetting indexes.
func (self DataFrame) Subset(start, end int) DataFrame {
	if self.Err != nil {
		return self
	}
	columns := make([]stat.Series, self.ncols)
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

// Select 选择一段记录
func (self DataFrame) SelectRows(p stat.ScopeLimit) DataFrame {
	columns := []stat.Series{}
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
