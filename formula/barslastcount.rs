//! barslastcount —— 与 Go 的 `barslastcount.go` 对应。
/// BARSLASTCOUNT 连续满足条件的周期数。
pub fn barslastcount(condition: &[f64]) -> Vec<f64> {
    let mut out = Vec::with_capacity(condition.len());
    let mut streak = 0usize;
    for value in condition {
        if *value != 0.0 {
            streak += 1;
            out.push(streak as f64);
        } else {
            streak = 0;
            out.push(0.0);
        }
    }
    out
}
