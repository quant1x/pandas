package pandas

import (
	"fmt"
	"gitee.com/quant1x/pandas/stat"
)

// Col returns a copy of the Series with the given column name contained in the DataFrame.
// 选取一列
func (self DataFrame) Col(colname string, args ...bool) stat.Series {
	inplace := false
	if len(args) >= 1 {
		inplace = args[0]
	}
	if self.Err != nil {
		return NewSeriesWithType(stat.SERIES_TYPE_INVAILD, "")
	}
	// Check that colname exist on dataframe
	idx := findInStringSlice(colname, self.Names())
	if idx < 0 {
		return NewSeriesWithType(stat.SERIES_TYPE_INVAILD, "")
	}
	if inplace {
		return self.columns[idx]
	}
	return self.columns[idx].Copy()
}

func (self DataFrame) ColAsNDArray(colname string) stat.Series {
	if self.Err != nil {
		return stat.NewSeries[stat.DType]()
	}
	// Check that colname exist on dataframe
	idx := findInStringSlice(colname, self.Names())
	if idx < 0 {
		return stat.NewSeries[stat.DType]()
	}
	vs := self.columns[idx].DTypes()
	return stat.NewSeries[stat.DType](vs...)
}

// SetNames changes the column names of a DataFrame to the ones passed as an
// argument.
// 修改全部的列名
func (self DataFrame) SetNames(colnames ...string) error {
	if len(colnames) != self.ncols {
		return fmt.Errorf("setting names: wrong dimensions")
	}
	for k, s := range colnames {
		self.columns[k].Rename(s)
	}
	return nil
}

// SetName 修改一个series的名称
func (self DataFrame) SetName(from string, to string) {
	for _, s := range self.columns {
		if s.Name() == from {
			s.Rename(to)
		}
	}
}
