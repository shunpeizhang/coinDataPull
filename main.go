package  main


import (
	"time"
	"flag"
	"runtime"
	"github.com/golang/glog"
	"coinDataPull/handleModules/controlModule/rpcSvr"
	"coinDataPull/handleModules/controlModule/pullDataModule"
	"coinDataPull/handleModules/baseModule"
	"coinDataPull/handleModules/controlModule/algorithmTestModel"
)

func main(){
	var confPath string
	var rpcFlag int
	var pullDataFlag int
	flag.StringVar(&confPath, "conf", "./", "If non-empty, load config file in this directory")
	flag.IntVar(&rpcFlag, "rpc", 0, "rpc")
	flag.IntVar(&pullDataFlag, "pullData", 0, "pullData")

	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	err := baseModule.Init(confPath)
	if nil != err{
		glog.Error(err.Error())
		return
	}

	if 1 == rpcFlag{
		rpcSvr.Init()
		rpcSvr.Run()
	}else if 1 == pullDataFlag{
		if !pullDataModule.Init(){
			glog.Error("commonHandle.Init failed!")
			return
		}
		pullDataModule.Start()
	}else{
		algorithmTestModel.Init()
		algorithmTestModel.RunTest()
	}

	for true{
		time.Sleep(time.Duration(1000 * 1000 * 1000 * 1000))
	}
}










