package pandas

import (
	"gitee.com/quant1x/pandas/stat"
	"github.com/viterin/vek"
)

func (r RollingAndExpandingMixin) Count() (s Series) {
	if r.series.Type() != SERIES_TYPE_BOOL {
		panic("不支持非bool序列")
	}
	values := make([]stat.DType, r.series.Len())
	for i, block := range r.getBlocks() {
		if block.Len() == 0 {
			values[i] = 0
			continue
		}
		bs := block.Values().([]bool)
		values[i] = stat.DType(vek.Count(bs))
	}
	s = NewSeries(SERIES_TYPE_DTYPE, r.series.Name(), values)
	return
}
