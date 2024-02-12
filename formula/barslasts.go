package formula

import "gitee.com/quant1x/pandas/stat"

// BARSLASTS 倒数第N次成立时距今的周期数.
//
//	用法:
//	BARSLASTS(X,N):X倒数第N满足到现在的周期数,N支持变量
//	go实现暂时不支持N为序列, 意义不大
func BARSLASTS(S stat.Series, N int) stat.Series {
	v := __bars_lasts(S, N)
	x := stat.ToSeries(v...)
	return x
}

func __bars_lasts(S stat.Series, N int) []stat.DType {
	v := __bars_last(S)
	x := stat.ToSeries(v...)
	m := x
	for i := 1; i < N; i++ {
		//第二次位置:=REF(BARSLAST(条件),第一次位置+1)+第一次位置+1;{倒数第二次条件距今位置}
		pos := x.Add(1)
		m = REF(x, pos).Add(pos)
		x = m
	}
	return m.DTypes()
}

func __bars_last(S stat.Series) []stat.DType {
	fs := S.DTypes()
	as := stat.Repeat[stat.DType](1, S.Len())
	bs := stat.Repeat[stat.DType](0, S.Len())
	x := stat.Where(fs, as, bs)
	M := []stat.DType{0}
	M = append(M, x...)
	for i := 1; i < len(M); i++ {
		if int(M[i]) != 0 {
			M[i] = 0
		} else {
			M[i] = M[i-1] + 1
		}
	}
	return M[1:]
}
