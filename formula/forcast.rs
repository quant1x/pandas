//! forcast —— 与 Go 的 `forcast.go` 对应。

use quant1x_num::{Window, rolling_apply};

use crate::slope::regression;
/// FORCAST 窗口内线性回归在最后一根上的预测值。
pub fn forcast(x: &[f64], window: i64) -> Vec<f64> {
    rolling_apply(x, &Window::Constant(window), |_, block| {
        let (slope, intercept) = regression(block);
        if slope.is_nan() {
            return f64::NAN;
        }
        intercept + slope * (block.len() - 1) as f64
    })
}
