//! if —— 与 Go 的 `if.go` 对应。
/// IF / IFF 条件非零取 a，否则取 b；按最长对齐，缺位为 NaN。
pub fn if_select(condition: &[f64], a: &[f64], b: &[f64]) -> Vec<f64> {
    let length = condition.len().max(a.len()).max(b.len());
    let mut out = Vec::with_capacity(length);
    for index in 0..length {
        let choose_a = condition.get(index).copied().unwrap_or(0.0) != 0.0;
        out.push(if choose_a {
            a.get(index).copied().unwrap_or(f64::NAN)
        } else {
            b.get(index).copied().unwrap_or(f64::NAN)
        });
    }
    out
}

/// IFN 条件为零取 a，否则取 b。
pub fn ifn(condition: &[f64], a: &[f64], b: &[f64]) -> Vec<f64> {
    if_select(condition, b, a)
}
