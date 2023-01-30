package pandas

import (
	"fmt"
	"testing"
)

func TestNewSeriesFloat64(t *testing.T) {
	var s1 Series
	s1 = NewSeriesFloat64("sales", nil, 50.3, 23.4, 56.2)
	fmt.Println(s1)
	expected := 4

	if s1.NRows() != expected {
		t.Errorf("wrong val: expected: %v actual: %v", expected, s1.NRows())
	}
	s2 := s1.Shift(2)
	fmt.Println(s2)

	s3 := s1.Repeat(1, 2)

	fmt.Println(s3)
}
