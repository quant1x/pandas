//! sum —— 与 Go 的 `sum.go` 对应。

use quant1x_num::{Window, rolling_apply, sum as sum_value};
/// SUM 最近 N 个周期的和。
pub fn sum(x: &[f64], window: i64) -> Vec<f64> {
    rolling_apply(x, &Window::Constant(window), |_, block| sum_value(block))
}
