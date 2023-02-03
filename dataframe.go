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
		if s.Type() == SERIES_TYPE_INT {
			d = NewSeriesInt64(s.Name(), s.Values())
		} else if s.Type() == SERIES_TYPE_BOOL {
			d = NewSeriesBool(s.Name(), s.Values())
		} else if s.Type() == SERIES_TYPE_STRING {
			d = NewSeriesString(s.Name(), s.Values())
		} else {
			d = NewSeriesFloat64(s.Name(), s.Values())
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
func (df DataFrame) Dims() (int, int) {
	return df.Nrow(), df.Ncol()
}

// Nrow returns the number of rows on a DataFrame.
func (df DataFrame) Nrow() int {
	return df.nrows
}

// Ncol returns the number of columns on a DataFrame.
func (df DataFrame) Ncol() int {
	return df.ncols
}

// Returns error or nil if no error occured
func (df *DataFrame) Error() error {
	return df.Err
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
func (df DataFrame) Types() []string {
	coltypes := make([]string, df.ncols)
	for i, s := range df.columns {
		coltypes[i] = s.Type().String()
	}
	return coltypes
}

// Records return the string record representation of a DataFrame.
func (df DataFrame) Records() [][]string {
	var records [][]string
	records = append(records, df.Names())
	if df.ncols == 0 || df.nrows == 0 {
		return records
	}
	var tRecords [][]string
	for _, col := range df.columns {
		tRecords = append(tRecords, col.Records())
	}
	records = append(records, transposeRecords(tRecords)...)
	return records
}

// Getters/Setters for DataFrame fields
// ====================================

// Names returns the name of the columns on a DataFrame.
func (df DataFrame) Names() []string {
	colnames := make([]string, df.ncols)
	for i, s := range df.columns {
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
func fixColnames(colnames []string) {
	// Find duplicated and missing colnames
	dupnamesidx := make(map[string][]int)
	var missingnames []int
	for i := 0; i < len(colnames); i++ {
		a := colnames[i]
		if a == "" {
			missingnames = append(missingnames, i)
			continue
		}
		// for now, dupnamesidx contains the indices of *all* the columns
		// the columns with unique locations will be removed after this loop
		dupnamesidx[a] = append(dupnamesidx[a], i)
	}
	// NOTE: deleting a map key in a range is legal and correct in Go.
	for k, places := range dupnamesidx {
		if len(places) < 2 {
			delete(dupnamesidx, k)
		}
	}
	// Now: dupnameidx contains only keys that appeared more than once

	// Autofill missing column names
	counter := 0
	for _, i := range missingnames {
		proposedName := fmt.Sprintf("X%d", counter)
		for findInStringSlice(proposedName, colnames) != -1 {
			counter++
			proposedName = fmt.Sprintf("X%d", counter)
		}
		colnames[i] = proposedName
		counter++
	}

	// Sort map keys to make sure it always follows the same order
	var keys []string
	for k := range dupnamesidx {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Add a suffix to the duplicated colnames
	for _, name := range keys {
		idx := dupnamesidx[name]
		if name == "" {
			name = "X"
		}
		counter := 0
		for _, i := range idx {
			proposedName := fmt.Sprintf("%s_%d", name, counter)
			for findInStringSlice(proposedName, colnames) != -1 {
				counter++
				proposedName = fmt.Sprintf("%s_%d", name, counter)
			}
			colnames[i] = proposedName
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

// Read/Write Methods
// =================

// LoadOption is the type used to configure the load of elements
type LoadOption func(*loadOptions)

type loadOptions struct {
	// Specifies which is the default type in case detectTypes is disabled.
	defaultType Type

	// If set, the type of each column will be automatically detected unless
	// otherwise specified.
	detectTypes bool

	// If set, the first row of the tabular structure will be used as column
	// names.
	hasHeader bool

	// The names to set as columns names.
	names []string

	// Defines which values are going to be considered as NaN when parsing from string.
	nanValues []string

	// Defines the csv delimiter
	delimiter rune

	// EnablesLazyQuotes
	lazyQuotes bool

	// Defines the comment delimiter
	comment rune

	// The types of specific columns can be specified via column name.
	types map[string]Type
}

// DefaultType sets the defaultType option for loadOptions.
func DefaultType(t Type) LoadOption {
	return func(c *loadOptions) {
		c.defaultType = t
	}
}

// DetectTypes sets the detectTypes option for loadOptions.
func DetectTypes(b bool) LoadOption {
	return func(c *loadOptions) {
		c.detectTypes = b
	}
}

// HasHeader sets the hasHeader option for loadOptions.
func HasHeader(b bool) LoadOption {
	return func(c *loadOptions) {
		c.hasHeader = b
	}
}

// Names sets the names option for loadOptions.
func Names(names ...string) LoadOption {
	return func(c *loadOptions) {
		c.names = names
	}
}

// NaNValues sets the nanValues option for loadOptions.
func NaNValues(nanValues []string) LoadOption {
	return func(c *loadOptions) {
		c.nanValues = nanValues
	}
}

// WithTypes sets the types option for loadOptions.
func WithTypes(coltypes map[string]Type) LoadOption {
	return func(c *loadOptions) {
		c.types = coltypes
	}
}

// WithDelimiter sets the csv delimiter other than ',', for example '\t'
func WithDelimiter(b rune) LoadOption {
	return func(c *loadOptions) {
		c.delimiter = b
	}
}

// WithLazyQuotes sets csv parsing option to LazyQuotes
func WithLazyQuotes(b bool) LoadOption {
	return func(c *loadOptions) {
		c.lazyQuotes = b
	}
}

// WithComments sets the csv comment line detect to remove lines
func WithComments(b rune) LoadOption {
	return func(c *loadOptions) {
		c.comment = b
	}
}

func parseType(s string) (Type, error) {
	switch s {
	case "float", "float64", "float32":
		return SERIES_TYPE_FLOAT, nil
	case "int", "int64", "int32", "int16", "int8":
		return SERIES_TYPE_INT, nil
	case "string":
		return SERIES_TYPE_STRING, nil
	case "bool":
		return SERIES_TYPE_BOOL, nil
	}
	return SERIES_TYPE_INVAILD, fmt.Errorf("type (%s) is not supported", s)
}
