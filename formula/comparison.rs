//! comparison —— 与 Go 的 `comparison.go` 对应。
/// GT 逐元素大于。
pub fn gt(a: &[f64], b: &[f64]) -> Vec<bool> {
    quant1x_num::greater(a, b)
}

/// GTE 逐元素大于等于。
pub fn gte(a: &[f64], b: &[f64]) -> Vec<bool> {
    quant1x_num::greater_equal(a, b)
}

/// LT 逐元素小于。
pub fn lt(a: &[f64], b: &[f64]) -> Vec<bool> {
    quant1x_num::less(a, b)
}

/// LTE 逐元素小于等于。
pub fn lte(a: &[f64], b: &[f64]) -> Vec<bool> {
    quant1x_num::less_equal(a, b)
}

/// EQ 逐元素相等；NaN 与自身不相等。
pub fn eq(a: &[f64], b: &[f64]) -> Vec<bool> {
    quant1x_num::equal(a, b)
}

/// NEQ 逐元素不等。
pub fn neq(a: &[f64], b: &[f64]) -> Vec<bool> {
    quant1x_num::not_equal(a, b)
}

/// AND 逐元素逻辑与。
pub fn and(a: &[bool], b: &[bool]) -> Vec<bool> {
    quant1x_num::logic_and(a, b)
}

/// OR 逐元素逻辑或。
pub fn or(a: &[bool], b: &[bool]) -> Vec<bool> {
    quant1x_num::logic_or(a, b)
}

/// NOT 逐元素逻辑非。
pub fn not(a: &[bool]) -> Vec<bool> {
    quant1x_num::logic_not(a)
}
