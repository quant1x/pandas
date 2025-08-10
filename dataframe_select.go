package pandas

import (
	"fmt"

	"gitee.com/quant1x/num"
)

// Col returns a copy of the Series with the given column name contained in the DataFrame.
//
//	选取一列, Col方法的目的是保持现有的name等字段
func (this DataFrame) Col(colname string, args ...bool) Series {
	inplace := false
	if len(args) >= 1 {
		inplace = args[0]
	}
	if this.Err != nil {
		return NewSeriesWithType(SERIES_TYPE_INVAILD, "")
	}
	// Check that colname exist on dataframe
	idx := findInStringSlice(colname, this.Names())
	if idx < 0 {
		return NewSeriesWithType(SERIES_TYPE_INVAILD, "")
	}
	if inplace {
		return this.columns[idx]
	}
	return this.columns[idx].Copy()
}

// ColAsNDArray 切换成vector的无名Series
//
//	目的是为了计算统一数据类型
func (this DataFrame) ColAsNDArray(colname string) Series {
	if this.Err != nil {
		return ToVector[num.DType]()
	}
	// Check that colname exist on dataframe
	idx := findInStringSlice(colname, this.Names())
	if idx < 0 {
		return ToVector[num.DType]()
	}
	vs := this.columns[idx].DTypes()
	return ToVector[num.DType](vs...)
}

// SetNames changes the column names of a DataFrame to the ones passed as an
// argument.
// 修改全部的列名
func (this DataFrame) SetNames(colNames ...string) error {
	if len(colNames) != this.ncols {
		return fmt.Errorf("setting names: wrong dimensions")
	}
	for k, s := range colNames {
		this.columns[k].Rename(s)
	}
	return nil
}

// SetName 修改一个series的名称
func (this DataFrame) SetName(from string, to string) {
	for _, s := range this.columns {
		if s.Name() == from {
			s.Rename(to)
		}
	}
}
