package pandas

// Remove 删除一段范围内的记录
func (self DataFrame) Remove(p ScopeLimit) DataFrame {
	rowLen := self.Nrow()
	start, end, err := p.Limits(rowLen)
	if err != nil {
		return self
	}
	columns := []Series{}
	for i := range self.columns {
		ht := self.columns[i].Subset(0, start, true)
		tail := self.columns[i].Subset(end+1, rowLen).Values()
		ht.Append(tail)
		columns = append(columns, ht)
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
