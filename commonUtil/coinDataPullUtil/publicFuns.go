package coinDataPullUtil

import (
	"time"
	"errors"
)

//得到给定时间对应的整体时间，向之前取整
//输入参数单位为秒
func GetHourTime(timePoint int64) int64{
	t := time.Unix(timePoint - 3600 * 8, 0)
	resultT := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, time.UTC)

	return resultT.Unix()
}

//得到数据中的最大值与最小值 (max, min, error)
func GetMaxAndMin(data []float64) (float64, float64, error){
	dataLen := len(data)
	if 0 == dataLen{
		return 0, 0, errors.New("0 == dataLen")
	}

	var max, min float64
	isInit := false
	for iPos := 0; dataLen > iPos; iPos++{
		if !isInit{
			max = (data)[iPos]
			min = (data)[iPos]
		}else{
			if (data)[iPos] > max{
				max = (data)[iPos]
			}

			if (data)[iPos] < min{
				min = (data)[iPos]
			}
		}
	}

	return max, min, nil
}


