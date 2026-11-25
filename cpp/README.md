# pandas/cpp —— 公式族 C++ 版（B2）

`quant1x_pandas_cpp` 是 `pandas/formula` 的 C++ 对齐实现起点。首批 13 个公式：ABS / REF / MA / EMA / SMA / HHV / LLV / SUM / STD / COUNT / WMA / DMA / CROSS，
序列用 `std::vector<double>`，不依赖 num / xtensor，保证在 pandas 仓内可独立构建。

## 构建与测试

```bash
cmake -S cpp -B cpp/build
cmake --build cpp/build
ctest --test-dir cpp/build -C Debug --output-on-failure   # VS 多配置生成器需要 -C
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

## 口径

- EMA / SMA 与 Go 的 EWM `Adjust=false` 一致：`y_t=(1-alpha)*y_{t-1}+alpha*x_t`。
- 空输入返回空序列；MA 的窗口为 0 或超过长度时全部 NaN。
- 首批只覆盖 f64；整型与 Series 容器后续补。
- `STD` 的 C++ 符号命名为 `stddev`，与 num 的 C++ 侧一致。

## 后续

- 先补 HHV / LLV / CROSS / SUM / STD / COUNT / WMA / DMA 等公式。
- 再评估是否把序列类型切到 num 的 C++ 包（`num::Vector`），届时与本仓的 `num_core` 安装包对接。
