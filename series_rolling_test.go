package pandas

import (
	"gitee.com/quant1x/num"
	"gitee.com/quant1x/num/labs"
	"slices"
	"testing"
)

func TestRollingAndExpandingMixin_GetBlocks(t *testing.T) {
	type fields struct {
		Window num.Window[num.DType]
		Series Series
	}
	tests := []struct {
		name       string
		fields     fields
		wantBlocks []Series
	}{
		{
			name: "string",
			fields: fields{
				Window: num.Window[num.DType]{
					V: []num.DType{2, 2, 2, 2},
				},
				Series: ToSeries[string]("1", "2", "3", "a"),
			},
			wantBlocks: []Series{
				ToSeries[string](),
				ToSeries[string]("1", "2"),
				ToSeries[string]("2", "3"),
				ToSeries[string]("3", "a"),
			},
		},
		{
			name: "string-const",
			fields: fields{
				Window: num.Window[num.DType]{
					V: []num.DType{2, 2, 2, 2},
				},
				Series: ToSeries[string]("1", "2", "2", "3"),
			},
			wantBlocks: []Series{
				ToSeries[string](),
				ToSeries[string]("1", "2"),
				ToSeries[string]("2", "2"),
				ToSeries[string]("2", "3"),
			},
		},
		{
			name: "string-vector-and-const",
			fields: fields{
				Window: num.Window[num.DType]{
					V: []num.DType{2, 2, 2},
					C: 2,
				},
				Series: ToSeries[string]("1", "2", "3", "a"),
			},
			wantBlocks: []Series{
				ToSeries[string](),
				ToSeries[string]("1", "2"),
				ToSeries[string]("2", "3"),
				ToSeries[string]("3", "a"),
			},
		},
		{
			name: "string-vector-error",
			fields: fields{
				Window: num.Window[num.DType]{
					V: []num.DType{2, 2, 2},
					C: 2,
				},
				Series: ToSeries[string]("1", "2", "3", "a"),
			},
			wantBlocks: []Series{
				ToSeries[string](),
				ToSeries[string]("1", "2"),
				ToSeries[string]("2", "3"),
				ToSeries[string]("3", "a"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RollingAndExpandingMixin{
				Window: tt.fields.Window,
				Series: tt.fields.Series,
			}
			if gotBlocks := r.GetBlocks(); !labs.DeepEqual(gotBlocks, tt.wantBlocks) {
				t.Errorf("GetBlocks() = %v, want %v", gotBlocks, tt.wantBlocks)
			}
		})
	}
}

func TestRollingAndExpandingMixin_Count(t *testing.T) {
	type fields struct {
		Window num.Window[num.DType]
		Series Series
	}
	tests := []struct {
		name   string
		fields fields
		wantS  Series
	}{
		{
			name: "float32",
			fields: fields{
				Window: num.Window[num.DType]{
					V: []num.DType{2, 2, 2},
				},
				Series: ToSeries[float32](1, 2, 3),
			},
			wantS: ToSeries[num.DType](num.Float64NaN(), 2, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RollingAndExpandingMixin{
				Window: tt.fields.Window,
				Series: tt.fields.Series,
			}
			if gotS := r.Count(); !labs.DeepEqual(gotS.Values(), tt.wantS.Values()) {
				t.Errorf("Count() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestRollingAndExpandingMixin_Min(t *testing.T) {
	type fields struct {
		Window num.Window[num.DType]
		Series Series
	}
	tests := []struct {
		name   string
		fields fields
		want   Series
	}{
		{
			name: "float32",
			fields: fields{
				Window: num.Window[num.DType]{
					V: []num.DType{2, 2, 2},
				},
				Series: ToSeries[float32](1, 2, 3),
			},
			want: ToSeries[float32](num.Float32NaN(), 1, 2),
		},
		{
			name: "float64",
			fields: fields{
				Window: num.Window[num.DType]{
					V: []num.DType{2, 2, 2},
				},
				Series: ToSeries[float64](1, 2, 3),
			},
			want: ToSeries[float64](num.Float64NaN(), 1, 2),
		},
		{
			name: "string",
			fields: fields{
				Window: num.Window[num.DType]{
					V: []num.DType{2, 2, 2, 2},
				},
				Series: ToSeries[string]("1", "2", "2", "3"),
			},
			want: ToSeries[string]("NaN", "1", "2", "2"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RollingAndExpandingMixin{
				Window: tt.fields.Window,
				Series: tt.fields.Series,
			}
			if got := r.Min(); !labs.DeepEqual(got.Values(), tt.want.Values()) {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRollingAndExpandingMixin_Mean(t *testing.T) {
	type fields struct {
		Window num.Window[num.DType]
		Series Series
	}
	tests := []struct {
		name   string
		fields fields
		want   Series
	}{
		{
			name: "float64",
			fields: fields{
				Window: num.Window[num.DType]{
					C: 5,
				},
				Series: ToSeries[float32](1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
			},
			want: ToSeries[float64](num.Float64NaN(), num.Float64NaN(), num.Float64NaN(), num.Float64NaN(), 3, 4, 5, 6, 7, 8),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RollingAndExpandingMixin{
				Window: tt.fields.Window,
				Series: tt.fields.Series,
			}
			if got := r.Mean(); !labs.DeepEqual(got.Values(), tt.want.Values()) {
				t.Errorf("Mean() = %v, want %v", got, tt.want)
			}
		})
	}
}

const (
	rollingAndExpandingMixinPeriod = 5
)

func BenchmarkRollingAndExpandingMixin_Sum_init(b *testing.B) {
	testDataOnce.Do(initTestData)
}

func BenchmarkRollingAndExpandingMixin_Sum_release(b *testing.B) {
	testDataOnce.Do(initTestData)
	f64s := slices.Clone(testDataFloat64)
	s := SliceToSeries(f64s)
	for i := 0; i < b.N; i++ {
		s.Rolling(rollingAndExpandingMixinPeriod).Sum()
	}
}

func BenchmarkRollingAndExpandingMixin_Sum_v1(b *testing.B) {
	testDataOnce.Do(initTestData)
	f64s := slices.Clone(testDataFloat64)
	s := SliceToSeries(f64s)
	for i := 0; i < b.N; i++ {
		s.Rolling(rollingAndExpandingMixinPeriod).v1Sum()
	}
}

func BenchmarkRollingAndExpandingMixin_Sum_v2(b *testing.B) {
	testDataOnce.Do(initTestData)
	f64s := slices.Clone(testDataFloat64)
	s := SliceToSeries(f64s)
	for i := 0; i < b.N; i++ {
		s.Rolling(rollingAndExpandingMixinPeriod).v2Sum()
	}
}

func BenchmarkRollingAndExpandingMixin_Sum_v3(b *testing.B) {
	testDataOnce.Do(initTestData)
	f64s := slices.Clone(testDataFloat64)
	s := SliceToSeries(f64s)
	for i := 0; i < b.N; i++ {
		s.Rolling(rollingAndExpandingMixinPeriod).v3Sum()
	}
}

func BenchmarkRollingAndExpandingMixin_Std_init(b *testing.B) {
	testDataOnce.Do(initTestData)
}

func BenchmarkRollingAndExpandingMixin_Std_release(b *testing.B) {
	testDataOnce.Do(initTestData)
	f64s := slices.Clone(testDataFloat64)
	s := SliceToSeries(f64s)
	for i := 0; i < b.N; i++ {
		s.Rolling(rollingAndExpandingMixinPeriod).Std()
	}
}

func BenchmarkRollingAndExpandingMixin_Std_v1(b *testing.B) {
	testDataOnce.Do(initTestData)
	f64s := slices.Clone(testDataFloat64)
	s := SliceToSeries(f64s)
	for i := 0; i < b.N; i++ {
		s.Rolling(rollingAndExpandingMixinPeriod).v1Std()
	}
}

func BenchmarkRollingAndExpandingMixin_Std_v2(b *testing.B) {
	testDataOnce.Do(initTestData)
	f64s := slices.Clone(testDataFloat64)
	s := SliceToSeries(f64s)
	for i := 0; i < b.N; i++ {
		s.Rolling(rollingAndExpandingMixinPeriod).v2Std()
	}
}
