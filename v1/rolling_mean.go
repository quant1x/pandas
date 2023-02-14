package v1

import "gitee.com/quant1x/pandas/stat"

// Mean returns the rolling mean.
func (r RollingAndExpandingMixin) Mean() (s Series) {
	var d []stat.DType
	for _, block := range r.getBlocks() {
		d = append(d, block.Mean())
	}
	s = NewSeries(SERIES_TYPE_DTYPE, r.series.Name(), d)
	return
}
