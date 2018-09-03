package main

import (
	"net/http"
	"example.com/demochain/core"
	"encoding/json"
	"io"
)

var blockchain *core.Blockchain

func run(){
	http.HandleFunc("/blockchain/get",blockchainGetHandler)
	http.HandleFunc("/blockchain/write",blockchainWriteHandler)

	//启动对端口的监听
	http.ListenAndServe("localhost:8888",nil)
}

func blockchainGetHandler(w http.ResponseWriter,r *http.Request){
	bytes,err := json.Marshal(blockchain)
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	io.WriteString(w,string(bytes))
}

func blockchainWriteHandler(w http.ResponseWriter,r *http.Request){
	blockData := r.URL.Query().Get("data")
	blockchain.SendData(blockData)
	blockchainGetHandler(w,r)
}

func main(){
	blockchain = core.NewBlockchain()
	run()
}
