//! llv —— 与 Go 的 `llv.go` 对应。

use quant1x_num::{Window, min as min_value, rolling_apply};
/// LLV 最近 N 个周期的最小值。
pub fn llv(x: &[f64], window: i64) -> Vec<f64> {
    rolling_apply(x, &Window::Constant(window), |_, block| min_value(block))
}
