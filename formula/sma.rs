//! sma —— 与 Go 的 `sma.go` 对应。

use crate::ema::ewm;
/// SMA 中国式平滑：alpha = M/N（M=0 视为 1）。
pub fn sma(x: &[f64], n: i64, m: i64) -> Vec<f64> {
    let weight = if m == 0 { 1 } else { m };
    let alpha = weight as f64 / n.max(1) as f64;
    ewm(x, alpha)
}

/// SMA 的序列周期形态：Go 取序列最后一个值作为 N（不是逐位置），M=0 视为 1。
pub fn sma_series(x: &[f64], periods: &[f64], m: i64) -> Vec<f64> {
    let n = periods.last().copied().unwrap_or(1.0);
    sma(x, n as i64, m)
}
