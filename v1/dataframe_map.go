package v1

import (
	"fmt"
	"sort"
)

// LoadMaps creates a new DataFrame based on the given maps. This function assumes
// that every map on the array represents a row of observations.
func LoadMaps(maps []map[string]interface{}, options ...LoadOption) DataFrame {
	if len(maps) == 0 {
		return DataFrame{Err: fmt.Errorf("load maps: empty array")}
	}
	inStrSlice := func(i string, s []string) bool {
		for _, v := range s {
			if v == i {
				return true
			}
		}
		return false
	}
	// Detect all colnames
	var colnames []string
	for _, v := range maps {
		for k := range v {
			if exists := inStrSlice(k, colnames); !exists {
				colnames = append(colnames, k)
			}
		}
	}
	sort.Strings(colnames)
	records := make([][]string, len(maps)+1)
	records[0] = colnames
	for k, m := range maps {
		row := make([]string, len(colnames))
		for i, colname := range colnames {
			element := ""
			val, ok := m[colname]
			if ok {
				element = fmt.Sprint(val)
			}
			row[i] = element
		}
		records[k+1] = row
	}
	return LoadRecords(records, options...)
}
