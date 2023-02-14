package v1

import (
	"fmt"
	"testing"
)

func TestReadExcel(t *testing.T) {
	filename := "../testfiles/test-excel-r01.xlsx"
	df := ReadExcel(filename)
	fmt.Println(df)
	toFile := "../testfiles/test-excel-w01.xlsx"
	err := df.WriteExcel(toFile)
	if err != nil {
		t.Errorf("write excel=%s, failed", toFile)
	}
}
