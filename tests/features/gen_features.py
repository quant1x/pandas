#!/usr/bin/env python3
"""生成特征 fixture（M8 的第一步：训练与推理共用同一份特征样本）。

输入是一段确定性的 OHLCV（无随机数），输出特征列只在已落地的公式口径上组合：
MA / EMA / HHV / LLV / STD / REF / CROSS。列顺序即 `features.csv` 的表头顺序。

CSV 列：date,open,high,low,close,ma5,ma20,ema12,wma10,hhv20,llv20,std20,hhvbars20,ret1,bias5,cross_up,pred
`pred` 由 `model.csv` 的冻结权重按表头顺序线性组合得出（M8 的推理侧口径）。
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


def weighted_average(x: list[float], window: int) -> list[float]:
    out = [math.nan] * len(x)
    denominator = window * (window + 1) / 2
    for i in range(window - 1, len(x)):
        block = x[i + 1 - window:i + 1]
        out[i] = sum((j + 1) * v for j, v in enumerate(block)) / denominator
    return out


def bars_since_max(x: list[float], window: int) -> list[float]:
    out = [math.nan] * len(x)
    for i in range(window - 1, len(x)):
        block = x[i + 1 - window:i + 1]
        extreme = max(block)
        position = max(j for j, v in enumerate(block) if v == extreme)
        out[i] = float(len(block) - 1 - position)
    return out


def bias(x: list[float], average: list[float]) -> list[float]:
    out = [math.nan] * len(x)
    for i in range(len(x)):
        if not math.isnan(average[i]):
            out[i] = x[i] / average[i] - 1.0
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


# 冻结模型权重：研究侧拟合后固定下来的参数，推理侧只做线性组合。
# 求和顺序固定为 features.csv 的表头顺序，最后加 bias —— 三端必须一致。
MODEL_WEIGHTS = [
    ("bias", 0.01),
    ("ma5", -0.05),
    ("ma20", 0.04),
    ("ema12", 0.03),
    ("wma10", -0.02),
    ("hhv20", 0.01),
    ("llv20", -0.01),
    ("std20", 0.5),
    ("hhvbars20", 0.002),
    ("ret1", 2.0),
    ("bias5", -1.5),
    ("cross_up", 0.8),
]

FEATURE_ORDER = ["ma5", "ma20", "ema12", "wma10", "hhv20", "llv20", "std20",
                 "hhvbars20", "ret1", "bias5", "cross_up"]


def predict(weights: dict[str, float], features: dict[str, float]) -> float:
    total = weights["bias"]
    for name in FEATURE_ORDER:
        value = features[name]
        if isinstance(value, float) and math.isnan(value):
            return math.nan
        total += weights[name] * value
    return total


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
    wma10 = weighted_average(close, 10)
    hhv20 = rolling_extreme(high, 20, True)
    llv20 = rolling_extreme(low, 20, False)
    std20 = rolling_std(close, 20)
    hhvbars20 = bars_since_max(high, 20)
    bias5 = bias(close, ma5)
    ret1 = change_rate(close)
    cross_up = crossing(ma5, ma20)

    weights = dict(MODEL_WEIGHTS)
    columns = {
        "ma5": ma5, "ma20": ma20, "ema12": ema12, "wma10": wma10,
        "hhv20": hhv20, "llv20": llv20, "std20": std20, "hhvbars20": hhvbars20,
        "ret1": ret1, "bias5": bias5, "cross_up": cross_up,
    }
    pred = [predict(weights, {name: values[i] for name, values in columns.items()})
            for i in range(BARS)]

    header = ["date", "open", "high", "low", "close",
              "ma5", "ma20", "ema12", "wma10", "hhv20", "llv20", "std20", "hhvbars20",
              "ret1", "bias5", "cross_up", "pred"]
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
            fmt(wma10[i]),
            fmt(hhv20[i]),
            fmt(llv20[i]),
            fmt(std20[i]),
            fmt(hhvbars20[i]),
            fmt(ret1[i]),
            fmt(bias5[i]),
            fmt(cross_up[i]),
            fmt(pred[i]),
        ])

    model_path = os.path.join(os.path.dirname(os.path.abspath(__file__)), "model.csv")
    with open(model_path, "w", encoding="utf-8", newline="") as handle:
        writer = csv.writer(handle, lineterminator="\n")
        writer.writerow(["term", "weight"])
        writer.writerows([[name, fmt(weight)] for name, weight in MODEL_WEIGHTS])

    path = os.path.join(os.path.dirname(os.path.abspath(__file__)), "features.csv")
    with open(path, "w", encoding="utf-8", newline="") as handle:
        writer = csv.writer(handle, lineterminator="\n")
        writer.writerow(header)
        writer.writerows(rows)
    print(f"wrote {path}")


if __name__ == "__main__":
    main()
