//! filter —— 与 Go 的 `filter.go` 对应。
/// FILTER 条件为真的位置，把之后 N 个位置置零。
pub fn filter(x: &[f64], n: i64) -> Vec<f64> {
    let mut out = x.to_vec();
    let length = out.len();
    for index in 0..length {
        if out[index] == 0.0 {
            continue;
        }
        let end = (index + 1 + n.max(0) as usize).min(length);
        for value in out.iter_mut().take(end).skip(index + 1) {
            *value = 0.0;
        }
    }
    out
}
