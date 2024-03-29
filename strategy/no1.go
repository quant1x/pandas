package main

import (
	"gitee.com/quant1x/data/cache"
	"gitee.com/quant1x/data/security"
	pandas "gitee.com/quant1x/pandas"
	. "gitee.com/quant1x/pandas/formula"
	"github.com/mymmsc/gox/logger"
	"github.com/mymmsc/gox/util/treemap"
)

// FormulaNo1 3天内5天线上穿10天线，10天线上穿20天线的个股
//
//	count(cross(MAV1(c,5),MAV1(c,10)),3)>=1 and count(cross(MAV1(c,10),MAV1(c,20)),3)>=1
type FormulaNo1 struct {
}

func (this *FormulaNo1) Code() int {
	return 1
}

func (this *FormulaNo1) Name() string {
	return "1号策略"
}

// Evaluate 评估K线数据
func (this *FormulaNo1) Evaluate(fullCode string, info *security.StaticBasic, result *treemap.Map) {
	N := MaximumResultDays

	//fmt.Printf("%s\n", fullCode)
	filename := cache.GetKLineFilename(fullCode)
	df := pandas.ReadCSV(filename)
	if df.Err != nil {
		return
	}
	df.SetNames("date", "open", "high", "low", "close", "volume")
	// 收盘价序列
	CLOSE := df.Col("close")
	days := CLOSE.Len()

	// 取5、10、20日均线
	ma5 := MAV1(CLOSE, 5)
	ma10 := MAV1(CLOSE, 10)
	ma20 := MAV1(CLOSE, 20)
	if len(ma5) != days || len(ma10) != days || len(ma20) != days {
		logger.Errorf("均线, 数据没对齐")
	}
	// 两个金叉
	c1 := V2CROSS(ma5, ma10)
	c2 := V2CROSS(ma10, ma20)
	if len(c1) != days || len(c2) != days {
		logger.Errorf("金叉, 数据没对齐")
	}
	// 两个统计
	r1 := V1COUNT(c1, N)
	r2 := V1COUNT(c2, N)
	if len(r1) != days || len(r2) != days {
		logger.Errorf("统计, 数据没对齐")
	}
	// 横向对比
	d := AND(r1, r2)
	if len(d) != days {
		logger.Errorf("横向对比, 数据没对齐")
	}

	//cc1 := CompareGte(r1, 1)
	rLen := len(d)
	if rLen > 1 && d[rLen-1] {
		buy := ma10[days-1]
		sell := buy * 1.05
		date := df.Col("date").Values().([]string)[days-1]
		result.Put(fullCode, ResultInfo{Code: fullCode,
			Name:         info.Name,
			Date:         date,
			Buy:          buy,
			Sell:         sell,
			StrategyCode: this.Code(),
			StrategyName: this.Name()})
	}
}
