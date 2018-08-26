package  main


import (
	"time"
	"flag"
	"runtime"
	"coinDataPull/handleModules/baseModule"
	"github.com/golang/glog"
	"coinDataPull/handleModules/controlModule/algorithmTestModel"
)

func main(){
	var confPath string
	flag.StringVar(&confPath, "conf", "./", "If non-empty, load config file in this directory")

	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	err := baseModule.Init(confPath)
	if nil != err{
		glog.Error(err.Error())
		return
	}

	//if !pullDataModule.Init(){
	//	glog.Error("commonHandle.Init failed!")
	//	return
	//}
	//
	//pullDataModule.Start()

	algorithmTestModel.Init()
	algorithmTestModel.RunTest()

	for true{
		time.Sleep(time.Duration(1000 * 1000 * 1000 * 1000))
	}
}










