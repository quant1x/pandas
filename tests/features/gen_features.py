#!/usr/bin/env python3
"""生成特征 fixture（M8 的第一步：训练与推理共用同一份特征样本）。

输入是一段确定性的 OHLCV（无随机数），输出特征列只在已落地的公式口径上组合：
MA / EMA / HHV / LLV / STD / REF / CROSS。列顺序即 `features.csv` 的表头顺序。

CSV 列：date,open,high,low,close,ma5,ma20,ema12,hhv20,llv20,std20,ret1,cross_up
缺失写 `nan`；文件 UTF-8 无 BOM、LF、末尾一个换行，重复生成逐字节一致。
"""

from __future__ import annotations

import csv
import math
import os

BARS = 40


def price_path() -> list[float]:
    """确定性价格路径：缓慢上行后回撤，保证 ma5/ma20 至少交叉一次。"""
    path = []
    for i in range(BARS):
        base = 10.0 + 0.08 * i
        wave = 0.6 * math.sin(i / 2.5)
        dip = -1.8 * math.exp(-((i - 22) ** 2) / 40.0)
        path.append(round(base + wave + dip, 2))
    return path


def high_low(close: list[float]) -> tuple[list[float], list[float]]:
    highs = [round(value + 0.15 + 0.02 * (i % 3), 2) for i, value in enumerate(close)]
    lows = [round(value - 0.12 - 0.02 * ((i + 1) % 3), 2) for i, value in enumerate(close)]
    return highs, lows


def moving_average(x: list[float], window: int) -> list[float]:
    out = [math.nan] * len(x)
    for i in range(window - 1, len(x)):
        out[i] = sum(x[i + 1 - window:i + 1]) / window
    return out


def exponential_average(x: list[float], span: int) -> list[float]:
    alpha = 2.0 / (span + 1.0)
    out = [x[0]]
    for value in x[1:]:
        out.append((1.0 - alpha) * out[-1] + alpha * value)
    return out


def rolling_extreme(x: list[float], window: int, pick_max: bool) -> list[float]:
    out = [math.nan] * len(x)
    for i in range(window - 1, len(x)):
        block = x[i + 1 - window:i + 1]
        out[i] = max(block) if pick_max else min(block)
    return out


def rolling_std(x: list[float], window: int) -> list[float]:
    out = [math.nan] * len(x)
    for i in range(window - 1, len(x)):
        block = x[i + 1 - window:i + 1]
        mean = sum(block) / window
        out[i] = math.sqrt(sum((v - mean) ** 2 for v in block) / window)
    return out


def change_rate(x: list[float]) -> list[float]:
    out = [math.nan] * len(x)
    for i in range(1, len(x)):
        out[i] = x[i] / x[i - 1] - 1.0
    return out


def crossing(fast: list[float], slow: list[float]) -> list[float]:
    out = [0.0] * len(fast)
    for i in range(1, len(fast)):
        if fast[i - 1] < slow[i - 1] and fast[i] > slow[i]:
            out[i] = 1.0
    return out


def fmt(value: float) -> str:
    if isinstance(value, float) and math.isnan(value):
        return "nan"
    return repr(float(value))


def main() -> None:
    close = price_path()
    high, low = high_low(close)
    open_prices = [close[0]] + close[:-1]
    ma5 = moving_average(close, 5)
    ma20 = moving_average(close, 20)
    ema12 = exponential_average(close, 12)
    hhv20 = rolling_extreme(high, 20, True)
    llv20 = rolling_extreme(low, 20, False)
    std20 = rolling_std(close, 20)
    ret1 = change_rate(close)
    cross_up = crossing(ma5, ma20)

    header = ["date", "open", "high", "low", "close",
              "ma5", "ma20", "ema12", "hhv20", "llv20", "std20", "ret1", "cross_up"]
    rows = []
    for i in range(BARS):
        rows.append([
            f"2024-01-{i + 1:02d}",
            fmt(open_prices[i]),
            fmt(high[i]),
            fmt(low[i]),
            fmt(close[i]),
            fmt(ma5[i]),
            fmt(ma20[i]),
            fmt(ema12[i]),
            fmt(hhv20[i]),
            fmt(llv20[i]),
            fmt(std20[i]),
            fmt(ret1[i]),
            fmt(cross_up[i]),
        ])

    path = os.path.join(os.path.dirname(os.path.abspath(__file__)), "features.csv")
    with open(path, "w", encoding="utf-8", newline="") as handle:
        writer = csv.writer(handle, lineterminator="\n")
        writer.writerow(header)
        writer.writerows(rows)
    print(f"wrote {path}")


if __name__ == "__main__":
    main()
