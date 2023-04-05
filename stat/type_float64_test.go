package stat

import (
	"fmt"
	"github.com/shopspring/decimal"
	"testing"
)

func TestDecimal(t *testing.T) {
	f := float64(1.234567891)
	fmt.Println(Decimal(f, 2))
	fmt.Println(Decimal(f, 3))
	fmt.Println(Decimal(f, 4))
	fmt.Println(Decimal(f, 5))
}

func TestDecimalV3(t *testing.T) {
	// 对于保留小数的处理
	pi := decimal.NewFromFloat(3.1415926535897932384626)
	pi1 := pi.Round(5)    // 对pi值四舍五入保留3位小数
	fmt.Println(pi1)      // 3.142
	pi2 := pi.Truncate(3) // 对pi值保留3位小数之后直接舍弃
	fmt.Println(pi2)      // 3.141
}
