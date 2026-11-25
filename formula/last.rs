//! last —— 与 Go 的 `last.go` 对应。

use quant1x_num::{Window, rolling_apply};
/// LAST 条件持续存在：在 A+1 的窗口里，检查除最近 B 根外的所有位置是否都非零。
///
/// 与 Go 的 `LAST(X,A,B)` 同口径：窗口不足时返回 NaN；B 大于等于窗口长度时视为成立。
pub fn last(condition: &[f64], a: i64, b: i64) -> Vec<f64> {
    rolling_apply(condition, &Window::Constant(a + 1), |_, block| {
        let skip = b.max(0) as usize;
        if skip >= block.len() {
            return 1.0;
        }
        let end = block.len() - skip;
        if block[..end].iter().all(|value| *value != 0.0) {
            1.0
        } else {
            0.0
        }
    })
}
