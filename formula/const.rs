//! const —— 与 Go 的 `const.go` 对应。
/// CONST 取序列最后值为常量序列（Go 先转 float32，精度按 f32）。
pub fn const_series(x: &[f64]) -> Vec<f64> {
    let last = x.last().copied().unwrap_or(f64::NAN) as f32;
    vec![f64::from(last); x.len()]
}
