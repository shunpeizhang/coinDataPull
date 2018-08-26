package coinDataPullUtil

import "time"

//得到给定时间对应的整体时间，向之前取整
//输入参数单位为秒
func GetHourTime(timePoint int64) int64{
	t := time.Unix(timePoint - 3600 * 8, 0)
	resultT := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, time.UTC)

	return resultT.Unix()
}




