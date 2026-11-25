# pandas-rust —— 公式族 Rust 版（B3）

`quant1x_pandas` 是 `pandas/formula` 的 Rust 对齐实现。当前覆盖 41 个基础公式，
签名与最小参数集对齐 Go 端：

| 公式 | Rust | 最小参数 |
| --- | --- | --- |
| ABS | `abs(&[f64]) -> Vec<f64>` | 序列 |
| REF | `ref_value(&[f64], n: i64) -> Vec<f64>` | 序列、周期 |
| MA | `ma(&[f64], window: i64) -> Vec<f64>` | 序列、窗口 |
| EMA | `ema(&[f64], span: i64) -> Vec<f64>` | 序列、周期 |
| SMA | `sma(&[f64], n: i64, m: i64) -> Vec<f64>` | 序列、周期、权重 |
| HHV / LLV | `hhv` / `llv(&[f64], window: i64)` | 序列、窗口 |
| SUM / STD | `sum` / `std(&[f64], window: i64)` | 序列、窗口 |
| COUNT | `count(&[bool], window: i64) -> Vec<f64>` | 条件序列、窗口 |
| WMA | `wma(&[f64], window: i64)` | 序列、窗口 |
| DMA | `dma(&[f64], a: &[f64])` | 序列、平滑周期（常量或逐位置） |
| CROSS | `cross(&[f64], &[f64]) -> Vec<bool>` | 两条序列 |
| HHVBARS / LLVBARS | `hhvbars` / `llvbars(&[f64], window: i64)` | 序列、窗口 |
| SLOPE / FORCAST | `slope` / `forcast(&[f64], window: i64)` | 序列、窗口 |
| AVEDEV | `avedev(&[f64], window: i64)` | 序列、窗口 |
| FILTER | `filter(&[f64], n: i64)` | 序列、屏蔽长度 |
| MAX / MIN | `max_series` / `min_series(&[f64], &[f64])` | 两条序列 |
| IF / IFF / IFN | `if_select` / `ifn(&[f64], &[f64], &[f64])` | 条件、两支 |
| BARSLAST / BARSLASTCOUNT | `barslast` / `barslastcount(&[f64])` | 条件序列 |
| BARSSINCEN | `barssincen(&[f64], window: i64)` | 条件序列、窗口 |
| GT / GTE / LT / LTE / EQ / NEQ | `gt` / `gte` / `lt` / `lte` / `eq` / `neq(&[f64], &[f64])` | 两条序列 |
| AND / OR / NOT | `and` / `or` / `not(&[bool], …)` | 布尔序列 |
| CONST | `const_series(&[f64])` | 序列（按 f32 精度取末值） |
| REF（逐位置） | `ref_series(&[f64], &[f64])` | 序列、逐位置周期 |
| EMA（逐位置） | `ema_series(&[f64], &[f64])` | 序列、逐位置周期 |
| SMA（序列周期） | `sma_series(&[f64], &[f64], m: i64)` | 序列、周期序列、权重 |
| SQRT | `sqrt(&[f64])` | 序列 |
| LAST | `last(&[f64], a: i64, b: i64)` | 条件序列、起止周期 |
| BARSLASTS | `barslasts(&[f64], n: i64)` | 条件序列、次数 |

## 构建与测试

```bash
# 要求 num 仓与本目录的父目录并排（path 依赖 ../../num）
cargo test
cargo clippy --all-targets -- -D warnings
cargo fmt --check
```

- `quant1x-num` 通过 path 依赖本地 num 仓；发布前需改成版本依赖，由发布流程处理。
- 版本占位 `0.0.0`，与 num 的 Rust crate 一致，不在代码里维护发布版本。

## 口径

- EMA：`alpha = 2/(N+1)`，`y0 = x0`，`y_t = (1-alpha)*y_{t-1} + alpha*x_t`（Go EWM `Adjust=false`）。
- SMA：`alpha = M/N`，`M=0` 视为 1。
- MA：只有完整窗口才计算，窗口不足为 NaN（复用 num 的 `rolling_apply`）。
- REF：等价于 `num::shift`，越界为 NaN。
- HHV / LLV / SUM / STD / WMA：完整窗口才计算，窗口不足为 NaN。
- COUNT：窗口不足计 0（与 Go 的空窗口计数一致）。
- DMA：`alpha = 1/A`，`A` 可逐位置；`1/A` 为 NaN 或 Inf 时按 1。
- CROSS：前一根 `S1<S2` 且当前 `S1>S2` 时为真，返回布尔序列。
- HHVBARS / LLVBARS：窗口内极值到当前的周期数，多个极值取最近一次。
- SLOPE / FORCAST：窗口内对下标做一元线性回归；SLOPE 取斜率，FORCAST 取最后一根的拟合值。
- AVEDEV：窗口内各值到均值的绝对偏差的平均。
- FILTER：扫描到非零值时，把其后 N 个位置置零；按处理后序列继续扫描（与 Go 一致）。
- MAX / MIN：逐元素取大 / 取小，短的一侧按 NaN 补齐。
- IF / IFF：条件非零取 A，否则取 B；IFN 取反。
- BARSLAST：上一次条件成立到当前的周期数，从未成立为 -1；BARSLASTCOUNT：连续成立计数。
- BARSSINCEN：窗口内第一次条件成立到当前的周期数；窗口内没有成立时为 0。
- GT/GTE/LT/LTE/EQ/NEQ 与 AND/OR/NOT 复用 num 的比较/逻辑实现。
- CONST 先转 float32 再复制（与 Go 的 Float32s 一致），大整数可能丢精度。
- `ref_series` / `ema_series` 按位置取周期；`sma_series` 与 Go 一致只取周期序列最后一个值。
- `LAST`：在 A+1 窗口内检查除最近 B 根外的位置是否都非零；窗口不足为 NaN，B 越界视为成立。
- `BARSLASTS`：按 Go 的递推实现（BARSLAST 计数 + 逐位置 REF 迭代 N-1 次），已用 Go 输出对拍。

剩余公式（HHV / LLV / CROSS / STD / SUM / COUNT 等）与 C++ 版本（B2）、Python 研究脚本（B4）
按 M5 计划继续补。
