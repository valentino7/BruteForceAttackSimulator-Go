package main

import (
	"fmt"
)

func main(){
	for i:=0;i!=999999;i++{
		//fmt.Printf("%d\033[30D",i)

		fmt.Printf("%d\n\033[A",i)
	}
}
