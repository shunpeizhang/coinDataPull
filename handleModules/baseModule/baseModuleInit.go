package baseModule

import (
	"coinDataPull/handleModules/baseModule/coinDataPullConfig"
	"github.com/golang/glog"
	"errors"
	"coinDataPull/commonUtil/coinDataPullUtil"
)

func Init(confPath string) error{
	err := coinDataPullConfig.LoadConfigInfo(confPath)
	if nil != err{
		glog.Error("config.LoadConfig failed!")
		return err
	}

	if !coinDataPullUtil.InitMysqlInterface(){
		glog.Error("coinDataPullUtil.InitMysqlInterface failed!")
		return errors.New("coinDataPullUtil.InitMysqlInterface failed!")
	}

	return nil
}




