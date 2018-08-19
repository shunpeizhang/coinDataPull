package ta_lib


/*
#cgo CFLAGS: -I/usr/ta-lib/include/ta-lib
#cgo LDFLAGS: /usr/ta-lib/lib/libta_lib.a -lm
#include <stdio.h>
#include "ta_libc.h"

void ta_MACD(int iLen, double *inReal, double *outMACD, double *outMACDSignal, double *outMACDHist, int *outBeg, int *outNbElement, int *ret)
{
	TA_RetCode retCode = TA_MACD(0, iLen - 1, inReal, 12, 26, 9, outBeg, outNbElement, outMACD, outMACDSignal, outMACDHist);
	*ret = (int)retCode;
}

*/
import "C"
import (
	"coinDataPull/thirdLib/huobiapi/models"
	"unsafe"
	"fmt"
)

func init(){
	C.TA_Initialize()
}


//return: macd, signal, hist
func MACD(data []models.KLineData) ([]float64, []float64, []float64, error) {
	var outBeg int32
	var outNbElement int32
	var ret int32

	iDataLen := len(data)
	outMACD := [iDataLen]float64{}
	outMACDSignal := [iDataLen]float64{}
	outMACDHist := [iDataLen]float64{}

	inReal := [iDataLen]float64{}
	for iPos := 0; iDataLen > iPos; iPos++{
		inReal[iPos] = data[iPos].Close
	}

	C.ta_MACD(iDataLen, (*C.double)(unsafe.Pointer(&inReal)),
		(*C.double)(unsafe.Pointer(&outMACD)),
		(*C.double)(unsafe.Pointer(&outMACDSignal)),
		(*C.double)(unsafe.Pointer(&outMACDHist)),
		(*C.int)(unsafe.Pointer(&outBeg)),
		(*C.int)(unsafe.Pointer(&outNbElement)),
		(*C.int)(unsafe.Pointer(&ret)))

	fmt.Println(ret)
	fmt.Println(outMACD)

	return outMACD[0:], outMACDSignal[0:], outMACDHist[0:], nil
}





