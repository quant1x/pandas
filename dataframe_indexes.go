package pandas

import (
	"fmt"
)

func parseSelectIndexes(l int, indexes SelectIndexes, colnames []string) ([]int, error) {
	var idx []int
	switch indexes.(type) {
	case []int:
		idx = indexes.([]int)
	case int:
		idx = []int{indexes.(int)}
	case []bool:
		bools := indexes.([]bool)
		if len(bools) != l {
			return nil, fmt.Errorf("indexing error: index dimensions mismatch")
		}
		for i, b := range bools {
			if b {
				idx = append(idx, i)
			}
		}
	case string:
		s := indexes.(string)
		i := findInStringSlice(s, colnames)
		if i < 0 {
			return nil, fmt.Errorf("can't select columns: column name %q not found", s)
		}
		idx = append(idx, i)
	case []string:
		xs := indexes.([]string)
		for _, s := range xs {
			i := findInStringSlice(s, colnames)
			if i < 0 {
				return nil, fmt.Errorf("can't select columns: column name %q not found", s)
			}
			idx = append(idx, i)
		}
	//case Series:
	//	s := indexes.(Series)
	//	//if err := s.Err; err != nil {
	//	//	return nil, fmt.Errorf("indexing error: new values has errors: %v", err)
	//	//}
	//	//if s.HasNaN() {
	//	//	return nil, fmt.Errorf("indexing error: indexes contain NaN")
	//	//}
	//	switch s.Type() {
	//	case SERIES_TYPE_INT64:
	//		return s.Int32s()
	//	case series.Bool:
	//		bools, err := s.Bool()
	//		if err != nil {
	//			return nil, fmt.Errorf("indexing error: %v", err)
	//		}
	//		return parseSelectIndexes(l, bools, colnames)
	//	case series.String:
	//		xs := indexes.(series.Series).Records()
	//		return parseSelectIndexes(l, xs, colnames)
	//	default:
	//		return nil, fmt.Errorf("indexing error: unknown indexing mode")
	//	}
	default:
		return nil, fmt.Errorf("indexing error: unknown indexing mode")
	}
	return idx, nil
}

// SelectIndexes are the supported indexes used for the DataFrame.Select method. Currently supported are:
//
//	int              // Matches the given index number
//	[]int            // Matches all given index numbers
//	[]bool           // Matches all columns marked as true
//	string           // Matches the column with the matching column name
//	[]string         // Matches all columns with the matching column names
//	Series [Int]     // Same as []int
//	Series [Bool]    // Same as []bool
//	Series [String]  // Same as []string
type SelectIndexes any

// Select the given DataFrame columns
func (this DataFrame) Select(indexes SelectIndexes) DataFrame {
	if this.Err != nil {
		return this
	}
	idx, err := parseSelectIndexes(this.ncols, indexes, this.Names())
	if err != nil {
		return DataFrame{Err: fmt.Errorf("can't select columns: %v", err)}
	}
	columns := make([]Series, len(idx))
	for k, i := range idx {
		if i < 0 || i >= this.ncols {
			return DataFrame{Err: fmt.Errorf("can't select columns: index out of range")}
		}
		columns[k] = this.columns[i].Copy()
	}
	nrows, ncols, err := checkColumnsDimensions(columns...)
	if err != nil {
		return DataFrame{Err: err}
	}
	this = DataFrame{
		columns: columns,
		ncols:   ncols,
		nrows:   nrows,
	}
	colnames := this.Names()
	fixColnames(colnames)
	for i, colname := range colnames {
		this.columns[i].Rename(colname)
	}
	return this
}
