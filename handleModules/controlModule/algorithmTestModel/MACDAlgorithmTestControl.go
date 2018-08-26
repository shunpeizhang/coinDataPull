package algorithmTestModel

import (
	"coinDataPull/handleModules/algorithmModule/MACDAlgorithm"
	"os"
)

func Init(){
	MACDAlgorithm.Init("tb_coinData_60min", 0)
}


func RunTest(){
	for iPos := int64(1527508800); 1535266800 > iPos; iPos = iPos + 3600{
		MACDAlgorithm.Heartbeat(iPos)
	}

	os.Exit(0)
}








