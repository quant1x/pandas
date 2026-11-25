# tests/features —— 特征 fixture（M8 第一步）

同一份 OHLCV 输入与期望特征同表，Go / C++ / Rust 各自读取、重算、逐值对拍。
训练与推理共用这份样本：只要三端都过，"同一份特征在训练侧与生产侧一致"就有可执行的判据。

## 文件

| 文件 | 内容 |
| --- | --- |
| `gen_features.py` | 生成器：确定性价格路径 + 特征计算 |
| `features.csv` | 输入 OHLCV 与期望特征同表，40 根 Bar |
| `features_test.go` | Go 侧对拍（`go test ./tests/features/`） |
| `../../cpp/tests/features_test.cpp` | C++ 侧对拍（ctest 的 `features_fixture_test`） |
| `../../pandas-rust/tests/features.rs` | Rust 侧对拍（`cargo test --test features`） |

## 列

```text
date,open,high,low,close,ma5,ma20,ema12,hhv20,llv20,std20,ret1,cross_up
```

- 前五列是输入；其余是期望值，缺失写 `nan`；
- 特征只用已四端落地的公式口径组合：`ma` / `ema` / `hhv` / `llv` / `std`（总体标准差）/ `ref` / `cross`；
- `ret1 = close / ref(close,1) - 1`，`cross_up = cross(ma5, ma20)`（0/1）；
- 这是 fixture 的内部格式，不进入 D1 交换层。

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
