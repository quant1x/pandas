package pandas

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestCsv(t *testing.T) {
	csvStr := `
Country,Date,Age,Amount,Id,close
"United States",2012-02-01,50,112.1,01234,1.231111
"United States",2012-02-01,32,321.31,54320,1.232222
"United Kingdom",2012-02-01,17,18.2,12345,1.233333
"United States",2012-02-01,32,321.31,54320,1.234444
"United Kingdom",2012-02-01,NA,18.2,12345,1.235555
"United States",2012-02-01,32,321.31,54320,1.236666
"United States",2012-02-01,32,321.31,54320,1.237777
Spain,2012-02-01,66,555.42,00241,1.23
`
	df := ReadCSV(strings.NewReader(csvStr))
	fmt.Println(df)
	filename := "./testfiles/test-tutorials-w01.csv"
	_ = df.WriteCSV(filename)
	buf := new(bytes.Buffer)
	_ = df.WriteCSV(buf)
	df = ReadCSV(filename)
	fmt.Println(df)
	_ = df.SetNames("a", "b", "c", "d", "e")
	//s1 := df.Col("d")
	//fmt.Println(s1)
	//
	//closes := df.Col("d")
	//ma5 := closes.RollingV1(5).Mean()
	//dframe.NewSeries(closes, dframe.Floats, "")
	//fmt.Println(ma5)
	d := df.Col("d")
	fmt.Println(d)
	_ = csvStr

}

type T1 struct {
	X []int64 `json:"x"`
}

func TestEwm(t *testing.T) {
	//a := make(map[string][]int, 8)
	t01 := map[string]int64{
		"x": 1,
	}
	fmt.Println(t01)
	t02 := map[string][]int64{
		"x": {1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	fmt.Println(t02)
	text := `{"x":[1,2,3,4,5,6,7,8,9]}`
	reader := strings.NewReader(text)
	parser := json.NewDecoder(reader)
	var t1 T1
	a1 := parser.Decode(&t1)
	fmt.Println(a1, t1)
	var t2 map[string][]int
	a2 := parser.Decode(&t2)
	fmt.Println(a2, t2)
	df := ReadCSV(reader)
	fmt.Println(df)
}
