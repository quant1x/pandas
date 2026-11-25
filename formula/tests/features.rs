//! 读取 tests/features/features.csv，用 quant1x_pandas 重算特征列并逐值对拍（M8 第一步）。

use std::collections::HashMap;
use std::fs;

use quant1x_pandas::{cross, ema, hhv, hhvbars, llv, ma, std as rolling_std, wma};

const FEATURES: [&str; 11] = [
    "ma5",
    "ma20",
    "ema12",
    "wma10",
    "hhv20",
    "llv20",
    "std20",
    "hhvbars20",
    "ret1",
    "bias5",
    "cross_up",
];

fn parse_value(text: &str) -> f64 {
    let text = text.trim();
    if text == "nan" {
        f64::NAN
    } else {
        text.parse::<f64>().expect("非法数值")
    }
}

fn index_of(header: &[&str], name: &str) -> usize {
    header
        .iter()
        .position(|value| *value == name)
        .unwrap_or_else(|| panic!("表头缺少 {name}"))
}

/// 读取冻结模型权重（term,weight）。
fn load_model(path: &str) -> HashMap<String, f64> {
    let text = fs::read_to_string(path).expect("读取模型");
    let mut weights = HashMap::new();
    for line in text.lines().skip(1) {
        if line.is_empty() {
            continue;
        }
        let fields: Vec<&str> = line.split(',').collect();
        weights.insert(fields[0].to_string(), parse_value(fields[1]));
    }
    weights
}

fn close_enough(actual: f64, expected: f64) -> bool {
    if actual.is_nan() && expected.is_nan() {
        return true;
    }
    if actual == expected {
        return true;
    }
    (actual - expected).abs() <= 1e-12 * 1.0f64.max(expected.abs())
}

#[test]
fn feature_fixtures_match_rust() {
    let path = "tests/features/features.csv";
    let text = fs::read_to_string(path).expect("读取 fixture");
    let mut lines = text.lines();
    let header: Vec<&str> = lines.next().expect("表头").split(',').collect();

    let close_col = index_of(&header, "close");
    let high_col = index_of(&header, "high");
    let low_col = index_of(&header, "low");

    let mut close = Vec::new();
    let mut high = Vec::new();
    let mut low = Vec::new();
    let mut want: Vec<Vec<f64>> = vec![Vec::new(); FEATURES.len()];
    let pred_col = index_of(&header, "pred");
    let mlp_col = index_of(&header, "pred_mlp");
    let mut want_pred = Vec::new();
    let mut want_mlp = Vec::new();

    for line in lines {
        if line.is_empty() {
            continue;
        }
        let fields: Vec<&str> = line.split(',').collect();
        close.push(parse_value(fields[close_col]));
        high.push(parse_value(fields[high_col]));
        low.push(parse_value(fields[low_col]));
        for (slot, name) in FEATURES.iter().enumerate() {
            want[slot].push(parse_value(fields[index_of(&header, name)]));
        }
        want_pred.push(parse_value(fields[pred_col]));
        want_mlp.push(parse_value(fields[mlp_col]));
    }

    let ma5 = ma(&close, 5);
    let ma20 = ma(&close, 20);
    let ema12 = ema(&close, 12);
    let wma10 = wma(&close, 10);
    let hhv20 = hhv(&high, 20);
    let llv20 = llv(&low, 20);
    let std20 = rolling_std(&close, 20);
    let hhvbars20 = hhvbars(&high, 20);
    let cross_up: Vec<f64> = cross(&ma5, &ma20)
        .into_iter()
        .map(|flag| if flag { 1.0 } else { 0.0 })
        .collect();

    let mut ret1 = vec![f64::NAN; close.len()];
    for index in 1..close.len() {
        ret1[index] = close[index] / close[index - 1] - 1.0;
    }

    let mut bias5 = vec![f64::NAN; close.len()];
    for index in 0..close.len() {
        if !ma5[index].is_nan() {
            bias5[index] = close[index] / ma5[index] - 1.0;
        }
    }

    let computed = [
        ma5, ma20, ema12, wma10, hhv20, llv20, std20, hhvbars20, ret1, bias5, cross_up,
    ];
    for (slot, name) in FEATURES.iter().enumerate() {
        assert_eq!(computed[slot].len(), want[slot].len(), "{name}: 长度");
        for (index, (actual, expected)) in computed[slot].iter().zip(want[slot].iter()).enumerate()
        {
            if !close_enough(*actual, *expected) {
                panic!("{name}: 第 {index} 位 {actual} != {expected}");
            }
        }
    }
    let weights = load_model("tests/features/model.csv");
    let mut pred = vec![f64::NAN; close.len()];
    for (index, slot) in pred.iter_mut().enumerate() {
        let mut total = weights["bias"];
        let mut missing = false;
        for (position, name) in FEATURES.iter().enumerate() {
            let value = computed[position][index];
            if value.is_nan() {
                missing = true;
                break;
            }
            total += weights[*name] * value;
        }
        if !missing {
            *slot = total;
        }
    }
    for (index, (actual, expected)) in pred.iter().zip(want_pred.iter()).enumerate() {
        assert!(
            close_enough(*actual, *expected),
            "pred: 第 {index} 位 {actual} != {expected}"
        );
    }
    let mlp_weights = load_model("tests/features/model_mlp.csv");
    let mut pred_mlp = vec![f64::NAN; close.len()];
    for (index, slot) in pred_mlp.iter_mut().enumerate() {
        let mut values = Vec::with_capacity(FEATURES.len());
        let mut missing = false;
        for (position, _) in FEATURES.iter().enumerate() {
            let value = computed[position][index];
            if value.is_nan() {
                missing = true;
                break;
            }
            values.push(value);
        }
        if missing {
            continue;
        }
        let mut hidden = [0.0f64; 4];
        for (j, activation) in hidden.iter_mut().enumerate() {
            let mut total = mlp_weights[&format!("b1_{j}")];
            for (i, value) in values.iter().enumerate() {
                total += mlp_weights[&format!("w1_{i}_{j}")] * value;
            }
            *activation = total.tanh();
        }
        let mut total = mlp_weights["b2"];
        for (j, activation) in hidden.iter().enumerate() {
            total += mlp_weights[&format!("w2_{j}")] * activation;
        }
        *slot = total;
    }
    for (index, (actual, expected)) in pred_mlp.iter().zip(want_mlp.iter()).enumerate() {
        assert!(
            close_enough(*actual, *expected),
            "pred_mlp: 第 {index} 位 {actual} != {expected}"
        );
    }

    assert!(close.len() >= 20, "fixture 行数不足");
}
