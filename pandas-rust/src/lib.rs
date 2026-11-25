//! pandas 公式族的 Rust 对齐实现（B3）。
//!
//! 首批覆盖 ABS / REF / MA / EMA / SMA；序列用 `&[f64]`，与 Go 端的最小参数集一致。
//! EMA / SMA 按 Go 的 EWM `Adjust=false` 口径：`y0=x0`、`y_t=(1-alpha)*y_{t-1}+alpha*x_t`。

use quant1x_num::{
    Window, abs as abs_vector, max as max_value, min as min_value, rolling_apply, shift,
    std as std_value, sum as sum_value,
};

/// ABS 逐元素绝对值。
pub fn abs(x: &[f64]) -> Vec<f64> {
    abs_vector(x)
}

/// REF 引用前 N 个周期的值；越界为 NaN。
pub fn ref_value(x: &[f64], n: i64) -> Vec<f64> {
    shift(x, n)
}

/// MA 简单移动平均：完整窗口才计算，窗口不足为 NaN。
pub fn ma(x: &[f64], window: i64) -> Vec<f64> {
    rolling_apply(x, &Window::Constant(window), |n, block| {
        block.iter().sum::<f64>() / n
    })
}

/// EMA 指数移动平均：alpha = 2/(N+1)，首值等于源序列首值。
pub fn ema(x: &[f64], span: i64) -> Vec<f64> {
    let alpha = 2.0 / (span.max(1) as f64 + 1.0);
    ewm(x, alpha)
}

/// SMA 中国式平滑：alpha = M/N（M=0 视为 1）。
pub fn sma(x: &[f64], n: i64, m: i64) -> Vec<f64> {
    let weight = if m == 0 { 1 } else { m };
    let alpha = weight as f64 / n.max(1) as f64;
    ewm(x, alpha)
}

/// HHV 最近 N 个周期的最大值。
pub fn hhv(x: &[f64], window: i64) -> Vec<f64> {
    rolling_apply(x, &Window::Constant(window), |_, block| max_value(block))
}

/// LLV 最近 N 个周期的最小值。
pub fn llv(x: &[f64], window: i64) -> Vec<f64> {
    rolling_apply(x, &Window::Constant(window), |_, block| min_value(block))
}

/// SUM 最近 N 个周期的和。
pub fn sum(x: &[f64], window: i64) -> Vec<f64> {
    rolling_apply(x, &Window::Constant(window), |_, block| sum_value(block))
}

/// STD 最近 N 个周期的总体标准差。
pub fn std(x: &[f64], window: i64) -> Vec<f64> {
    rolling_apply(x, &Window::Constant(window), |_, block| std_value(block))
}

/// COUNT 最近 N 个周期内为真的个数；窗口不足时为 0（与 Go 的空窗口计数一致）。
pub fn count(conditions: &[bool], window: i64) -> Vec<f64> {
    let mut out = vec![0.0; conditions.len()];
    if window <= 0 {
        return out;
    }
    let size = window as usize;
    for index in 0..conditions.len() {
        if index + 1 < size {
            continue;
        }
        let start = index + 1 - size;
        out[index] = conditions[start..=index]
            .iter()
            .filter(|value| **value)
            .count() as f64;
    }
    out
}

/// WMA 加权移动平均：窗口内按 1..N 递增加权，分母为 N(N+1)/2。
pub fn wma(x: &[f64], window: i64) -> Vec<f64> {
    rolling_apply(x, &Window::Constant(window), |_, block| {
        let mut total = 0.0;
        let mut weights = 0.0;
        for (index, value) in block.iter().enumerate() {
            let weight = (index + 1) as f64;
            total += weight * value;
            weights += weight;
        }
        total / weights
    })
}

/// DMA 动态移动平均：alpha = 1/A（逐位置或常量），NaN 视为 1。
pub fn dma(x: &[f64], a: &[f64]) -> Vec<f64> {
    if x.is_empty() {
        return Vec::new();
    }
    let mut out = Vec::with_capacity(x.len());
    let mut previous = x[0];
    out.push(previous);
    for index in 1..x.len() {
        let period = if a.len() == 1 { a[0] } else { a[index] };
        let mut alpha = 1.0 / period;
        if !alpha.is_finite() {
            alpha = 1.0;
        }
        previous = alpha * x[index] + (1.0 - alpha) * previous;
        out.push(previous);
    }
    out
}

