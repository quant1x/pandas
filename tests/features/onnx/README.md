# tests/features/onnx —— 特征模型导出（M8 后半的前置）

`mlp.onnx` 是 `../model_mlp.csv` 那份两层小网络（11 → 4 → 1，tanh）的 ONNX 表达：

```text
features[*,11] → Gemm(11×4) → Tanh → Gemm(4×1) → Reshape → output[*]
```

权重与 `model_mlp.csv` 逐项相同，全程 `double`——fixture 的 `pred_mlp` 是 f64 口径，
用 f32 会引入约 1e-7 的差，无法与 1e-12 容差对上。

## 生成与复核

```bash
pip install -r tests/features/onnx/requirements.txt
python tests/features/onnx/export_mlp.py
```

脚本会重写 `mlp.onnx`，再用 `onnxruntime` 跑一遍 `../features.csv` 的输入，
与 `pred_mlp` 列逐值比对（相对 1e-12）；不一致即非零退出。
导出结果确定：同一份权重连跑两次 `mlp.onnx` 的 md5 不变。

`mlp.onnx` 是产物，随 fixture 一起入库（766 字节）。

## 生产侧运行时

| 语言 | 状态 | 接法 |
| --- | --- | --- |
| Rust | **已接并复核** | `ort` crate（`load-dynamic`，构建期不下载二进制），见 `pandas-rust/tests/onnx.rs` |
| C++ | 待接 | 需 onnxruntime C API 头文件与导入库（官方包或 `ort-sys` 自带头） |
| Go | 待接 | 需 cgo + 同一份头文件与库，`CGO_ENABLED=1` |

Rust 侧的运行方式（把 `ORT_DYLIB_PATH` 指向任意一份 onnxruntime 动态库）：

```bash
ORT_DYLIB_PATH=/path/to/onnxruntime.dll cargo test --manifest-path pandas-rust/Cargo.toml --test onnx
```

未设置该变量时用例**跳过**（CI 上通常没有这份库）；设置后它会读 `features.csv` 的 11 个特征，
跑 `mlp.onnx`，与 `pred_mlp` 列逐值比对（相对 1e-12，NaN 对 NaN 视为相等）。

判据不变：输入喂给运行时，输出必须与 `pred_mlp` 一致。C++ / Go 接上后同样按这条判。

## 为什么仓库里没有动态库

onnxruntime 的二进制体积大、且按平台分发，不随仓库提交；`ort` 以 `load-dynamic` 方式在运行时加载，
所以构建不受影响、也不需要联网拉二进制。
