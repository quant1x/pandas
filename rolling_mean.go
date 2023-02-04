package pandas

// Mean returns the rolling mean.
func (r RollingAndExpandingMixin) Mean() (s Series) {
	var d []float64
	for _, block := range r.getBlocks() {
		d = append(d, block.Mean())
	}
	s = NewSeries(SERIES_TYPE_FLOAT, r.series.Name(), d)
	return
}
