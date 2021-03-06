package pullDataModule

import (
	"coinDataPull/handleModules/baseModule/coinDataPullModel"
	"github.com/robfig/cron"
	"fmt"
	"github.com/golang/glog"
)


var crontab *cron.Cron
var isInit bool

func Init() bool{
	crontab = cron.New()

	//添加定时处理
	{
		spec := fmt.Sprintf("5 */1 * * * ?")
		crontab.AddFunc(spec, coinDataPull_1min)

		spec = fmt.Sprintf("5 */5 * * * ?")
		crontab.AddFunc(spec, coinDataPull_5min)

		spec = fmt.Sprintf("5 */15 * * * ?")
		crontab.AddFunc(spec, coinDataPull_15min)

		spec = fmt.Sprintf("5 */30 * * * ?")
		crontab.AddFunc(spec, coinDataPull_30min)

		spec = fmt.Sprintf("5 0 */1 * * ?")
		crontab.AddFunc(spec, coinDataPull_60min)

		spec = fmt.Sprintf("5 0 0 */1 * ?")
		crontab.AddFunc(spec, coinDataPull_1day)
	}

	isInit = true

	return true
}

func Start() bool{
	if isInit{
		crontab.Start()

		return true
	}

	return false
}




func coinDataPull_1min(){
	coinDataPull_commonHandle(coinDataPullModel.COINSRCTIMETYPE_1MIN)
}

func coinDataPull_5min(){
	coinDataPull_commonHandle(coinDataPullModel.COINSRCTIMETYPE_5MIN)
}

func coinDataPull_15min(){
	coinDataPull_commonHandle(coinDataPullModel.COINSRCTIMETYPE_15MIN)
}

func coinDataPull_30min(){
	coinDataPull_commonHandle(coinDataPullModel.COINSRCTIMETYPE_30MIN)
}

func coinDataPull_60min(){
	coinDataPull_commonHandle(coinDataPullModel.COINSRCTIMETYPE_60MIN)
}

func coinDataPull_1day(){
	coinDataPull_commonHandle(coinDataPullModel.COINSRCTIMETYPE_1DAY)
}


//处理数据获取
func coinDataPull_commonHandle(coinSrcTimeType int32){
	glog.Info("coinSrcTimeType:", coinSrcTimeType)
	fmt.Println("coinSrcTimeType:", coinSrcTimeType)

	needGetCount := coinDataPullModel.COINDATAGET_COUNTPRE * (1 + coinDataPullModel.COINDATAGET_DATAAddRATE)
	for iCoinTye := coinDataPullModel.COINTYPE_BTC_USDT; coinDataPullModel.COINTYPE_MAX > iCoinTye; iCoinTye++{
		GetCoinDataAndpersist(int32(iCoinTye), coinSrcTimeType, uint32(needGetCount))

		glog.Info("handle finish iCoinTye:", iCoinTye)
	}
}













