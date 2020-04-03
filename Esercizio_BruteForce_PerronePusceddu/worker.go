package main

import (
	"./rpc_join"
	"./rpc_worker"
	"net/rpc"
	"log"
	"fmt"
	"net"
	"net/http"
	"sync"
	"os"
)


/* synchronous_call_worker
   performs synchronous call to join the master. */
func synchronous_call_worker(client rpc.Client,args rpc_join.Workers){
	// Synchronous call
	// reply will store the RPC result
	var pwdResult rpc_join.Result
	w := &rpc_join.Workers {args.Ip , args.Port_hack}
	// Call remote procedure
	err := client.Call("Join.Join_master", w, &pwdResult)
	if err != nil || pwdResult != 1 {
		log.Fatal("Error in Join.Join_master: ", err)
	}

	fmt.Println("join effettuata")
}

func connect_to_master() *rpc.Client{
	client, err := rpc.DialHTTP("tcp", "localhost:15001")
	if err != nil {
		log.Fatal("Error in dialing: ", err)
	}
	return client
}

func start_services(l_worker net.Listener,w sync.WaitGroup){
	defer w.Done()
	http.Serve(l_worker, nil)
}
/* service_work
   start offering forcing and closer services */
func service_work(port string,w sync.WaitGroup ) {
	//register forcing service
	forcing := new(rpc_worker.Forcing)
	rpc.Register(forcing)

	//register closer service
	closer := new(rpc_worker.Closer)
	rpc.Register(closer)

	rpc.HandleHTTP()

	l_worker, e := net.Listen("tcp", ":"+port)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go start_services(l_worker, w)

}


func main() {
	
	var wg sync.WaitGroup
	
	// generating casual port number

	str_port:= os.Args[1]

	//worker offers forcing and closer services
	service_work(str_port , wg)
	wg.Add(1)

	var args rpc_join.Workers
	args.Ip = "127.0.0.1"
	args.Port_hack = str_port
	// worker found master
	client:= connect_to_master()
	defer client.Close()
	//rpc to join the master
	synchronous_call_worker(*client,args)

	wg.Wait()

}
