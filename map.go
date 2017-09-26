package main

import "fmt"

func main() {

	 x:=map[string]int{
		 "abc":1,
		 "xza":2,
	 }

	for i,v:=range x{
		fmt.Println(i,v)
	}


}
