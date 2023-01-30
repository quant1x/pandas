package pandas

import (
	"sync"
)

type SeriesFrame struct {
	valFormatter ValueToStringFormatter
	lock         sync.RWMutex
	name         string
	nilCount     int
}

func NewSeriesFrame(name string) SeriesFrame {
	return SeriesFrame{
		name: name,
	}
}
