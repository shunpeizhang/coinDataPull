package coinDataPullModel

import "coinDataPull/thirdLib/huobiapi/models"

const (
	COINTYPE_BTC_USDT = iota
	COINTYPE_ETH_USDT
	COINTYPE_BCH_USDT
	COINTYPE_ETC_USDT
	COINTYPE_LTC_USDT
	COINTYPE_EOS_USDT
	COINTYPE_XRP_USDT
	COINTYPE_OMG_USDT
	COINTYPE_DASH_USDT
	COINTYPE_ZEC_USDT
	COINTYPE_ADA_USDT
	COINTYPE_STEEM_USDT
	COINTYPE_IOTA_USDT

	COINTYPE_MAX
)

const (
	COINSRCTIMETYPE_1MIN = iota
	COINSRCTIMETYPE_5MIN
	COINSRCTIMETYPE_15MIN
	COINSRCTIMETYPE_30MIN
	COINSRCTIMETYPE_60MIN
	COINSRCTIMETYPE_1DAY

	COINSRCTIMETYPE_MAX
)


//取源数据相关宏定义
const (
	COINDATAGET_LIMTCOUNT = 2000  //单次取数据最大上限
	COINDATAGET_COUNTPRE = 100    //程序每次最多取多少数据比较好
	COINDATAGET_DATAAddRATE = 1 //每次请求冗余多少数据比较好(本来需要取100，但取110比较好，防止中间有丢失)

)

//
const(
	MACD_CAL_MAX_COUNT = 100
)



//MACD结果数据结构
type STMACDResultInfo struct {
	SourceData [MACD_CAL_MAX_COUNT]models.KLineData          //源数据

	OutMACD [MACD_CAL_MAX_COUNT]float64             //macd
	OutMACDSignal [MACD_CAL_MAX_COUNT]float64       //diff
	OutMACDHist [MACD_CAL_MAX_COUNT]float64         //dea
	OutBeg int32                                      //有效数起点
	OutNbElement int32                                //合法数数量
}

type STRSIResultItemInfo struct{
	Rsi [MACD_CAL_MAX_COUNT]float64                 //结果
	OutBeg int32                                      //有效数起点
	OutNbElement int32                                //合法数数量
}

//RSI数据结构
type STRSIResultInfo struct{
	SourceData [MACD_CAL_MAX_COUNT]models.KLineData          //源数据

	Rsi1 STRSIResultItemInfo                 //5的结果
	Rsi2 STRSIResultItemInfo                 //10的结果
	Rsi3 STRSIResultItemInfo                 //14的结果
}

//KDJ数据结构
type STKDJResultInfo struct {
	SourceData [MACD_CAL_MAX_COUNT]models.KLineData          //源数据

	K [MACD_CAL_MAX_COUNT]float64                   //k值
	D [MACD_CAL_MAX_COUNT]float64                   //d值
	OutBeg int32                                      //有效数起点
	OutNbElement int32                                //合法数数量
}












