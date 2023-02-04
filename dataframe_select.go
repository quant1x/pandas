package pandas

import "fmt"

// Col returns a copy of the Series with the given column name contained in the DataFrame.
// 选取一列
func (df DataFrame) Col(colname string) Series {
	if df.Err != nil {
		return NewSeriesWithType(SERIES_TYPE_INVAILD, "")
	}
	// Check that colname exist on dataframe
	idx := findInStringSlice(colname, df.Names())
	if idx < 0 {
		return NewSeriesWithType(SERIES_TYPE_INVAILD, "")
	}
	return df.columns[idx].Copy()
}

// SetNames changes the column names of a DataFrame to the ones passed as an
// argument.
// 修改全部的列名
func (df DataFrame) SetNames(colnames ...string) error {
	if len(colnames) != df.ncols {
		return fmt.Errorf("setting names: wrong dimensions")
	}
	for k, s := range colnames {
		df.columns[k].Rename(s)
	}
	return nil
}

// SetName 修改一个series的名称
func (df DataFrame) SetName(from string, to string) {
	for _, s := range df.columns {
		if s.Name() == from {
			s.Rename(to)
		}
	}
}
