//! barslasts —— 与 Go 的 `barslasts.go` 对应。
/// BARSLASTS 倒数第 N 次条件成立到当前的周期数。
///
/// 按 Go 的递推实现：先算 BARSLAST 计数序列，再迭代 `N-1` 次
/// `m[i] = m[i - (m[i]+1)] + (m[i]+1)`，越界为 NaN。
pub fn barslasts(condition: &[f64], n: i64) -> Vec<f64> {
    let mut current = barslast_zero(condition);
    for _ in 1..n {
        let mut next = vec![f64::NAN; current.len()];
        for (index, value) in current.iter().enumerate() {
            if value.is_nan() {
                continue;
            }
            let step = *value + 1.0;
            let source = index as f64 - step;
            if source < 0.0 {
                continue;
            }
            let source = source as usize;
            if source >= current.len() || current[source].is_nan() {
                continue;
            }
            next[index] = current[source] + step;
        }
        current = next;
    }
    current
}

pub(crate) fn barslast_zero(condition: &[f64]) -> Vec<f64> {
    let mut out = Vec::with_capacity(condition.len());
    let mut previous = 0.0;
    for value in condition {
        if *value != 0.0 {
            previous = 0.0;
        } else {
            previous += 1.0;
        }
        out.push(previous);
    }
    out
}
