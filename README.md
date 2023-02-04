pandas
===
[![Sourcegraph](https://sourcegraph.com/github.com/quant1x/pandas/-/badge.svg)](https://sourcegraph.com/github.com/quant1x/pandas?badge)
[![Build Status](https://api.travis-ci.com/repos/quant1x/pandas.png)](https://travis-ci.com/quant1x/pandas)
[![codecov](https://codecov.io/gh/quant1x/pandas/branch/master/graph/badge.svg)](https://codecov.io/gh/quant1x/pandas)
![Golang 1.11.4+](https://img.shields.io/badge/Golang-1.20+-orange.svg?style=flat)
![tag](https://img.shields.io/github/tag/quant1x/pandas.svg?style=flat)
![license](https://img.shields.io/github/license/quant1x/pandas.svg)

## 1. 介绍
golang版本的pandas

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

## 4. 参考的的代码:
- https://github.com/go-gota/gota
- https://github.com/WinPooh32/series
- https://github.com/rocketlaunchr/dataframe-go.git
