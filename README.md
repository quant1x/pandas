pandas
===
[![Sourcegraph](https://sourcegraph.com/github.com/quant1x/pandas/-/badge.svg)](https://sourcegraph.com/github.com/quant1x/pandas?badge)
[![Build Status](https://api.travis-ci.com/repos/quant1x/pandas.png)](https://travis-ci.com/quant1x/pandas)
[![codecov](https://codecov.io/gh/quant1x/pandas/branch/master/graph/badge.svg)](https://codecov.io/gh/quant1x/pandas)
![Golang 1.27.3+](https://img.shields.io/badge/Golang-1.27.3+-orange.svg?style=flat)
![tag](https://img.shields.io/github/tag/quant1x/pandas.svg?style=flat)
![license](https://img.shields.io/github/license/quant1x/pandas.svg)

## 1. 介绍
golang版本的pandas

## 1.1 多语言实现与验收

除 Go 主实现外，公式族在四个方向对齐，口径以 `formula/`（Go）与 `tests/formula/cases.csv` 的期望值为准：

| 目录 | 语言 | 内容 |
| --- | --- | --- |
| `formula/` | Go | 通达信公式（B1），权威实现 |
| `cpp/` | C++ | 39 个公式 + 特征对拍（B2） |
| `pandas-rust/` | Rust | 41 个公式 + 特征与 ONNX 推理用例（B3） |
| `zig/` | Zig | 39 个公式（B6，扩展线） |
| `scripts/` | Python | 研究侧因子报告（B4，不参与交付） |

一键验收（生成器幂等 + Go / C++ / Rust / Zig + 可选 ONNX）：

```bash
python tests/run_all.py
```

样本格式、口径与新用例流程见 [tests/README.md](tests/README.md)；`.numb` 二进制格式见
`num` 仓的 `docs/DATA-FORMAT.md`。

## 2. 功能/模块划分

### 2.1 特性列表
| 模块        | 一级功能          | 二级功能         | 进展情况                                |
|:----------|:--------------|:-------------|:------------------------------------|
| dataframe | dataframe     | new          | [√]                                 |
| dataframe | 类型约束          | string       | [√]                                 |
| dataframe | 类型约束          | bool         | [√]                                 |
| dataframe | 类型约束          | int64        | [√]                                 |
| dataframe | 类型约束          | float64      | [√]                                 |
| dataframe | 泛型类型          | 支持全部的基础类型    | [√]                                 |
| dataframe | 泛型类型          | 自动检测类型       | [√] 优先级:string > bool > float > int |
| dataframe | align         | series长度自动对齐 | [√]                                 |
| dataframe | col           | 选择           | [√]                                 |
| dataframe | col           | 新增1列         | [√]                                 |
| dataframe | row           | 删除多行         | [√]                                 |
| dataframe | name          | 改名, 支持单一列改名  | [√]                                 |
| series    | series        | new          | [√] series的列元素类型和reflect.Kind保持一致   |
| series    | 伪泛型           | 构建           | [√] 再新建series完成之后类型就确定了             |
| series    | SeriesBool    | bool类型       | [√]                                 |
| series    | SeriesString  | string类型     | [√]                                 |
| series    | SeriesInt64   | int64类型      | [√]                                 |
| series    | SeriesFloat64 | float64类型    | [√]                                 |
|series | rolling       | 支持序列化参数      | [√]                                 |



## 3. 示例

### 3.1. dataframe
### 3.2. series

## 4. 参考的代码:
- https://github.com/go-gota/gota
- https://github.com/WinPooh32/series
- https://github.com/rocketlaunchr/dataframe-go.git
