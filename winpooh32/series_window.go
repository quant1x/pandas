package winpooh32

import (
	"sort"

	"gitee.com/quant1x/pandas/winpooh32/math"
)

type BaseWindow struct {
	len  int
	data Series
}

func (w BaseWindow) Sum() Series {
	return w.Apply(Sum)
}

func (w BaseWindow) Mean() Series {
	return w.Apply(Mean)
}

func (w BaseWindow) Min() Series {
	return w.Apply(Min)
}

func (w BaseWindow) Max() Series {
	return w.Apply(Max)
}

func (w BaseWindow) Skew(ma Series) Series {
	return w.Apply(Skew)
}

func (w BaseWindow) Median() Series {
	return w.applyMedian()
}

func (w BaseWindow) Variance(ma Series, ddof int) Series {
	return w.applyVar(Variance, ma, ddof)
}

func (w BaseWindow) Std(ma Series, ddof int) Series {
	return w.applyVar(Std, ma, ddof)
}

func (w BaseWindow) Apply(agg AggregateFunc) Series {
	var (
		clone  = w.data.Clone()
		values = clone.Values()
		period = w.len
	)

	for i := 0; i < w.len-1; i++ {
		values[i] = math.NaN()
	}

	w.data.RollData(period, func(l int, r int) {
		slice := w.data.Slice(l, r)
		values[r-1] = agg(slice)
	})

	return clone
}

func (w BaseWindow) applyVar(varfn func(data Series, mean DType, ddof int) DType, ma Series, ddof int) Series {
	var (
		clone  = w.data.Clone()
		values = clone.Values()
		period = w.len
	)

	total := period - 1

	for i := total; i < len(values); i++ {
		p := i + 1
		v := w.data.Slice(p-period, p)
		values[i] = varfn(v, ma.values[p-1], ddof)
	}

	for i := 0; i < total; i++ {
		values[i] = math.NaN()
	}

	return clone
}

func (w BaseWindow) applyMedian() Series {
	var (
		clone  = w.data.Clone()
		values = clone.Values()
		tmp    = make([]DType, 0, w.len)
		period = w.len
	)

	for i := 0; i < w.len-1; i++ {
		values[i] = math.NaN()
	}

	w.data.RollData(period, func(l int, r int) {
		slice := w.data.Slice(l, r)

		tmp = append(tmp[:0], slice.values...)
		sort.Sort(DTypeSlice(tmp))

		values[r-1] = Median(Series{values: tmp})
	})

	return clone
}
