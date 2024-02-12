package stat

import "gitee.com/quant1x/num"

func (r RollingAndExpandingMixin) Sum() Series {
	var d []num.DType
	for _, block := range r.GetBlocks() {
		d = append(d, block.Sum())
	}
	//s := pandas.NewSeries(SERIES_TYPE_DTYPE, r.series.Name(), d)
	s := r.Series.Empty(SERIES_TYPE_DTYPE)
	s.Rename(r.Series.Name())
	s = s.Append(d)
	return s
}
