package pandas

import (
	"fmt"
	"gitee.com/quant1x/num"
)

// LoadRecords creates a new DataFrame based on the given records.
// 这个方法是从本地缓存文件读取数据的第二步, 数据从形式上只能是字符串
func LoadRecords(records [][]string, options ...LoadOption) DataFrame {
	// Set the default load options
	cfg := loadOptions{
		defaultType: SERIES_TYPE_STRING,
		detectTypes: true,
		hasHeader:   true,
		nanValues:   num.PossibleNaOfString,
	}

	// Set any custom load options
	for _, option := range options {
		option(&cfg)
	}

	if len(records) == 0 {
		return DataFrame{Err: fmt.Errorf("load records: empty DataFrame")}
	}
	if cfg.hasHeader && len(records) <= 1 {
		return DataFrame{Err: fmt.Errorf("load records: empty DataFrame")}
	}
	if cfg.names != nil && len(cfg.names) != len(records[0]) {
		if len(cfg.names) > len(records[0]) {
			return DataFrame{Err: fmt.Errorf("load records: too many column names")}
		}
		return DataFrame{Err: fmt.Errorf("load records: not enough column names")}
	}

	// Extract headers
	headers := make([]string, len(records[0]))
	if cfg.hasHeader {
		headers = records[0]
		records = records[1:]
	}
	if cfg.names != nil {
		headers = cfg.names
	}

	types := make([]Type, len(headers))
	rawcols := make([][]string, len(headers))
	for i, colname := range headers {
		rawcol := make([]string, len(records))
		for j := 0; j < len(records); j++ {
			rawcol[j] = records[j][i]
			// 收敛NaN的情况, 统一替换为NaN
			if findInStringSlice(rawcol[j], cfg.nanValues) != -1 {
				rawcol[j] = "NaN"
			}
		}
		rawcols[i] = rawcol

		t, ok := cfg.types[colname]
		if !ok {
			t = cfg.defaultType
			if cfg.detectTypes {
				if l, err := findTypeByString(rawcol); err == nil {
					t = l
				}
			}
		}
		types[i] = t
	}

	columns := make([]Series, len(headers))
	for i, colname := range headers {
		cols := rawcols[i]
		col := NewSeriesWithType(types[i], colname, cols)
		columns[i] = col
	}
	nrows, ncols, err := checkColumnsDimensions(columns...)
	if err != nil {
		return DataFrame{Err: err}
	}
	df := DataFrame{
		columns: columns,
		ncols:   ncols,
		nrows:   nrows,
	}

	colnames := df.Names()
	fixColnames(colnames)
	for i, colname := range colnames {
		df.columns[i].Rename(colname)
	}
	return df
}
