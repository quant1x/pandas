package v1

import "gitee.com/quant1x/pandas/stat"

func (r RollingAndExpandingMixin) Sum() Series {
	var d []stat.DType
	for _, block := range r.getBlocks() {
		d = append(d, block.Sum())
	}
	s := NewSeries(SERIES_TYPE_DTYPE, r.series.Name(), d)
	return s
}
