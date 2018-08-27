package MACDAlgorithmControl

import (
	"coinDataPull/commonUtil/coinDataPullUtil"
	"github.com/golang/glog"
	"errors"
	"coinDataPull/handleModules/baseModule/dbInterface"
	"coinDataPull/thirdLib/ta_lib"
	"os"
	"coinDataPull/handleModules/algorithmModule/MACDAlgorithm/statusTypeHandle"
	"coinDataPull/thirdLib/huobiapi/models"
	"fmt"
	"coinDataPull/handleModules/algorithmModule/MACDAlgorithm"
)

var stMACDAlgorithmInfo MACDAlgorithm.STMACDAlgorithmInfo


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
		if MACDAlgorithm.MACDAlgorithmStatus_invalide == stMACDAlgorithmInfo.MACDAlgorithmStatus{
			//检测是否可以开始
			if statusTypeHandle.CanStart(stAllNormResultInfo){
				stMACDAlgorithmInfo.MACDAlgorithmStatus = MACDAlgorithm.MACDAlgorithmStatus_start
				stMACDAlgorithmInfo.StartTimePoint = realTime
				continue
			}
		}else if MACDAlgorithm.MACDAlgorithmStatus_start == stMACDAlgorithmInfo.MACDAlgorithmStatus{
			//检测是否重置
			if statusTypeHandle.IsNeedReset(stAllNormResultInfo){
				stMACDAlgorithmInfo.Reset()
				break
			}

			//检测是否可以买入
			if statusTypeHandle.CanBuy(stAllNormResultInfo){
				stMACDAlgorithmInfo.MACDAlgorithmStatus = MACDAlgorithm.MACDAlgorithmStatus_buy
				stMACDAlgorithmInfo.BuyTimePoint = realTime
				stMACDAlgorithmInfo.BuyKlinePoint = *lineData
				break
			}
		}else if MACDAlgorithm.MACDAlgorithmStatus_buy == stMACDAlgorithmInfo.MACDAlgorithmStatus{
			//是否过红线或可以卖出
			if statusTypeHandle.IsRedLineOut(&stMACDAlgorithmInfo) || statusTypeHandle.CanSell(stAllNormResultInfo){
				stMACDAlgorithmInfo.MACDAlgorithmStatus = MACDAlgorithm.MACDAlgorithmStatus_sell
				stMACDAlgorithmInfo.SellTimePoint = realTime
				stMACDAlgorithmInfo.SellKlinePoint = *lineData

				continue
			}
		}else if MACDAlgorithm.MACDAlgorithmStatus_sell == stMACDAlgorithmInfo.MACDAlgorithmStatus{
			//卖出
			{
				fmt.Println("sell ===================================================")
				fmt.Println("startTimePoint: ", stMACDAlgorithmInfo.StartTimePoint)
				fmt.Println("buyKlinePoint: ", stMACDAlgorithmInfo.BuyKlinePoint, " close:", stMACDAlgorithmInfo.BuyKlinePoint.Close)
				fmt.Println("sellTimePoint: ", stMACDAlgorithmInfo.SellTimePoint, " close:", stMACDAlgorithmInfo.SellKlinePoint.Close)
				fmt.Println("sell +++++++++++++++++++++++++++++++++++++++++++++++++++")
			}
			os.Exit(1)

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
func getTimePointNormData(curPoint int64) (*models.KLineData, *MACDAlgorithm.STAllNormResultInfo, error){
	//得到此时间点的数据
	lineData, err := dbInterface.Select_CoinDataByID(stMACDAlgorithmInfo.TableName, stMACDAlgorithmInfo.CoinType, curPoint)
	if nil != err{
		glog.Error("dbInterface.Select_CoinDataByID failed! curPoint:", curPoint)
		return nil, nil, err
	}

	//得到对应标准数据
	stAllNormResultInfo := new(MACDAlgorithm.STAllNormResultInfo)
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
		//fmt.Println("OutMACD:", stAllNormResultInfo.MacdInfo.OutMACD)
		//fmt.Println("OutMACDSignal:", stAllNormResultInfo.MacdInfo.OutMACDSignal)

		//计算RSI
		err = ta_lib.RSIAll(datas, &stAllNormResultInfo.RsiInfo)
		if nil != err{
			glog.Error(err.Error())
			return nil, nil, err
		}
		//fmt.Println("Rsi:", stAllNormResultInfo.RsiInfo.Rsi1.Rsi)

		//计算KDJ
		err = ta_lib.KDJ(datas, &stAllNormResultInfo.KdjInfo)
		if nil != err{
			glog.Error(err.Error())
			return nil, nil, err
		}
		//fmt.Println("k: ", stAllNormResultInfo.KdjInfo.K)
		//fmt.Println("D: ", stAllNormResultInfo.KdjInfo.D)
		//os.Exit(1)
	}

	return lineData, stAllNormResultInfo, nil
}












