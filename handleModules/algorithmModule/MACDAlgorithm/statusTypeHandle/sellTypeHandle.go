package statusTypeHandle

import (
	"coinDataPull/handleModules/baseModule/coinDataPullModel"
	"coinDataPull/handleModules/algorithmModule/MACDAlgorithm"
	"coinDataPull/handleModules/baseModule/waveAnalyse"
	"fmt"
)

func CanSell(data *MACDAlgorithm.STAllNormResultInfo)bool{
	fmt.Println("CanSell")

	if handle_sellType_KDJCross(&data.KdjInfo) ||
		handle_sellType_RSICross(&data.RsiInfo) ||
		handle_sellType_MACDCross(&data.MacdInfo){
			return true
	}

	return false
}


//MACDAlgorithm_sellType_KDJCross
func handle_sellType_KDJCross(data *coinDataPullModel.STKDJResultInfo) bool{
	//从买入点开始到目前，只要出现交叉就算
	if waveAnalyse.WaveAnalyse_IsCross(data.K[:], data.D[:], 0){
		return true
	}

	return false
}

//MACDAlgorithm_sellType_RSICross
func handle_sellType_RSICross(data *coinDataPullModel.STRSIResultInfo) bool{
	//从买入点开始到目前，只要出现交叉就算
	if waveAnalyse.WaveAnalyse_IsCross(data.Rsi1.Rsi[:], data.Rsi2.Rsi[:], 0){
		return true
	}

	return false
}

//MACDAlgorithm_sellType_MACDCross
func handle_sellType_MACDCross(data *coinDataPullModel.STMACDResultInfo) bool{
	//从买入点开始到目前，只要出现交叉就算
	if waveAnalyse.WaveAnalyse_IsCross(data.OutMACDSignal[:], data.OutMACDSignal[:], 0){
		return true
	}

	return false
}





