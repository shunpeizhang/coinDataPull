package verify

import (
	"coinDataPull/coinDataPullModel"
	"coinDataPull/thirdLib/huobiapi/models"
	"coinDataPull/dao"
	"coinDataPull/ta_lib"
	"github.com/golang/glog"
	"fmt"
)

func VerifyMACD(){
	iCurPos := 0//当前处理位置

	for true{
		//检测数据是否足够
		if len(dao.G_coinAllData.CoinData) <= iCurPos + coinDataPullModel.MACD_CAL_MAX_COUNT{
			break
		}

		//取一定量的数据
		datas := [coinDataPullModel.MACD_CAL_MAX_COUNT]models.KLineData{}
		for iPos := 0; coinDataPullModel.MACD_CAL_MAX_COUNT > iPos; iPos++{
			datas[iPos] = dao.G_coinAllData.CoinData[iCurPos + iPos]
		}

		//计算CDMA
		outMACD, outMACDSignal, outMACDHist, outBeg, outNbElement, err := ta_lib.MACD(datas)
		if nil != err{
			glog.Error("ta_lib.MACD failed")
			return
		}
		fmt.Println(outBeg)

		//找到买入点
		iByPos := -1
		iSellPos := -1
		for i := 0; 1 > i; i++{
			iStart := -1

			for iPos := int32(0); outNbElement > iPos; iPos++{
				if 0.00001 > outMACD[iPos] || 0.00001 > outMACDSignal[iPos]{
					continue
				}

				if -1 == iStart{
					if outMACDSignal[iPos] > outMACD[iPos]{
						iStart = int(iPos)
					}else{
						continue
					}
				}

				if -1 == iByPos{
					if outMACDSignal[iPos] <= outMACD[iPos]{
						iByPos = int(iPos)
					}
					continue
				}

				if outMACDHist[iPos] < outMACDHist[iPos - 1]{
					iSellPos = int(iPos)
					fmt.Println("============================================")
					fmt.Println("iCurPos", iCurPos, " iStart:", iStart, " iByPos:", iByPos, " iSellPos:", iSellPos)
					fmt.Println("buyPrice:", dao.G_coinAllData.CoinData[iByPos].Close, " sellPrice:", dao.G_coinAllData.CoinData[iSellPos].Close)
					fmt.Println("++++++++++++++++++++++++++++++++++++++++++++")

					iCurPos = iCurPos + iSellPos
					break
				}
			}

			if -1 == iStart || -1 == iByPos || -1 == iSellPos{
				iCurPos = iCurPos + coinDataPullModel.MACD_CAL_MAX_COUNT
				break
			}
		}
	}




}











