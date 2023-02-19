package indicator

import (
	"fmt"
	"gitee.com/quant1x/data/cache"
	"testing"
)

func TestMACD(t *testing.T) {
	df := cache.KLine("sz002528")
	fmt.Println(df)
	df1 := MACD(df, 5, 13, 3)
	fmt.Println(df1)
}
