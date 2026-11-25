#!/usr/bin/env python3
"""研究侧因子报告（B4）。

只读 `tests/features/features.csv`，对每个特征列算分布统计，输出 `tests/features/report.csv`：

```text
column,count,nan_ratio,mean,std,min,max,autocorr1
```

约束（B4 的定位）：

- **不 import 生产实现**（`formula/` / `cpp/` / `pandas-rust/`），自带一份 NumPy 计算，做交叉视角；
- 产物只用于研究对照，不参与交付链路；重复执行逐字节一致。
"""

from __future__ import annotations

import csv
import math
import os

import numpy as np

FEATURES = ["ma5", "ma20", "ema12", "wma10", "hhv20", "llv20", "std20",
            "hhvbars20", "ret1", "bias5", "cross_up", "pred", "pred_mlp"]


def autocorr(values: np.ndarray, lag: int = 1) -> float:
    """一阶自相关；样本不足或方差为 0 时返回 NaN。"""
    if values.size <= lag:
        return math.nan
    left, right = values[:-lag], values[lag:]
    left = left - left.mean()
    right = right - right.mean()
    denominator = math.sqrt(float((left * left).sum()) * float((right * right).sum()))
    if denominator == 0:
        return math.nan
    return float((left * right).sum()) / denominator


def fmt(value: float) -> str:
    if isinstance(value, float) and math.isnan(value):
        return "nan"
    return repr(float(value))


def main() -> None:
    here = os.path.dirname(os.path.abspath(__file__))
    features_path = os.path.join(here, "..", "tests", "features", "features.csv")
    with open(features_path, encoding="utf-8", newline="") as handle:
        rows = list(csv.DictReader(handle))

    output = []
    for name in FEATURES:
        column = np.array([float(row[name]) if row[name] != "nan" else np.nan for row in rows])
        finite = column[~np.isnan(column)]
        count = int(finite.size)
        nan_ratio = (column.size - count) / column.size if column.size else math.nan
        if count == 0:
            output.append([name, "0", fmt(nan_ratio), "nan", "nan", "nan", "nan", "nan"])
            continue
        output.append([
            name,
            str(count),
            fmt(nan_ratio),
            fmt(float(finite.mean())),
            fmt(float(finite.std(ddof=0))),
            fmt(float(finite.min())),
            fmt(float(finite.max())),
            fmt(autocorr(finite)),
        ])

    report_path = os.path.join(here, "..", "tests", "features", "report.csv")
    with open(report_path, "w", encoding="utf-8", newline="") as handle:
        writer = csv.writer(handle, lineterminator="\n")
        writer.writerow(["column", "count", "nan_ratio", "mean", "std", "min", "max", "autocorr1"])
        writer.writerows(output)
    print(f"wrote {report_path}（{len(output)} 列）")


if __name__ == "__main__":
    main()
