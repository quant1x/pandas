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

## 三侧运行时（未接）

ONNX 文件已就绪，生产侧还差运行时。需要在 issue 里定方案，建议三侧共用同一份
官方 release 动态库，版本天然一致：

| 语言 | 接入方式 | 备注 |
| --- | --- | --- |
| C++ | CMake `FetchContent` 拉官方 win/linux 包（带 SHA256），链 C API | 不把二进制打进仓库 |
| Rust | `ort` crate + `load-dynamic` 特性，运行时指向同一份动态库 | 避免 crate 自带第二份运行时 |
| Go | `onnxruntime_go` + cgo，`ONNXRUNTIME_LIB` 指向同一路径 | 需 `CGO_ENABLED=1` |

接上后判据不变：`features.csv` 的输入喂给三侧运行时，输出必须与 `pred_mlp` 一致。
