package formula

import (
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/pandas"
)

// LAST LAST(X,A,B):持续存在.
//
//	A 支持序列化, B不支持
//	例如:
//	LAST(CLOSE>OPEN,10,5)
//	表示从前10日到前5日内一直阳线
//	若A为0,表示从第一天开始,B为0,表示到最后日止
func LAST(X pandas.Series, A, B int) pandas.Series {
	s := X.Rolling(A + 1).Apply(func(S pandas.Series, N num.DType) num.DType {
		s := S.DTypes()
		s = api.Reverse(s)
		T := s[B:]
		n := len(T)
		for _, v := range T {
			if v != 0 {
				n--
			}
		}
		if n == 0 {
			return 1
		} else {
			return 0
		}
	})
	d := s.Values().([]num.DType)
	num.Fill(d, 1, true)
	return s
}
