package main

import (
	"fmt"
	"bufio"
	"os"
	"errors"
	"./rpc_struct"
	"log"
	"net/rpc"
)


func check_arg(str string) error{

	for i:=range str{
		if str[i]< 48 || str[i]> 57 {
			err:= errors.New("Error in character")
			return err
		}
	}
	return nil
}


func synchronous_call(client *rpc.Client,str string){
	// Synchronous call
	// reply will store the RPC result
	var pwdResult rpc_struct.Result


	// Call remote procedure
	err := client.Call("Hack.BruteForce", str, &pwdResult)
	if err != nil {
		log.Fatal("Error in Hack.BruteForce: ", err)
	}
	fmt.Println("Hack.BruteForce: The password is: ", pwdResult)
}

func connect_server() *rpc.Client{
	client, err := rpc.DialHTTP("tcp", "localhost:15000")
	if err != nil {
		log.Fatal("Error in dialing: ", err)
	}
	return client
}


func main(){


	fmt.Println("Enter 8 numeric characters: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		//chiamata remota
		if len(scanner.Text())!=8 || check_arg(scanner.Text())!=nil {
			fmt.Println("Error, enter 8 numeric characters:")
		}else{
			client :=connect_server()
			//chiamata remota
			synchronous_call(client,scanner.Text())
			client.Close()
			fmt.Print("Enter 8 numeric characters: ")
		}
	}
}
