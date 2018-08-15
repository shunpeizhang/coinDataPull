package  main

import (
	"flag"
	"runtime"
	"github.com/golang/glog"
	"coinDataPull/commonHandle"
	"time"
	"coinDataPull/coinDataPullConfig"
	"coinDataPull/coinDataPullUtil"
)

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


	for true{
		time.Sleep(time.Duration(1000 * 1000 * 1000 * 1000))
	}
}









