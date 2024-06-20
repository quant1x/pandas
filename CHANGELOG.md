# Changelog
All notable changes to this project will be documented in this file.

## [Unreleased]

## [1.4.11] - 2024-07-05
### Changed
- 更新依赖库num版本到0.3.5

## [1.4.10] - 2024-06-20
### Changed
- 更新依赖库num版本到0.3.4

## [1.4.9] - 2024-06-14
### Changed
- 更新依赖库版本
- update changelog

## [1.4.8] - 2024-05-16
### Changed
- 更新依赖库版本num到0.3.2
- update changelog
- update changelog

## [1.4.7] - 2024-05-11
### Changed
- 更新依赖库版本num到0.3.1
- update changelog

## [1.4.6] - 2024-05-11
### Changed
- 更新依赖库版本num到0.3.0
- update changelog

## [1.4.5] - 2024-04-16
### Changed
- 更新依赖库版本num到0.2.8
- 更新依赖库版本num到0.2.9
- update changelog

## [1.4.4] - 2024-03-30
### Changed
- 更新依赖库版本
- update changelog

## [1.4.3] - 2024-03-28
### Changed
- 更新依赖库版本
- update changelog

## [1.4.2] - 2024-03-21
### Changed
- 更新依赖库版本
- update changelog

## [1.4.1] - 2024-03-17
### Changed
- 更新依赖库版本
- update changelog

## [1.4.0] - 2024-03-12
### Changed
- 更新gox版本及go版本
- update changelog

## [1.3.9] - 2024-03-12
### Changed
- 更新num版本, 删除对excel的支持, 精简功能
- update changelog

## [1.3.8] - 2024-02-26
### Changed
- 增加测试数据
- 调整copy内部为clone, 增加按照索引赋值的函数
- 删除多余的空白行
- 调整EWM数据为clone
- 增加测试数据
- 优化Series的Rolling函数
- 优化Series的Sum方法
- 优化RollingAndExpandingMixin的Sum函数
- 优化RollingAndExpandingMixin的Std函数
- 优化RollingAndExpandingMixin的Mean函数
- 优化RollingAndExpandingMixin的Min函数
- 优化RollingAndExpandingMixin的Max函数
- 优化RollingAndExpandingMixin的Count函数
- 优化series子集的处理方式
- 优化sum函数
- 优化std函数
- 调整sma各阶段实现函数的名称
- 优化Ref函数
- 优化ma函数
- 优化llv函数
- 优化hhv函数
- 优化ema函数
- 优化cross函数
- 更新num版本
- update changelog

## [1.3.7] - 2024-02-24
### Changed
- 适配num新版本0.1.5
- update changelog

## [1.3.6] - 2024-02-24
### Changed
- 增加试验中的测试代码
- update changelog

## [1.3.5] - 2024-02-19
### Changed
- 新增example_test.go代码示例源文件, example_test.go是相同package下示例代码的固定文件名
- 调整Apply2函数的第二个参数名为inplace, 明确用法
- 更新README中go版本的提示
- 调整dataframe string类方法的源文件名
- 微调dataframe部分方法
- vector追加any类型数据, 允许传入Series
- 调整部分测试代码
- 新增多种不同需求的切片转Series函数
- update changelog

## [1.3.4] - 2024-02-19
### Changed
- 调整pandas关联度高的Series实现名为NDFrame
- 调整单元测试函数名
- 微调部分代码
- 调整vector为私有结构体,对外只提供Series接口
- update changelog

## [1.3.3] - 2024-02-18
### Changed
- 修正series接口Reverse函数的注释
- NDFrame结构体删除锁
- NDFrame标记为不推荐
- 修订Rolling参数类型错误
- 增加上证指数的测试数据
- 删除废弃的测试数据
- 修订README
- 修订.gitignore
- 重构NDArray
- 重构vector
- 抽象series接口
- 调整series方法
- dataframe适配series
- 优化部分底层公式函数
- 更新num版本
- update changelog

## [1.3.2] - 2024-02-12
### Changed
- 调整NDFrame源文件名前缀
- 调整函数名
- interface{}改为any
- 修订max作为局部变量名和内置函数max的冲突
- 移除gonum.org代码
- 切片到series
- 优化部分函数的切片内存
- 删除废弃的代码
- 删除废弃的代码
- 删除废弃的代码
- 补充公开方法的注释
- 调整函数名
- 调整NDFrame赋值函数名
- 修订Rename方法参数未使用的告警信息
- 统一NDArray接收器名
- 优化检测泛型any的元素类型函数
- 优化部分函数的内存开销
- 调整入参名
- 拆分stat中计算类函数
- 优化部分代码
- 调整包路径
- update changelog

## [1.3.1] - 2024-02-08
### Changed
- 统一结构体方法接收器
- update changelog

## [1.3.0] - 2024-01-28
### Changed
- 更新依赖库版本
- 新增EMA增量计算函数
- update changelog

## [1.2.9] - 2024-01-28
### Changed
- 更新gox版本号
- update changelog

## [1.2.8] - 2024-01-13
### Changed
- 更新gox版本号
- update changelog

## [1.2.7] - 2023-12-30
### Changed
- 更新gox版本号
- update changelog

## [1.2.6] - 2023-12-23
### Changed
- 更新gox版本号
- update changelog

## [1.2.5] - 2023-12-14
### Changed
- 更新gox版本号
- update changelog

## [1.2.4] - 2023-12-13
### Changed
- series增加Neq和Not方法
- update changelog

## [1.2.3] - 2023-12-12
### Changed
- RollingWindow支持series
- update changelog

## [1.2.2] - 2023-12-04
### Changed
- 更新依赖库gox版本
- update changelog

## [1.2.1] - 2023-12-03
### Changed
- 更新依赖库gox版本号, 调整gocsv的引用到pkg
- update changelog

## [1.2.0] - 2023-12-02
### Changed
- 更新依赖库版本号
- update changelog

## [1.1.9] - 2023-12-02
### Changed
- 测试REF是否支持周期0为当期的数据
- 暴露一个类型转换函数
- update changelog

## [1.1.8] - 2023-10-29
### Changed
- 更新依赖版本
- update changelog

## [1.1.7] - 2023-10-22
### Changed
- 更新依赖版本
- update changelog

## [1.1.6] - 2023-10-08
### Changed
- 增加测试MA的增量计算函数
- 更新依赖版本
- update changelog

## [1.1.5] - 2023-09-15
### Changed
- 更新依赖库版本
- update changelog

## [1.1.4] - 2023-09-15
### Changed
- 更新依赖库

## [1.1.3] - 2023-09-12
### Changed
- 更换golang.org/x/exp/slices为系统标准库
- update changelog

## [1.1.2] - 2023-09-10
### Changed
- 升级依赖库版本
- update changelog

## [1.1.1] - 2023-08-13
### Changed
- 升级go版本到1.21.0
- update changelog

## [1.1.0] - 2023-08-01
### Changed
- 增加逻辑操作OR
- update changelog

## [1.0.9] - 2023-07-24
### Changed
- 修正金叉一定要高于, 不包含等于
- update changelog

