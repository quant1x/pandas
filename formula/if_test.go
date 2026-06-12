package formula

import (
	"fmt"
	"testing"

	"github.com/quant1x/pandas"
)

func TestIF(t *testing.T) {
	S := pandas.NewSeriesWithType(pandas.SERIES_TYPE_FLOAT32, "", []float32{1, 1, 1})
	A := pandas.NewSeriesWithType(pandas.SERIES_TYPE_FLOAT32, "", []float32{11, 12, 13})
	B := pandas.NewSeriesWithType(pandas.SERIES_TYPE_FLOAT32, "", []float32{21, 22, 23})
	fmt.Println(IF(S, A, B))
}

func TestIFF(t *testing.T) {
	S := pandas.NewSeriesWithType(pandas.SERIES_TYPE_FLOAT32, "", []float32{1, 1, 1})
	A := pandas.NewSeriesWithType(pandas.SERIES_TYPE_FLOAT32, "", []float32{11, 12, 13})
	B := pandas.NewSeriesWithType(pandas.SERIES_TYPE_FLOAT32, "", []float32{21, 22, 23})
	fmt.Println(IFF(S, A, B))
	fmt.Println(IFF(S, 1, 0))
}

func TestIFN(t *testing.T) {
	S := pandas.NewSeriesWithType(pandas.SERIES_TYPE_FLOAT32, "", []float32{1, 0, 1})
	A := pandas.NewSeriesWithType(pandas.SERIES_TYPE_FLOAT32, "", []float32{11, 12, 13})
	B := pandas.NewSeriesWithType(pandas.SERIES_TYPE_FLOAT32, "", []float32{21, 22, 23})
	fmt.Println(IFN(S, A, B))
	fmt.Println(IFN(S, A, 0))
}
