package pandas

import (
	"slices"

	"gitee.com/quant1x/num"
)

func (this DataFrame) align(ss ...Series) []Series {
	defaultValue := []Series{}
	sLen := len(ss)
	if sLen == 0 {
		return defaultValue
	}
	ls := make([]float32, sLen)
	for i, v := range ss {
		ls[i] = float32(v.Len())
	}

	maxLength := num.Max(ls)
	if maxLength <= 0 {
		return defaultValue
	}
	cols := make([]Series, sLen)
	for i, v := range ss {
		vt := v.Type()
		vn := v.Name()
		vs := v.Values()
		// 声明any的ns变量用于接收逻辑分支的输出
		// 切片数据不能直接对齐, 需要根据类型指定Nil和NaN默认值
		var ns any
		if vt == SERIES_TYPE_BOOL {
			ns = num.Align(vs.([]bool), num.Nil2Bool, int(maxLength))
		} else if vt == SERIES_TYPE_INT64 {
			ns = num.Align(vs.([]int64), num.Nil2Int64, int(maxLength))
		} else if vt == SERIES_TYPE_STRING {
			ns = num.Align(vs.([]string), num.Nil2String, int(maxLength))
		} else if vt == SERIES_TYPE_FLOAT32 {
			ns = num.Align(vs.([]float32), num.Float32NaN(), int(maxLength))
		} else if vt == SERIES_TYPE_FLOAT64 {
			ns = num.Align(vs.([]float64), num.Float64NaN(), int(maxLength))
		}
		cols[i] = NewSeriesWithType(vt, vn, ns)
	}
	return cols
}

// Join 默认右连接, 加入一个series
func (this DataFrame) Join(S ...Series) DataFrame {
	sNum := len(S)
	if sNum == 0 {
		return this
	}

	cols := slices.Clone(this.columns)
	for _, series := range S {
		if series.Len() < 0 {
			continue
		}
		cols = append(cols, series)
	}

	cols = this.align(cols...)
	df := NewDataFrame(cols...)
	this = df
	return this
}
