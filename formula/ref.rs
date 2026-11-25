//! ref —— 与 Go 的 `ref.go` 对应。

use quant1x_num::shift;
/// REF 引用前 N 个周期的值；越界为 NaN。
pub fn ref_value(x: &[f64], n: i64) -> Vec<f64> {
    shift(x, n)
}

/// REF 的逐位置窗口形态：第 i 位引用 i - periods[i] 的值，越界为 NaN。
pub fn ref_series(x: &[f64], periods: &[f64]) -> Vec<f64> {
    x.iter()
        .enumerate()
        .map(|(index, _)| {
            let period = periods.get(index).copied().unwrap_or(f64::NAN);
            if period.is_nan() || period < 0.0 {
                return f64::NAN;
            }
            let source = index as i64 - period as i64;
            if source < 0 || source as usize >= x.len() {
                f64::NAN
            } else {
                x[source as usize]
            }
        })
        .collect()
}
