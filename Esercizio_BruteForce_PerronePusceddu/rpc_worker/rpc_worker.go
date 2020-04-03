package rpc_worker

import (
	"strconv"
	"fmt"
	"sync"
)

// Forcing service for RPC
type Forcing string
type Closer string
type Result int
type Arguments struct{
	Password string
	Min int
	Max int
}


var B bool
var Mux sync.Mutex

// convert : convert a non 8 character string to an 8 character string
func convert(test string) string{
	var temp string
	word := test
	length := len( word )
	if length  < 8 {
		for j := 0 ; j < 8-length ; j ++ {
			temp += "0"
		}
		temp += word
	} else {
		temp = word
	}
	return temp
}

/* test_password : 
	return true and password found or return false and "notfound" string*/
func test_password (test int, psw string, string_max string)(bool,string){
	final_string := convert( strconv.Itoa(test) )
	fmt.Printf("forcing. . . %s/%s \n\033[A",final_string,string_max)
	if final_string == psw {
		return true , final_string
	}
	return false ,"notfound"
}

/*check_finish_bruteforce : check if other worker already found the password, in that case the worker stop the elaboration*/
func check_finish_bruteforce(check *bool) {
	Mux.Lock()
	if B == true {
		B = false
		*check = true
	}
	Mux.Unlock()
}

// procedure to stop the service of forcing if another worker found the password
func (t *Closer) Close_service(C bool, Result *int) error {
	Mux.Lock()
	B=C
	Mux.Unlock()
	*Result=1
	return nil
}
// procedure to find a password from a predefinite space of password
func (t *Forcing) Forcing_work(args Arguments, Psw *string) error {

	str_max:=strconv.Itoa(args.Max-1)
	str_max= convert(str_max)
	fmt.Println("Range of values: ",args.Min,args.Max)
	var check bool
	for i := args.Min ; i< args.Max ; i++ {
		control, s := test_password(i,args.Password,str_max)

		check_finish_bruteforce(&check)
		if check==true {
			break
		}
		
		if  control {
			fmt.Printf("\n")
			*Psw = s
			fmt.Println("Forcing finisced, passoword found :",*Psw)
			return nil
		}

	}
	*Psw = "notfound"
	fmt.Printf("\n")
	fmt.Println("Forcing finisced, password not found :",*Psw)
	return nil
}
