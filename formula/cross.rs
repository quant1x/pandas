//! cross —— 与 Go 的 `cross.go` 对应。
/// CROSS 金叉：前一根 S1 < S2 且当前 S1 > S2。
pub fn cross(s1: &[f64], s2: &[f64]) -> Vec<bool> {
    assert_eq!(s1.len(), s2.len(), "序列长度不一致");
    let mut out = vec![false; s1.len()];
    for index in 1..s1.len() {
        out[index] = s1[index - 1] < s2[index - 1] && s1[index] > s2[index];
    }
    out
}
