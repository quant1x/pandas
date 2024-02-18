package pandas

import (
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/num"
	"strings"
)

// NDArray series多属性封装实现
type NDArray struct {
	typ      Type   // values元素类型
	rows     int    // 行数
	nilCount int    // nil和nan的元素有多少, 这种统计在bool和int64类型中不会大于0, 只对float64及string有效
	name     string // 名称
	data     any    // for Vector
}

func (this *NDArray) String() string {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	return s.String()
}

func (this *NDArray) Name() string {
	this.name = defaultSeriesName(this.name)
	return this.name
}

func (this *NDArray) Rename(name string) {
	this.name = strings.TrimSpace(name)
}

func (this *NDArray) Type() Type {
	return this.typ
}

// 转成vector
func (this *NDArray) asSeries() (Series, bool) {
	s, ok := this.data.(Series)
	return s, ok
}

func (this *NDArray) Values() any {
	//return this.data
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	return s.Values()
}

func (this *NDArray) toSeries(v Series) Series {
	return &NDArray{
		typ:      v.Type(),
		rows:     v.Len(),
		nilCount: 0,
		name:     this.Name(),
		data:     v,
	}
}

func (this *NDArray) NaN() any {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	return s.NaN()
}

func (this *NDArray) Float32s() []float32 {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	return s.Float32s()
}

func (this *NDArray) Float64s() []float64 {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	return s.Float64s()
}

func (this *NDArray) DTypes() []num.DType {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	return s.DTypes()
}

func (this *NDArray) Ints() []int {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	return s.Ints()
}

func (this *NDArray) Int32s() []int32 {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	return s.Int32s()
}

func (this *NDArray) Int64s() []int64 {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	return s.Int64s()
}

func (this *NDArray) Strings() []string {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	return s.Strings()
}

func (this *NDArray) Bools() []bool {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	return s.Bools()
}

func (this *NDArray) Len() int {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	return s.Len()
}

func (this *NDArray) Less(i, j int) bool {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	return s.Less(i, j)
}

func (this *NDArray) Swap(i, j int) {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	s.Swap(i, j)
}

func (this *NDArray) Empty(t ...Type) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Empty(t...)
	return this.toSeries(v)
}

func (this *NDArray) Copy() Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Copy()
	return this.toSeries(v)
}

func (this *NDArray) Reverse() Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Reverse()
	return this.toSeries(v)
}

func (this *NDArray) Select(r api.ScopeLimit) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Select(r)
	return this.toSeries(v)
}

func (this *NDArray) Append(values ...any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Append(values...)
	return this.toSeries(v)
}

func (this *NDArray) Concat(x Series) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Concat(x)
	return this.toSeries(v)
}

func (this *NDArray) Records(round ...bool) []string {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Records(round...)
	return v
}

func (this *NDArray) IndexOf(index int, opt ...any) any {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.IndexOf(index, opt...)
	return v
}

func (this *NDArray) Subset(start, end int, opt ...any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Subset(start, end, opt...)
	return this.toSeries(v)
}

func (this *NDArray) Repeat(x any, repeats int) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Repeat(x, repeats)
	return this.toSeries(v)
}

func (this *NDArray) FillNa(x any, inplace bool) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.FillNa(x, inplace)
	return this.toSeries(v)
}

func (this *NDArray) Ref(periods any) (s Series) {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Ref(periods)
	return this.toSeries(v)
}

func (this *NDArray) Shift(periods int) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Shift(periods)
	return this.toSeries(v)
}

func (this *NDArray) Rolling(param any) RollingAndExpandingMixin {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Rolling(param)
	return v
}

func (this *NDArray) Apply(f func(idx int, v any)) {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	s.Apply(f)
}

func (this *NDArray) Apply2(f func(idx int, v any) any, args ...bool) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Apply2(f, args...)
	return this.toSeries(v)
}

func (this *NDArray) Logic(f func(idx int, v any) bool) []bool {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Logic(f)
	return v
}

func (this *NDArray) EWM(alpha EW) ExponentialMovingWindow {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.EWM(alpha)
	return v
}

func (this *NDArray) Mean() num.DType {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Mean()
	return v
}

func (this *NDArray) StdDev() num.DType {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.StdDev()
	return v
}

func (this *NDArray) Max() any {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Max()
	return v
}

func (this *NDArray) ArgMax() int {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.ArgMax()
	return v
}

func (this *NDArray) Min() any {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Min()
	return v
}

func (this *NDArray) ArgMin() int {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.ArgMin()
	return v
}

func (this *NDArray) Diff(param any) (s Series) {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Diff(param)
	return this.toSeries(v)
}

func (this *NDArray) Std() num.DType {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Std()
	return v
}

func (this *NDArray) Sum() num.DType {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Sum()
	return v
}

func (this *NDArray) Add(x any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Add(x)
	return this.toSeries(v)
}

func (this *NDArray) Sub(x any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Sub(x)
	return this.toSeries(v)
}

func (this *NDArray) Mul(x any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Mul(x)
	return this.toSeries(v)
}

func (this *NDArray) Div(x any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Div(x)
	return this.toSeries(v)
}

func (this *NDArray) Eq(x any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Eq(x)
	return this.toSeries(v)
}

func (this *NDArray) Neq(x any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Neq(x)
	return this.toSeries(v)
}

func (this *NDArray) Gt(x any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Gt(x)
	return this.toSeries(v)
}

func (this *NDArray) Gte(x any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Gte(x)
	return this.toSeries(v)
}

func (this *NDArray) Lt(x any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Lt(x)
	return this.toSeries(v)
}

func (this *NDArray) Lte(x any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Lte(x)
	return this.toSeries(v)
}

func (this *NDArray) And(x any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.And(x)
	return this.toSeries(v)
}

func (this *NDArray) Or(x any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Or(x)
	return this.toSeries(v)
}

func (this *NDArray) Not() Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Not()
	return this.toSeries(v)
}
