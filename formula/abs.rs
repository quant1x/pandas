//! abs —— 与 Go 的 `abs.go` 对应。

use quant1x_num::abs as abs_vector;
/// ABS 逐元素绝对值。
pub fn abs(x: &[f64]) -> Vec<f64> {
    abs_vector(x)
}
