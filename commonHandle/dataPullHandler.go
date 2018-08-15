package commonHandle

import (
	"coinDataPull/thirdLib/huobiapi/services"
	"coinDataPull/coinDataPullConfig"
	"github.com/golang/glog"
	"github.com/syndtr/goleveldb/leveldb/errors"

	"coinDataPull/thirdLib/huobiapi/models"
	"coinDataPull/coinDataPullModel"
	"coinDataPull/dbInterface"
)

//处理某类型某时间的数据获取,并存储到数据库
func GetCoinDataAndpersist(coinType int32, coinSrcTimeType int32, getCount uint32) error{
	//参数检测
	{
		_, res := coinDataPullConfig.COINTYPE_NAME_MAP[coinType]
		if !res{
			glog.Error("coinType not exist! coinType:", coinType)
			return errors.New("coinType not exist")
		}

		_, res = coinDataPullConfig.COINSRCTIMETYPE_TABLENAME_MAP[coinSrcTimeType]
		if !res{
			glog.Error("coinSrcTimeType not exist! coinType:", coinType)
			return errors.New("coinSrcTimeType not exist")
		}

		_, res = coinDataPullConfig.COINSRCTIMETYPE_PERIOD_MAP[coinSrcTimeType]
		if !res{
			glog.Error("coinSrcTimeType not exist! coinType:", coinType)
			return errors.New("coinSrcTimeType not exist")
		}

		if coinDataPullModel.COINDATAGET_LIMTCOUNT < getCount{
			glog.Error("coinDataPullModel.COINDATAGETLIMTCOUNT < getCount error! getCount:", getCount)
			return errors.New("coinDataPullModel.COINDATAGETLIMTCOUNT < getCount error!")
		}
	}

	//取数据
	var ret models.KLineReturn
	ret = services.GetKLine(coinDataPullConfig.COINTYPE_NAME_MAP[coinType], coinDataPullConfig.COINSRCTIMETYPE_PERIOD_MAP[coinSrcTimeType], int(getCount))
	if "ok" != ret.Status{
		glog.Error("services.GetKLine failed! ", "coinType:", coinType, " coinSrcTimeType:", coinSrcTimeType, " getCount:", getCount)
		return errors.New("services.GetKLine failed! ")
	}

	//将数据持久化
	retCount := len(ret.Data)
	for iPos := 0; retCount > iPos; iPos++{
		err := dbInterface.Insert_CoinKLineData(&ret.Data[iPos], coinDataPullConfig.COINSRCTIMETYPE_TABLENAME_MAP[coinSrcTimeType], coinType)
		if nil != err{
			//如果失败，继承处理下一个，因为可能之前数据在表中已经存在
			continue
		}
	}

	return nil
}












