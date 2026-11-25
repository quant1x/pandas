#!/usr/bin/env python3
"""把 model_mlp.csv 的两层小网络导出为 ONNX，并用 onnxruntime 复核 pred_mlp 列。

依赖（研究侧，不入库）：`pip install onnx onnxruntime`。
用法（在 pandas/ 目录下）：

    python tests/features/onnx/export_mlp.py

产物：`mlp.onnx`（约 2 KB）。脚本同时读回 `../features.csv` 的 `pred_mlp` 列逐值比对，
不一致直接非零退出 —— 导出正确是接三侧运行时的前提。

全程用 double：fixture 的 `pred_mlp` 是 f64 口径，用 f32 会引入 ~1e-7 的差。
"""

from __future__ import annotations

import csv
import math
import os
import sys

import numpy as np
import onnx
import onnxruntime as ort
from onnx import TensorProto, helper, numpy_helper

HIDDEN = 4
FEATURES = ["ma5", "ma20", "ema12", "wma10", "hhv20", "llv20", "std20",
            "hhvbars20", "ret1", "bias5", "cross_up"]
TOLERANCE = 1e-12


def load_weights(path: str) -> dict[str, float]:
    with open(path, encoding="utf-8", newline="") as handle:
        rows = list(csv.reader(handle))[1:]
    return {name: float(value) for name, value in rows if name}


def load_features(path: str) -> tuple[np.ndarray, np.ndarray]:
    with open(path, encoding="utf-8", newline="") as handle:
        rows = list(csv.DictReader(handle))
    values = np.array([[float(row[name]) if row[name] != "nan" else np.nan for name in FEATURES]
                       for row in rows], dtype=np.float64)
    expected = np.array([float(row["pred_mlp"]) if row["pred_mlp"] != "nan" else np.nan
                         for row in rows], dtype=np.float64)
    return values, expected


def build_model(weights: dict[str, float]) -> onnx.ModelProto:
    w1 = np.array([[weights[f"w1_{i}_{j}"] for j in range(HIDDEN)]
                   for i in range(len(FEATURES))], dtype=np.float64)
    b1 = np.array([weights[f"b1_{j}"] for j in range(HIDDEN)], dtype=np.float64)
    w2 = np.array([[weights[f"w2_{j}"]] for j in range(HIDDEN)], dtype=np.float64)
    b2 = np.array([weights["b2"]], dtype=np.float64)

    nodes = [
        helper.make_node("Gemm", ["features", "w1", "b1"], ["hidden_pre"], name="fc1"),
        helper.make_node("Tanh", ["hidden_pre"], ["hidden"], name="tanh"),
        helper.make_node("Gemm", ["hidden", "w2", "b2"], ["output_2d"], name="fc2"),
        helper.make_node("Reshape", ["output_2d", "shape_out"], ["output"], name="flatten"),
    ]
    initializers = [
        numpy_helper.from_array(w1, "w1"),
        numpy_helper.from_array(b1, "b1"),
        numpy_helper.from_array(w2, "w2"),
        numpy_helper.from_array(b2, "b2"),
        numpy_helper.from_array(np.array([-1], dtype=np.int64), "shape_out"),
    ]
    graph = helper.make_graph(
        nodes,
        "mlp_11_4_1",
        [helper.make_tensor_value_info("features", TensorProto.DOUBLE, ["n", len(FEATURES)])],
        [helper.make_tensor_value_info("output", TensorProto.DOUBLE, ["n"])],
        initializer=initializers,
    )
    model = helper.make_model(graph, opset_imports=[helper.make_opsetid("", 17)])
    model.ir_version = 8
    model.producer_name = "quant1x-pandas"
    model.producer_version = "1"
    onnx.checker.check_model(model)
    return model


def main() -> int:
    here = os.path.dirname(os.path.abspath(__file__))
    weights = load_weights(os.path.join(here, "..", "model_mlp.csv"))
    model = build_model(weights)

    model_path = os.path.join(here, "mlp.onnx")
    onnx.save_model(model, model_path)

    features, expected = load_features(os.path.join(here, "..", "features.csv"))
    session = ort.InferenceSession(model_path, providers=["CPUExecutionProvider"])
    actual = session.run(["output"], {"features": features})[0].astype(np.float64)

    if actual.shape != expected.shape:
        print(f"shape {actual.shape} != {expected.shape}")
        return 1
    for index, (got, want) in enumerate(zip(actual, expected)):
        if math.isnan(got) and math.isnan(want):
            continue
        if abs(got - want) <= TOLERANCE * max(1.0, abs(want)):
            continue
        print(f"第 {index} 位 {got} != {want}")
        return 1

    print(f"wrote {model_path}；onnxruntime 与 pred_mlp 逐值一致（{len(expected)} 行）")
    return 0


if __name__ == "__main__":
    sys.exit(main())