## [1.0.8] - 2023-07-21
### Changed
- 更新依赖库版本号
- update changelog

## [1.0.7] - 2023-07-15
### Changed
- 增加通达信函数BARSLASTS, 计算倒数第N次条件到现在的周期数
- 新增增加通达信函数BARSLASTS词条
- update changelog

## [1.0.6] - 2023-07-08
### Changed
- 更新依赖库版本号
- update changelog

## [1.0.5] - 2023-06-16
### Changed
- 拆分group和filter, 为后续优化做准备
- update changelog

## [1.0.4] - 2023-06-16
### Changed
- 引用api中的struct默认tag name
- update changelog

## [1.0.3] - 2023-06-16
### Changed
- 更新依赖库版本
- 删除废弃的功能
- update changelog

## [1.0.2] - 2023-06-14
### Changed
- 收敛ScopeLimit到gox工具库
- update changelog

## [1.0.1] - 2023-06-10
### Changed
- 旧版本的转换函数设置为不推荐使用, 并提示正确方法
- update changelog

## [1.0.0] - 2023-06-10
### Changed
- 统一CSV文件的tag为dataframe
- update changelog

## [0.9.37] - 2023-06-10
### Changed
- 调整csv组件依赖库
- update changelog

## [0.9.36] - 2023-06-07
### Changed
- 增加结构体切片保存到csv文件的函数
- update changelog

## [0.9.35] - 2023-06-06
### Changed
- 增加CSV转结构体切片的函数
- update changelog

## [0.9.34] - 2023-06-06
### Changed
- 更新依赖库版本
- 删除浮点操作的函数
- update changelog

## [0.9.33] - 2023-05-13
### Changed
- 更新依赖库版本
- update changelog

## [0.9.32] - 2023-05-12
### Changed
- 更新依赖库版本
- update changelog

## [0.9.31] - 2023-05-12
### Changed
- 更新依赖库版本
- update changelog

## [0.9.30] - 2023-05-10
### Changed
- 更新依赖库版本
- update changelog

## [0.9.29] - 2023-05-10
### Changed
- 更新依赖库版本
- update changelog

## [0.9.28] - 2023-05-10
### Changed
- 去掉读取文件时的错误日志的数据, 日志应该有应用层处理
- update changelog

## [0.9.27] - 2023-05-07
### Changed
- 调整package
- 升级gox版本号
- update changelog

## [0.9.26] - 2023-05-06
### Changed
- 更新依赖库版本号
- update changelog

## [0.9.25] - 2023-05-06
### Changed
- 调整文件名
- update changelog

## [0.9.24] - 2023-05-06
### Changed
- 梳理依赖
- update changelog

## [0.9.23] - 2023-05-06
### Changed
- 增加斜率计算
- update changelog

## [0.9.22] - 2023-05-05
### Changed
- 梳理依赖
- 梳理依赖
- 梳理依赖
- update changelog

## [0.9.21] - 2023-05-05
### Changed
- 更新版本号
- update changelog

## [0.9.20] - 2023-05-04
### Changed
- 收敛输出浮点的函数
- update changelog

## [0.9.19] - 2023-05-04
### Changed
- 修订NaN字符串解释异常的bug
- 浮点显示的时候, 小数点后3位四舍五入
- update changelog

## [0.9.18] - 2023-05-04
### Changed
- 支持是否四舍五入
- 支持是否四舍五入
- update changelog

## [0.9.17] - 2023-04-26
### Changed
- 更新依赖库版本
- update changelog

## [0.9.16] - 2023-04-26
### Changed
- 升级gox版本
- update changelog
- 更新依赖库版本

## [0.9.15] - 2023-04-19
### Changed
- !84 #I6W559 实现dataframe分组能力
* 实现分组功能
* 修复print在记录集少于maxRows序号错1位的bug
* dataframe支持从后到前的选择一个子集
- update changelog

## [0.9.14] - 2023-04-15
### Changed
- update modules
- update changelog

## [0.9.13] - 2023-04-09
### Changed
- 增加float64保留小数点, 目的是去掉科学计数法带来的精度问题
- docs: update changelog
- add changelog
- docs: update changelog
- test change
- docs: update changelog
- docs: update changelog

## [0.9.12] - 2023-04-05
### Changed
- 增加float64保留小数点, 目的是去掉科学计数法带来的精度问题

## [0.9.11] - 2023-03-24
### Changed
- 修订COUNT函数

## [0.9.10] - 2023-03-24
### Changed
- 修订MA为正式版本
- 修订IFF公式, 支持入参是常量

## [0.9.9] - 2023-03-18
### Changed
- 调整版本升级带来的变化

## [0.9.8] - 2023-03-18
### Changed
- 调整版本升级带来的变化

## [0.9.7] - 2023-03-11
### Changed
- 结构体支持无符号整整

## [0.9.6] - 2023-03-10
### Changed
- 修订版本号

## [0.9.5] - 2023-03-10
### Changed
- 修正版本号

## [0.9.4] - 2023-03-10
### Changed
- 修订版本号

## [0.9.3] - 2023-03-10
### Changed
- 升级data版本号
- 调整版本升级带来的变化

## [0.9.2] - 2023-03-09
### Changed
- 升级gox版本
- 修正写csv时对路径的检测

## [0.9.1] - 2023-03-09
### Changed
- dataframe.join支持多个series

## [0.9.0] - 2023-03-09
### Changed
- 删除AVX2默认打开的选项
- 补全pandas.sereis未实现的逻辑和算术方法

## [0.8.9] - 2023-03-08
### Changed
- 修订Empty对于类型的处理

## [0.8.8] - 2023-03-08
### Changed
- 修订数据版本

## [0.8.7] - 2023-03-08
### Changed
- 修订代码库

## [0.8.6] - 2023-03-08
### Changed
- Merge branch '0.8.x'

## [0.8.5] - 2023-03-08
### Changed
- 删除废弃的代码

## [0.8.4] - 2023-03-08
### Changed
- !83 #I6L2WO 支持ARM框架
* 支持arm64
* 修订SHIFT对N的series类型支持

## [0.8.3] - 2023-03-08
### Changed
- 修订SHIFT对N的series类型支持

## [0.8.2] - 2023-03-08
### Changed
- !82 #I6KYJ8 部分formula返回值改造成series
* 部分函数返回值从数组改成series

## [0.8.1] - 2023-03-05
### Changed
- 收敛and函数

## [0.8.0] - 2023-03-05
### Changed
- 强制float全部是float64

## [0.7.10] - 2023-03-05
### Changed
- 调整NDArray的and逻辑方法
- 收敛and函数

## [0.7.9] - 2023-03-05
### Changed
- series增加Bools方法
- 归类逻辑函数
- 修订部分注释
- 调整NDArray的4个逻辑方法

## [0.7.8] - 2023-03-04
### Changed
- 置信区间和Z分值的转换函数

## [0.7.7] - 2023-03-04
### Changed
- 更新data版本
- !81 #I6JM6X 新增置信区间和Z分值的转换函数
* 置信区间和Z分值的转换函数

