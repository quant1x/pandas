//! ONNX 推理对拍（M8）：用 ort 动态加载 onnxruntime，跑 `onnx/mlp.onnx` 与 `pred_mlp` 比较。
//!
//! 需要把 `ORT_DYLIB_PATH` 指向 onnxruntime 动态库（`ort` 以 `load-dynamic` 方式加载，
//! 构建期不从网络下载二进制）。未设置该变量时跳过——CI 上通常没有这份库。

use std::fs;
use std::path::PathBuf;

use ort::session::Session;
use ort::value::Tensor;

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

fn features_root() -> PathBuf {
    PathBuf::from(env!("CARGO_MANIFEST_DIR")).join("../tests/features")
}

fn parse(text: &str) -> f64 {
    let text = text.trim();
    if text == "nan" {
        f64::NAN
    } else {
        text.parse::<f64>().expect("非法数值")
    }
}

#[test]
fn mlp_onnx_matches_pred_mlp() {
    if std::env::var("ORT_DYLIB_PATH")
        .map(|value| value.is_empty())
        .unwrap_or(true)
    {
        eprintln!("跳过：未设置 ORT_DYLIB_PATH");
        return;
    }

    let root = features_root();
    let text = fs::read_to_string(root.join("features.csv")).expect("读取 features.csv");
    let mut lines = text.lines();
    let header: Vec<&str> = lines.next().expect("表头").split(',').collect();
    let index = |name: &str| header.iter().position(|item| *item == name).expect("列名");

    let mut inputs: Vec<f64> = Vec::new();
    let mut expected: Vec<f64> = Vec::new();
    let mut rows = 0usize;
    for line in lines {
        if line.is_empty() {
            continue;
        }
        let fields: Vec<&str> = line.split(',').collect();
        for name in FEATURES {
            inputs.push(parse(fields[index(name)]));
        }
        expected.push(parse(fields[index("pred_mlp")]));
        rows += 1;
    }
    assert!(rows >= 20, "样本太少");

    let environment = ort::init().build().expect("创建环境");
    let mut session = Session::builder(&environment)
        .expect("builder")
        .commit_from_file(root.join("onnx/mlp.onnx"))
        .expect("加载模型");

    let tensor =
        Tensor::from_array((vec![rows as i64, FEATURES.len() as i64], inputs)).expect("构造输入");
    let outputs = session
        .run(ort::inputs!["features" => tensor])
        .expect("推理");
    let (shape, data) = outputs["output"]
        .try_extract_tensor::<f64>()
        .expect("取输出");
    assert_eq!(shape.len(), 1);
    assert_eq!(data.len(), expected.len(), "输出长度");

    for (position, (got, want)) in data.iter().zip(expected.iter()).enumerate() {
        if got.is_nan() && want.is_nan() {
            continue;
        }
        let tolerance = 1e-12 * 1.0f64.max(want.abs());
        assert!(
            (got - want).abs() <= tolerance,
            "第 {position} 位 {got} != {want}"
        );
    }
}
