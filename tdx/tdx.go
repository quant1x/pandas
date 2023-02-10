package tdx

import (
	"gitee.com/quant1x/gotdx/proto"
	"gitee.com/quant1x/gotdx/quotes"
	"gitee.com/quant1x/pandas"
	"gitee.com/quant1x/pandas/stat"
	"strings"
)

var (
	stdApi *quotes.StdApi = nil
)

func prepare() *quotes.StdApi {
	if stdApi == nil {
		std_api, err := quotes.NewStdApi()
		if err != nil {
			return nil
		}
		stdApi = std_api
	}
	return stdApi
}

func startsWith(str string, prefixs []string) bool {
	if len(str) == 0 || len(prefixs) == 0 {
		return false
	}
	for _, prefix := range prefixs {
		if strings.HasPrefix(str, prefix) {
			return true
		}
	}
	return false
}

// 判断股票ID对应的证券市场匹配规则
//
// ['50', '51', '60', '90', '110'] 为 sh
// ['00', '12'，'13', '18', '15', '16', '18', '20', '30', '39', '115'] 为 sz
// ['5', '6', '9'] 开头的为 sh， 其余为 sz
func getStockMarket(symbol string) string {
	//:param string: False 返回市场ID，否则市场缩写名称
	//:param symbol: 股票ID, 若以 'sz', 'sh' 开头直接返回对应类型，否则使用内置规则判断
	//:return 'sh' or 'sz'

	market := "sh"
	if startsWith(symbol, []string{"sh", "sz", "SH", "SZ"}) {
		market = strings.ToLower(symbol[0:2])
	} else if startsWith(symbol, []string{"50", "51", "60", "68", "90", "110", "113", "132", "204"}) {
		market = "sh"
	} else if startsWith(symbol, []string{"00", "12", "13", "18", "15", "16", "18", "20", "30", "39", "115", "1318"}) {
		market = "sz"
	} else if startsWith(symbol, []string{"5", "6", "9", "7"}) {
		market = "sh"
	} else if startsWith(symbol, []string{"4", "8"}) {
		market = "bj"
	}
	return market
}

func getStockMarketId(symbol string) uint8 {
	market := getStockMarket(symbol)
	marketId := proto.MarketShangHai
	if market == "sh" {
		marketId = proto.MarketShangHai
	} else if market == "sz" {
		marketId = proto.MarketShenZhen
	} else if market == "bj" {
		marketId = proto.MarketBeiJing
	}
	//# logger.debug(f"market => {market}")

	return marketId
}

// GetKLine 获取日K线
func GetKLine(code string, start uint16, count uint16) pandas.DataFrame {
	api := prepare()

	marketId := getStockMarketId(code)
	data, _ := api.GetKLine(marketId, code, proto.KLINE_TYPE_RI_K, start, count)
	df := pandas.LoadStructs(data.List)
	df = df.Select([]string{"Open", "Close", "High", "Low", "Vol", "Amount", "DateTime"})
	err := df.SetNames("open", "close", "high", "low", "vol", "amount", "date")
	if err != nil {
		return pandas.DataFrame{}
	}

	return df
}

// GetKLine 获取日K线
func GetKLineAll(code string) pandas.DataFrame {
	api := prepare()

	marketId := getStockMarketId(code)
	history := make([]quotes.SecurityBar, 0)
	count := uint16(800)
	step := uint16(800)
	start := uint16(0)
	hs := make([]quotes.SecurityBarsReply, 0)
	for {
		data, err := api.GetKLine(marketId, code, proto.KLINE_TYPE_RI_K, uint16(start), uint16(count))
		if err != nil {
			panic("接口异常")
		}
		hs = append(hs, (*data))
		if data.Count < count {
			// 已经是最早的记录
			// 需要排序
			break
		}
		start += step
	}
	hs = stat.Reverse(hs)
	for _, v := range hs {
		history = append(history, v.List...)
	}

	//data, _ := api.GetKLine(marketId, code, proto.KLINE_TYPE_RI_K, start, count)
	df := pandas.LoadStructs(history)
	df = df.Select([]string{"Open", "Close", "High", "Low", "Vol", "Amount", "DateTime"})
	err := df.SetNames("open", "close", "high", "low", "vol", "amount", "date")
	if err != nil {
		return pandas.DataFrame{}
	}

	return df
}
