//! wma —— 与 Go 的 `wma.go` 对应。

use quant1x_num::{Window, rolling_apply};
/// WMA 加权移动平均：窗口内按 1..N 递增加权，分母为 N(N+1)/2。
pub fn wma(x: &[f64], window: i64) -> Vec<f64> {
    rolling_apply(x, &Window::Constant(window), |_, block| {
        let mut total = 0.0;
        let mut weights = 0.0;
        for (index, value) in block.iter().enumerate() {
            let weight = (index + 1) as f64;
            total += weight * value;
            weights += weight;
        }
        total / weights
    })
}