## [0.7.6] - 2023-03-03
### Changed
- !80 #I6JB7Q 修复reverse对series的影响
* 修复reverse对series的影响

## [0.7.5] - 2023-03-03
### Changed
- !79 #I6JATS and支持布尔数组
* and增加识别[]bool

## [0.7.4] - 2023-03-03
### Changed
- !78 #I6JAQ1 WMA返回series
* wma返回series

## [0.7.3] - 2023-03-03
### Changed
- 升级quant1x.data版本

## [0.7.2] - 2023-03-01
### Changed
- !77 #I6IUM9 剔除go-clone
* #I6IUM9 剔除go-clone

## [0.7.1] - 2023-02-28
### Changed
- !76 #I6I6T7 修复显示超过最小行数时省略号也占1个行号的bug
* #I6I6T7 修复显示超过最小行数时省略号也占1个行号的bug
* 更新版本

## [0.7.0] - 2023-02-25
### Changed
- !75 #I6HES2 修订shape函数
* #I6HES2 修订shape函数
* !74 #I6HES3 实现numpy.dot函数
- !74 #I6HES3 实现numpy.dot函数
* #I6HES3 新增dot函数
- 合并dot和shape代码

## [0.6.23] - 2023-02-23
### Changed
- 修订依赖库版本

## [0.6.22] - 2023-02-22
### Changed
- !73 #I6GRQ5 给series添加Strings方法
* #I6GRQ5 给series添加Strings方法

## [0.6.21] - 2023-02-21
### Changed
- 修正版本号

## [0.6.20] - 2023-02-21
### Changed
- 修订data package路径

## [0.6.19] - 2023-02-21
### Changed
- 更新版本号

## [0.6.18] - 2023-02-21
### Changed
- !72 #I6GHRR 实现IndexOf方法，支持修改
* #I6GHRR 实现IndexOf方法, 支持修改
* 修订IndexOf函数的返回值
* series增加按切片下标取一行记录

## [0.6.17] - 2023-02-20
### Changed
- 选择列时支持可能会被修改, 直接返回对象而不是副本

## [0.6.16] - 2023-02-20
### Changed
- !71 #I6G7FP dataframe & series 实现concat方法和apply2, 支持替换元素的值
* #I6G7FP dataframe & series 实现concat方法和apply2, 支持替换元素的值
* series 新增concat方法

## [0.6.15] - 2023-02-20
### Changed
- !70 #I6G5HO 调整本地缓存数据的获取方式
* #I6G5HO 调整本地缓存数据的获取方式
* 调整data和gox版本
- 提供显式类型的apply方法

## [0.6.14] - 2023-02-19
### Changed
- 替换进度条工具库

## [0.6.13] - 2023-02-19
### Changed
- !69 #I6G4T0 修复读csv文件可能存在文件名是空串的bug
* #I6G4T0 修复读csv文件可能存在文件名是空串的bug
* 修订字符串转换缺失的数据类型

## [0.6.12] - 2023-02-19
### Changed
- 修订判断字符串的一处错误

## [0.6.11] - 2023-02-19
### Changed
- 修订版本交叉的问题

## [0.6.10] - 2023-02-19
### Changed
- fix data version

## [0.6.9] - 2023-02-18
### Changed
- fix package

## [0.6.8] - 2023-02-18
### Changed
- 修订package引用不一致的问题

## [0.6.7] - 2023-02-18
### Changed
- !68 修订部分问题, 清理package
* 删除早期的测试代码
* 删除早期的测试代码
* 删除tdx和v1版本的dataframe
* 删除data包, 引入gitee.com/quant1x/data工具包
* 更新gox版本
* 更新gitee脚本, 只做pull相关操作, 方便在github脚本中同步
* 调整部分函数的判断逻辑
* series 补充逻辑判断的方法
* 补充新扩展的接口, 暂未实现
* 修订注释, 增加测试代码
* 修复浮点0.xx会认为是字符串的bug
* 公式返回值序列化调整
* 删除无用的代码
* 删除无用的代码
* 删除无用的代码
* 删除无用的代码
* 修定参数名
* 修定参数名
* 修正注释
* 修复序列版本不支持负值向前移动
* 修复序列版本不支持负值向前移动
* 修正对齐函数, 长度相等直接返回
* 修订兼容性

## [0.6.6] - 2023-02-17
### Changed
- !53 #I6EMN5 实现乘法
* #I6EMN5 实现乘法
* #I6EMN5 实现乘法
- !51 #I6EI4S 修复策略运行慢的问题, 原因是rolling滑动窗口数据副本不是引用而克隆的问题，导致内存激增而GC等待
* 修复进度条乱闪的问题, 原因在于协程的进度不同, 导致控制台进度字符条跳跃式修正而闪动
* 增加判断channel是否关闭的函数
* 整理第三方库
* 调整测试参数
* 扩充数据类型
* 删除无用的代码
* 规范注释
* 修复泛型默认值的bug
* 规范注释
* 规范注释
* 滑动窗口的切片做引用处理, 不能clone
* 允许在任何时候关闭AVX2加速开关
* rolling时切片采用引用的方式
- !52 #I6EMN2 实现加法
* #I6EMN2 实现add函数
* 扩充数据类型
- #I6EMN5 实现乘法
- 修订测试用例
- !54 #I6EMN8 实现除法
* #I6EMN8 实现除法
- !55 #I6ENFJ 实现All函数
* #I6ENFJ 实现All函数
- !56 #I6ENG2 实现Count函数
* #I6ENG2 实现Count函数
- !57 #I6ENH7 实现any函数
* #I6ENH7 实现any函数
- !58 #I6EVCO 实现NDArray
* #I6EVCO 实现了NDArray
* 标注即将废弃的代码
* 标注即将废弃的代码
* 增加bool2int函数
* 扩充数据类型
* 删除废弃的代码
* 初步完成NDArray泛型数组的基础功能
* NDArray增加排序机制
* 调整文件分类
* 扩充数据类型
* 调整文件分类
* 调整文件分类
* 调整scope limit的package
* 调整scope limit的package
- !59 #I6F1P2 优化序列处理方式
* #I6F1P2 优化序列处理方式
* 删除num
* 召回第一版代码
* 调整package
- !60 #I6FC9Q 增加读取CPU信息
* #I6FC9Q 增加CPU和内存等信息的读取
- 修正内存变量名
- 调整函数的文件路径
- !61 #I6CYP6 实现通达信LLVBARS函数
* #I6CYP6 实现通达信LLVBARS函数
* 源文件改名
* ewm功能源文件改名
* 预备一个准备合并函数的"数字的"源文件
* 去掉废弃的代码
- !62 #I6CYP5 实现通达信HHVBARS函数
* #I6CYP5 实现通达信HHVBARS函数
- !63 #I6EV12 实现通达信FILTER函数
* #I6EV12 实现通达信FILTER函数
- !64 #I6CC23 实现KDJ指标
* #I6CC23 实现KDJ指标
* dataframe增加colAsNDArray泛型series的方法
* 修改ema返回值为stat.Series
* 调整部分函数的公开和私有属性
* fixed: rolling里面series没有重新赋值的问题
* 扩充数据类型
* 增加加载日线行情数据的函数
* 调整类型转换的源文件分类
* series 增加四则运算方法
* 调整函数名
* 调整方法的顺序
* 调整方法的顺序
* 调整方法的顺序
* 调整方法的顺序
* 调整方法的顺序
* 增加int32类型
- 修订KDJ注释
- !65 #I6CC21 实现MACD指标
* #I6CC21 实现MACD指标
- !66 #I6CC24 实现RSI指标
* #I6CC24 实现RSI指标
* 调整slice从any转float, 对于相同类型直接返回, 不clone
* 调整公式函数返回值为series
- !67 #I6CC6A 实现情绪指标BRAR
* #I6CC6A 实现情绪指标BRAR
* 修正SUM返回值为series
* 修订NDArray取值的问题

