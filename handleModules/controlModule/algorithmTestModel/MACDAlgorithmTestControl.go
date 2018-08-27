package algorithmTestModel

import (
	"os"
	"coinDataPull/handleModules/algorithmModule/MACDAlgorithm/MACDAlgorithmControl"
)

func Init(){
	MACDAlgorithmControl.Init("tb_coinData_30min", 0)
}


func RunTest(){
	//for iPos := int64(1527508800); 1535266800 > iPos; iPos = iPos + 3600{
	for iPos := int64(1531004410); 1535266800 > iPos; iPos = iPos + 3600{
		MACDAlgorithmControl.Heartbeat(iPos)
	}

	os.Exit(0)
}








