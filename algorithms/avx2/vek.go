package avx2

//
//import "github.com/viterin/vek"
//
//// 初始化 avx2
//// 可以参考另一个实现库 gonum.org/v1/gonum/stat
//func init() {
//	// 开启加速选项
//	vek.SetAcceleration(true)
//}
//
//func AddScalar(x []float64, a float64) { vek.AddNumber_Inplace(x, a) }
//func SubScalar(x []float64, a float64) { vek.SubNumber_Inplace(x, a) }
//func MulScalar(x []float64, a float64) { vek.MulNumber_Inplace(x, a) }
//func DivScalar(x []float64, a float64) { vek.DivNumber_Inplace(x, a) }
//
//func Add(x []float64, y []float64)     { vek.Add_Inplace(x, y) }
//func Sub(x []float64, y []float64)     { vek.Sub_Inplace(x, y) }
//func Mul(x []float64, y []float64)     { vek.Mul_Inplace(x, y) }
//func Div(x []float64, y []float64)     { vek.Div_Inplace(x, y) }
//func Minimum(x []float64, y []float64) { vek.Minimum_Inplace(x, y) }
//func Maximum(x []float64, y []float64) { vek.Maximum_Inplace(x, y) }
//func Pow(x []float64, y []float64)     { vek.Pow_Inplace(x, y) }
//
//func Sqrt(x []float64)  { vek.Sqrt_Inplace(x) }
//func Abs(x []float64)   { vek.Abs_Inplace(x) }
//func Round(x []float64) { vek.Round_Inplace(x) }
//func Ceil(x []float64)  { vek.Ceil_Inplace(x) }
//func Floor(x []float64) { vek.Floor_Inplace(x) }
//
//func Min(x []float64) float64    { return vek.Min(x) }
//func Max(x []float64) float64    { return vek.Max(x) }
//func Mean(x []float64) float64   { return vek.Mean(x) }
//func Median(x []float64) float64 { return vek.Median(x) }
//
//func Dot(x []float64, y []float64) float64 { return vek.Dot(x, y) }
//
//func ArgMin(x []float64) int { return vek.ArgMin(x) }
//func ArgMax(x []float64) int { return vek.ArgMax(x) }
//
//// func Repeat(dst []float64, a float64, n int) []float64 { return vek.Repeat_Into(dst, a, n) }
//func Repeat(a float64, n int) []float64            { return vek.Repeat(a, n) }
//func RepeatAll(dst []float64, a float64) []float64 { return vek.Repeat_Into(dst, a, len(dst)) }
//func ToInt64(dst []int64, x []float64)             { vek.ToInt64_Into(dst, x) }
//func ToInt32(dst []int32, x []float64)             { vek.ToInt32_Into(dst, x) }
//func ToFloat64(dst []float64, x []float64)         { panic("not implemented!") }
//func ToFloat32(dst []float32, x []float64)         { vek.ToFloat32_Into(dst, x) }
//
//// float32 exclusive.
//func Exp(x []float64)   { panic("not implemented!") }
//func Cos(x []float64)   { panic("not implemented!") }
//func Sin(x []float64)   { panic("not implemented!") }
//func Log(x []float64)   { panic("not implemented!") }
//func Log2(x []float64)  { panic("not implemented!") }
//func Log10(x []float64) { panic("not implemented!") }
