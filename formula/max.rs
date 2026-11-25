//! max —— 与 Go 的 `max.go` 对应。
/// MAX 两条序列逐元素取大；短的一侧按 NaN 补齐（与 num::maximum 一致）。
pub fn max_series(x: &[f64], y: &[f64]) -> Vec<f64> {
    quant1x_num::maximum(x, y)
}
