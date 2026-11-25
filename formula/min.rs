//! min —— 与 Go 的 `min.go` 对应。
/// MIN 两条序列逐元素取小。
pub fn min_series(x: &[f64], y: &[f64]) -> Vec<f64> {
    quant1x_num::minimum(x, y)
}
