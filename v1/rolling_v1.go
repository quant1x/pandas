package v1

// RollingWindowV1 is used for rolling window calculations.
// Deprecated: 使用RollingAndExpandingMixin
type RollingWindowV1 struct {
	window int
	series Series
}

// RollingV1 滑动窗口
// Deprecated: 使用RollingAndExpandingMixin
func (self *NDFrame) RollingV1(window int) RollingWindowV1 {
	return RollingWindowV1{
		window: window,
		series: self,
	}
}

func (r RollingWindowV1) getBlocks() (blocks []Series) {
	for i := 1; i <= r.series.Len(); i++ {
		if i < r.window {
			blocks = append(blocks, r.series.Empty())
			continue
		}

		start := i - r.window
		end := i
		blocks = append(blocks, r.series.Subset(start, end))
	}

	return
}

// Mean returns the rolling mean.
func (r RollingWindowV1) Mean() (s Series) {
	var d []float64
	for _, block := range r.getBlocks() {
		d = append(d, block.Mean())
	}
	s = NewSeriesFloat64(r.series.Name(), d)
	return
}

// StdDev returns the rolling mean.
func (r RollingWindowV1) StdDev() (s Series) {
	var d []float64
	for _, block := range r.getBlocks() {
		d = append(d, block.StdDev())
	}
	s = NewSeriesFloat64(r.series.Name(), d)

	return
}

func (r RollingWindowV1) Max() any {
	var fs []float64
	var is []int64
	var ss []string
	for _, block := range r.getBlocks() {
		//d = append(d, block.Max())
		v := block.Max()
		switch val := v.(type) {
		case float64:
			fs = append(fs, val)
		case int64:
			is = append(is, val)
		case string:
			ss = append(ss, val)
		}
	}
	if len(ss) > 0 {
		return ss
	} else if len(is) > 0 {
		return is
	} else {
		return fs
	}
}

func (r RollingWindowV1) Min() any {
	var fs []float64
	var is []int64
	var ss []string
	for _, block := range r.getBlocks() {
		v := block.Min()
		switch val := v.(type) {
		case float64:
			fs = append(fs, val)
		case int64:
			is = append(is, val)
		case string:
			ss = append(ss, val)
		}
	}
	if len(ss) > 0 {
		return ss
	} else if len(is) > 0 {
		return is
	} else {
		return fs
	}
}
