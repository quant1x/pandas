formula
===

***通达信指标函数实现***

- [通达信指标公式函数大全](https://www.chanluns.com/tdxfun/)

| 函数名      | 功能                   | 示例                | 固定参数 | 序列参数 |
|:---------|:---------------------|:------------------|:-----|:-----|
| ABS      | 绝对值                  | ABS(X)            | [√]  | [√]  |
| REF      | 引用N周期前的值             | REF(CLOSE, 5)     | [√]  | [√]  |
| IF       | 逻辑判断                 | IF(CLOSE>10,1,0)  | [√]  | [√]  |
| IFF      | 逻辑判断                 | IFF(CLOSE>10,1,2) | [√]  | [√]  |
| IFN      | 逻辑判断                 | IFN(CLOSE>10,1,2) | [√]  | [√]  |
| HHV      | 计算N周期内最高             | HHV(HIGH,5)       | [√]  | [√]  |
| LLV      | 计算N周期内最低             | LLV(HLOW,5)       | [√]  | [√]  |
| MAX      | 计算AB最大值              | MAX(CLOSE,HIGH)   | [√]  | [√]  |
| MIN      | 计算AB最小值              | MIN(CLOSE,HIGH)   | [√]  | [√]  |
| MA       | 计算N周期的移动平均值, 简称均线    | MA(CLOSE,5)       | [√]  | [√]  |
| SMA      | 计算N周期的简单移动平均值        | SMA(CLOSE,5, 1)   | [X]  | [X]  |
| STD      | 计算N周期内的标准差           | STD(CLOSE,N)      | [√]  | [√]  |
| SUM      | 求总和, 如果N=0则从第一个有效值开始 | SUM(CLOSE,N)      | [X]  | [X]  |
| BARSLAST | 上一次条件成立到当前的周期数       | BARSLAST(X)       | [X]  | [X]  |
