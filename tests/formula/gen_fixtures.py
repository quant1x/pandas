#!/usr/bin/env python3
"""生成跨语言公式 fixture（ABS/REF/MA/EMA/SMA 与 HHV/LLV/SUM/STD/WMA/SQRT）。

CSV 列：formula,x,y,n,m,expected。列表用 `;` 分隔，`nan` 表示缺失，未用字段留空。
"""

from __future__ import annotations

import csv
import math
import os


def fmt(value: float) -> str:
    if isinstance(value, float) and math.isnan(value):
        return "nan"
    return repr(float(value))


def series(values) -> str:
    return ";".join(fmt(v) for v in values)


def ema(x, span: int):
    alpha = 2.0 / (span + 1.0)
    out = [x[0]]
    for value in x[1:]:
        out.append((1.0 - alpha) * out[-1] + alpha * value)
    return out


def sma(x, n: int, m: int):
    alpha = (1 if m == 0 else m) / n
    out = [x[0]]
    for value in x[1:]:
        out.append((1.0 - alpha) * out[-1] + alpha * value)
    return out


def ma(x, window: int):
    out = [float("nan")] * len(x)
    for i in range(window - 1, len(x)):
        out[i] = sum(x[i + 1 - window:i + 1]) / window
    return out


def ref(x, n: int):
    out = [float("nan")] * len(x)
    for i in range(len(x)):
        source = i - n
        if 0 <= source < len(x):
            out[i] = x[source]
    return out


def rolling(x, window: int, apply):
    """完整窗口才计算，窗口不足为 NaN。"""
    out = [float("nan")] * len(x)
    if window <= 0 or window > len(x):
        return out
    for i in range(window - 1, len(x)):
        out[i] = apply(x[i + 1 - window:i + 1])
    return out


def wma(x, window: int):
    """权重 1..N，分母 N(N+1)/2。"""
    return rolling(x, window, lambda block: sum((i + 1) * v for i, v in enumerate(block))
                   / (window * (window + 1) / 2))


def std(x, window: int):
    """总体标准差：两遍算法，分母为 N。"""
    def population(block):
        mean = sum(block) / len(block)
        return math.sqrt(sum((v - mean) ** 2 for v in block) / len(block))

    return rolling(x, window, population)


def main() -> None:
    path = os.path.join(os.path.dirname(os.path.abspath(__file__)), "cases.csv")
    rows = []
    x = [1.0, -2.0, 3.0]
    rows.append(["abs", series(x), "", "", "", series([abs(v) for v in x])])

    ref_x = [1.0, 2.0, 3.0, 4.0, 5.0]
    rows.append(["ref", series(ref_x), "", 2, "", series(ref(ref_x, 2))])

    ma_x = [1.0, 2.0, 3.0, 4.0, 5.0]
    rows.append(["ma", series(ma_x), "", 3, "", series(ma(ma_x, 3))])

    ema_x = [1.0, 2.0, 3.0, 4.0]
    rows.append(["ema", series(ema_x), "", 4, "", series(ema(ema_x, 4))])

    sma_x = [1.0, 2.0, 3.0]
    rows.append(["sma", series(sma_x), "", 3, 1, series(sma(sma_x, 3, 1))])

    window_x = [1.0, 3.0, 2.0, 5.0, 4.0]
    rows.append(["hhv", series(window_x), "", 3, "", series(rolling(window_x, 3, max))])
    rows.append(["llv", series(window_x), "", 3, "", series(rolling(window_x, 3, min))])

    sum_x = [1.0, 2.0, 3.0, 4.0, 5.0]
    rows.append(["sum", series(sum_x), "", 2, "", series(rolling(sum_x, 2, sum))])

    std_x = [2.0, 4.0, 4.0, 4.0, 5.0, 5.0, 7.0, 9.0]
    rows.append(["std", series(std_x), "", 4, "", series(std(std_x, 4))])

    wma_x = [1.0, 2.0, 3.0, 4.0]
    rows.append(["wma", series(wma_x), "", 3, "", series(wma(wma_x, 3))])

    sqrt_x = [1.0, 4.0, 9.0, 16.0]
    rows.append(["sqrt", series(sqrt_x), "", "", "", series([math.sqrt(v) for v in sqrt_x])])

    with open(path, "w", encoding="utf-8", newline="") as handle:
        writer = csv.writer(handle, lineterminator="\n")
        writer.writerow(["formula", "x", "y", "n", "m", "expected"])
        writer.writerows(rows)
    print(f"wrote {path}")


if __name__ == "__main__":
    main()
