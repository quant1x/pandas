package tests

import (
	"fmt"
	"gitee.com/quant1x/pandas"
	"strings"
	"testing"
)

func TestMean(t *testing.T) {
	csvStr := `
Country,Date,Age,Amount,Id
"United States",2012-02-01,50,112.1,01234
"United States",2012-02-01,32,321.31,54320
"United Kingdom",2012-02-01,17,18.2,12345
"United States",2012-02-01,32,321.31,54320
"United Kingdom",2012-02-01,NA,18.2,12345
"United States",2012-02-01,32,321.31,54320
"United States",2012-02-01,32,321.31,54320
Spain,2012-02-01,66,555.42,00241
`
	df := pandas.ReadCSV(strings.NewReader(csvStr))
	fmt.Println(df)
	df.SetNames("a", "b", "c", "d", "e")
	s1 := df.Col("d")
	fmt.Println(s1)

	closes := df.Col("d")
	//closes.Median()
	ma5 := closes.Rolling(5).Mean()
	fmt.Println(ma5)

	e1 := closes.EWM(pandas.Alpha{Span: 5, At: pandas.AlphaSpan}, false, false).Mean()
	fmt.Println(e1)
}
