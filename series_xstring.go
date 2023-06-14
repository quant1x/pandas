package pandas

import (
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/pandas/stat"
)

type SeriesString struct {
	stat.NDArray[string]
	name string
}

func (self SeriesString) Name() string {
	return self.name
}

func (self SeriesString) Rename(name string) {
	self.name = name
}

func (self SeriesString) Type() stat.Type {
	return self.NDArray.Type()
}

func (self SeriesString) Values() any {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) NaN() any {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) Floats() []float32 {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) DTypes() []stat.DType {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) Ints() []stat.Int {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) Len() int {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) Less(i, j int) bool {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) Swap(i, j int) {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) Empty(t ...stat.Type) stat.Series {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) Copy() stat.Series {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) Records() []string {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) Subset(start, end int, opt ...any) stat.Series {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) Repeat(x any, repeats int) stat.Series {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) Shift(periods int) stat.Series {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) Rolling(param any) stat.RollingAndExpandingMixin {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) Mean() stat.DType {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) StdDev() stat.DType {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) FillNa(v any, inplace bool) stat.Series {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) Max() any {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) Min() any {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) Select(r api.ScopeLimit) stat.Series {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) Append(values ...any) stat.Series {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) Apply(f func(idx int, v any)) {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) Logic(f func(idx int, v any) bool) []bool {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) Diff(param any) (s stat.Series) {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) Ref(param any) (s stat.Series) {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) Std() stat.DType {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) Sum() stat.DType {
	//TODO implement me
	panic("implement me")
}

func (self SeriesString) EWM(alpha stat.EW) stat.ExponentialMovingWindow {
	//TODO implement me
	panic("implement me")
}
