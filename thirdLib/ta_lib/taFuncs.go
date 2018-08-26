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
	"coinDataPull/handleModules/baseModule/coinDataPullModel"
	"github.com/golang/glog"
	"errors"
)

func init(){
	C.TA_Initialize()
}


//return: macd, signal, hist
//func MACD(data [coinDataPullModel.MACD_CAL_MAX_COUNT]models.KLineData) (*[coinDataPullModel.MACD_CAL_MAX_COUNT]float64, *[coinDataPullModel.MACD_CAL_MAX_COUNT]float64, *[coinDataPullModel.MACD_CAL_MAX_COUNT]float64, int32, int32, error) {
//	var outBeg int32
//	var outNbElement int32
//
//	outMACD := [coinDataPullModel.MACD_CAL_MAX_COUNT]float64{}
//	outMACDSignal := [coinDataPullModel.MACD_CAL_MAX_COUNT]float64{}
//	outMACDHist := [coinDataPullModel.MACD_CAL_MAX_COUNT]float64{}
//
//	inReal := [coinDataPullModel.MACD_CAL_MAX_COUNT]float64{}
//	for iPos := 0; coinDataPullModel.MACD_CAL_MAX_COUNT > iPos; iPos++{
//		inReal[iPos] = data[iPos].Close
//	}
//
//	retCode := C.TA_MACD(C.int(0), C.int(coinDataPullModel.MACD_CAL_MAX_COUNT - 1),
//		(*C.double)(unsafe.Pointer(&inReal)),
//		C.int(12), C.int(26), C.int(9),
//		(*C.int)(unsafe.Pointer(&outBeg)),
//			(*C.int)(unsafe.Pointer(&outNbElement)),
//		(*C.double)(unsafe.Pointer(&outMACD)),
//		(*C.double)(unsafe.Pointer(&outMACDSignal)),
//		(*C.double)(unsafe.Pointer(&outMACDHist)))
//	if 0 != retCode{
//		glog.Error("C.TA_MACD failed!")
//		return nil, nil, nil, 0, 0, errors.New("C.TA_MACD failed!")
//	}
//
//	fmt.Println(retCode)
//	fmt.Println(outMACD)
//
//	return &outMACD, &outMACDSignal, &outMACDHist, outBeg, outNbElement, nil
//}

func MACD(data *[coinDataPullModel.MACD_CAL_MAX_COUNT]models.KLineData, stMACDResultInfo *coinDataPullModel.STMACDResultInfo) error {
	for iPos := 0; coinDataPullModel.MACD_CAL_MAX_COUNT > iPos; iPos++{
		stMACDResultInfo.SourceData[iPos] = data[iPos]
	}

	inReal := [coinDataPullModel.MACD_CAL_MAX_COUNT]float64{}
	for iPos := 0; coinDataPullModel.MACD_CAL_MAX_COUNT > iPos; iPos++{
		inReal[iPos] = data[iPos].Close
	}

	retCode := C.TA_MACD(C.int(0), C.int(coinDataPullModel.MACD_CAL_MAX_COUNT - 1),
		(*C.double)(unsafe.Pointer(&inReal)),
		C.int(12), C.int(26), C.int(9),
		(*C.int)(unsafe.Pointer(&stMACDResultInfo.OutBeg)),
		(*C.int)(unsafe.Pointer(&stMACDResultInfo.OutNbElement)),
		(*C.double)(unsafe.Pointer(&stMACDResultInfo.OutMACD)),
		(*C.double)(unsafe.Pointer(&stMACDResultInfo.OutMACDSignal)),
		(*C.double)(unsafe.Pointer(&stMACDResultInfo.OutMACDHist)))
	if 0 != retCode{
		glog.Error("C.TA_MACD failed!")
		return errors.New("C.TA_MACD failed!")
	}

	return nil
}

