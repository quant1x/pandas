package tdx

import (
	"fmt"
	"testing"
)

func TestGetKLine(t *testing.T) {
	code := "000002"
	data := GetKLine(code, 0, 5)
	df := data
	//date := df.Col("date")
	//t1 := date.Map(func(element pandas.Element) pandas.Element {
	//	e := element.String()[0:10]
	//	element.Set(e)
	//	return element
	//})
	//df = df.Mutate(t1)
	fmt.Println(df)
	len := df.Nrow()
	df1 := df.Select([]int{0})
	df.Col("date")
	fmt.Println(df1)
	//df = df.Concat(df1)
	fmt.Println(df)

	df = GetKLineAll(code)
	fmt.Println(df)
	_ = len
}
