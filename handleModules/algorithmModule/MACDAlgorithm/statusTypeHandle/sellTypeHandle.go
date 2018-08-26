package statusTypeHandle

import (
	"coinDataPull/handleModules/baseModule/coinDataPullModel"
	"coinDataPull/handleModules/algorithmModule/MACDAlgorithm"
)

func CanSell(data *MACDAlgorithm.STAllNormResultInfo)bool{


	return false
}


//MACDAlgorithm_sellType_KDJCross
func handle_sellType_KDJCross(data *coinDataPullModel.STKDJResultInfo) bool{
	//从买入点开始到目前，只要出现交叉就算

	return false
}

//MACDAlgorithm_sellType_RSICross
func handle_sellType_RSICross(data *coinDataPullModel.STRSIResultInfo) bool{
	//从买入点开始到目前，只要出现交叉就算

	return false
}

//MACDAlgorithm_sellType_MACDCross
func handle_sellType_MACDCross(data *coinDataPullModel.STMACDResultInfo) bool{
	//从买入点开始到目前，只要出现交叉就算

	return false
}





