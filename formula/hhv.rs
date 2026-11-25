//! hhv —— 与 Go 的 `hhv.go` 对应。

use quant1x_num::{Window, max as max_value, rolling_apply};
/// HHV 最近 N 个周期的最大值。
pub fn hhv(x: &[f64], window: i64) -> Vec<f64> {
    rolling_apply(x, &Window::Constant(window), |_, block| max_value(block))
}
