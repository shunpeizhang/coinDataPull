package ta_lib


/*
#cgo CFLAGS: -I/usr/ta-lib/include/ta-lib
#cgo LDFLAGS: /usr/ta-lib/lib/libta_lib.a -lm
#include <stdio.h>
#include "ta_libc.h"

*/
import "C"
import (
	"coinDataPull/thirdLib/huobiapi/models"
	"unsafe"
	"fmt"
	"coinDataPull/coinDataPullModel"
	"github.com/golang/glog"
	"errors"
)

func init(){
	C.TA_Initialize()
}


//return: macd, signal, hist
func MACD(data [coinDataPullModel.MACD_CAL_MAX_COUNT]models.KLineData) (*[coinDataPullModel.MACD_CAL_MAX_COUNT]float64, *[coinDataPullModel.MACD_CAL_MAX_COUNT]float64, *[coinDataPullModel.MACD_CAL_MAX_COUNT]float64, int32, int32, error) {
	var outBeg int32
	var outNbElement int32

	outMACD := [coinDataPullModel.MACD_CAL_MAX_COUNT]float64{}
	outMACDSignal := [coinDataPullModel.MACD_CAL_MAX_COUNT]float64{}
	outMACDHist := [coinDataPullModel.MACD_CAL_MAX_COUNT]float64{}

	inReal := [coinDataPullModel.MACD_CAL_MAX_COUNT]float64{}
	for iPos := 0; coinDataPullModel.MACD_CAL_MAX_COUNT > iPos; iPos++{
		inReal[iPos] = data[iPos].Close
	}

	retCode := C.TA_MACD(C.int(0), C.int(coinDataPullModel.MACD_CAL_MAX_COUNT - 1),
		(*C.double)(unsafe.Pointer(&inReal)),
		C.int(12), C.int(26), C.int(9),
		(*C.int)(unsafe.Pointer(&outBeg)),
			(*C.int)(unsafe.Pointer(&outNbElement)),
		(*C.double)(unsafe.Pointer(&outMACD)),
		(*C.double)(unsafe.Pointer(&outMACDSignal)),
		(*C.double)(unsafe.Pointer(&outMACDHist)))
	if 0 != retCode{
		glog.Error("C.TA_MACD failed!")
		return nil, nil, nil, 0, 0, errors.New("C.TA_MACD failed!")
	}

	fmt.Println(retCode)
	fmt.Println(outMACD)

	return &outMACD, &outMACDSignal, &outMACDHist, outBeg, outNbElement, nil
}











