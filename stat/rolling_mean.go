package stat

import "gitee.com/quant1x/num"

// Mean returns the rolling mean.
func (r RollingAndExpandingMixin) Mean() (s Series) {
	var d []num.DType
	for _, block := range r.GetBlocks() {
		d = append(d, block.Mean())
	}
	s = r.Series.Empty(SERIES_TYPE_DTYPE)
	s.Rename(r.Series.Name())
	s = s.Append(d)
	return
}
