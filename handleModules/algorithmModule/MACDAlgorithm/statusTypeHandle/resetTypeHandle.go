package statusTypeHandle

import (
	"coinDataPull/handleModules/baseModule/coinDataPullModel"
	"coinDataPull/handleModules/algorithmModule/MACDAlgorithm"
	"coinDataPull/handleModules/baseModule/waveAnalyse"
)

func IsNeedReset(data *MACDAlgorithm.STAllNormResultInfo) bool{
	if handle_resetType_MACDAcross(&data.MacdInfo) || handle_resetType_RSIAcross(&data.RsiInfo){
		return true
	}

	return false
}


//MACDAlgorithm_resetType_MACDAcross macd出现向下交叉
func handle_resetType_MACDAcross(data *coinDataPullModel.STMACDResultInfo) bool{
	//判断macd线交叉diff线
	if waveAnalyse.WaveAnalyse_IsCross(data.OutMACD[:], data.OutMACDSignal[:], -1){
		//fmt.Println("handle_resetType_MACDAcross IsCross: =========================")
		//fmt.Println("OutMACD: ", data.OutMACD)
		//fmt.Println("OutMACDSignal: ", data.OutMACDSignal)
		//fmt.Println("handle_resetType_MACDAcross IsCross: +++++++++++++++++++++++++")
		//os.Exit(1)

		return true
	}

	return false
}

//MACDAlgorithm_resetType_RSIAcross rsi出现向下交叉
func handle_resetType_RSIAcross(data *coinDataPullModel.STRSIResultInfo) bool{
	//判断rsi5线交叉rsi10线
	if waveAnalyse.WaveAnalyse_IsCross(data.Rsi1.Rsi[:], data.Rsi2.Rsi[:], -1){
		//fmt.Println("handle_resetType_RSIAcross IsCross: =========================")
		//fmt.Println("OutMACD: ", data.Rsi1.Rsi)
		//fmt.Println("OutMACDSignal: ", data.Rsi2.Rsi)
		//fmt.Println("handle_resetType_RSIAcross IsCross: +++++++++++++++++++++++++")
		//os.Exit(1)

		return true
	}

	return false
}






