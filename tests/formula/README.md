# tests/formula —— 跨语言公式 fixture

同一份 CSV 由 NumPy 生成，Go / C++ / Rust / Zig 各自读取并计算，结果必须一致。当前覆盖
ABS / REF / MA / EMA / SMA / HHV / LLV / SUM / STD / WMA / SQRT / COUNT / FILTER /
HHVBARS / LLVBARS / CROSS / DMA / MAX / MIN，对应 M5 的 B1 / B2 / B3 / B6 对拍入口。
布尔序列在 CSV 里写作 0/1；`CROSS` 的输出同样落成 0/1。

## 格式

`cases.csv` 列：

```text
formula,x,y,n,m,expected
```

- 列表用 `;` 分隔，`nan` 表示缺失，未用字段留空；
- `n` / `m` 是标量参数（窗口、权重）；`y` 留给双序列公式；
- 这是 fixture 的内部格式，不进入 D1 交换层。

## 生成

```bash
python tests/formula/gen_fixtures.py
```

固定输入、无随机数，重复执行逐字节一致。

## 四端对拍

```bash
# Go（pandas/formula）
go test ./formula/ -run TestFormulaFixtures

# Rust（pandas-rust）
cargo test --manifest-path pandas-rust/Cargo.toml --test fixtures

# C++
cmake -S cpp -B cpp/build
cmake --build cpp/build
ctest --test-dir cpp/build -C Debug -R formula_fixture_test

# Zig（pandas/zig）
zig test zig/formula_test.zig
```

## 新增用例

1. 在 `gen_fixtures.py` 里追加一行（formula 名、输入、参数、期望）；
2. 四个语言的测试各补一个分支（Go / C++ / Rust / Zig）；
3. 重跑生成器与四端测试。

期望值由 Python 侧生成；与 Go 口径有差异时先在 issue 里记录，不直接改 fixture 迁就实现。