## [0.6.5] - 2023-02-12
### Changed
- !50 #I6CYO6 实现了1号策略在新框架下的运行
* #I6CYO6 实现了1号策略
* 去掉无用的代码
* 实现了滑动窗口
* diff代码有问题, 后面解决
* 优化代码
* 优化代码
* 优化代码
* 优化代码
* 优化代码
* 优化代码
* 优化代码
* 优化代码
* 优化代码
* 优化代码
* 优化代码
* 优化代码
* 暴露BaseType
* 新增zeros函数
* 调整代码层次
* 增加测试代码
* 新增一个lambda表达式操作方式的工具库
* 扩充数据类型
* 扩充数据类型
* 优化代码
* 屏蔽废弃的代码
* 扩充数据类型
* 调整函数的分类
* 扩充数据类型
* 扩充数据类型
* 范围限制结构改名
* 扩充数据类型
* 扩充数据类型
* 优化类型处理函数
* 增加打印异常堆栈信息的函数
* 处理默认值
* 调整代码结构
* 收敛处理流程
* 增加测试代码
* 收敛处理流程
* 统一收敛数据类型
* 增加测试代码
* 增加测试代码
* 增加测试代码
* 增加测试代码
* 规范代码
* 调整AVX2初始化的位置
* 增加测试代码
* 规范注释
* 增加测试点, 提高覆盖率

## [0.6.4] - 2023-02-10
### Changed
- !22 #I6C2TX 读写excel, 日期和代码字段的格式不太准确, 留到下期解决
* #I6C2TX 读写excel, 日期和代码字段的格式不太准确, 留到下期解决
* 预留quant1x系统目录
* 添加测试excel文件
* 调整写csv文件的文件名
* 约定只忽略test-*-w*格式的文件名, 这部分文件是测试中输出的文件
* 删除废弃的代码
* 拆分options
* 拆分options
* 拆分options
* 增加一个从gonum mat.Matrix的加载数据的函数
* 增加一个加载Map的函数
- !23 #I6CXFA fixed: 读写excel文件，时间和代码的格式处理不符合逻辑
* 修复excel读写的问题:
* 验证相同字符长度的浮点是否会判断为string
- !24 #I6CY6M MA函数
* 新增float32数据类型
* 新增float32数据类型
* 调整shift函数为独立一个源文件
* 暴露为公共函数
* 从变量名上明确int是64位的int64
* 删除废弃的代码
- !25 #I6CXOV 增加float32类型的series, 增加REF函数
* Merge branch 'ref' of https://gitee.com/quant1x/pandas into ref
* 实现REF函数
* 实现REF函数
* 新增float32数据类型
* 新增float32数据类型
* 调整shift函数为独立一个源文件
* 暴露为公共函数
* 从变量名上明确int是64位的int64
* 删除废弃的代码
- 调整功能归属
- 调整功能归属
- 调整功能归属
- !26 #I6CYOP 实现了HHV函数
* Merge branch 'hhv' of https://gitee.com/quant1x/pandas into hhv
* #I6CYOP 增加HHV函数
* #I6CYPC 增加HHV函数
* rolling增加max方法
* 新增不支持类型的异常
* 规制函数的归属到内建源文件
* 增加类型强制转换函数
* 增加泛型类型, 强化类型约束
* 选择记录集时使用copy选项
* 增加sort.Interface需要的除Len()之外的Less()和Swap两个方法
* 优化max
* 增加一个输出NaN的方法
* 测试字符串数组的排序方法
* 增加sort.Interface需要的除Len()之外的Less()和Swap两个方法
* 修订注释
- !27 #I6CYOQ 实现LLV函数
* #I6CYOQ 实现LLV函数
* 修复max方法在记录为空时的一个bug
* 修正注释
- !28 #I6CYPC 实现MAX函数
* #I6CYPC 实现MAX函数
* 增加加速开关的初始化
- !29 #I6CYPE 实现MIN函数, 发现avx2的max比对的bug,NaN比任何float都大,也比任何float都小
* #I6CYPE 实现MIN函数
* 修复avx2 max的bug
* vek的float数组最大值对比有bug, NaN比任何float都大, 也比任何float都小, 所以最大值比对的时候, 正确的应该是Na…
- !30 #I6CYPG 实现了IF, IFF, IFN三个函数
* #I6CYPG 实现IF, IFF, IFN三个函数
* 删除无用的代码
- !31 #I6CYPA 实现ABS绝对值函数
* #I6CYPA 实现ABS绝对值函数
- !32 #I6CYP9 实现了STD标准差计算函数及series.rolling的算法
* #I6CYP9 新增Std标准差方法及stat函数, 原gonum的stat.StdDev计算有变差, 可能是理解的有问题, 再同等输入参数…
* 修复一处any数组类型检测只跑了一层的bug,改为递归调用
* series接口增加Std方法
* 修复一处any数组类型检测只跑了一层的bug,改为递归调用
- !33 #I6CYP8 实现了SUM函数
* #I6CYP8 实现了SUM函数
* 早期版本的rolling函数改名为后缀加V1, 逐步会淘汰
* 新增README文档, 标注通达信指标公式函数的实现情况
- !34 #I6CYP3 实现WMA函数
* #I6CYP3 实现WMA函数, 返回数据类型有待统一
* 早期版本的rolling函数改名为后缀加V1, 逐步会淘汰
- !35 #I6CYP7 实现CONST函数
* #I6CYP7 实现CONST函数
* 补全0级和1级全部函数
* 补全0级和1级全部函数
- !36 #I6CYOV 实现了SMA。注意，通达信的SMA是个伪序列，不要使用，通达信在序列参数的情况下只会用最后一个
* #I6CYOV SMA实现了, 同时发现一个问题, 通达信实现的SMA时, 如果入参是序列, 其实是取的最后一个值
* EWM增加回调函数, 预备给扩展功能
* Rolling的空block要返回NaN
* 为了测试SMA,BARSLAST必须要先实现, 为了给SMA提供序列换参数, 以便验证, python那边还没实现
* 扩展repeat支持的类型
* 修订String方法, 能打印前后两段数据
* 修订String方法, 能打印前后两段数据
* #I6CYOV 实现SMA函数, 这个函数不支持序列换参数
- !37 #I6DLC4 实现BARSLAST函数
* #I6DLC4 实现BARSLAST函数
* 统一了数据类型以及切换接口
- !38 #I6DOBI 实现BARSSINCEN函数
* #I6DOBI 实现BARSSINCEN函数
* 增加3个函数, ArgMax, ArgMin, Median, 其中Median未加验证, 先带上.
- !39 #I6CYOT 实现了EMA函数
* #I6CYOT 实现EMA函数. 通达信的EMA函数的参数序列化, 验证有效.
* NDFrame.DTypes的克隆功能在EWM里面存在这个, 目前不影响使用, 后续优化
- !40 #I6CYP2 实现了DMA函数, 这个函数没有歧义
* #I6CYP2 实现了DMA函数, 这个函数没有歧义
* 如果不在预计范围内的类型, 抛异常
* 使用加速版本, NaN约定必须要处理, 否则可能出现无法预知的错误
* 修订注释
* fixed: 修复泛型类型扩展int32和int64, 会出现类型转换异常
* fixed: 修复泛型类型扩展int32和int64, 会出现类型转换异常
* fixed: 修复泛型类型扩展int32和int64, 会出现类型转换异常
* 修复泛型类型扩展int32和int64, 会出现类型转换异常, 新增加一个普通函数在这里
* 修正注释
* fixed: 修复泛型类型扩展int32和int64, 会出现类型转换异常
* fixed: 修复泛型类型扩展int32和int64, 会出现类型转换异常
* fixed: 修复泛型类型扩展int32和int64, 会出现类型转换异常
* fixed: 修复泛型类型扩展int32和int64, 会出现类型转换异常
* fixed: 修复Sum一处bug, 泛型类型扩展int32和int64, 会出现类型转换异常
* 新增fill函数, 支持默认值
* 新增泛型类型的接口定义
* 增加int64类型的处理函数
* 增加string类型的处理函数
* 调整alpha常量名
- #I6CYP2 修订DMA函数状态为已实现
- !41 #I6CYPB 实现SQRT求平方根函数
* #I6CYPB 实现SQRT求平方根函数
* 修订README新增SQRT求平方根函数
- !42 #I6CYP0 实现了COUNT, 增补了序列逻辑比较的4个函数
* #I6CYP0 实现COUNT函数
- !43 #I6CYP1 实现了LAST函数, A参数支持序列化, B不支持
* #I6CYP1 实现了LAST函数, A参数支持序列化, B不支持
- !44 #I6E1HG 实现BARSLASTCOUNT函数
* #I6E1HG 实现BARSLASTCOUNT函数
- !45 #I6E1KS 实现CROSS
* #I6E1KS 实现CROSS
- !46 #I6CYOY 实现SLOPE函数
* #I6CYOY 实现SLOPE函数
* 增加一个序列数生成函数
* 最小二乘法好多golang库实现和python numpy结果不一样, 原因在于范德蒙德矩阵(Vandermonde matrix)的生成,…
- !47 #I6CYOX 实现AVEDEV函数
* #I6CYOX 实现AVEDEV函数
* 新增mean函数
* 删除无用的代码
* 新增固定参数或者切片转成确定的切片参数
* 增加sub函数
* 完善sum函数
- !48 #I6CYOZ 实现FORCAST函数
* #I6CYOZ 实现FORCAST函数
* 简化测试代码
- !49 #I6CYOF 实现了通达信数据接口对接
* #I6CYOF 实现了通达信数据接口对接
* 增加索引对象和选择列的处理方法
* 修改原select选择记录的方法名

