package pandas

import (
	"gitee.com/quant1x/gox/api"
)

// Remove 删除一段范围内的记录
func (this DataFrame) Remove(p api.ScopeLimit) DataFrame {
	rowLen := this.Nrow()
	start, end, err := p.Limits(rowLen)
	if err != nil {
		return this
	}
	columns := []Series{}
	for i := range this.columns {
		ht := this.columns[i].Subset(0, start, true)
		tail := this.columns[i].Subset(end+1, rowLen, false).Values()
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