/// CROSS 金叉：前一根 S1 < S2 且当前 S1 > S2。
pub fn cross(s1: &[f64], s2: &[f64]) -> Vec<bool> {
    assert_eq!(s1.len(), s2.len(), "序列长度不一致");
    let mut out = vec![false; s1.len()];
    for index in 1..s1.len() {
        out[index] = s1[index - 1] < s2[index - 1] && s1[index] > s2[index];
    }
    out
}

/// HHVBARS 最近 N 周期内最高点到当前的周期数（0 表示当前就是最高点）。
pub fn hhvbars(x: &[f64], window: i64) -> Vec<f64> {
    rolling_apply(x, &Window::Constant(window), |_, block| {
        let maximum = max_value(block);
        let position = block
            .iter()
            .rposition(|value| *value == maximum)
            .unwrap_or(0);
        (block.len() - 1 - position) as f64
    })
}

/// LLVBARS 最近 N 周期内最低点到当前的周期数。
pub fn llvbars(x: &[f64], window: i64) -> Vec<f64> {
    rolling_apply(x, &Window::Constant(window), |_, block| {
        let minimum = min_value(block);
        let position = block
            .iter()
            .rposition(|value| *value == minimum)
            .unwrap_or(0);
        (block.len() - 1 - position) as f64
    })
}

/// SLOPE 窗口内线性回归斜率；窗口不足 2 个点时为 NaN。
pub fn slope(x: &[f64], window: i64) -> Vec<f64> {
    rolling_apply(x, &Window::Constant(window), |_, block| {
        let (slope, _) = regression(block);
        slope
    })
}

/// FORCAST 窗口内线性回归在最后一根上的预测值。
pub fn forcast(x: &[f64], window: i64) -> Vec<f64> {
    rolling_apply(x, &Window::Constant(window), |_, block| {
        let (slope, intercept) = regression(block);
        if slope.is_nan() {
            return f64::NAN;
        }
        intercept + slope * (block.len() - 1) as f64
    })
}

fn regression(block: &[f64]) -> (f64, f64) {
    let n = block.len();
    if n < 2 {
        return (f64::NAN, f64::NAN);
    }
    let nf = n as f64;
    let mut sum_x = 0.0;
    let mut sum_y = 0.0;
    let mut sum_xy = 0.0;
    let mut sum_xx = 0.0;
    for (index, value) in block.iter().enumerate() {
        let x = index as f64;
        sum_x += x;
        sum_y += value;
        sum_xy += x * value;
        sum_xx += x * x;
    }
    let denominator = nf * sum_xx - sum_x * sum_x;
    if denominator.abs() < 1e-12 {
        return (f64::NAN, f64::NAN);
    }
    let slope = (nf * sum_xy - sum_x * sum_y) / denominator;
    let intercept = (sum_y - slope * sum_x) / nf;
    (slope, intercept)
}

/// AVEDEV 窗口内平均绝对偏差。
pub fn avedev(x: &[f64], window: i64) -> Vec<f64> {
    rolling_apply(x, &Window::Constant(window), |_, block| {
        let average = quant1x_num::mean(block);
        block
            .iter()
            .map(|value| (value - average).abs())
            .sum::<f64>()
            / block.len() as f64
    })
}

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

/// MAX 两条序列逐元素取大；短的一侧按 NaN 补齐（与 num::maximum 一致）。
pub fn max_series(x: &[f64], y: &[f64]) -> Vec<f64> {
    quant1x_num::maximum(x, y)
}

/// MIN 两条序列逐元素取小。
pub fn min_series(x: &[f64], y: &[f64]) -> Vec<f64> {
    quant1x_num::minimum(x, y)
}

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

