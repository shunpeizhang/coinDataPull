package waveAnalyse

import (
	"coinDataPull/commonUtil/coinDataPullUtil"
	"github.com/golang/glog"
	"math"
	"fmt"
)

/*
判断数据出现交叉
参数：
	first对second是否出交叉
	isUp = 1，如果first向上交叉second，那么返回true,其他返回false
	isUp = -1, 如果first向下交叉second，那么返回true,其他返回false
	isUp = 0, 只要交叉就行
*/
func WaveAnalyse_IsCross(first []float64, second []float64, isUp int8) bool{
	dataLen := len(first)

	if 1 == isUp{
		if (first)[dataLen - 1] > (second)[dataLen - 1]{
			if (first)[dataLen - 2] <= (second)[dataLen - 2]{
				return true
			}
		}
	}else if 0 == isUp{
		if (first)[dataLen - 1] < (second)[dataLen - 1]{
			if (first)[dataLen - 2] >= (second)[dataLen - 2]{
				return true
			}
		}

		if (first)[dataLen - 1] > (second)[dataLen - 1]{
			if (first)[dataLen - 2] <= (second)[dataLen - 2]{
				return true
			}
		}
	}else if -1 == isUp{
		if (first)[dataLen - 1] < (second)[dataLen - 1]{
			if (first)[dataLen - 2] >= (second)[dataLen - 2]{
				return true
			}
		}
	}else{
		return false
	}

	return false
}

/*
	之前有交叉，到目前一直在维持, 不需要开口越来越大
	first在second之上
	diff为最后一个数据需要的y间距
*/
func WaveAnalyse_IsCrossAndContinue(first []float64, second []float64, diff float64) bool{
	dataLen := len(first)

	if (first)[dataLen - 1] > (second)[dataLen - 1] && diff <= (first)[dataLen - 1] - (second)[dataLen - 1]{
		return true
	}

	return false
}

/*
	之前有交叉，到目前一直在维持, 不需要开口越来越大
	first在second之上
	disX为x间距
*/
func WaveAnalyse_IsCrossAndContinueByXDistance(first []float64, second []float64, disX int) bool{
	dataLen := len(first)

	if dataLen <= disX{
		disX = dataLen - 1
	}

	for iPos := 0; disX > iPos; iPos++{
		if (first)[dataLen - 1 - iPos] < (second)[dataLen - 1 - iPos]{
			return false
		}
	}

	return true
}

/*
	之前有交叉，到目前一直在维持, 不需要开口越来越大
	first在second之上
	disX为x间距
	disY为最后一个数据需要的y间距
*/
func WaveAnalyse_IsCrossAndContinueByXAndYDistance(first []float64, second []float64, disX int, disY float64) bool{
	dataLen := len(first)

	if dataLen <= disX{
		disX = dataLen - 1
	}

	for iPos := 0; disX > iPos; iPos++{
		if (first)[dataLen - 1 - iPos] < (second)[dataLen - 1 - iPos]{
			return false
		}
	}

	if disY > (first)[dataLen - 1] - (second)[dataLen - 1]{
		return false
	}

	return true
}


/*
速率计算

1:11窗口大小关系
degree是度数

*/
func WaveAnalyse_speedRate(data []float64, degree float64) bool{
	dataLen := len(data)

	//得到最大最小值
	max, min, err := coinDataPullUtil.GetMaxAndMin(data)
	if nil != err{
		glog.Error(err.Error())
		return false
	}
	fmt.Println("GetMaxAndMin ", "  max:", max, " min:", min)
	fmt.Println("GetMaxAndMin ", "  data:", data)

	//得到单位长度的对应值
	unitValue := 10.0 / (max - min)
	yDistance := ((data)[dataLen - 1] - (data)[dataLen - 2]) * unitValue
	xDistance := 110.0 / 100.0

	//计算tan值
	targetValue := math.Tan(3.14 * degree / 180.0)
	if targetValue <= yDistance / xDistance{
		fmt.Println("unitValue:", unitValue,  "  yDistance:", yDistance, " xDistance:", xDistance, " targetValue:", targetValue)
		return true
	}

	return false
}









