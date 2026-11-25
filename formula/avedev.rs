//! avedev —— 与 Go 的 `avedev.go` 对应。

use quant1x_num::{Window, rolling_apply};
/// AVEDEV 窗口内平均绝对偏差。
pub fn avedev(x: &[f64], window: i64) -> Vec<f64> {
    rolling_apply(x, &Window::Constant(window), |_, block| {
        let average = quant1x_num::mean(block);
        block
            .iter()
            .map(|value| (value - average).abs())
            .sum::<f64>()
            / block.len() as f64
    })
}
