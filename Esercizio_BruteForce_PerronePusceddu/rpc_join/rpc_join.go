package rpc_join

import (
	"fmt"
	"net/rpc"
	"log"
	"sync"
)

// Join service for RPC
type Join string

//Result of RPC call
type Result int

type Workers struct{
	Ip,Port_hack string
}


// Slc -> slice of clients : contains the workers joined at master
var Slc []rpc.Client
var Mux sync.Mutex


func connect_worker(ip string, port string) *rpc.Client{
	fmt.Println(ip,port)
	client, err := rpc.DialHTTP("tcp", ip+":"+port)

	if err != nil {
		log.Fatal("Error in dialing: ", err)
	}
	return client
}

//procedure to join the server
func (t *Join) Join_master(w Workers, Result *int) error {

	client:=connect_worker(w.Ip,w.Port_hack)
	Mux.Lock()
	Slc=append(Slc,*client)
	Mux.Unlock()
	fmt.Println("bind effettuato ")
	*Result = 1
	return nil
}

