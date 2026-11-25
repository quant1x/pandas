//! count —— 与 Go 的 `count.go` 对应。
/// COUNT 最近 N 个周期内为真的个数；窗口不足时为 0（与 Go 的空窗口计数一致）。
pub fn count(conditions: &[bool], window: i64) -> Vec<f64> {
    let mut out = vec![0.0; conditions.len()];
    if window <= 0 {
        return out;
    }
    let size = window as usize;
    for index in 0..conditions.len() {
        if index + 1 < size {
            continue;
        }
        let start = index + 1 - size;
        out[index] = conditions[start..=index]
            .iter()
            .filter(|value| **value)
            .count() as f64;
    }
    out
}
