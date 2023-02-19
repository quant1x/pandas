package indicator

import (
	"fmt"
	"gitee.com/quant1x/data/cache"
	"testing"
)

func TestKDJ(t *testing.T) {
	df := cache.KLine("sz002528")
	fmt.Println(df)
	df1 := KDJ(df, 9, 3, 3)
	fmt.Println(df1)
}
