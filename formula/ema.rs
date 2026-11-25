//! ema —— 与 Go 的 `ema.go` 对应。
/// EMA 指数移动平均：alpha = 2/(N+1)，首值等于源序列首值。
pub fn ema(x: &[f64], span: i64) -> Vec<f64> {
    let alpha = 2.0 / (span.max(1) as f64 + 1.0);
    ewm(x, alpha)
}

/// EMA 的逐位置周期形态：第 i 位用 spans[i] 计算 alpha=2/(span+1)。
pub fn ema_series(x: &[f64], spans: &[f64]) -> Vec<f64> {
    if x.is_empty() {
        return Vec::new();
    }
    let mut out = Vec::with_capacity(x.len());
    let mut previous = x[0];
    out.push(previous);
    for (index, value) in x.iter().enumerate().skip(1) {
        let span = spans.get(index).copied().unwrap_or(1.0);
        let span = if span.is_nan() || span <= 0.0 {
            1.0
        } else {
            span
        };
        let alpha = 2.0 / (span + 1.0);
        previous = (1.0 - alpha) * previous + alpha * value;
        out.push(previous);
    }
    out
}

pub(crate) fn ewm(x: &[f64], alpha: f64) -> Vec<f64> {
    if x.is_empty() {
        return Vec::new();
    }
    let mut out = Vec::with_capacity(x.len());
    let mut previous = x[0];
    out.push(previous);
    for value in &x[1..] {
        previous = (1.0 - alpha) * previous + alpha * value;
        out.push(previous);
    }
    out
}
