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
		handle_canBuyType_RSISpeedRateOk(&data.RsiInfo)&&
		handle_canBuyType_RSILowMaxLimit(&data.RsiInfo){
			return true
	}

	return false
}






//MACDAlgorithm_canBuyType_MACDMouthOpen
func handle_canBuyType_MACDMouthOpen(data *coinDataPullModel.STMACDResultInfo) bool{
	return true

	//是否macd值 - diff值达到要求
	diff := data.OutMACD[coinDataPullModel.MACD_CAL_MAX_COUNT - 1] - data.OutMACDSignal[coinDataPullModel.MACD_CAL_MAX_COUNT - 1]
	if 4.0 < diff{
		//fmt.Println("handle_canBuyType_MACDMouthOpen IsCross: =========================")
		//fmt.Println("OutMACD: ", data.OutMACD)
		//fmt.Println("OutMACDSignal: ", data.OutMACDSignal)
		//fmt.Println("handle_canBuyType_MACDMouthOpen IsCross: +++++++++++++++++++++++++")
		//os.Exit(1)

		return true
	}

	return false
}

//MACDAlgorithm_canBuyType_RSIOk   RSI满足条件  之前有交叉，到目前一直在维持
func handle_canBuyType_RSIOk(data *coinDataPullModel.STRSIResultInfo) bool{
	if waveAnalyse.WaveAnalyse_IsCrossAndContinue(data.Rsi1.Rsi[:], data.Rsi2.Rsi[:], MACDAlgorithm.WAVEDistanceDiff_RSI){
		//fmt.Println("handle_canBuyType_RSIOk IsCross: =========================")
		//fmt.Println("Rsi1: ", data.Rsi1.Rsi)
		//fmt.Println("Rsi2: ", data.Rsi2.Rsi)
		//fmt.Println("handle_canBuyType_RSIOk IsCross: +++++++++++++++++++++++++")
		//os.Exit(1)

		return true
	}

	return false
}

//MACDAlgorithm_canBuyType_KDJUpCross   KDJ满足条件  之前有交叉，到目前一直在维持
func handle_canBuyType_KDJUpCross(data *coinDataPullModel.STKDJResultInfo) bool{
	if waveAnalyse.WaveAnalyse_IsCrossAndContinue(data.K[:], data.D[:], MACDAlgorithm.WAVEDistanceDiff_KDJ){
		//fmt.Println("handle_canBuyType_KDJUpCross IsCross: =========================")
		//fmt.Println("k: ", data.K)
		//fmt.Println("d: ", data.D)
		//fmt.Println("handle_canBuyType_KDJUpCross IsCross: +++++++++++++++++++++++++")
		//os.Exit(1)

		return true
	}

	return false
}

//MACDAlgorithm_canBuyType_RSISpeedRateOk
func handle_canBuyType_RSISpeedRateOk(data *coinDataPullModel.STRSIResultInfo) bool{
	//当前速率是否达到要求
	if waveAnalyse.WaveAnalyse_speedRate(int(data.Rsi1.OutBeg), data.Rsi1.Rsi[:], 70){
		//fmt.Println("handle_canBuyType_RSISpeedRateOk IsCross: =========================")
		//fmt.Println("Rsi1: ", data.Rsi1.Rsi)
		//fmt.Println("handle_canBuyType_RSISpeedRateOk IsCross: +++++++++++++++++++++++++")
		//os.Exit(1)

		return true
	}

	return false
}

//MACDAlgorithm_canBuyType_RSILowMaxLimit
func handle_canBuyType_RSILowMaxLimit(data *coinDataPullModel.STRSIResultInfo) bool{
	dataLen := len(data.Rsi1.Rsi)

	if 85 <= data.Rsi1.Rsi[dataLen - 1]{
		return false
	}

	return true
}


