package pandas

// Subset returns a subset of the rows of the original DataFrame based on the
// Series subsetting indexes.
func (df DataFrame) Subset(start, end int) DataFrame {
	if df.Err != nil {
		return df
	}
	columns := make([]Series, df.ncols)
	for i, column := range df.columns {
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
