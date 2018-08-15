package dbInterface

import (
	"coinDataPull/thirdLib/huobiapi/models"
	"fmt"
	"coinDataPull/coinDataPullUtil"
	"github.com/golang/glog"
)


//保存kline数据到数据库
func Insert_CoinKLineData(data *models.KLineData, tableName string, coinType int32) error{
	sql := fmt.Sprintf("insert into coin_data.%v(coinType, ID, Amount, Count, Open, Close, Low, High, Vol) values(%v, %v, %v, %v, %v, %v, %v, %v, %v)",
		tableName, coinType, data.ID, data.Amount, data.Count, data.Open, data.Close, data.Low, data.High, data.Vol)
	_, err := coinDataPullUtil.GetMysqlDB().Exec(sql)
	if nil != err{
		//glog.Error("Exec failed!", sql, err.Error())
		return nil
	}

	glog.Info(sql)

	return nil
}
























