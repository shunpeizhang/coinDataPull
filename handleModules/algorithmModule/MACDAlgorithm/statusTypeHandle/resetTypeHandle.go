package statusTypeHandle

import "coinDataPull/handleModules/baseModule/coinDataPullModel"

func IsNeedReset(data *coinDataPullModel.STAllNormResultInfo) bool{

	return false
}



//MACDAlgorithm_resetType_MACDAcross macd出现向下交叉
func handle_resetType_MACDAcross(data *coinDataPullModel.STMACDResultInfo) bool{

	return false
}

//MACDAlgorithm_resetType_RSIAcross rsi出现向下交叉
func handle_resetType_RSIAcross(data *coinDataPullModel.STRSIResultInfo) bool{

	return false
}