## [0.6.3] - 2023-02-04
### Changed
- #I6CC11 暴露apply方法为公共方法
- !16 #I6CC1Z series泛型方法
* #I6CC1Z 实现series 泛型构造方法, 早期源自gota的seriesXXX可以不用了。
* 增加选择列、改名的测试代码
* 增加选择列、改名功能
* 备注LoadRecords的函数特点
* 这段代码是实验, 考虑一个问题, 反射和.(type)哪个更快
* 增加注释
* 实现一个泛型帧
* 增加注释, 注明导入数据时的优先级
- !15 #I6CCAE 支持前后两个方向选择记录集
* #I6CCAE 支持范围选择, 从前和从后两个方向, 支持默认值
* 独立min功能出来一个源文件
* 独立max功能出来一个源文件
* 独立apply功能出来一个源文件
* 增加select方法
* 增加全局的私有变量
* abs增加注释
* 添加range的视线代码
- 增加注释, 注明导入数据时的优先级
- 实现一个泛型帧
- Merge branch '0.6.x' of https://gitee.com/quant1x/pandas into 0.6.x
- !17 #I6BMTE 实现了rolling的序列化版本
* #I6BMTE 实现序列换的rolling
* where函数调整为public函数
* 增加float的泛化类型
* 修复截断slice的初始化bug
* 增加float32和float64的泛型repeat函数
* 增加一个带错误号的异常结构
* dataframe使用any参数, 去掉原gota的用法
* fixed: 修复截取series的长度错误的bug
* 预备扩展的代码, 先占个位置
* 拆分rolling方法
* 备注泛型数据的扩展性问题
- !18 #I6CLQH dataframe新增右连接方法, 增加个series对齐方法
* #I6CLQH dataframe 增加对齐操作, 新增右连接方法
* 增加一个泛型计算最大的函数
* 扩展可以移动操作的泛型类型
* 修改变量名
- !19 #I6CB66 实现删除记录功能
* #I6CB66 实现dataframe/series删除记录的功能, 衍生出append批量增加记录的功能
* 统一方法定义方式
* 调整append函数的文件
* 补充功能实现列表
- !20 #I6CC11 暴露apply方法私有方法为公共方法
* #I6CC11 暴露apply方法为公共方法
- Merge branch '0.6.x' of https://gitee.com/quant1x/pandas into 0.6.x
- !21 #I6CC1K 实现DIFF功能
* #I6CC1K 实现DIFF功能, 包括固定参数和序列化参数
* 独立mean方法

