package indicator

import (
	"fmt"
	"gitee.com/quant1x/pandas/data/cache"
	"testing"
)

func TestBRAR(t *testing.T) {
	df := cache.KLine("sz002528")
	fmt.Println(df)
	df1 := BRAR(df, 26)
	fmt.Println(df1)
}
