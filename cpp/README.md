# pandas/cpp —— 公式族 C++ 版（B2）

`quant1x_pandas_cpp` 是 `pandas/formula` 的 C++ 对齐实现起点。当前 39 个公式：ABS / REF / MA / EMA / SMA / HHV / LLV / SUM / STD / COUNT / WMA / DMA / CROSS / MAX / MIN / IF / IFN / BARSLAST / BARSLASTCOUNT / BARSSINCEN / FILTER / SQRT / GT / GTE / LT / LTE / EQ / NEQ / AND / OR / NOT / LAST / BARSLASTS / HHVBARS / LLVBARS / AVEDEV / SLOPE / FORCAST / CONST，
序列用 `std::vector<double>`，不依赖 num / xtensor，保证在 pandas 仓内可独立构建。

## 构建与测试

```bash
cmake -S cpp -B cpp/build
cmake --build cpp/build
ctest --test-dir cpp/build -C Debug --output-on-failure   # VS 多配置生成器需要 -C
# 其中 formula_fixture_test 读取 ../tests/formula/cases.csv
```

`cpp/build/` 是构建产物，不入库。

## 公式

| 公式 | C++ | 说明 |
| --- | --- | --- |
| ABS | `quant1x::pandas::abs` | 逐元素绝对值 |
| REF | `ref(x, n)` | 前 N 周期；N 为负向后引用；越界 NaN |
| MA | `ma(x, window)` | 完整窗口，窗口不足 NaN |
| EMA | `ema(x, span)` | alpha=2/(N+1)，`y0=x0` |
| SMA | `sma(x, n, m)` | alpha=M/N，M=0 视为 1 |
| HHV / LLV | `hhv` / `llv(x, window)` | 滚动最大 / 最小 |
| SUM | `sum(x, window)` | 滚动求和 |
| STD | `stddev(x, window)` | 滚动总体标准差（避开 C++ 的 `std`） |
| COUNT | `count(condition, window)` | 非零计数；窗口不足 0 |
| WMA | `wma(x, window)` | 权重 1..N |
| DMA | `dma(x, a)` | alpha=1/A，A 可常量/逐位置 |
| CROSS | `cross(s1, s2)` | 前一根 S1<S2 且当前 S1>S2 |
| MAX / MIN | `max_series` / `min_series` | 逐元素取大 / 小 |
| IF / IFF / IFN | `if_select` / `ifn` | 条件选择 |
| BARSLAST / BARSLASTCOUNT | `barslast` / `barslastcount` | 距上次成立 / 连续成立 |
| BARSSINCEN | `barssincen(condition, window)` | 窗口内首次成立到当前 |
| FILTER | `filter(x, n)` | 条件后 N 位清零 |
| SQRT | `sqrt(x)` | 逐元素平方根 |
| GT / GTE / LT / LTE / EQ / NEQ | `gt` / `gte` / `lt` / `lte` / `eq` / `neq` | 比较，缺位 NaN |
| AND / OR / NOT | `logic_and` / `logic_or` / `logic_not` | 布尔逻辑 |
| LAST | `last(condition, a, b)` | A+1 窗口排除最近 B 根 |
| BARSLASTS | `barslasts(condition, n)` | 倒数第 N 次成立 |
| HHVBARS / LLVBARS | `hhvbars` / `llvbars(x, window)` | 极值点到当前的周期数，并列取最近 |
| AVEDEV | `avedev(x, window)` | 窗口内与均值之差的绝对值的平均 |
| SLOPE | `slope(x, window)` | 窗口内线性回归斜率 |
| FORCAST | `forcast(x, window)` | 线性回归在窗口最后一根的预测值 |
| CONST | `const_series(x)` | 取最后值为常量序列（按 f32 精度） |

## 口径

- EMA / SMA 与 Go 的 EWM `Adjust=false` 一致：`y_t=(1-alpha)*y_{t-1}+alpha*x_t`。
- 空输入返回空序列；MA 的窗口为 0 或超过长度时全部 NaN。
- 首批只覆盖 f64；整型与 Series 容器后续补。
- `STD` 的 C++ 符号命名为 `stddev`，与 num 的 C++ 侧一致。

## 后续

- 补齐 B3 剩下的条件与逻辑公式（IF / BARSLAST 系列已有，后续对齐整型与 Series 容器）。
- 再评估是否把序列类型切到 num 的 C++ 包（`num::Vector`），届时与本仓的 `num_core` 安装包对接。
