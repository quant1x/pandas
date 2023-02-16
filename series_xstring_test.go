package pandas

import (
	"fmt"
	"testing"
)

func TestSeriesString_Type(t *testing.T) {
	s := new(SeriesString)
	fmt.Println(s.Type())
	s1 := s.NDArray
	fmt.Println(s1)
}
