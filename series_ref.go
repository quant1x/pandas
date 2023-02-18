package pandas

import (
	"gitee.com/quant1x/pandas/stat"
)

func (self *NDFrame) Ref(param any) (s stat.Series) {

	switch values := self.values.(type) {
	case []bool:
		d := stat.Shift[bool](values, param)
		return NewSeries(stat.SERIES_TYPE_BOOL, self.Name(), d)
	case []string:
		d := stat.Shift[string](values, param)
		return NewSeries(stat.SERIES_TYPE_STRING, self.Name(), d)
	case []int64:
		d := stat.Shift[int64](values, param)
		return NewSeries(stat.SERIES_TYPE_INT32, self.Name(), d)
	case []float32:
		d := stat.Shift[float32](values, param)
		return NewSeries(stat.SERIES_TYPE_FLOAT32, self.Name(), d)
	default: //case []float64:
		d := stat.Shift[float64](values.([]float64), param)
		return NewSeries(stat.SERIES_TYPE_FLOAT64, self.Name(), d)
	}
}
