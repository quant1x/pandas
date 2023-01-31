package pandas

import (
	"github.com/huandu/go-clone"
	"math"
	"sync"
)

type SeriesFrame struct {
	valFormatter ValueToStringFormatter
	lock         sync.RWMutex
	name         string
	nilCount     int
	//elements     any
}

func NewSeries(t Type, name string, vals ...interface{}) *Series {
	var series Series
	if t == SERIES_TYPE_BOOL {
		series = NewSeriesBool(name, vals...)
	} else if t == SERIES_TYPE_INT {
		series = NewSeriesInt64(name, vals...)
	} else if t == SERIES_TYPE_STRING {
		series = NewSeriesString(name, vals...)
	} else {
		series = NewSeriesFloat64(name, vals...)
	}
	return &series
}

func Shift[T ~int64 | ~float64 | ~bool | ~string](s *Series, periods int, cbNan func() T) *Series {
	var d Series
	d = clone.Clone(*s).(Series)
	if periods == 0 {
		return &d
	}

	values := d.Values().([]T)

	var (
		naVals []T
		dst    []T
		src    []T
	)

	if shlen := int(math.Abs(float64(periods))); shlen < len(values) {
		if periods > 0 {
			naVals = values[:shlen]
			dst = values[shlen:]
			src = values
		} else {
			naVals = values[len(values)-shlen:]
			dst = values[:len(values)-shlen]
			src = values[shlen:]
		}

		copy(dst, src)
	} else {
		naVals = values
	}
	for i := range naVals {
		naVals[i] = cbNan()
	}
	_ = naVals
	return &d
}

func (self *SeriesFrame) Name() string {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesFrame) Rename(n string) {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesFrame) Type() Type {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesFrame) Len() int {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesFrame) Values() any {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesFrame) Empty() Series {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesFrame) Records() []string {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesFrame) Subset(start, end int) *Series {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesFrame) Repeat(x any, repeats int) *Series {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesFrame) Shift(periods int) *Series {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesFrame) Rolling(window int) RollingWindow {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesFrame) Mean() float64 {
	//TODO implement me
	panic("implement me")
}

func (self *SeriesFrame) StdDev() float64 {
	//TODO implement me
	panic("implement me")
}
