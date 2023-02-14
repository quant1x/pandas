package v1

import "gonum.org/v1/gonum/mat"

// LoadMatrix loads the given Matrix as a DataFrame
// TODO: Add Loadoptions
func LoadMatrix(mat mat.Matrix) DataFrame {
	nrows, ncols := mat.Dims()
	columns := make([]Series, ncols)
	for i := 0; i < ncols; i++ {
		floats := make([]float64, nrows)
		for j := 0; j < nrows; j++ {
			floats[j] = mat.At(j, i)
		}
		columns[i] = NewSeries(SERIES_TYPE_FLOAT64, "", floats)
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
