//! dma —— 与 Go 的 `dma.go` 对应。
/// DMA 动态移动平均：alpha = 1/A（逐位置或常量），NaN 视为 1。
pub fn dma(x: &[f64], a: &[f64]) -> Vec<f64> {
    if x.is_empty() {
        return Vec::new();
    }
    let mut out = Vec::with_capacity(x.len());
    let mut previous = x[0];
    out.push(previous);
    for index in 1..x.len() {
        let period = if a.len() == 1 { a[0] } else { a[index] };
        let mut alpha = 1.0 / period;
        if !alpha.is_finite() {
            alpha = 1.0;
        }
        previous = alpha * x[index] + (1.0 - alpha) * previous;
        out.push(previous);
    }
    out
}
