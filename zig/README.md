# pandas/zig —— 公式族 Zig 版（B6）

`pandas/zig` 是 `pandas/formula` 的 Zig 扩展实现起点，与 Go / C++ / Rust 共用同一份公式口径。
当前 35 个公式：ABS / REF / MA / EMA / SMA / HHV / LLV / SUM / STD / WMA / SQRT /
COUNT / FILTER / HHVBARS / LLVBARS / CROSS / DMA / MAX / MIN / GT / GTE / LT / LTE /
EQ / NEQ / AND / OR / NOT / IF / IFN / BARSLAST / BARSLASTCOUNT / BARSLASTS /
BARSSINCEN / LAST，
序列用 `[]const f64` + `std.mem.Allocator`，返回调用方负责释放的切片。
35 个全部进 `tests/formula/cases.csv` 四端对拍。

Zig 版本基线 0.17.1（`std.Io` 新 IO 接口，`std.ArrayList` 为 unmanaged 形态）。

## 运行测试

在 `pandas/` 目录下：

```bash
zig test zig/formula_test.zig
```

也可以在 `pandas/zig/` 目录下执行 `zig test formula_test.zig`，fixture 路径两种工作目录都能解析。
本目录不提供 `build.zig`：单文件测试直接由 `zig test` 编译，产物走默认缓存目录。

## 公式

| 公式 | Zig | 说明 |
| --- | --- | --- |
| ABS | `abs(allocator, x)` | 逐元素绝对值 |
| REF | `refValue(allocator, x, n)` | 前 N 周期；N 为负向后引用；越界 NaN |
| MA | `ma(allocator, x, window)` | 完整窗口，窗口不足 NaN |
| EMA | `ema(allocator, x, span)` | alpha=2/(N+1)，`y0=x0` |
| SMA | `sma(allocator, x, n, m)` | alpha=M/N，M=0 视为 1 |
| HHV / LLV | `hhv` / `llv(allocator, x, window)` | 滚动最大 / 最小 |
| SUM | `sum(allocator, x, window)` | 滚动求和 |
| STD | `stddev(allocator, x, window)` | 滚动总体标准差（避开 `std` 命名） |
| WMA | `wma(allocator, x, window)` | 权重 1..N，分母 N(N+1)/2 |
| SQRT | `sqrt(allocator, x)` | 逐元素平方根 |
| COUNT | `count(allocator, condition, window)` | 窗口内非零个数；窗口不足计 0 |
| FILTER | `filter(allocator, condition, window)` | 条件成立后 N 位清零 |
| HHVBARS / LLVBARS | `hhvbars` / `llvbars(allocator, x, window)` | 极值点到当前的周期数，并列取最近 |
| CROSS | `cross(allocator, s1, s2)` | 金叉，返回 `[]bool` |
| DMA | `dma(allocator, x, a)` | alpha=1/A，A 可常量或逐位置 |
| MAX / MIN | `maxSeries` / `minSeries(allocator, x, y)` | 逐元素取大 / 小，要求等长 |
| GT / GTE / LT / LTE / EQ / NEQ | `gt` / `gte` / `lt` / `lte` / `eq` / `neq(allocator, x, y)` | 逐元素比较，返回 `[]bool`，要求等长 |
| AND / OR / NOT | `logicAnd` / `logicOr` / `logicNot(allocator, x[, y])` | 逻辑族，`and` / `or` 是关键字故加前缀 |
| IF / IFN | `ifSelect` / `ifn(allocator, condition, a, b)` | 条件选择，`if` 是关键字故加前缀；三条序列等长 |
| BARSLAST | `barslast(allocator, condition)` | 上一次成立到当前的周期数；从未成立为 -1 |
| BARSLASTCOUNT | `barslastcount(allocator, condition)` | 连续成立的周期数 |
| BARSLASTS | `barslasts(allocator, condition, n)` | 倒数第 N 次成立到当前的周期数 |
| BARSSINCEN | `barssincen(allocator, condition, window)` | N 周期内第一次成立到当前的周期数 |
| LAST | `last(allocator, condition, a, b)` | A+1 窗口排除最近 B 根后全非零 |

## 口径

- EMA / SMA 与 Go 的 EWM `Adjust=false` 一致：`y_t=(1-alpha)*y_{t-1}+alpha*x_t`；NaN 位置沿用前值，首值为 NaN 时按 0 起算。
- 空输入返回空序列；滚动族的窗口为 0 或超过长度时全部 NaN，只算完整窗口。
- 首批只覆盖 f64 与一维序列；整型与容器封装后续补。
- 函数名按 Zig 的 camelCase（`refValue` / `maxSeries`），对应合同名仍是 `REF` / `MAX`；`STD` 落为 `stddev`，与 C++ 侧命名一致。
- `CROSS` 与比较 / 逻辑族返回 `[]bool`，fixture 对拍时转 0/1；输入为 0/1 序列的公式（`COUNT` / `FILTER` / 逻辑族）按非零判断。
- 双序列公式要求等长，长度不一致返回 `error.LengthMismatch`；Go 侧按最长对齐、缺位补 0 的口径不在 Zig 里复制。
- `LAST` 的窗口不足位取 NaN，与 Go / C++ / Rust 一致（滚动族的统一口径）。

## 后续

- 公式族已与 C++ 的 35 个对齐且全部入表；后续补整型重载与容器封装。
- 评估改用 `pandas/` 根目录的 `build.zig` 统一驱动 Go / C++ / Rust / Zig 四端 fixture 对拍。
