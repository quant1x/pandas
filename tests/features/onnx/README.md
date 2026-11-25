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

三侧都已接上，判据一致：`features.csv` 的 11 个特征 → `mlp.onnx` → 与 `pred_mlp` 逐列一致（相对 1e-12）。

| 语言 | 入口 | 依赖来源 |
| --- | --- | --- |
| Rust | `formula/tests/onnx.rs` | `ort` crate（`load-dynamic`） |
| C++ | `cpp/tests/onnx_test.cpp` | onnxruntime SDK 的 `include/` + `onnxruntime.lib` |
| Go | `tests/features/onnx.go` / `onnx_test.go`（`-tags onnx`） | 同一份头文件；运行时按绝对路径加载 dll |

### SDK 从哪来

不用 GitHub 也能拿到：NuGet 上的 `Microsoft.ML.OnnxRuntime` 包就是官方 C/C++ SDK（含
`build/native/include/onnxruntime_c_api.h`、`runtimes/win-x64/native/onnxruntime.{lib,dll}`）：

```bash
curl -L -o ort.nupkg   https://api.nuget.org/v3-flatcontainer/microsoft.ml.onnxruntime/1.32.0/microsoft.ml.onnxruntime.1.32.0.nupkg
unzip -q ort.nupkg "build/native/include/*" "runtimes/win-x64/native/*" -d ort-sdk
```

（pip 的 `onnxruntime` wheel 只有 dll、没有头文件，所以它帮不上 C++ / Go。）

### 三种运行方式

```bash
# Rust
ORT_DYLIB_PATH=<ort-sdk>/runtimes/win-x64/native/onnxruntime.dll cargo test --test onnx

# C++（配置期给 SDK，构建后自动把 dll 拷到可执行文件旁）
cmake -S cpp -B cpp/build -DPANDAS_WITH_ONNX=ON       -DPANDAS_ORT_INCLUDE_DIR=<ort-sdk>/build/native/include       -DPANDAS_ORT_LIB=<ort-sdk>/runtimes/win-x64/native/onnxruntime.lib       -DPANDAS_ORT_DLL=<ort-sdk>/runtimes/win-x64/native/onnxruntime.dll
cmake --build cpp/build --config Debug
ctest --test-dir cpp/build -C Debug -R onnx_fixture_test

# Go（只需头文件；运行时按 ORT_DYLIB_PATH 绝对路径加载）
CGO_ENABLED=1 CGO_CFLAGS="-I<ort-sdk>/build/native/include"   ORT_DYLIB_PATH=<ort-sdk>/runtimes/win-x64/native/onnxruntime.dll   go test -tags onnx ./tests/features/ -run TestONNX
```

三侧都遵循「没有 SDK 就跳过」：Rust 未设 `ORT_DYLIB_PATH`、Go 未设 `ORT_DYLIB_PATH`、
C++ 未开 `PANDAS_WITH_ONNX` 时不参与构建 / 直接 SKIP，CI 不受影响。

## 为什么仓库里没有动态库

onnxruntime 的二进制体积大、且按平台分发，不随仓库提交；`ort` 以 `load-dynamic` 方式在运行时加载，
所以构建不受影响、也不需要联网拉二进制。
