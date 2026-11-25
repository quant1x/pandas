//! slope —— 与 Go 的 `slope.go` 对应。

use quant1x_num::{Window, rolling_apply};
/// SLOPE 窗口内线性回归斜率；窗口不足 2 个点时为 NaN。
pub fn slope(x: &[f64], window: i64) -> Vec<f64> {
    rolling_apply(x, &Window::Constant(window), |_, block| {
        let (slope, _) = regression(block);
        slope
    })
}

pub(crate) fn regression(block: &[f64]) -> (f64, f64) {
    let n = block.len();
    if n < 2 {
        return (f64::NAN, f64::NAN);
    }
    let nf = n as f64;
    let mut sum_x = 0.0;
    let mut sum_y = 0.0;
    let mut sum_xy = 0.0;
    let mut sum_xx = 0.0;
    for (index, value) in block.iter().enumerate() {
        let x = index as f64;
        sum_x += x;
        sum_y += value;
        sum_xy += x * value;
        sum_xx += x * x;
    }
    let denominator = nf * sum_xx - sum_x * sum_x;
    if denominator.abs() < 1e-12 {
        return (f64::NAN, f64::NAN);
    }
    let slope = (nf * sum_xy - sum_x * sum_y) / denominator;
    let intercept = (sum_y - slope * sum_x) / nf;
    (slope, intercept)
}
