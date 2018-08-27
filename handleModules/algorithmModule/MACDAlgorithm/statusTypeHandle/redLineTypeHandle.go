package statusTypeHandle

import (
	"coinDataPull/handleModules/algorithmModule/MACDAlgorithm"
	"fmt"
)

//是否已经越过红线，需要卖出
func IsRedLineOut(data *MACDAlgorithm.STMACDAlgorithmInfo) bool{
	fmt.Println("IsRedLineOut")

	return false
}


//MACDAlgorithm_redLine_topGrow
func handle_redLine_topGrow(data *MACDAlgorithm.STMACDAlgorithmInfo) bool{



	return false
}

//MACDAlgorithm_redLine_lowPoint
func handle_redLine_lowPoint(data *MACDAlgorithm.STMACDAlgorithmInfo) bool{

	return false
}




