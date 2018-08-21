package  main


import (
	"time"
	"flag"
	"runtime"
	"coinDataPull/coinDataPullConfig"
	"github.com/golang/glog"
	"coinDataPull/coinDataPullUtil"
	"coinDataPull/commonHandle"
)

//
//func test(){
//	//var data = [coinDataPullModel.MACD_CAL_MAX_COUNT]models.KLineData{}
//
//	//ta_lib.MACD(data)
//
//	os.Exit(0)
//}
//


func main(){
	var confPath string
	flag.StringVar(&confPath, "conf", "./", "If non-empty, load config file in this directory")

	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	err := coinDataPullConfig.LoadConfigInfo(confPath)
	if nil != err{
		glog.Error("config.LoadConfig failed!")
		return
	}

	if !coinDataPullUtil.InitMysqlInterface(){
		glog.Error("coinDataPullUtil.InitMysqlInterface failed!")
		return
	}

	if !commonHandle.Init(){
		glog.Error("commonHandle.Init failed!")
		return
	}

	commonHandle.Start()


	//test()
	//dao.Init()
	//verify.VerifyMACD()

	for true{
		time.Sleep(time.Duration(1000 * 1000 * 1000 * 1000))
	}
}










