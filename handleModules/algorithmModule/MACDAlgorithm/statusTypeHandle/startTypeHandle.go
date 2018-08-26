package statusTypeHandle

import (
	"coinDataPull/handleModules/baseModule/coinDataPullModel"
	"coinDataPull/handleModules/algorithmModule/MACDAlgorithm"
)


func CanStart(data *MACDAlgorithm.STAllNormResultInfo) bool{

	return false
}




//MACDAlgorithm_canStart_MACDCross
func handle_canStart_MACDCross(data *coinDataPullModel.STMACDResultInfo) bool{
	//判断macd线是否向上交叉diff线


	return false
}








