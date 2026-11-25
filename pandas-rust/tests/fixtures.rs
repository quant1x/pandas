//! 读取 tests/formula/cases.csv，与 B3 实现逐值对拍。

use std::fs;

use quant1x_pandas::{
    abs, and, avedev, barslast, barslastcount, barslasts, barssincen, const_series, count, cross,
    dma, ema, eq, filter, forcast, gt, gte, hhv, hhvbars, if_select, ifn, last, llv, llvbars, lt,
    lte, ma, max_series, min_series, neq, not, or, ref_value, slope, sma, sqrt, std, sum, wma,
};

fn parse_list(text: &str) -> Vec<f64> {
    if text.is_empty() {
        return Vec::new();
    }
    text.split(';')
        .map(|item| {
            let item = item.trim();
            if item == "nan" {
                f64::NAN
            } else {
                item.parse::<f64>().expect("非法数值")
            }
        })
        .collect()
}

/// 布尔序列落成 0/1，供 fixture 比对。
fn flags_to_float(flags: Vec<bool>) -> Vec<f64> {
    flags
        .into_iter()
        .map(|flag| if flag { 1.0 } else { 0.0 })
        .collect()
}

/// 0/1 序列转布尔序列，非零即真。
fn to_bools(values: &[f64]) -> Vec<bool> {
    values.iter().map(|value| *value != 0.0).collect()
}

fn parse_optional(text: &str) -> i64 {
    if text.is_empty() {
        0
    } else {
        text.parse::<i64>().expect("非法整数")
    }
}

#[test]
fn formula_fixtures_match_rust() {
    let path = "../tests/formula/cases.csv";
    let text = fs::read_to_string(path).expect("读取 fixture");
    let mut lines = text.lines();
    let header = lines.next().expect("表头");
    assert_eq!(header, "formula,x,y,n,m,expected");

    let mut checked = 0;
    for line in lines {
        if line.is_empty() {
            continue;
        }
        let fields: Vec<&str> = line.split(',').collect();
        assert_eq!(fields.len(), 6, "列数不对: {line}");
        let name = fields[0];
        let x = parse_list(fields[1]);
        let y = parse_list(fields[2]);
        let n = parse_optional(fields[3]);
        let m = parse_optional(fields[4]);
        let want = parse_list(fields[5]);

        let got = match name {
            "abs" => abs(&x),
            "ref" => ref_value(&x, n),
            "ma" => ma(&x, n),
            "ema" => ema(&x, n),
            "sma" => sma(&x, n, m),
            "hhv" => hhv(&x, n),
            "llv" => llv(&x, n),
            "sum" => sum(&x, n),
            "std" => std(&x, n),
            "wma" => wma(&x, n),
            "sqrt" => sqrt(&x),
            "count" => count(&x.iter().map(|v| *v != 0.0).collect::<Vec<bool>>(), n),
            "filter" => filter(&x, n),
            "hhvbars" => hhvbars(&x, n),
            "llvbars" => llvbars(&x, n),
            "cross" => cross(&x, &y)
                .into_iter()
                .map(|flag| if flag { 1.0 } else { 0.0 })
                .collect(),
            "dma" => dma(&x, &y),
            "max" => max_series(&x, &y),
            "min" => min_series(&x, &y),
            "gt" => flags_to_float(gt(&x, &y)),
            "gte" => flags_to_float(gte(&x, &y)),
            "lt" => flags_to_float(lt(&x, &y)),
            "lte" => flags_to_float(lte(&x, &y)),
            "eq" => flags_to_float(eq(&x, &y)),
            "neq" => flags_to_float(neq(&x, &y)),
            "and" => flags_to_float(and(&to_bools(&x), &to_bools(&y))),
            "or" => flags_to_float(or(&to_bools(&x), &to_bools(&y))),
            "not" => flags_to_float(not(&to_bools(&x))),
            "barslast" => barslast(&x),
            "barslastcount" => barslastcount(&x),
            "barslasts" => barslasts(&x, n),
            "barssincen" => barssincen(&x, n),
            "if" => if_select(&x, &y, &vec![n as f64; y.len()]),
            "ifn" => ifn(&x, &y, &vec![n as f64; y.len()]),
            "avedev" => avedev(&x, n),
            "slope" => slope(&x, n),
            "forcast" => forcast(&x, n),
            "const" => const_series(&x),
            "last" => last(&x, n, m),
            other => panic!("未覆盖的公式 {other}"),
        };
        assert_eq!(got.len(), want.len(), "{name}: 长度");
        for (index, (actual, expected)) in got.iter().zip(want.iter()).enumerate() {
            if actual.is_nan() && expected.is_nan() {
                continue;
            }
            if actual == expected
                || (*actual - *expected).abs() <= 1e-12 * 1.0f64.max(expected.abs())
            {
                continue;
            }
            panic!("{name}: 第 {index} 位 {actual} != {expected}");
        }
        checked += 1;
    }
    assert!(checked >= 39, "fixture 行数不足");
}
