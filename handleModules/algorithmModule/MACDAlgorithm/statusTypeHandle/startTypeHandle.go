package statusTypeHandle

import (
	"coinDataPull/handleModules/baseModule/coinDataPullModel"
	"coinDataPull/handleModules/algorithmModule/MACDAlgorithm"
	"coinDataPull/handleModules/baseModule/waveAnalyse"
	"fmt"
)


func CanStart(data *MACDAlgorithm.STAllNormResultInfo) bool{
	fmt.Println("CanStart")

	return handle_canStart_MACDCross(&data.MacdInfo)
}


//MACDAlgorithm_canStart_MACDCross
func handle_canStart_MACDCross(data *coinDataPullModel.STMACDResultInfo) bool{
	//判断macd线是否向上交叉diff线
	if waveAnalyse.WaveAnalyse_IsCross(data.OutMACD[:], data.OutMACDSignal[:], 1){
		return true
	}

	return false
}








