//! hhvbars —— 与 Go 的 `hhvbars.go` 对应。

use quant1x_num::{Window, max as max_value, rolling_apply};
/// HHVBARS 最近 N 周期内最高点到当前的周期数（0 表示当前就是最高点）。
pub fn hhvbars(x: &[f64], window: i64) -> Vec<f64> {
    rolling_apply(x, &Window::Constant(window), |_, block| {
        let maximum = max_value(block);
        let position = block
            .iter()
            .rposition(|value| *value == maximum)
            .unwrap_or(0);
        (block.len() - 1 - position) as f64
    })
}
