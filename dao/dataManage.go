package dao

import (
	"coinDataPull/thirdLib/huobiapi/models"
	"coinDataPull/coinDataPullModel"
	"coinDataPull/dbInterface"
	"coinDataPull/coinDataPullConfig"
	"github.com/golang/glog"
	"fmt"
)

type STCoinAllData struct {
	CoinData []models.KLineData
}

var G_coinAllData STCoinAllData


func Init() error{
	result, err := dbInterface.Select_CoinAllKLineData(coinDataPullConfig.COINSRCTIMETYPE_TABLENAME_MAP[coinDataPullModel.COINSRCTIMETYPE_15MIN],0)
	if nil != err{
		glog.Error("Select_CoinAllKLineData failed!")
		return err
	}

	iCount := len(result)
	for iPos := 0; iCount > iPos; iPos++{
		G_coinAllData.CoinData = append(G_coinAllData.CoinData, result[iPos])
	}
	fmt.Println(G_coinAllData)

	return nil
}




