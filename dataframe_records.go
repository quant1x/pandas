package pandas

import (
	"fmt"
	"strconv"
)

// LoadRecords creates a new DataFrame based on the given records.
func LoadRecords(records [][]string, options ...LoadOption) DataFrame {
	// Set the default load options
	cfg := loadOptions{
		defaultType: SERIES_TYPE_STRING,
		detectTypes: true,
		hasHeader:   true,
		nanValues:   []string{"NA", "NaN", "<nil>"},
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
			if findInStringSlice(rawcol[j], cfg.nanValues) != -1 {
				rawcol[j] = "NaN"
			}
		}
		rawcols[i] = rawcol

		t, ok := cfg.types[colname]
		if !ok {
			t = cfg.defaultType
			if cfg.detectTypes {
				if l, err := findType(rawcol); err == nil {
					t = l
				}
			}
		}
		types[i] = t
	}

	columns := make([]Series, len(headers))
	for i, colname := range headers {
		cols := rawcols[i]
		col := NewSeries(types[i], colname, cols)
		//if col.Err != nil {
		//	return DataFrame{Err: col.Err}
		//}
		columns[i] = *col
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

func findType(arr []string) (Type, error) {
	var hasFloats, hasInts, hasBools, hasStrings bool
	for _, str := range arr {
		if str == "" || str == "NaN" {
			continue
		}
		if _, err := strconv.Atoi(str); err == nil {
			hasInts = true
			continue
		}
		if _, err := strconv.ParseFloat(str, 64); err == nil {
			hasFloats = true
			continue
		}
		if str == "true" || str == "false" {
			hasBools = true
			continue
		}
		hasStrings = true
	}

	switch {
	case hasStrings:
		return SERIES_TYPE_STRING, nil
	case hasBools:
		return SERIES_TYPE_BOOL, nil
	case hasFloats:
		return SERIES_TYPE_FLOAT, nil
	case hasInts:
		return SERIES_TYPE_INT, nil
	default:
		return SERIES_TYPE_STRING, fmt.Errorf("couldn't detect type")
	}
}
