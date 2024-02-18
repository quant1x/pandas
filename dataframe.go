package pandas

import (
	"fmt"
	"sort"
)

// DataFrame 以gota的DataFrame的方法为主, 兼顾新流程, 避免单元格元素结构化
type DataFrame struct {
	columns []Series
	ncols   int
	nrows   int

	// deprecated: Use Error() instead
	Err error
}

// NewDataFrame is the generic DataFrame constructor
func NewDataFrame(se ...Series) DataFrame {
	if se == nil || len(se) == 0 {
		return DataFrame{Err: fmt.Errorf("empty DataFrame")}
	}

	columns := make([]Series, len(se))
	for i, s := range se {
		var d Series
		if s.Type() == SERIES_TYPE_INT64 {
			d = NewSeriesWithType(SERIES_TYPE_INT64, s.Name(), s.Values())
		} else if s.Type() == SERIES_TYPE_BOOL {
			d = NewSeriesWithType(SERIES_TYPE_BOOL, s.Name(), s.Values())
		} else if s.Type() == SERIES_TYPE_STRING {
			d = NewSeriesWithType(SERIES_TYPE_STRING, s.Name(), s.Values())
		} else if s.Type() == SERIES_TYPE_FLOAT32 {
			d = NewSeriesWithType(SERIES_TYPE_FLOAT32, s.Name(), s.Values())
		} else {
			d = NewSeriesWithType(SERIES_TYPE_FLOAT64, s.Name(), s.Values())
		}
		columns[i] = d
	}
	nrows, ncols, err := checkColumnsDimensions(columns...)
	if err != nil {
		return DataFrame{Err: err}
	}

	// Fill DataFrame base structure
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

// Dims retrieves the dimensions of a DataFrame.
func (this DataFrame) Dims() (int, int) {
	return this.Nrow(), this.Ncol()
}

// Nrow returns the number of rows on a DataFrame.
func (this DataFrame) Nrow() int {
	return this.nrows
}

// Ncol returns the number of columns on a DataFrame.
func (this DataFrame) Ncol() int {
	return this.ncols
}

// Returns error or nil if no error occured
func (this DataFrame) Error() error {
	return this.Err
}

// 检查列的尺寸
func checkColumnsDimensions(se ...Series) (nrows, ncols int, err error) {
	ncols = len(se)
	nrows = -1
	if se == nil || ncols == 0 {
		err = fmt.Errorf("no Series given")
		return
	}
	for i, s := range se {
		//if s.Err != nil {
		//	err = fmt.Errorf("error on series %d: %v", i, s.Err)
		//	return
		//}
		if nrows == -1 {
			nrows = s.Len()
		}
		if nrows != s.Len() {
			err = fmt.Errorf("arguments have different dimensions")
			return
		}
		_ = i
	}
	return
}

// Types returns the types of the columns on a DataFrame.
func (this DataFrame) Types() []string {
	coltypes := make([]string, this.ncols)
	for i, s := range this.columns {
		coltypes[i] = s.Type().String()
	}
	return coltypes
}

// Records return the string record representation of a DataFrame.
func (this DataFrame) Records(round ...bool) [][]string {
	needRound := false
	if len(round) > 0 {
		needRound = round[0]
	}
	var records [][]string
	records = append(records, this.Names())
	if this.ncols == 0 || this.nrows == 0 {
		return records
	}
	var tRecords [][]string
	for _, col := range this.columns {
		tRecords = append(tRecords, col.Records(needRound))
	}
	records = append(records, transposeRecords(tRecords)...)
	return records
}

// Getters/Setters for DataFrame fields
// ====================================

// Names returns the name of the columns on a DataFrame.
func (this DataFrame) Names() []string {
	colnames := make([]string, this.ncols)
	for i, s := range this.columns {
		colnames[i] = s.Name()
	}
	return colnames
}

func transposeRecords(x [][]string) [][]string {
	n := len(x)
	if n == 0 {
		return x
	}
	m := len(x[0])
	y := make([][]string, m)
	for i := 0; i < m; i++ {
		z := make([]string, n)
		for j := 0; j < n; j++ {
			z[j] = x[j][i]
		}
		y[i] = z
	}
	return y
}

// fixColnames assigns a name to the missing column names and makes it so that the
// column names are unique.
func fixColnames(colNames []string) {
	// Find duplicated and missing colNames
	dupNamesIdx := make(map[string][]int)
	var missingnames []int
	for i := 0; i < len(colNames); i++ {
		a := colNames[i]
		if a == "" {
			missingnames = append(missingnames, i)
			continue
		}
		// for now, dupNamesIdx contains the indices of *all* the columns
		// the columns with unique locations will be removed after this loop
		dupNamesIdx[a] = append(dupNamesIdx[a], i)
	}
	// NOTE: deleting a map key in a range is legal and correct in Go.
	for k, places := range dupNamesIdx {
		if len(places) < 2 {
			delete(dupNamesIdx, k)
		}
	}
	// Now: dupnameidx contains only keys that appeared more than once

	// Autofill missing column names
	counter := 0
	for _, i := range missingnames {
		proposedName := fmt.Sprintf("X%d", counter)
		for findInStringSlice(proposedName, colNames) != -1 {
			counter++
			proposedName = fmt.Sprintf("X%d", counter)
		}
		colNames[i] = proposedName
		counter++
	}

	// Sort map keys to make sure it always follows the same order
	var keys []string
	for k := range dupNamesIdx {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Add a suffix to the duplicated colNames
	for _, name := range keys {
		idx := dupNamesIdx[name]
		if name == "" {
			name = "X"
		}
		counter := 0
		for _, i := range idx {
			proposedName := fmt.Sprintf("%s_%d", name, counter)
			for findInStringSlice(proposedName, colNames) != -1 {
				counter++
				proposedName = fmt.Sprintf("%s_%d", name, counter)
			}
			colNames[i] = proposedName
			counter++
		}
	}
}

func findInStringSlice(str string, s []string) int {
	for i, e := range s {
		if e == str {
			return i
		}
	}
	return -1
}
