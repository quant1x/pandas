//! llvbars —— 与 Go 的 `llvbars.go` 对应。

use quant1x_num::{Window, min as min_value, rolling_apply};
/// LLVBARS 最近 N 周期内最低点到当前的周期数。
pub fn llvbars(x: &[f64], window: i64) -> Vec<f64> {
    rolling_apply(x, &Window::Constant(window), |_, block| {
        let minimum = min_value(block);
        let position = block
            .iter()
            .rposition(|value| *value == minimum)
            .unwrap_or(0);
        (block.len() - 1 - position) as f64
    })
}
