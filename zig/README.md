# pandas/zig —— 公式族 Zig 版（B6）

`pandas/zig` 是 `pandas/formula` 的 Zig 扩展实现起点，与 Go / C++ / Rust 共用同一份公式口径。
当前 11 个公式：ABS / REF / MA / EMA / SMA / HHV / LLV / SUM / STD / WMA / SQRT，
序列用 `[]const f64` + `std.mem.Allocator`，返回调用方负责释放的 `[]f64`。

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

## 口径

- EMA / SMA 与 Go 的 EWM `Adjust=false` 一致：`y_t=(1-alpha)*y_{t-1}+alpha*x_t`；NaN 位置沿用前值，首值为 NaN 时按 0 起算。
- 空输入返回空序列；滚动族的窗口为 0 或超过长度时全部 NaN，只算完整窗口。
- 首批只覆盖 f64 与一维序列；整型与容器封装后续补。
- 函数名按 Zig 的 camelCase（`refValue`），对应合同名仍是 `REF`；`STD` 落为 `stddev`，与 C++ 侧命名一致。

## 后续

- 对齐 B2 / B3 的其余公式（COUNT / DMA / CROSS / IF / BARSLAST 等）。
- 评估改用 `pandas/` 根目录的 `build.zig` 统一驱动 Go / C++ / Rust / Zig 四端 fixture 对拍。
