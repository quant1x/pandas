package pandas

func (r RollingAndExpandingMixin) Sum() Series {
	var d []float64
	for _, block := range r.getBlocks() {
		d = append(d, block.Sum())
	}
	s := NewSeries(SERIES_TYPE_FLOAT64, r.series.Name(), d)
	return s
}
