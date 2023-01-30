package df

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := NewSeries(x, Float, "x")
	fmt.Println(s)
	df := NewFrame(s)
	fmt.Println(df)
	a := df.Col("x")
	b := a.Rolling(2).Mean()
	fmt.Println(b)
	data := MakeValues(a.Float())
	//data := series.MakeValues(a.Float())
	w := data.Rolling(2)
	d0 := w.Mean()
	d0.Sort()
	fmt.Println(d0)
	fmt.Println(data)
	d1 := data.Cumsum()
	fmt.Println(d1.Values())
	_ = x
}
