package coinDataPullConfig

import (
	"os"
	"fmt"
	"errors"
	"gopkg.in/yaml.v2"
	"github.com/golang/glog"
	"coinDataPull/handleModules/baseModule/coinDataPullModel"
)

//mysql配制
type MysqlConfigInfo struct {
	Host string
	Port int
	UserName string
	Passwd string
	MysqlConnMaxIdle int         //mysql连接池最大连接数量
	MysqlConnMaxOpen int         //mysql连接池最大已连接的数量
}


//配制信息
type ConfigInfo struct{
	MysqlConfig MysqlConfigInfo
}


var G_configInfo ConfigInfo
var COINTYPE_NAME_MAP map[int32]string    //保存类型与名称之间的映射关系
var COINSRCTIMETYPE_TABLENAME_MAP map[int32]string  //保存时间与表名的对应关系
var COINSRCTIMETYPE_PERIOD_MAP map[int32]string  //保存时间与时间名

func init(){
	COINTYPE_NAME_MAP = make(map[int32]string)
	COINSRCTIMETYPE_TABLENAME_MAP = make(map[int32]string)
	COINSRCTIMETYPE_PERIOD_MAP = make(map[int32]string)

	COINTYPE_NAME_MAP[coinDataPullModel.COINTYPE_BTC_USDT] = "btcusdt"
	COINTYPE_NAME_MAP[coinDataPullModel.COINTYPE_ETH_USDT] = "ethusdt"
	COINTYPE_NAME_MAP[coinDataPullModel.COINTYPE_BCH_USDT] = "bchusdt"
	COINTYPE_NAME_MAP[coinDataPullModel.COINTYPE_ETC_USDT] = "etcusdt"
	COINTYPE_NAME_MAP[coinDataPullModel.COINTYPE_LTC_USDT] = "ltcusdt"
	COINTYPE_NAME_MAP[coinDataPullModel.COINTYPE_EOS_USDT] = "eosusdt"
	COINTYPE_NAME_MAP[coinDataPullModel.COINTYPE_XRP_USDT] = "xrpusdt"
	COINTYPE_NAME_MAP[coinDataPullModel.COINTYPE_OMG_USDT] = "omgusdt"
	COINTYPE_NAME_MAP[coinDataPullModel.COINTYPE_DASH_USDT] = "dashusdt"
	COINTYPE_NAME_MAP[coinDataPullModel.COINTYPE_ZEC_USDT] = "zecusdt"
	COINTYPE_NAME_MAP[coinDataPullModel.COINTYPE_ADA_USDT] = "adausdt"
	COINTYPE_NAME_MAP[coinDataPullModel.COINTYPE_STEEM_USDT] = "steemusdt"
	COINTYPE_NAME_MAP[coinDataPullModel.COINTYPE_IOTA_USDT] = "iotausdt"

	COINSRCTIMETYPE_TABLENAME_MAP[coinDataPullModel.COINSRCTIMETYPE_1MIN] = "tb_coinData_1min"
	COINSRCTIMETYPE_TABLENAME_MAP[coinDataPullModel.COINSRCTIMETYPE_5MIN] = "tb_coinData_5min"
	COINSRCTIMETYPE_TABLENAME_MAP[coinDataPullModel.COINSRCTIMETYPE_15MIN] = "tb_coinData_15min"
	COINSRCTIMETYPE_TABLENAME_MAP[coinDataPullModel.COINSRCTIMETYPE_30MIN] = "tb_coinData_30min"
	COINSRCTIMETYPE_TABLENAME_MAP[coinDataPullModel.COINSRCTIMETYPE_60MIN] = "tb_coinData_60min"
	COINSRCTIMETYPE_TABLENAME_MAP[coinDataPullModel.COINSRCTIMETYPE_1DAY] = "tb_coinData_1day"

	COINSRCTIMETYPE_PERIOD_MAP[coinDataPullModel.COINSRCTIMETYPE_1MIN] = "1min"
	COINSRCTIMETYPE_PERIOD_MAP[coinDataPullModel.COINSRCTIMETYPE_5MIN] = "5min"
	COINSRCTIMETYPE_PERIOD_MAP[coinDataPullModel.COINSRCTIMETYPE_15MIN] = "15min"
	COINSRCTIMETYPE_PERIOD_MAP[coinDataPullModel.COINSRCTIMETYPE_30MIN] = "30min"
	COINSRCTIMETYPE_PERIOD_MAP[coinDataPullModel.COINSRCTIMETYPE_60MIN] = "60min"
	COINSRCTIMETYPE_PERIOD_MAP[coinDataPullModel.COINSRCTIMETYPE_1DAY] = "1day"
}



//加载配制
func LoadConfigInfo(configPath string) error{
	f, err := os.Open(configPath)
	if nil != err {
		fmt.Println("open failed!" + " configPath:" + configPath + ", errorInfo" + err.Error())
		return errors.New("Open configPath failed!")
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&G_configInfo)
	if nil != err{
		glog.Error(err.Error())
		return err
	}

	fmt.Println(G_configInfo)

	return nil
}















