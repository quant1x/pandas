package formula

import (
	"gitee.com/quant1x/pandas/stat"
	"gitee.com/quant1x/vek"
)

// CROSS
//
//	判断向上金叉穿越 V2CROSS(MA(C,5),MA(C,10))
//	判断向下死叉穿越 V2CROSS(MA(C,10),MA(C,5))
func CROSS(S1, S2 stat.Series) stat.Series {
	b1 := S1.Ref(1).Lt(S2.Ref(1))
	b2 := S1.Gte(S2)

	return b1.And(b2)
}

func V2CROSS(S1, S2 []stat.DType) []bool {
	r1 := S1
	r2 := S2
	r11 := REF2(S1, 1)
	r12 := REF2(S2, 1)

	b1 := CompareLt(r11, r12)
	b2 := CompareGte(r1, r2)

	c := vek.And(b1, b2)
	return c
}

func V1CROSS(S1, S2 stat.Series) []bool {
	r1 := S1.DTypes()
	r2 := S2.DTypes()
	r11 := S1.Ref(1).DTypes()
	r12 := S2.Ref(1).DTypes()

	b1 := CompareLt(r11, r12)
	b2 := CompareGte(r1, r2)

	c := vek.And(b1, b2)
	return c
}
