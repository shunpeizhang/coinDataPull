package coinDataPullUtil

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/golang/glog"
	"coinDataPull/handleModules/baseModuleles/baseModule/coinDataPullConfig"
)

var mysqlDB *sql.DB                         //mysql的客户端对象



//初始化
func InitMysqlInterface() bool{
	var err error
	mysqlAddr := fmt.Sprintf("%v:%v@tcp(%v:%v)/coin_data?charset=utf8", coinDataPullConfig.G_configInfo.MysqlConfig.UserName, coinDataPullConfig.G_configInfo.MysqlConfig.Passwd, coinDataPullConfig.G_configInfo.MysqlConfig.Host, coinDataPullConfig.G_configInfo.MysqlConfig.Port)
	mysqlDB, err = sql.Open("mysql", mysqlAddr)
	if nil != err{
		glog.Error("sql.Open failed! ", mysqlAddr)
		return false
	}

	mysqlDB.SetMaxIdleConns(coinDataPullConfig.G_configInfo.MysqlConfig.MysqlConnMaxIdle)
	mysqlDB.SetMaxOpenConns(coinDataPullConfig.G_configInfo.MysqlConfig.MysqlConnMaxOpen)

	return true
}

//得到mysql对象
func GetMysqlDB() *sql.DB {
	return mysqlDB
}