//RSI
func RSI(dayCount int32, data *[coinDataPullModel.MACD_CAL_MAX_COUNT]models.KLineData, stRSIResultItemInfo *coinDataPullModel.STRSIResultItemInfo) (error){
	inReal := [coinDataPullModel.MACD_CAL_MAX_COUNT]float64{}
	for iPos := 0; coinDataPullModel.MACD_CAL_MAX_COUNT > iPos; iPos++{
		inReal[iPos] = data[iPos].Close
	}

	retCode := C.TA_RSI(C.int(0), C.int(coinDataPullModel.MACD_CAL_MAX_COUNT - 1),
		(*C.double)(unsafe.Pointer(&inReal)),
		C.int(dayCount),
		(*C.int)(unsafe.Pointer(&stRSIResultItemInfo.OutBeg)),
		(*C.int)(unsafe.Pointer(&stRSIResultItemInfo.OutNbElement)),
		(*C.double)(unsafe.Pointer(&stRSIResultItemInfo.Rsi)))
	if 0 != retCode{
		glog.Error("C.RSI failed!", dayCount)
		return errors.New("C.RSI failed!")
	}

	return nil
}

func RSIAll(data *[coinDataPullModel.MACD_CAL_MAX_COUNT]models.KLineData, stRSIResultInfo *coinDataPullModel.STRSIResultInfo) error{
	for iPos := 0; coinDataPullModel.MACD_CAL_MAX_COUNT > iPos; iPos++{
		stRSIResultInfo.SourceData[iPos] = data[iPos]
	}

	//5
	err := RSI(5, data, &stRSIResultInfo.Rsi1)
	if nil != err{
		glog.Error("RSI 5 failed!")
		return err
	}

	//10
	err = RSI(10, data, &stRSIResultInfo.Rsi2)
	if nil != err{
		glog.Error("RSI 10 failed!")
		return err
	}

	//14
	err = RSI(14, data, &stRSIResultInfo.Rsi3)
	if nil != err{
		glog.Error("RSI 14 failed!")
		return err
	}

	return nil
}


//KDJ
func KDJ(data *[coinDataPullModel.MACD_CAL_MAX_COUNT]models.KLineData, stKDJResultInfo *coinDataPullModel.STKDJResultInfo) error{
	for iPos := 0; coinDataPullModel.MACD_CAL_MAX_COUNT > iPos; iPos++{
		stKDJResultInfo.SourceData[iPos] = data[iPos]
	}

	inHigh := [coinDataPullModel.MACD_CAL_MAX_COUNT]float64{}
	for iPos := 0; coinDataPullModel.MACD_CAL_MAX_COUNT > iPos; iPos++{
		inHigh[iPos] = data[iPos].High
	}

	inLow := [coinDataPullModel.MACD_CAL_MAX_COUNT]float64{}
	for iPos := 0; coinDataPullModel.MACD_CAL_MAX_COUNT > iPos; iPos++{
		inLow[iPos] = data[iPos].Low
	}

	inClose := [coinDataPullModel.MACD_CAL_MAX_COUNT]float64{}
	for iPos := 0; coinDataPullModel.MACD_CAL_MAX_COUNT > iPos; iPos++{
		inClose[iPos] = data[iPos].Close
	}

	retCode := C.TA_STOCH(C.int(0), C.int(coinDataPullModel.MACD_CAL_MAX_COUNT - 1),
		(*C.double)(unsafe.Pointer(&inHigh)),(*C.double)(unsafe.Pointer(&inLow)),(*C.double)(unsafe.Pointer(&inClose)),
		C.int(5),C.int(3),C.int(0),C.int(3),C.int(0),
		(*C.int)(unsafe.Pointer(&stKDJResultInfo.OutBeg)),
		(*C.int)(unsafe.Pointer(&stKDJResultInfo.OutNbElement)),
		(*C.double)(unsafe.Pointer(&stKDJResultInfo.K)),
		(*C.double)(unsafe.Pointer(&stKDJResultInfo.D)))
	if 0 != retCode{
		glog.Error("C.KDJ failed!")
		return errors.New("C.KDJ failed!")
	}

	return nil
}









