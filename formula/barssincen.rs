//! barssincen —— 与 Go 的 `barssincen.go` 对应。
/// BARSSINCEN N 周期内第一次条件成立到当前的周期数；没有成立时为 0。
pub fn barssincen(condition: &[f64], window: i64) -> Vec<f64> {
    let mut out = vec![0.0; condition.len()];
    if window <= 0 {
        return out;
    }
    let size = window as usize;
    for index in 0..condition.len() {
        if index + 1 < size {
            continue;
        }
        let start = index + 1 - size;
        if let Some(position) = condition[start..=index]
            .iter()
            .position(|value| *value != 0.0)
        {
            out[index] = (size - 1 - position) as f64;
        }
    }
    out
}
