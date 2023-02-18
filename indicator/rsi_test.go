package indicator

import (
	"fmt"
	"gitee.com/quant1x/pandas/internal"
	"testing"
)

func TestRSI(t *testing.T) {
	df := internal.KLine("sz002528")
	fmt.Println(df)
	df1 := RSI(df, 6, 12, 24)
	fmt.Println(df1)
}