## [0.6.2] - 2023-02-03
### Changed
- !7 #I6CCA6 series实现FillNa方法
* 补全bool/string/int64/float64的series
* #I6CCA6 series泛型版本实现fillna方法
* 修订一处注释错误
- 增加string类型NaN的判断条件
- 修订一处注释错误
- Merge branch '0.6.x' into a
- !8 #I6CCA7:dataframe实现FillNa方法
* #I6CCA7:dataframe实现FillNa方法
- !9 #I6C994: 支持dataframe选一段子集, 目前暂不支持索引, 只能按照行号区间来选择
* #I6C994: 支持dataframe选一段子集, 目前暂不支持索引, 只能按照行号区间来选择
- !10 #I6CC0Q 实现series最大值
* 实现了最大值的处理方式, 目前支持string,int64和float64
* 重新梳理类型转换函数
* 调整int64的转换函数
* 调整int64的转换函数
* 调整float64的转换函数
* 保留原case判断的位置, 代码注释掉
* 新增.gitignore
* 调整函数的相对位置
* 规范float64的nil全局变量名
* 修订解析float64字符串的函数
* 规范NaN的判断函数, 归于统一用法
* 解析string为float32失败后返回NaN
* 增加float64判断是否NaN的函数
* 增加float32判断是否NaN的函数
* 规范函数注释
* 修改float64转string的函数名
* 修改字符串解析float64的函数名称
* 增加bool转float32的函数
* 调整avx2的package路径
* math相关基础函数归于系统库
* math相关基础函数归于系统库
* 增加float64的切片处理方法
* 增加float32切片的处理方法
* 增加float32类型的转换
* 增加指针的判断函数
* 调整文件名与引用package同名
* 调整imports
* 调整imports
- !11 #I6CC0K series计算最小值
* 完成最小值的功能
* 调整类型的处理方式
* 又实现一个泛型构造方法,自动检测类型,从任意类型输入到series
* 修复泛型处理any时缺失nil处理的bug
* 调整series的type类型, 和reflect同步
* 预增float32的实现,理由有两个：
* 收敛float64的判断
* 收敛字符串出现NaN的判断条件
* 归档类型检测
- 修订github的代码同步脚本, 只push master分支和tags
- !12 #I6CC1C 新增stat.Where泛型函数, 用于做序列化三元判断
* #I6CC1C:新增Where的测试用例
* #I6CC1C:新增Where函数, slice元素类型泛型化
* 新增数据对齐功能的泛型函数
* 新增使用any类型作为默认参数的处理方式的参数检测函数
* 新增数据统计时需要的泛型类型声明
* 预备读切片的逻辑判断函数, 暂时用不到, 留给于总优化
* 新增单元测试时对float存在NaN的判断
* 归类默认函数
* 增加线性回归算法
* 为了不让GoLand折叠_string.go, 文件名改为_xstring.go
* 拆解print相关代码
* 修订README
- !13 #I6CC1H 实现序列化版本, 没有做优化处理, 留给于总实现吧
* 删除废弃的函数
- !14  #I6CC1H 实现序列化版本, 没有做优化处理, 留给于总实现吧
* #I6CC1H 实现abs绝对值的序列化版本
* 删除废弃的函数

## [0.6.1] - 2023-02-01
### Changed
- 实现了rolling固定参数的用法
- 增加Mean的一般用法
- fixed: 修订如果series长度为0还会继续计算的bug
- #I6C2X4实现string的series
- !2 #I6C2X4实现string的series
Merge pull request !2 from 王布衣/series-string
- 增加写csv功能
- 完成读写CSV文件的功能
- 增加csv函数读写功能变相默认参数的注释
- 新增参考类库https://github.com/rocketlaunchr/dataframe-go.git
- 删除gota代码
- add LICENSE.

Signed-off-by: 王布衣 <wangfengxy@sina.cn>
- 修订float64 NaN的判断条件
- 清理gota残余关联代码
- 清理go mod
- !1 实现csv文件的读写功能
Merge pull request !1 from 王布衣/csv
- !3 #I6C3A0 
* fixed: 修订如果series长度为0还会继续计算的bug
* 增加Mean的一般用法
* !2 #I6C2X4实现string的series
* #I6C2X4实现string的series
- Merge branch 'frame' of https://gitee.com/quant1x/pandas into frame
- Merge branch 'ewm' of https://gitee.com/quant1x/pandas into ewm
- !4  #I6C6UU 实现series复制功能
* 增加copy功能
* 增加泛型序列的实现
* 调整series接口的返回值, 去掉指针
* 调整series接口的返回值, 去掉指针
* repeat函数迁移到generic文件
* 删除废弃的源文件
* 修订any强制转换string的测试用例
* 测试代码增加golang的struct继承, 抽象的实现方法
- !5 #I6C70D 实现了rolling固定参数
Merge pull request !5 from 王布衣/rolling
- !6 #I6C8QK 新增3个脚本, 用于测试和同步仓库
* #I6C8QK 新增3个可执行脚本, 用于测试和同步代码

## [0.6.0] - 2023-01-31
### Changed
- 暂时提交代码, 然而优化的路线还没结束
- 调整package路径
- 完成float64的第一步
- 调整avx2路径
- 增加rolling功能, 实现mean函数
- 恢复gota原本的mean方法
- 新增series int64的初始化,方法都没实现
- 修复package
- 细化指数加权EWM计算方式
- 增加建议版本的dataframe
- 存在一个bug, 先放过
- 修订repeat用法
- 增加string,bool,int64的series实现

## [0.5.2] - 2023-01-28
### Changed
- 增加开启AVX2加速开关
- 增加加速测试的代码

## [0.5.1] - 2023-01-28
### Changed
- 修订注释

## [0.5.0] - 2023-01-28
### Changed
- 解决和series基准测试函数的冲突
- 解决和series基准测试函数的冲突
- 解决和series基准测试函数的冲突
- #I6BAJF fixed

## [0.2.2] - 2023-01-27
### Changed
- 修订README

## [0.2.1] - 2023-01-27
### Changed
- 修订README

## [0.2.0] - 2023-01-16
### Changed
- 调整package
- 调整package
- 调整package结构, 精简代码

## [0.1.0] - 2023-01-10
### Changed
- Initial commit
- 引入github.com/WinPooh32/series代码
- 引入github.com/go-gota/gota代码
- 引入github.com/go-gota/gota代码
- 修订WinHoop32的adjust=True一处bug, 没有计算第一个元素
- 优化package引入


