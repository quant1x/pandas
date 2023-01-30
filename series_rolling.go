package pandas

// RollingWindow is used for rolling window calculations.
type RollingWindow struct {
	window int
	series Series
}

func (r RollingWindow) getBlocks() (blocks []Series) {
	for i := 1; i <= r.series.Len(); i++ {
		if i < r.window {
			blocks = append(blocks, r.series.Empty())
			continue
		}

		start := i - r.window
		end := i
		blocks = append(blocks, *r.series.Subset(start, end))
	}

	return
}

// Mean returns the rolling mean.
func (r RollingWindow) Mean() (s Series) {
	var d []float64
	for _, block := range r.getBlocks() {
		d = append(d, block.Mean())
	}
	s = NewSeriesFloat64(r.series.Name(), d)
	return
}

// StdDev returns the rolling mean.
func (r RollingWindow) StdDev() (s Series) {
	var d []float64
	for _, block := range r.getBlocks() {
		d = append(d, block.StdDev())
	}
	s = NewSeriesFloat64(r.series.Name(), d)

	return
}
