package pandas

import (
	"fmt"
	"testing"
)

func TestDataFrameT0(t *testing.T) {
	var s1 Series
	s1 = NewSeriesFloat64("sales", nil, 50.3, 23.4, 56.2)
	fmt.Println(s1)
	expected := 4

	if s1.Len() != expected {
		t.Errorf("wrong val: expected: %v actual: %v", expected, s1.Len())
	}
	s2 := s1.Shift(2)
	df := NewDataFrame(s1, *s2)
	fmt.Println(df)

	_ = s2
}
