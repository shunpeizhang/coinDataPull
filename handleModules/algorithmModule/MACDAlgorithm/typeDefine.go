package MACDAlgorithm

import (
	"coinDataPull/thirdLib/huobiapi/models"
	"coinDataPull/handleModules/baseModule/coinDataPullModel"
)

//处理总体状态
const (
	MACDAlgorithmStatus_invalide = iota
	MACDAlgorithmStatus_start   //找到对应点了，但还不能买入
	MACDAlgorithmStatus_buy     //已经买入了
	MACDAlgorithmStatus_sell    //已经卖出了

	MACDAlgorithmStatus_max
)

//可以开始的条件类型
const (
	MACDAlgorithm_canStart_invalide = iota + 50
	MACDAlgorithm_canStart_MACDCross                  //MACD出现交叉

	MACDAlgorithm_canStart_max
)

//买入条件类型
const(
	MACDAlgorithm_canBuyType_invalide = iota + 100
	MACDAlgorithm_canBuyType_MACDMouthOpen       //MACD的开口大小达到条件
	MACDAlgorithm_canBuyType_RSIUpCross           //RSI满足条件  之前有交叉，到目前一直在维持
	MACDAlgorithm_canBuyType_KDJUpCross           //kdj之前有个向上交叉  之前有交叉，到目前一直在维持
	MACDAlgorithm_canBuyType_RSISpeedRateOk      //在购买点RSI的增长速率达到要求

	MACDAlgorithm_canBuyType_max
)

//开始但还没买入时，下面条件满足时需要重置状态
const(
	MACDAlgorithm_resetType_invalide = iota + 200
	MACDAlgorithm_resetType_MACDAcross              //MACD出现交叉
	MACDAlgorithm_resetType_RSIAcross               //RSI出现交叉

	MACDAlgorithm_resetType_max
)

//买入后的红线，超过这些红线时需要立即卖出
const(
	MACDAlgorithm_redLine_invalide = iota + 300
	MACDAlgorithm_redLine_topGrow                  //增长到这个点必须卖出了
	MACDAlgorithm_redLine_lowPoint                 //跌了这么多，必须卖出了

	MACDAlgorithm_redLine_max
)

//卖出点类型
const (
	MACDAlgorithm_sellType_invalide = iota + 400
	MACDAlgorithm_sellType_KDJCross                   //KDJ出现交叉
	MACDAlgorithm_sellType_RSICross                   //RSI出现交叉
	MACDAlgorithm_sellType_MACDCross                  //MACD出现交叉

	MACDAlgorithm_sellType_max
)

//常量
const (
	WAVEDistanceDiff_MACD = 4      //表明MACD之前有交叉，到目前一直在维持需当前的间距值
	WAVEDistanceDiff_RSI = 9       //表明RSI之前有交叉，到目前一直在维持需当前的间距值
	WAVEDistanceDiff_KDJ = 7       //表明KDJ之前有交叉，到目前一直在维持需当前的间距值
)


//MACDAlgorithm处理用到的数据结构
type STMACDAlgorithmInfo struct {
	MACDAlgorithmStatus int32        //当前处理状态

	MACDAlgorithm_canBuyTypes map[int32]int32    //如果处理在MACDAlgorithmStatus_start，此容器存储已满足的条件
	BuyKlinePoint models.KLineData               //买入点的kline值
	SellKlinePoint models.KLineData               //卖出点的kline值

	StartTimePoint int64                         //开始的时间点
	BuyTimePoint int64                           //买入的时间点
	SellTimePoint int64                          //卖出的时间点

	TableName string
	CoinType int32
}

func (this *STMACDAlgorithmInfo) Reset(){
	this.MACDAlgorithmStatus = MACDAlgorithmStatus_invalide
	this.MACDAlgorithm_canBuyTypes = make(map[int32]int32)

	this.StartTimePoint = 0
	this.BuyTimePoint = 0
	this.SellTimePoint = 0

	this.BuyKlinePoint.Reset()
	this.SellKlinePoint.Reset()
}



//所有指标数据
type STAllNormResultInfo struct{
	TimePoint int64                 //当前时间点

	MacdInfo coinDataPullModel.STMACDResultInfo
	RsiInfo coinDataPullModel.STRSIResultInfo
	KdjInfo coinDataPullModel.STKDJResultInfo
}










