# scripts —— 研究侧脚本（B4）

本目录放**不参与交付**的研究脚本：只读 fixture、产出对照报告。定位见 `docs/PLAN.md` 的 B4。

| 脚本 | 输入 | 输出 | 用途 |
| --- | --- | --- | --- |
| `factor_report.py` | `tests/features/features.csv` | `tests/features/report.csv` | 每个特征列的分布统计（覆盖数、缺失率、均值、标准差、极值、一阶自相关） |

约束：

- **不 import 生产实现**（`formula/` / `cpp/` / `pandas-rust/` / `zig/`），自带一份 NumPy 计算，
  用来提供「另一条路径」的视角；生产侧一致性由 `tests/formula` 与 `tests/features` 的对拍负责；
- 输出固定、无随机数，重复执行逐字节一致（`tests/run_all.py` 会校验）；
- 只写 `tests/features/report.csv`，不碰其它文件。

## 用法

```bash
python scripts/factor_report.py
```

`report.csv` 的列：

```text
column,count,nan_ratio,mean,std,min,max,autocorr1
```

`count` 是有效样本数（NaN 不计），`std` 为总体口径（ddof=0），`autocorr1` 在有效样本不足或方差为 0 时写 `nan`。
