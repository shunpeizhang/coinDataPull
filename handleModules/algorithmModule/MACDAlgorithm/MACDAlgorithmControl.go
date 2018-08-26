package MACDAlgorithm

import (
	"coinDataPull/handleModules/baseModule/coinDataPullUtil"
	"github.com/golang/glog"
	"errors"
	"coinDataPull/handleModules/baseModule/dbInterface"
	"coinDataPull/thirdLib/ta_lib"
	"os"
	"coinDataPull/handleModules/algorithmModule/MACDAlgorithm/statusTypeHandle"
	"coinDataPull/thirdLib/huobiapi/models"
	"fmt"
)

var stMACDAlgorithmInfo STMACDAlgorithmInfo


func Init(tableName string, coinType int32){
	stMACDAlgorithmInfo.Reset()

	stMACDAlgorithmInfo.TableName = tableName
	stMACDAlgorithmInfo.CoinType = coinType
}


//由外部总控制驱动运行
func Heartbeat(curPoint int64) error{
	//得到整点对应的时间
	realTime := coinDataPullUtil.GetHourTime(curPoint)
	if 60 <= curPoint - realTime{
		glog.Error("60 <= curPoint - realTime error!", curPoint, realTime)
		return errors.New("60 <= curPoint - realTime error!")
	}

	//得到计算需要的数据
	lineData, stAllNormResultInfo, err := getTimePointNormData(realTime)
	if nil != err{
		glog.Error(err.Error())
		return err
	}

	iMaxCount := 10
	for 0 < iMaxCount{
		iMaxCount = iMaxCount - 1
		if MACDAlgorithmStatus_invalide == stMACDAlgorithmInfo.MACDAlgorithmStatus{
			//检测是否可以开始
			if statusTypeHandle.CanStart(stAllNormResultInfo){
				stMACDAlgorithmInfo.MACDAlgorithmStatus = MACDAlgorithmStatus_start
				stMACDAlgorithmInfo.startTimePoint = realTime
				continue
			}
		}else if MACDAlgorithmStatus_start == stMACDAlgorithmInfo.MACDAlgorithmStatus{
			//检测是否重置
			if statusTypeHandle.IsNeedReset(stAllNormResultInfo){
				stMACDAlgorithmInfo.Reset()
				break
			}

			//检测是否可以买入
			if statusTypeHandle.CanBuy(stAllNormResultInfo){
				stMACDAlgorithmInfo.MACDAlgorithmStatus = MACDAlgorithmStatus_buy
				stMACDAlgorithmInfo.buyTimePoint = realTime
				stMACDAlgorithmInfo.buyKlinePoint = *lineData
				break
			}
		}else if MACDAlgorithmStatus_buy == stMACDAlgorithmInfo.MACDAlgorithmStatus{
			//是否过红线或可以卖出
			if statusTypeHandle.IsRedLineOut(&stMACDAlgorithmInfo) || statusTypeHandle.CanSell(stAllNormResultInfo){
				stMACDAlgorithmInfo.MACDAlgorithmStatus = MACDAlgorithmStatus_sell
				stMACDAlgorithmInfo.sellTimePoint = realTime
				stMACDAlgorithmInfo.sellKlinePoint = *lineData

				continue
			}
		}else if MACDAlgorithmStatus_sell == stMACDAlgorithmInfo.MACDAlgorithmStatus{
			//卖出
			{
				fmt.Println("sell ===================================================")
				fmt.Println("startTimePoint: ", stMACDAlgorithmInfo.startTimePoint)
				fmt.Println("buyKlinePoint: ", stMACDAlgorithmInfo.buyKlinePoint, " close:", stMACDAlgorithmInfo.buyKlinePoint.Close)
				fmt.Println("sellTimePoint: ", stMACDAlgorithmInfo.sellTimePoint, " close:", stMACDAlgorithmInfo.sellKlinePoint.Close)
				fmt.Println("sell +++++++++++++++++++++++++++++++++++++++++++++++++++")
			}

			//重置
			stMACDAlgorithmInfo.Reset()
		}else{
			glog.Error("MACDAlgorithmStatus error! status", stMACDAlgorithmInfo.MACDAlgorithmStatus)
			os.Exit(1)
		}

		break
	}

	return nil
}



//得到指定时间点的指标数据
func getTimePointNormData(curPoint int64) (*models.KLineData, *STAllNormResultInfo, error){
	//得到此时间点的数据
	lineData, err := dbInterface.Select_CoinDataByID(stMACDAlgorithmInfo.TableName, stMACDAlgorithmInfo.CoinType, curPoint)
	if nil != err{
		glog.Error("dbInterface.Select_CoinDataByID failed! curPoint:", curPoint)
		return nil, nil, err
	}

	//得到对应标准数据
	stAllNormResultInfo := new(STAllNormResultInfo)
	{
		//得到kline数据
		datas, err := dbInterface.Select_CoinKLineDataByIDLimit(stMACDAlgorithmInfo.TableName, stMACDAlgorithmInfo.CoinType, curPoint)
		if nil != err{
			glog.Error("dbInterface.Select_CoinDataByID failed! curPoint:", curPoint)
			return nil, nil, err
		}

		//计算MACD
		err = ta_lib.MACD(datas, &stAllNormResultInfo.MacdInfo)
		if nil != err{
			glog.Error(err.Error())
			return nil, nil, err
		}

		//计算RSI
		err = ta_lib.RSIAll(datas, &stAllNormResultInfo.RsiInfo)
		if nil != err{
			glog.Error(err.Error())
			return nil, nil, err
		}

		//计算KDJ
		err = ta_lib.KDJ(datas, &stAllNormResultInfo.KdjInfo)
		if nil != err{
			glog.Error(err.Error())
			return nil, nil, err
		}
	}

	return lineData, stAllNormResultInfo, nil
}












