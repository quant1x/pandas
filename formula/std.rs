//! std —— 与 Go 的 `std.go` 对应。

use quant1x_num::{Window, rolling_apply, std as std_value};
/// STD 最近 N 个周期的总体标准差。
pub fn std(x: &[f64], window: i64) -> Vec<f64> {
    rolling_apply(x, &Window::Constant(window), |_, block| std_value(block))
}
