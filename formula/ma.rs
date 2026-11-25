//! ma —— 与 Go 的 `ma.go` 对应。

use quant1x_num::{Window, rolling_apply};
/// MA 简单移动平均：完整窗口才计算，窗口不足为 NaN。
pub fn ma(x: &[f64], window: i64) -> Vec<f64> {
    rolling_apply(x, &Window::Constant(window), |n, block| {
        block.iter().sum::<f64>() / n
    })
}
