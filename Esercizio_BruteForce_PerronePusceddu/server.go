package main

import (
	"net/rpc"
	"log"
	"net"
	"./rpc_struct"
	"./rpc_join"
	"net/http"
)


func main() {

	//register hack and join services
	hack := new(rpc_struct.Hack)
	join := new(rpc_join.Join)
	rpc.Register(join)
	rpc.Register(hack)

	rpc.HandleHTTP()

	l_client, e := net.Listen("tcp", ":15000")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l_client, nil)


	l_join, e := net.Listen("tcp", ":15001")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l_join, nil)
}