[Unreleased]: https://gitee.com/quant1x/pandas.git/compare/v1.4.11...HEAD
[1.4.11]: https://gitee.com/quant1x/pandas.git/compare/v1.4.10...v1.4.11
[1.4.10]: https://gitee.com/quant1x/pandas.git/compare/v1.4.9...v1.4.10
[1.4.9]: https://gitee.com/quant1x/pandas.git/compare/v1.4.8...v1.4.9
[1.4.8]: https://gitee.com/quant1x/pandas.git/compare/v1.4.7...v1.4.8
[1.4.7]: https://gitee.com/quant1x/pandas.git/compare/v1.4.6...v1.4.7
[1.4.6]: https://gitee.com/quant1x/pandas.git/compare/v1.4.5...v1.4.6
[1.4.5]: https://gitee.com/quant1x/pandas.git/compare/v1.4.4...v1.4.5
[1.4.4]: https://gitee.com/quant1x/pandas.git/compare/v1.4.3...v1.4.4
[1.4.3]: https://gitee.com/quant1x/pandas.git/compare/v1.4.2...v1.4.3
[1.4.2]: https://gitee.com/quant1x/pandas.git/compare/v1.4.1...v1.4.2
[1.4.1]: https://gitee.com/quant1x/pandas.git/compare/v1.4.0...v1.4.1
[1.4.0]: https://gitee.com/quant1x/pandas.git/compare/v1.3.9...v1.4.0
[1.3.9]: https://gitee.com/quant1x/pandas.git/compare/v1.3.8...v1.3.9
[1.3.8]: https://gitee.com/quant1x/pandas.git/compare/v1.3.7...v1.3.8
[1.3.7]: https://gitee.com/quant1x/pandas.git/compare/v1.3.6...v1.3.7
[1.3.6]: https://gitee.com/quant1x/pandas.git/compare/v1.3.5...v1.3.6
[1.3.5]: https://gitee.com/quant1x/pandas.git/compare/v1.3.4...v1.3.5
[1.3.4]: https://gitee.com/quant1x/pandas.git/compare/v1.3.3...v1.3.4
[1.3.3]: https://gitee.com/quant1x/pandas.git/compare/v1.3.2...v1.3.3
[1.3.2]: https://gitee.com/quant1x/pandas.git/compare/v1.3.1...v1.3.2
[1.3.1]: https://gitee.com/quant1x/pandas.git/compare/v1.3.0...v1.3.1
[1.3.0]: https://gitee.com/quant1x/pandas.git/compare/v1.2.9...v1.3.0
[1.2.9]: https://gitee.com/quant1x/pandas.git/compare/v1.2.8...v1.2.9
[1.2.8]: https://gitee.com/quant1x/pandas.git/compare/v1.2.7...v1.2.8
[1.2.7]: https://gitee.com/quant1x/pandas.git/compare/v1.2.6...v1.2.7
[1.2.6]: https://gitee.com/quant1x/pandas.git/compare/v1.2.5...v1.2.6
[1.2.5]: https://gitee.com/quant1x/pandas.git/compare/v1.2.4...v1.2.5
[1.2.4]: https://gitee.com/quant1x/pandas.git/compare/v1.2.3...v1.2.4
[1.2.3]: https://gitee.com/quant1x/pandas.git/compare/v1.2.2...v1.2.3
[1.2.2]: https://gitee.com/quant1x/pandas.git/compare/v1.2.1...v1.2.2
[1.2.1]: https://gitee.com/quant1x/pandas.git/compare/v1.2.0...v1.2.1
[1.2.0]: https://gitee.com/quant1x/pandas.git/compare/v1.1.9...v1.2.0
[1.1.9]: https://gitee.com/quant1x/pandas.git/compare/v1.1.8...v1.1.9
[1.1.8]: https://gitee.com/quant1x/pandas.git/compare/v1.1.7...v1.1.8
[1.1.7]: https://gitee.com/quant1x/pandas.git/compare/v1.1.6...v1.1.7
[1.1.6]: https://gitee.com/quant1x/pandas.git/compare/v1.1.5...v1.1.6
[1.1.5]: https://gitee.com/quant1x/pandas.git/compare/v1.1.4...v1.1.5
[1.1.4]: https://gitee.com/quant1x/pandas.git/compare/v1.1.3...v1.1.4
[1.1.3]: https://gitee.com/quant1x/pandas.git/compare/v1.1.2...v1.1.3
[1.1.2]: https://gitee.com/quant1x/pandas.git/compare/v1.1.1...v1.1.2
[1.1.1]: https://gitee.com/quant1x/pandas.git/compare/v1.1.0...v1.1.1
[1.1.0]: https://gitee.com/quant1x/pandas.git/compare/v1.0.9...v1.1.0
[1.0.9]: https://gitee.com/quant1x/pandas.git/compare/v1.0.8...v1.0.9
[1.0.8]: https://gitee.com/quant1x/pandas.git/compare/v1.0.7...v1.0.8
[1.0.7]: https://gitee.com/quant1x/pandas.git/compare/v1.0.6...v1.0.7
[1.0.6]: https://gitee.com/quant1x/pandas.git/compare/v1.0.5...v1.0.6
[1.0.5]: https://gitee.com/quant1x/pandas.git/compare/v1.0.4...v1.0.5
[1.0.4]: https://gitee.com/quant1x/pandas.git/compare/v1.0.3...v1.0.4
[1.0.3]: https://gitee.com/quant1x/pandas.git/compare/v1.0.2...v1.0.3
[1.0.2]: https://gitee.com/quant1x/pandas.git/compare/v1.0.1...v1.0.2
[1.0.1]: https://gitee.com/quant1x/pandas.git/compare/v1.0.0...v1.0.1
[1.0.0]: https://gitee.com/quant1x/pandas.git/compare/v0.9.37...v1.0.0
[0.9.37]: https://gitee.com/quant1x/pandas.git/compare/v0.9.36...v0.9.37
[0.9.36]: https://gitee.com/quant1x/pandas.git/compare/v0.9.35...v0.9.36
[0.9.35]: https://gitee.com/quant1x/pandas.git/compare/v0.9.34...v0.9.35
[0.9.34]: https://gitee.com/quant1x/pandas.git/compare/v0.9.33...v0.9.34
[0.9.33]: https://gitee.com/quant1x/pandas.git/compare/v0.9.32...v0.9.33
[0.9.32]: https://gitee.com/quant1x/pandas.git/compare/v0.9.31...v0.9.32
[0.9.31]: https://gitee.com/quant1x/pandas.git/compare/v0.9.30...v0.9.31
[0.9.30]: https://gitee.com/quant1x/pandas.git/compare/v0.9.29...v0.9.30
[0.9.29]: https://gitee.com/quant1x/pandas.git/compare/v0.9.28...v0.9.29
[0.9.28]: https://gitee.com/quant1x/pandas.git/compare/v0.9.27...v0.9.28
[0.9.27]: https://gitee.com/quant1x/pandas.git/compare/v0.9.26...v0.9.27
[0.9.26]: https://gitee.com/quant1x/pandas.git/compare/v0.9.25...v0.9.26
[0.9.25]: https://gitee.com/quant1x/pandas.git/compare/v0.9.24...v0.9.25
[0.9.24]: https://gitee.com/quant1x/pandas.git/compare/v0.9.23...v0.9.24
[0.9.23]: https://gitee.com/quant1x/pandas.git/compare/v0.9.22...v0.9.23
[0.9.22]: https://gitee.com/quant1x/pandas.git/compare/v0.9.21...v0.9.22
[0.9.21]: https://gitee.com/quant1x/pandas.git/compare/v0.9.20...v0.9.21
[0.9.20]: https://gitee.com/quant1x/pandas.git/compare/v0.9.19...v0.9.20
[0.9.19]: https://gitee.com/quant1x/pandas.git/compare/v0.9.18...v0.9.19
[0.9.18]: https://gitee.com/quant1x/pandas.git/compare/v0.9.17...v0.9.18
[0.9.17]: https://gitee.com/quant1x/pandas.git/compare/v0.9.16...v0.9.17
[0.9.16]: https://gitee.com/quant1x/pandas.git/compare/v0.9.15...v0.9.16
[0.9.15]: https://gitee.com/quant1x/pandas.git/compare/v0.9.14...v0.9.15
[0.9.14]: https://gitee.com/quant1x/pandas.git/compare/v0.9.13...v0.9.14
[0.9.13]: https://gitee.com/quant1x/pandas.git/compare/v0.9.12...v0.9.13
[0.9.12]: https://gitee.com/quant1x/pandas.git/compare/v0.9.11...v0.9.12
[0.9.11]: https://gitee.com/quant1x/pandas.git/compare/v0.9.10...v0.9.11
[0.9.10]: https://gitee.com/quant1x/pandas.git/compare/v0.9.9...v0.9.10
[0.9.9]: https://gitee.com/quant1x/pandas.git/compare/v0.9.8...v0.9.9
[0.9.8]: https://gitee.com/quant1x/pandas.git/compare/v0.9.7...v0.9.8
[0.9.7]: https://gitee.com/quant1x/pandas.git/compare/v0.9.6...v0.9.7
[0.9.6]: https://gitee.com/quant1x/pandas.git/compare/v0.9.5...v0.9.6
[0.9.5]: https://gitee.com/quant1x/pandas.git/compare/v0.9.4...v0.9.5
[0.9.4]: https://gitee.com/quant1x/pandas.git/compare/v0.9.3...v0.9.4
[0.9.3]: https://gitee.com/quant1x/pandas.git/compare/v0.9.2...v0.9.3
[0.9.2]: https://gitee.com/quant1x/pandas.git/compare/v0.9.1...v0.9.2
[0.9.1]: https://gitee.com/quant1x/pandas.git/compare/v0.9.0...v0.9.1
[0.9.0]: https://gitee.com/quant1x/pandas.git/compare/v0.8.9...v0.9.0
[0.8.9]: https://gitee.com/quant1x/pandas.git/compare/v0.8.8...v0.8.9
[0.8.8]: https://gitee.com/quant1x/pandas.git/compare/v0.8.7...v0.8.8
[0.8.7]: https://gitee.com/quant1x/pandas.git/compare/v0.8.6...v0.8.7
[0.8.6]: https://gitee.com/quant1x/pandas.git/compare/v0.8.5...v0.8.6
[0.8.5]: https://gitee.com/quant1x/pandas.git/compare/v0.8.4...v0.8.5
[0.8.4]: https://gitee.com/quant1x/pandas.git/compare/v0.8.3...v0.8.4
[0.8.3]: https://gitee.com/quant1x/pandas.git/compare/v0.8.2...v0.8.3
[0.8.2]: https://gitee.com/quant1x/pandas.git/compare/v0.8.1...v0.8.2
[0.8.1]: https://gitee.com/quant1x/pandas.git/compare/v0.8.0...v0.8.1
[0.8.0]: https://gitee.com/quant1x/pandas.git/compare/v0.7.10...v0.8.0
[0.7.10]: https://gitee.com/quant1x/pandas.git/compare/v0.7.9...v0.7.10
[0.7.9]: https://gitee.com/quant1x/pandas.git/compare/v0.7.8...v0.7.9
[0.7.8]: https://gitee.com/quant1x/pandas.git/compare/v0.7.7...v0.7.8
[0.7.7]: https://gitee.com/quant1x/pandas.git/compare/v0.7.6...v0.7.7
[0.7.6]: https://gitee.com/quant1x/pandas.git/compare/v0.7.5...v0.7.6
[0.7.5]: https://gitee.com/quant1x/pandas.git/compare/v0.7.4...v0.7.5
[0.7.4]: https://gitee.com/quant1x/pandas.git/compare/v0.7.3...v0.7.4
[0.7.3]: https://gitee.com/quant1x/pandas.git/compare/v0.7.2...v0.7.3
[0.7.2]: https://gitee.com/quant1x/pandas.git/compare/v0.7.1...v0.7.2
[0.7.1]: https://gitee.com/quant1x/pandas.git/compare/v0.7.0...v0.7.1
[0.7.0]: https://gitee.com/quant1x/pandas.git/compare/v0.6.23...v0.7.0
[0.6.23]: https://gitee.com/quant1x/pandas.git/compare/v0.6.22...v0.6.23
[0.6.22]: https://gitee.com/quant1x/pandas.git/compare/v0.6.21...v0.6.22
[0.6.21]: https://gitee.com/quant1x/pandas.git/compare/v0.6.20...v0.6.21
[0.6.20]: https://gitee.com/quant1x/pandas.git/compare/v0.6.19...v0.6.20
[0.6.19]: https://gitee.com/quant1x/pandas.git/compare/v0.6.18...v0.6.19
[0.6.18]: https://gitee.com/quant1x/pandas.git/compare/v0.6.17...v0.6.18
[0.6.17]: https://gitee.com/quant1x/pandas.git/compare/v0.6.16...v0.6.17
[0.6.16]: https://gitee.com/quant1x/pandas.git/compare/v0.6.15...v0.6.16
[0.6.15]: https://gitee.com/quant1x/pandas.git/compare/v0.6.14...v0.6.15
[0.6.14]: https://gitee.com/quant1x/pandas.git/compare/v0.6.13...v0.6.14
[0.6.13]: https://gitee.com/quant1x/pandas.git/compare/v0.6.12...v0.6.13
[0.6.12]: https://gitee.com/quant1x/pandas.git/compare/v0.6.11...v0.6.12
[0.6.11]: https://gitee.com/quant1x/pandas.git/compare/v0.6.10...v0.6.11
[0.6.10]: https://gitee.com/quant1x/pandas.git/compare/v0.6.9...v0.6.10
[0.6.9]: https://gitee.com/quant1x/pandas.git/compare/v0.6.8...v0.6.9
[0.6.8]: https://gitee.com/quant1x/pandas.git/compare/v0.6.7...v0.6.8
[0.6.7]: https://gitee.com/quant1x/pandas.git/compare/v0.6.6...v0.6.7
[0.6.6]: https://gitee.com/quant1x/pandas.git/compare/v0.6.5...v0.6.6
[0.6.5]: https://gitee.com/quant1x/pandas.git/compare/v0.6.4...v0.6.5
[0.6.4]: https://gitee.com/quant1x/pandas.git/compare/v0.6.3...v0.6.4
[0.6.3]: https://gitee.com/quant1x/pandas.git/compare/v0.6.2...v0.6.3
[0.6.2]: https://gitee.com/quant1x/pandas.git/compare/v0.6.1...v0.6.2
[0.6.1]: https://gitee.com/quant1x/pandas.git/compare/v0.6.0...v0.6.1
[0.6.0]: https://gitee.com/quant1x/pandas.git/compare/v0.5.2...v0.6.0
[0.5.2]: https://gitee.com/quant1x/pandas.git/compare/v0.5.1...v0.5.2
[0.5.1]: https://gitee.com/quant1x/pandas.git/compare/v0.5.0...v0.5.1
[0.5.0]: https://gitee.com/quant1x/pandas.git/compare/v0.2.2...v0.5.0
[0.2.2]: https://gitee.com/quant1x/pandas.git/compare/v0.2.1...v0.2.2
[0.2.1]: https://gitee.com/quant1x/pandas.git/compare/v0.2.0...v0.2.1
[0.2.0]: https://gitee.com/quant1x/pandas.git/compare/v0.1.0...v0.2.0

[0.1.0]: https://gitee.com/quant1x/pandas.git/releases/tag/v0.1.0
