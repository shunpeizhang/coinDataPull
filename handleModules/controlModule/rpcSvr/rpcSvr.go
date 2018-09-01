package rpcSvr

import (
	"net/rpc"
	"net"
	"log"
	"net/http"
	"coinDataPull/thirdLib/ta_lib"
	"fmt"
)

func Init(){
	ta_lib.Init()
}

func Run(){
	taLibInfo := new(ta_lib.STTaLibInfo)
	rpc.Register(taLibInfo)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":8081")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)



	fmt.Println("rpc server start success")
}


