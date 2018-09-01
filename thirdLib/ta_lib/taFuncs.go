package ta_lib

import "C"
import (
	"coinDataPull/thirdLib/huobiapi/models"
	"coinDataPull/handleModules/baseModule/coinDataPullModel"
	"log"
	"net/rpc"
)

const(
	rpc_server_addr = "47.97.202.23:8081"
)





func MACD(data *[coinDataPullModel.MACD_CAL_MAX_COUNT]models.KLineData, stMACDResultInfo *coinDataPullModel.STMACDResultInfo) error {
	client, err := rpc.DialHTTP("tcp", rpc_server_addr)
	if err != nil {
		log.Fatal("dialing:", err)
		return err
	}
	return client.Call("STTaLibInfo.MACD", data, stMACDResultInfo)
}

func RSIAll(data *[coinDataPullModel.MACD_CAL_MAX_COUNT]models.KLineData, stRSIResultInfo *coinDataPullModel.STRSIResultInfo) error{
	client, err := rpc.DialHTTP("tcp", rpc_server_addr)
	if err != nil {
		log.Fatal("dialing:", err)
		return err
	}
	return client.Call("STTaLibInfo.MACD", data, stRSIResultInfo)
}

//KDJ
func KDJ(data *[coinDataPullModel.MACD_CAL_MAX_COUNT]models.KLineData, stKDJResultInfo *coinDataPullModel.STKDJResultInfo) error{
	client, err := rpc.DialHTTP("tcp", rpc_server_addr)
	if err != nil {
		log.Fatal("dialing:", err)
		return err
	}
	return client.Call("STTaLibInfo.MACD", data, stKDJResultInfo)
}









