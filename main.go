package  main

import (
	"flag"
	"runtime"
	"commonCoinLib/dataSource/config"
	"github.com/golang/glog"
	"coinDataPull/commonHandle"
	"time"
)

func main(){
	var confPath string
	flag.StringVar(&confPath, "conf", "./", "If non-empty, load config file in this directory")

	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	err := config.LoadConfigInfo(confPath)
	if nil != err{
		glog.Error("config.LoadConfig failed!")
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









