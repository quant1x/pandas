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


def count(condition, window: int):
    """窗口内非零条件的个数；窗口不足计 0。"""
    out = [0.0] * len(condition)
    if window <= 0 or window > len(condition):
        return out
    for i in range(window - 1, len(condition)):
        out[i] = float(sum(1 for v in condition[i + 1 - window:i + 1] if v != 0))
    return out


def bars_since(condition, window: int, pick_max: bool):
    """极值点到当前的周期数：并列时取最近的一个。"""
    def distance(block):
        extreme = (max if pick_max else min)(block)
        position = max(i for i, v in enumerate(block) if v == extreme)
        return float(len(block) - 1 - position)

    return rolling(condition, window, distance)


def cross(s1, s2):
    """金叉：前一根 S1<S2 且当前 S1>S2。"""
    out = [0.0] * len(s1)
    for i in range(1, min(len(s1), len(s2))):
        if s1[i - 1] < s2[i - 1] and s1[i] > s2[i]:
            out[i] = 1.0
    return out


def dma(x, a):
    """动态移动平均：alpha = 1/A，非有限按 1，首值等于源序列首值。"""
    out = [x[0]]
    for i in range(1, len(x)):
        period = a[0] if len(a) == 1 else a[i]
        alpha = 1.0 / period
        if not math.isfinite(alpha):
            alpha = 1.0
        out.append(alpha * x[i] + (1.0 - alpha) * out[-1])
    return out


def filter_signal(condition, window: int):
    """条件成立后 N 位清零，扫描已处理的序列。"""
    out = list(condition)
    for i in range(len(out)):
        if out[i] == 0:
            continue
        for j in range(i + 1, min(i + 1 + window, len(out))):
            out[j] = 0.0
    return out



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

    condition = [1.0, 0.0, 1.0, 1.0, 0.0, 1.0]
    rows.append(["count", series(condition), "", 3, "", series(count(condition, 3))])
    rows.append(["filter", series(condition), "", 2, "", series(filter_signal(condition, 2))])

    bars_x = [1.0, 5.0, 3.0, 5.0, 2.0]
    rows.append(["hhvbars", series(bars_x), "", 3, "", series(bars_since(bars_x, 3, True))])
    rows.append(["llvbars", series(bars_x), "", 3, "", series(bars_since(bars_x, 3, False))])

    cross_x = [1.0, 2.0, 1.0, 3.0]
    cross_y = [1.5, 1.5, 2.5, 2.5]
    rows.append(["cross", series(cross_x), series(cross_y), "", "", series(cross(cross_x, cross_y))])

    dma_x = [1.0, 2.0, 3.0, 4.0]
    dma_a = [2.0, 2.0, 2.0, 2.0]
    rows.append(["dma", series(dma_x), series(dma_a), "", "", series(dma(dma_x, dma_a))])

    max_x = [1.0, 5.0, 3.0]
    max_y = [2.0, 4.0, 6.0]
    rows.append(["max", series(max_x), series(max_y), "", "", series([max(a, b) for a, b in zip(max_x, max_y)])])
    rows.append(["min", series(max_x), series(max_y), "", "", series([min(a, b) for a, b in zip(max_x, max_y)])])

    with open(path, "w", encoding="utf-8", newline="") as handle:
        writer = csv.writer(handle, lineterminator="\n")
        writer.writerow(["formula", "x", "y", "n", "m", "expected"])
        writer.writerows(rows)
    print(f"wrote {path}")


if __name__ == "__main__":
    main()
