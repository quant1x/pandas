package pandas

import (
	"gitee.com/quant1x/pandas/stat"
	"golang.org/x/exp/slices"
)

func (self DataFrame) align(ss ...stat.Series) []stat.Series {
	defaultValue := []stat.Series{}
	sLen := len(ss)
	if sLen == 0 {
		return defaultValue
	}
	ls := make([]float32, sLen)
	for i, v := range ss {
		ls[i] = float32(v.Len())
	}

	maxLength := stat.Max(ls)
	if maxLength <= 0 {
		return defaultValue
	}
	cols := make([]stat.Series, sLen)
	for i, v := range ss {
		vt := v.Type()
		vn := v.Name()
		vs := v.Values()
		// 声明any的ns变量用于接收逻辑分支的输出
		// 切片数据不能直接对齐, 需要根据类型指定Nil和NaN默认值
		var ns any
		if vt == stat.SERIES_TYPE_BOOL {
			ns = stat.Align(vs.([]bool), stat.Nil2Bool, int(maxLength))
		} else if vt == stat.SERIES_TYPE_INT64 {
			ns = stat.Align(vs.([]int64), stat.Nil2Int64, int(maxLength))
		} else if vt == stat.SERIES_TYPE_STRING {
			ns = stat.Align(vs.([]string), stat.Nil2String, int(maxLength))
		} else if vt == stat.SERIES_TYPE_FLOAT32 {
			ns = stat.Align(vs.([]float32), stat.Nil2Float32, int(maxLength))
		} else if vt == stat.SERIES_TYPE_FLOAT64 {
			ns = stat.Align(vs.([]float64), stat.Nil2Float64, int(maxLength))
		}
		cols[i] = NewSeries(vt, vn, ns)
	}
	return cols
}

// Join 默认右连接, 加入一个series
func (self DataFrame) Join(S ...stat.Series) DataFrame {
	sNum := len(S)
	if sNum == 0 {
		return self
	}

	cols := slices.Clone(self.columns)
	for _, series := range S {
		if series.Len() < 0 {
			continue
		}
		cols = append(cols, series)
	}

	cols = self.align(cols...)
	df := NewDataFrame(cols...)
	self = df
	return self
}
