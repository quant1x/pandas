package pandas

import (
	"gitee.com/quant1x/pandas/stat"
)

func (this *NDFrame) Ref(param any) (s stat.Series) {

	switch values := this.values.(type) {
	case []bool:
		d := stat.Shift[bool](values, param)
		return NewSeries(stat.SERIES_TYPE_BOOL, this.Name(), d)
	case []string:
		d := stat.Shift[string](values, param)
		return NewSeries(stat.SERIES_TYPE_STRING, this.Name(), d)
	case []int64:
		d := stat.Shift[int64](values, param)
		return NewSeries(stat.SERIES_TYPE_INT32, this.Name(), d)
	case []float32:
		d := stat.Shift[float32](values, param)
		return NewSeries(stat.SERIES_TYPE_FLOAT32, this.Name(), d)
	default: //case []float64:
		d := stat.Shift[float64](values.([]float64), param)
		return NewSeries(stat.SERIES_TYPE_FLOAT64, this.Name(), d)
	}
}
