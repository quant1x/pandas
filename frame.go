package pandas

import "github.com/WinPooh32/series"

type NDFrame struct {
	series.Data
}

func New(se ...Series) DataFrame {
	return NewFrame(se...)
}

func MakeValues(values []DType) NDFrame {
	v_len := len(values)
	indexs := make([]int64, v_len)
	for i := 0; i < v_len; i++ {
		indexs[i] = int64(i)
	}
	freq := int64(1)
	data := series.MakeData(freq, indexs, values)
	return NDFrame{data}
}

type Window = series.Window

func (self NDFrame) Rolling(n int) Window {
	return self.Data.Rolling(n)
}
