package pandas

import (
	"fmt"
	"testing"
)

func TestCreateSeries(t *testing.T) {
	t0 := []any{1, true, "abc", 3.45, NaN()}
	df0 := NewSeriesWithoutType("x", t0...)
	fmt.Printf("%+v\n", df0)

	s1 := NewSeriesWithoutType("sales", nil, 50.3, 23.4, 56.2)
	fmt.Println(s1)

	var values any
	values = make([]float64, 0)
	values = append(values.([]float64), 1)
	fmt.Printf("%+v\n", values)

	//s1 := []string{"1", "2"}
	//t1 := SERIES_TYPE_STRING
	//if t1 == SERIES_TYPE_STRING {
	//	s2 := reflect.MakeSlice(typeString, 0, 0).Interface()
	//	//st1 := reflect.TypeOf(s1)
	//	//s2 := reflect.New(st1).Interface()
	//	s2 = append(s2.([]string), "1")
	//	fmt.Printf("%+v\n", s1)
	//	fmt.Printf("%+v\n", s2)
	//}

}
