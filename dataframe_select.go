package pandas

import "fmt"

// Col returns a copy of the Series with the given column name contained in the DataFrame.
// 选取一列
func (self DataFrame) Col(colname string) Series {
	if self.Err != nil {
		return NewSeriesWithType(SERIES_TYPE_INVAILD, "")
	}
	// Check that colname exist on dataframe
	idx := findInStringSlice(colname, self.Names())
	if idx < 0 {
		return NewSeriesWithType(SERIES_TYPE_INVAILD, "")
	}
	return self.columns[idx].Copy()
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
