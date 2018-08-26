package dbInterface

import (
	"coinDataPull/thirdLib/huobiapi/models"
	"fmt"
	"github.com/golang/glog"
	"coinDataPull/commonUtil/coinDataPullUtil"
	"errors"
	"coinDataPull/handleModules/baseModule/coinDataPullModel"
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

//通过id查询数据
func Select_CoinDataByID(tableName string, coinType int32, id int64) (*models.KLineData, error){
	sql := fmt.Sprintf("SELECT ID, Amount, COUNT, OPEN, CLOSE, Low, High, Vol FROM coin_data.`%v` WHERE coinType = %v and ID = %v", tableName, coinType, id)
	rows, err := coinDataPullUtil.GetMysqlDB().Query(sql)
	if nil != err{
		glog.Error("Query failed!", sql, err.Error())
		return nil, err
	}

	data := new(models.KLineData)
	for rows.Next(){
		rows.Scan(&data.ID, &data.Amount, &data.Count, &data.Open, &data.Close, &data.Low, &data.High, &data.Vol)

		return data, nil
	}

	return nil, errors.New("not found!")
}

//取某个时间点之前的一定量数据
func Select_CoinKLineDataByIDLimit(tableName string, coinType int32, id int64) (*[coinDataPullModel.MACD_CAL_MAX_COUNT]models.KLineData, error){
	sql := fmt.Sprintf("SELECT ID, Amount, COUNT, OPEN, CLOSE, Low, High, Vol FROM (SELECT ID, Amount, COUNT, OPEN, CLOSE, Low, High, Vol FROM coin_data.`%v` WHERE ID <= %v AND coinType = %v ORDER BY id DESC LIMIT %v) info ORDER BY ID",
		tableName, id, coinType, coinDataPullModel.MACD_CAL_MAX_COUNT)
	rows, err := coinDataPullUtil.GetMysqlDB().Query(sql)
	if nil != err{
		glog.Error("Query failed!", sql, err.Error())
		return nil, err
	}
	fmt.Println(sql)

	result := [coinDataPullModel.MACD_CAL_MAX_COUNT]models.KLineData{}
	iPos := 0
	for rows.Next(){
		data := models.KLineData{}
		rows.Scan(&data.ID, &data.Amount, &data.Count, &data.Open, &data.Close, &data.Low, &data.High, &data.Vol)

		result[iPos] = data
		iPos++
	}

	if coinDataPullModel.MACD_CAL_MAX_COUNT != iPos{
		glog.Error("coinDataPullModel.MACD_CAL_MAX_COUNT != iPos error!")
		return nil, errors.New("coinDataPullModel.MACD_CAL_MAX_COUNT != iPos error!")
	}

	return &result, nil
}


func Insert_testData(start int32, stop int32){
	sql := fmt.Sprintf("INSERT INTO coin_data.`test_data` VALUE(%v, %v)", start, stop)
	_, err := coinDataPullUtil.GetMysqlDB().Exec(sql)
	if nil != err{
		glog.Error("Exec failed!", sql, err.Error())
		return
	}
}



















