package statusTypeHandle

import (
	"coinDataPull/handleModules/baseModule/coinDataPullModel"
	"coinDataPull/handleModules/algorithmModule/MACDAlgorithm"
	"coinDataPull/handleModules/baseModule/waveAnalyse"
)



//判断买入条件是否达到要求
func CanBuy(data *MACDAlgorithm.STAllNormResultInfo) bool{
	if handle_canBuyType_MACDMouthOpen(&data.MacdInfo) &&
		handle_canBuyType_RSIOk(&data.RsiInfo) &&
		handle_canBuyType_KDJUpCross(&data.KdjInfo) &&
		handle_canBuyType_RSISpeedRateOk(&data.RsiInfo){
			return true
	}

	return false
}






//MACDAlgorithm_canBuyType_MACDMouthOpen
func handle_canBuyType_MACDMouthOpen(data *coinDataPullModel.STMACDResultInfo) bool{
	//是否macd值 - diff值达到要求
	diff := data.OutMACD[coinDataPullModel.MACD_CAL_MAX_COUNT - 1] - data.OutMACDSignal[coinDataPullModel.MACD_CAL_MAX_COUNT - 1]
	if 4 < diff{
		return true
	}

	return false
}

//MACDAlgorithm_canBuyType_RSIOk   RSI满足条件  之前有交叉，到目前一直在维持
func handle_canBuyType_RSIOk(data *coinDataPullModel.STRSIResultInfo) bool{
	if waveAnalyse.WaveAnalyse_IsCrossAndContinue(data.Rsi1.Rsi[:], data.Rsi2.Rsi[:], MACDAlgorithm.WAVEDistanceDiff_RSI){
		return true
	}

	return false
}

//MACDAlgorithm_canBuyType_KDJUpCross   KDJ满足条件  之前有交叉，到目前一直在维持
func handle_canBuyType_KDJUpCross(data *coinDataPullModel.STKDJResultInfo) bool{
	if waveAnalyse.WaveAnalyse_IsCrossAndContinue(data.K[:], data.D[:], MACDAlgorithm.WAVEDistanceDiff_KDJ){
		return true
	}

	return false
}

//MACDAlgorithm_canBuyType_RSISpeedRateOk
func handle_canBuyType_RSISpeedRateOk(data *coinDataPullModel.STRSIResultInfo) bool{
	//当前速率是否达到要求
	if waveAnalyse.WaveAnalyse_speedRate(data.Rsi1.Rsi[:], 70){
		return true
	}

	return false
}




