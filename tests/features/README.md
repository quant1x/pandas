# tests/features —— 特征 fixture（M8 第一步）

同一份 OHLCV 输入与期望特征同表，Go / C++ / Rust 各自读取、重算、逐值对拍。
训练与推理共用这份样本：只要三端都过，"同一份特征在训练侧与生产侧一致"就有可执行的判据。

## 文件

| 文件 | 内容 |
| --- | --- |
| `gen_features.py` | 生成器：确定性价格路径 + 特征计算 |
| `features.csv` | 输入 OHLCV 与期望特征同表，40 根 Bar、11 个特征列 + `pred` + `pred_mlp` |
| `model.csv` | 冻结的线性模型参数（`term,weight`），推理侧只做线性组合 |
| `model_mlp.csv` | 冻结的两层小网络（11 → 4 → 1，tanh 隐层）参数 |
| `features_test.go` | Go 侧对拍（`go test ./tests/features/`） |
| `../../cpp/tests/features_test.cpp` | C++ 侧对拍（ctest 的 `features_fixture_test`） |
| `../../pandas-rust/tests/features.rs` | Rust 侧对拍（`cargo test --test features`） |
| `onnx/` | `model_mlp` 的 ONNX 导出与 onnxruntime 复核（三侧运行时未接，见其 README） |

## 列

```text
date,open,high,low,close,ma5,ma20,ema12,wma10,hhv20,llv20,std20,hhvbars20,ret1,bias5,cross_up,pred,pred_mlp
```

- 前五列是输入；其余是期望值，缺失写 `nan`；
- 特征只用已四端落地的公式口径组合：`ma` / `ema` / `wma` / `hhv` / `llv` / `std`（总体标准差）/
  `hhvbars` / `ref` / `cross`；
- `hhv20` / `llv20` / `hhvbars20` 取**最高价**窗口；`std20` / `bias5` / `wma10` 取收盘价；
- `ret1 = close / ref(close,1) - 1`，`bias5 = close / ma5 - 1`，`cross_up = cross(ma5, ma20)`（0/1）；
- 这是 fixture 的内部格式，不进入 D1 交换层。

`pred` 是推理侧口径：按 `model.csv` 的冻结权重做线性组合，求和顺序固定为特征列顺序，最后加 `bias`：

```text
pred = bias + Σ weight[term] × feature[term]
```

`pred_mlp` 是同一条链路的非线性版本（`model_mlp.csv`，11 → 4 → 1，tanh 隐层）：

```text
h_j = b1[j] + Σ_i W1[i][j] × f_i      （i 升序）
pred_mlp = b2 + Σ_j W2[j] × tanh(h_j)  （j 升序）
```

任一特征为 NaN 则两列都为 NaN。三端必须照抄求和顺序 —— 换顺序会改变浮点结果。
`pred_mlp` 是接 ONNX 之前的替身：结构与一次小网络推理相同，只是执行器是自己写的。

## 生成

```bash
python tests/features/gen_features.py
```

固定输入、无随机数，重复执行逐字节一致。

## 对拍

```bash
go test ./tests/features/
ctest --test-dir cpp/build -C Debug -R features_fixture_test
cargo test --manifest-path pandas-rust/Cargo.toml --test features
```

## 口径

- 滚动族窗口不足为 NaN（`ma` / `hhv` / `llv` / `std`），`ema` 首值等于源序列首值；
- 容差：相对 1e-12，NaN 对 NaN 视为相等；
- Zig 不接：`pandas/zig` 只有公式层，没有特征入口。

## 下一步

- 接 ONNX 推理侧（M8 的后半），同一模型在 C++ / Rust / Go 三侧输出一致；
- 特征集扩充时同步三端实现与 `gen_features.py`，不要只改一侧迁就。
