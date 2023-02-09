package formula

import (
	"gitee.com/quant1x/pandas"
	"github.com/viterin/vek"
)

// CROSS
//
//	判断向上金叉穿越 CROSS(MA(C,5),MA(C,10))
//	判断向下死叉穿越 CROSS(MA(C,10),MA(C,5))
func CROSS(S1, S2 pandas.Series) []bool {
	r1 := S1.DTypes()
	r2 := S2.DTypes()
	r11 := S1.Ref(1).DTypes()
	r12 := S2.Ref(1).DTypes()

	b1 := CompareLt(r11, r12)
	b2 := CompareGte(r1, r2)

	c := vek.And(b1, b2)
	return c
}