/// BARSSINCEN N 周期内第一次条件成立到当前的周期数；没有成立时为 0。
pub fn barssincen(condition: &[f64], window: i64) -> Vec<f64> {
    let mut out = vec![0.0; condition.len()];
    if window <= 0 {
        return out;
    }
    let size = window as usize;
    for index in 0..condition.len() {
        if index + 1 < size {
            continue;
        }
        let start = index + 1 - size;
        if let Some(position) = condition[start..=index]
            .iter()
            .position(|value| *value != 0.0)
        {
            out[index] = (size - 1 - position) as f64;
        }
    }
    out
}

/// CONST 取序列最后值为常量序列（Go 先转 float32，精度按 f32）。
pub fn const_series(x: &[f64]) -> Vec<f64> {
    let last = x.last().copied().unwrap_or(f64::NAN) as f32;
    vec![f64::from(last); x.len()]
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

/// EMA 的逐位置周期形态：第 i 位用 spans[i] 计算 alpha=2/(span+1)。
pub fn ema_series(x: &[f64], spans: &[f64]) -> Vec<f64> {
    if x.is_empty() {
        return Vec::new();
    }
    let mut out = Vec::with_capacity(x.len());
    let mut previous = x[0];
    out.push(previous);
    for (index, value) in x.iter().enumerate().skip(1) {
        let span = spans.get(index).copied().unwrap_or(1.0);
        let span = if span.is_nan() || span <= 0.0 {
            1.0
        } else {
            span
        };
        let alpha = 2.0 / (span + 1.0);
        previous = (1.0 - alpha) * previous + alpha * value;
        out.push(previous);
    }
    out
}

/// SMA 的序列周期形态：Go 取序列最后一个值作为 N（不是逐位置），M=0 视为 1。
pub fn sma_series(x: &[f64], periods: &[f64], m: i64) -> Vec<f64> {
    let n = periods.last().copied().unwrap_or(1.0);
    sma(x, n as i64, m)
}

fn ewm(x: &[f64], alpha: f64) -> Vec<f64> {
    if x.is_empty() {
        return Vec::new();
    }
    let mut out = Vec::with_capacity(x.len());
    let mut previous = x[0];
    out.push(previous);
    for value in &x[1..] {
        previous = (1.0 - alpha) * previous + alpha * value;
        out.push(previous);
    }
    out
}

#[cfg(test)]
mod tests {
    use super::*;

    fn assert_close(got: &[f64], want: &[f64]) {
        assert_eq!(got.len(), want.len());
        for (index, (actual, expected)) in got.iter().zip(want.iter()).enumerate() {
            if expected.is_nan() {
                assert!(actual.is_nan(), "第 {index} 位应为 NaN，实际 {actual}");
                continue;
            }
            assert!(
                (actual - expected).abs() <= 1e-12 * 1.0f64.max(expected.abs()),
                "第 {index} 位 {actual} != {expected}"
            );
        }
    }

    #[test]
    fn abs_basic() {
        assert_close(&abs(&[-1.5, 0.0, 2.5]), &[1.5, 0.0, 2.5]);
        assert!(abs(&[]).is_empty());
    }

    #[test]
    fn ref_basic() {
        assert_close(
            &ref_value(&[1.0, 2.0, 3.0, 4.0, 5.0], 1),
            &[f64::NAN, 1.0, 2.0, 3.0, 4.0],
        );
        assert_close(&ref_value(&[1.0, 2.0, 3.0], 2), &[f64::NAN, f64::NAN, 1.0]);
    }

    #[test]
    fn ma_basic() {
        assert_close(
            &ma(&[1.0, 2.0, 3.0, 4.0, 5.0], 3),
            &[f64::NAN, f64::NAN, 2.0, 3.0, 4.0],
        );
    }

    #[test]
    fn hhv_llv_basic() {
        assert_close(&hhv(&[1.0, 3.0, 2.0, 5.0], 2), &[f64::NAN, 3.0, 3.0, 5.0]);
        assert_close(&llv(&[1.0, 3.0, 2.0, 5.0], 2), &[f64::NAN, 1.0, 2.0, 2.0]);
    }

    #[test]
    fn sum_std_basic() {
        assert_close(&sum(&[1.0, 2.0, 3.0, 4.0], 2), &[f64::NAN, 3.0, 5.0, 7.0]);
        assert_close(&std(&[1.0, 2.0, 3.0, 4.0], 2), &[f64::NAN, 0.5, 0.5, 0.5]);
    }

    #[test]
    fn count_basic() {
        assert_close(&count(&[true, false, true, true], 2), &[0.0, 1.0, 1.0, 2.0]);
        assert_close(&count(&[true], 3), &[0.0]);
    }

    #[test]
    fn wma_basic() {
        let expected = (1.0 + 2.0 * 2.0 + 3.0 * 3.0) / 6.0;
        assert_close(&wma(&[1.0, 2.0, 3.0], 3), &[f64::NAN, f64::NAN, expected]);
    }

    #[test]
    fn dma_basic() {
        assert_close(
            &dma(&[1.0, 2.0, 3.0, 4.0], &[2.0]),
            &[1.0, 1.5, 2.25, 3.125],
        );
        assert_close(&dma(&[1.0, 2.0, 3.0], &[2.0, 4.0, 0.0]), &[1.0, 1.25, 3.0]);
    }

    #[test]
    fn cross_basic() {
        assert_eq!(
            cross(&[1.0, 3.0, 1.0, 3.0], &[2.0, 2.0, 2.0, 2.0]),
            vec![false, true, false, true]
        );
    }

    #[test]
    fn hhvbars_llvbars_basic() {
        assert_close(
            &hhvbars(&[1.0, 3.0, 2.0, 5.0], 2),
            &[f64::NAN, 0.0, 1.0, 0.0],
        );
        assert_close(
            &llvbars(&[1.0, 3.0, 2.0, 5.0], 2),
            &[f64::NAN, 1.0, 0.0, 1.0],
        );
    }

    #[test]
    fn slope_forcast_basic() {
        assert_close(
            &slope(&[1.0, 2.0, 3.0, 4.0], 3),
            &[f64::NAN, f64::NAN, 1.0, 1.0],
        );
        assert_close(
            &forcast(&[1.0, 2.0, 3.0, 4.0], 3),
            &[f64::NAN, f64::NAN, 3.0, 4.0],
        );
    }

    #[test]
    fn avedev_basic() {
        assert_close(
            &avedev(&[1.0, 2.0, 3.0, 4.0], 2),
            &[f64::NAN, 0.5, 0.5, 0.5],
        );
    }

    #[test]
    fn filter_basic() {
        assert_close(
            &filter(&[1.0, 0.0, 1.0, 0.0, 0.0], 1),
            &[1.0, 0.0, 1.0, 0.0, 0.0],
        );
        assert_close(&filter(&[1.0, 1.0, 1.0, 1.0], 2), &[1.0, 0.0, 0.0, 1.0]);
    }

    #[test]
    fn max_min_basic() {
        assert_close(
            &max_series(&[1.0, 5.0, 3.0], &[2.0, 4.0, 3.0]),
            &[2.0, 5.0, 3.0],
        );
        assert_close(
            &min_series(&[1.0, 5.0, 3.0], &[2.0, 4.0, 3.0]),
            &[1.0, 4.0, 3.0],
        );
    }

    #[test]
    fn if_select_basic() {
        assert_close(
            &if_select(
                &[1.0, 0.0, 2.0, 0.0],
                &[10.0, 20.0, 30.0, 40.0],
                &[1.0, 2.0, 3.0, 4.0],
            ),
            &[10.0, 2.0, 30.0, 4.0],
        );
        assert_close(
            &ifn(
                &[1.0, 0.0, 2.0, 0.0],
                &[10.0, 20.0, 30.0, 40.0],
                &[1.0, 2.0, 3.0, 4.0],
            ),
            &[1.0, 20.0, 3.0, 40.0],
        );
    }

    #[test]
    fn barslast_basic() {
        assert_close(
            &barslast(&[0.0, 1.0, 0.0, 0.0, 1.0, 0.0]),
            &[-1.0, 0.0, 1.0, 2.0, 0.0, 1.0],
        );
    }

    #[test]
    fn barslastcount_basic() {
        assert_close(
            &barslastcount(&[1.0, 1.0, 0.0, 1.0, 1.0, 1.0]),
            &[1.0, 2.0, 0.0, 1.0, 2.0, 3.0],
        );
    }

    #[test]
    fn comparison_basic() {
        assert_eq!(
            gt(&[1.0, 2.0, 3.0], &[2.0, 2.0, 2.0]),
            vec![false, false, true]
        );
        assert_eq!(
            gte(&[1.0, 2.0, 3.0], &[2.0, 2.0, 2.0]),
            vec![false, true, true]
        );
        assert_eq!(
            lt(&[1.0, 2.0, 3.0], &[2.0, 2.0, 2.0]),
            vec![true, false, false]
        );
        assert_eq!(
            lte(&[1.0, 2.0, 3.0], &[2.0, 2.0, 2.0]),
            vec![true, true, false]
        );
        assert_eq!(eq(&[1.0, 2.0], &[1.0, 3.0]), vec![true, false]);
        assert_eq!(neq(&[1.0, 2.0], &[1.0, 3.0]), vec![false, true]);
    }

    #[test]
    fn logic_basic() {
        assert_eq!(
            and(&[true, true, false], &[true, false, false]),
            vec![true, false, false]
        );
        assert_eq!(
            or(&[true, true, false], &[false, false, false]),
            vec![true, true, false]
        );
        assert_eq!(not(&[true, false]), vec![false, true]);
    }

    #[test]
    fn barssincen_basic() {
        assert_close(
            &barssincen(&[0.0, 1.0, 0.0, 0.0, 1.0], 3),
            &[0.0, 0.0, 1.0, 2.0, 0.0],
        );
    }

    #[test]
    fn const_series_basic() {
        assert_close(&const_series(&[1.0, 2.0, 3.0]), &[3.0, 3.0, 3.0]);
        assert!(const_series(&[]).is_empty());
    }

    #[test]
    fn ref_series_basic() {
        assert_close(
            &ref_series(&[1.0, 2.0, 3.0, 4.0], &[1.0, 2.0, 0.0, 1.0]),
            &[f64::NAN, f64::NAN, 3.0, 3.0],
        );
    }

    #[test]
    fn ema_series_basic() {
        let alpha1 = 2.0 / 3.0;
        let y1 = (1.0 - alpha1) * 1.0 + alpha1 * 2.0;
        let alpha2 = 2.0 / 5.0;
        let y2 = (1.0 - alpha2) * y1 + alpha2 * 3.0;
        assert_close(
            &ema_series(&[1.0, 2.0, 3.0], &[2.0, 2.0, 4.0]),
            &[1.0, y1, y2],
        );
    }

    #[test]
    fn sma_series_uses_last_period() {
        let alpha: f64 = 0.25;
        let y1 = (1.0 - alpha) * 1.0 + alpha * 2.0;
        let y2 = (1.0 - alpha) * y1 + alpha * 3.0;
        assert_close(
            &sma_series(&[1.0, 2.0, 3.0], &[2.0, 2.0, 4.0], 1),
            &[1.0, y1, y2],
        );
    }

    #[test]
    fn ema_matches_go() {
        assert_close(&ema(&[1.0, 2.0, 3.0, 4.0], 4), &[1.0, 1.4, 2.04, 2.824]);
    }

    #[test]
    fn sma_matches_ewm() {
        let values = [1.0, 2.0, 3.0, 4.0];
        let alpha: f64 = 1.0 / 3.0;
        let want = [
            1.0,
            (1.0 - alpha) * 1.0 + alpha * 2.0,
            (1.0 - alpha) * ((1.0 - alpha) * 1.0 + alpha * 2.0) + alpha * 3.0,
            (1.0 - alpha) * ((1.0 - alpha) * ((1.0 - alpha) * 1.0 + alpha * 2.0) + alpha * 3.0)
                + alpha * 4.0,
        ];
        assert_close(&sma(&values, 3, 1), &want);
    }
}
