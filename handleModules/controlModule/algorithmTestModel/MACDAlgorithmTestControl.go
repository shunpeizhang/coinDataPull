package algorithmTestModel

import (
	"os"
	"coinDataPull/handleModules/algorithmModule/MACDAlgorithm/MACDAlgorithmControl"
)

var add_tick int64

func Init(){
	//add_tick = 3600
	add_tick = 1800
	MACDAlgorithmControl.Init("tb_coinData_30min", 0)
}


func RunTest(){
	for iPos := int64(1527508800); 1535266800 > iPos; iPos = iPos + add_tick{
	//for iPos := int64(1531004410); 1535266800 > iPos; iPos = iPos + 3600{
		MACDAlgorithmControl.Heartbeat(iPos)
	}

	os.Exit(0)
}








