package pandas

import (
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/num"
	"strings"
)

// NDFrame series多属性封装实现
type NDFrame struct {
	typ      Type   // values元素类型
	rows     int    // 行数
	nilCount int    // nil和nan的元素有多少, 这种统计在bool和int64类型中不会大于0, 只对float64及string有效
	name     string // 名称
	data     any    // for Vector
}

func (this *NDFrame) String() string {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	return s.String()
}

func (this *NDFrame) Name() string {
	this.name = defaultSeriesName(this.name)
	return this.name
}

func (this *NDFrame) Rename(name string) {
	this.name = strings.TrimSpace(name)
}

func (this *NDFrame) Type() Type {
	return this.typ
}

// 转成vector
func (this *NDFrame) asSeries() (Series, bool) {
	s, ok := this.data.(Series)
	return s, ok
}

func (this *NDFrame) Values() any {
	//return this.data
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	return s.Values()
}

func (this *NDFrame) toSeries(v Series) Series {
	return &NDFrame{
		typ:      v.Type(),
		rows:     v.Len(),
		nilCount: 0,
		name:     this.Name(),
		data:     v,
	}
}

func (this *NDFrame) NaN() any {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	return s.NaN()
}

func (this *NDFrame) Float32s() []float32 {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	return s.Float32s()
}

func (this *NDFrame) Float64s() []float64 {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	return s.Float64s()
}

func (this *NDFrame) DTypes() []num.DType {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	return s.DTypes()
}

func (this *NDFrame) Ints() []int {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	return s.Ints()
}

func (this *NDFrame) Int32s() []int32 {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	return s.Int32s()
}

func (this *NDFrame) Int64s() []int64 {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	return s.Int64s()
}

func (this *NDFrame) Strings() []string {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	return s.Strings()
}

func (this *NDFrame) Bools() []bool {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	return s.Bools()
}

func (this *NDFrame) Len() int {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	return s.Len()
}

func (this *NDFrame) Less(i, j int) bool {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	return s.Less(i, j)
}

func (this *NDFrame) Swap(i, j int) {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	s.Swap(i, j)
}

func (this *NDFrame) Empty(t ...Type) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Empty(t...)
	return this.toSeries(v)
}

func (this *NDFrame) Copy() Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Copy()
	return this.toSeries(v)
}

func (this *NDFrame) Reverse() Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Reverse()
	return this.toSeries(v)
}

func (this *NDFrame) Select(r api.ScopeLimit) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Select(r)
	return this.toSeries(v)
}

func (this *NDFrame) Append(values ...any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Append(values...)
	return this.toSeries(v)
}

func (this *NDFrame) Concat(x Series) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Concat(x)
	return this.toSeries(v)
}

func (this *NDFrame) Records(round ...bool) []string {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Records(round...)
	return v
}

func (this *NDFrame) IndexOf(index int, opt ...any) any {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.IndexOf(index, opt...)
	return v
}

func (this *NDFrame) Subset(start, end int, opt ...any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Subset(start, end, opt...)
	return this.toSeries(v)
}

func (this *NDFrame) Repeat(x any, repeats int) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Repeat(x, repeats)
	return this.toSeries(v)
}

func (this *NDFrame) FillNa(x any, inplace bool) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.FillNa(x, inplace)
	return this.toSeries(v)
}

func (this *NDFrame) Ref(periods any) (s Series) {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Ref(periods)
	return this.toSeries(v)
}

func (this *NDFrame) Shift(periods int) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Shift(periods)
	return this.toSeries(v)
}

func (this *NDFrame) Rolling(param any) RollingAndExpandingMixin {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Rolling(param)
	return v
}

func (this *NDFrame) Apply(f func(idx int, v any)) {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	s.Apply(f)
}

func (this *NDFrame) Apply2(f func(idx int, v any) any, args ...bool) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Apply2(f, args...)
	return this.toSeries(v)
}

func (this *NDFrame) Logic(f func(idx int, v any) bool) []bool {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Logic(f)
	return v
}

func (this *NDFrame) EWM(alpha EW) ExponentialMovingWindow {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.EWM(alpha)
	return v
}

func (this *NDFrame) Mean() num.DType {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Mean()
	return v
}

func (this *NDFrame) StdDev() num.DType {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.StdDev()
	return v
}

func (this *NDFrame) Max() any {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Max()
	return v
}

func (this *NDFrame) ArgMax() int {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.ArgMax()
	return v
}

func (this *NDFrame) Min() any {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Min()
	return v
}

func (this *NDFrame) ArgMin() int {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.ArgMin()
	return v
}

func (this *NDFrame) Diff(param any) (s Series) {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Diff(param)
	return this.toSeries(v)
}

func (this *NDFrame) Std() num.DType {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Std()
	return v
}

func (this *NDFrame) Sum() num.DType {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Sum()
	return v
}

func (this *NDFrame) Add(x any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Add(x)
	return this.toSeries(v)
}

func (this *NDFrame) Sub(x any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Sub(x)
	return this.toSeries(v)
}

func (this *NDFrame) Mul(x any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Mul(x)
	return this.toSeries(v)
}

func (this *NDFrame) Div(x any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Div(x)
	return this.toSeries(v)
}

func (this *NDFrame) Eq(x any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Eq(x)
	return this.toSeries(v)
}

func (this *NDFrame) Neq(x any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Neq(x)
	return this.toSeries(v)
}

func (this *NDFrame) Gt(x any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Gt(x)
	return this.toSeries(v)
}

func (this *NDFrame) Gte(x any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Gte(x)
	return this.toSeries(v)
}

func (this *NDFrame) Lt(x any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Lt(x)
	return this.toSeries(v)
}

func (this *NDFrame) Lte(x any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Lte(x)
	return this.toSeries(v)
}

func (this *NDFrame) And(x any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.And(x)
	return this.toSeries(v)
}

func (this *NDFrame) Or(x any) Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Or(x)
	return this.toSeries(v)
}

func (this *NDFrame) Not() Series {
	s, ok := this.asSeries()
	if !ok {
		panic(num.TypeError(this.data))
	}
	v := s.Not()
	return this.toSeries(v)
}
