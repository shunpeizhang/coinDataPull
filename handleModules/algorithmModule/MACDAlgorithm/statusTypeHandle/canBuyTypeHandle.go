package statusTypeHandle

import (
	"coinDataPull/handleModules/baseModule/coinDataPullModel"
	"coinDataPull/handleModules/algorithmModule/MACDAlgorithm"
)



//判断买入条件是否达到要求
func CanBuy(data *MACDAlgorithm.STAllNormResultInfo) bool{


	return false
}






//MACDAlgorithm_canBuyType_MACDMouthOpen
func handle_canBuyType_MACDMouthOpen(data *coinDataPullModel.STMACDResultInfo) bool{
	//是否macd值 - diff值达到要求


	return false
}

//MACDAlgorithm_canBuyType_RSIOk   RSI满足条件  之前有交叉，到目前一直在维持
func handle_canBuyType_RSIOk(data *coinDataPullModel.STRSIResultInfo) bool{
	//之前一定范围内是否有交叉


	//是否到目前一直在维持


	return false
}

//MACDAlgorithm_canBuyType_KDJUpCross   KDJ满足条件  之前有交叉，到目前一直在维持
func handle_canBuyType_KDJUpCross(data *coinDataPullModel.STKDJResultInfo) bool{
	//之前一定范围内是否有交叉


	//是否到目前一直在维持


	return false
}

//MACDAlgorithm_canBuyType_RSISpeedRateOk
func handle_canBuyType_RSISpeedRateOk(data *coinDataPullModel.STRSIResultInfo) bool{
	//当前速率是否达到要求


	return false
}




