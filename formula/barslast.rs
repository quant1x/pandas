//! barslast —— 与 Go 的 `barslast.go` 对应。
/// BARSLAST 上一次条件成立到当前的周期数；从未成立时为 -1。
pub fn barslast(condition: &[f64]) -> Vec<f64> {
    let mut out = Vec::with_capacity(condition.len());
    let mut last_true: Option<usize> = None;
    for (index, value) in condition.iter().enumerate() {
        if *value != 0.0 {
            last_true = Some(index);
            out.push(0.0);
        } else {
            out.push(match last_true {
                Some(position) => (index - position) as f64,
                None => -1.0,
            });
        }
    }
    out
}
