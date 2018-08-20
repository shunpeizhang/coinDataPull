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

	fmt.Println(sql)

	return nil
}

//查询某种币的所有数据
func Select_CoinAllKLineData(tableName string, coinType int32) ([]models.KLineData, error){
	sql := fmt.Sprintf("SELECT ID, Amount, COUNT, OPEN, CLOSE, Low, High, Vol FROM coin_data.`%v` WHERE coinType = %v ORDER BY id", tableName, coinType)
	rows, err := coinDataPullUtil.GetMysqlDB().Query(sql)
	if nil != err{
		glog.Error("Query failed!", sql, err.Error())
		return nil, err
	}

	var result []models.KLineData
	for rows.Next(){
		data := models.KLineData{}
		rows.Scan(&data.ID, &data.Amount, &data.Count, &data.Open, &data.Close, &data.Low, &data.High, &data.Vol)

		result = append(result, data)
	}

	return result, nil
}


func Insert_testData(start int32, stop int32){
	sql := fmt.Sprintf("INSERT INTO coin_data.`test_data` VALUE(%v, %v)", start, stop)
	_, err := coinDataPullUtil.GetMysqlDB().Exec(sql)
	if nil != err{
		glog.Error("Exec failed!", sql, err.Error())
		return
	}
}



















