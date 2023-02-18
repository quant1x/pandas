package indicator

import (
	"fmt"
	"gitee.com/quant1x/pandas/internal"
	"testing"
)

func TestBRAR(t *testing.T) {
	df := internal.KLine("sz002528")
	fmt.Println(df)
	df1 := BRAR(df, 26)
	fmt.Println(df1)
}
