
package rpc_struct

import (
	"../rpc_join"
	"../rpc_worker"
	"fmt"
	"math"
	"reflect"
	"net/rpc"
	"strconv"
	"log"
)

// Hack service for RPC
type Hack string

//Result of RPC call
type Result string


// procedure for bruteForce a password
func (t *Hack) BruteForce(Psw string, Result *string) error {

	rpc_join.Mux.Lock()
	l := len(rpc_join.Slc)
	rpc_join.Mux.Unlock()
	var chans []chan *rpc.Call
	if l==0 {
		fmt.Println("No worker avaiable, retry later")
		*Result = "noworker"
	}else{
		fmt.Println("\n Workers Start")
		Password := make([]string,l)
		space_number := math.Pow(10,8)
		
		fmt.Println("Password value space:",strconv.Itoa(int(space_number)))

		//space for one worker
		var single_space float64 =  space_number / float64(l)
		control :=int(math.Ceil(single_space))

		for i := 0 ; i< l ; i++ {
			max:= (i+1)* control
			min:= i* control
			if max > int(space_number)-1{
				max = int(space_number)
			}
			args := rpc_worker.Arguments{Psw,min, max }
			fmt.Println("Number worker:", i,"; Worker space:",args.Min,args.Max)
			// asynchronous call
			c := rpc_join.Slc[i].Go("Forcing.Forcing_work", &args, &Password[i], nil)
			chans = append(chans, c.Done)
		}
		cases := make([]reflect.SelectCase, len(chans))
		for i, ch := range chans {
			cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
		}

		remaining := len(cases)
		for remaining > 0 {

			chosen, _, ok := reflect.Select(cases)
			if !ok {
				// The chosen channel has been closed, so zero out the channel to disable the case
				cases[chosen].Chan = reflect.ValueOf(nil)
				remaining -= 1
				continue
			}
			if Password[chosen] !="notfound"{
				*Result = Password[chosen]

				for i := 0 ; i< l ; i++ {
					if i == chosen{
						continue
					}
					args :=true
					var pwdResult rpc_worker.Result
					// Call remote procedure

					err := rpc_join.Slc[i].Call("Closer.Close_service", args, &pwdResult)
					if err != nil || pwdResult != 1 {
						fmt.Println("i",i)
						log.Fatal("Error in Closer: ", err)
					}
					
				}
				return nil
			}

		}
		*Result = "notfound"

	}
	return nil
}

