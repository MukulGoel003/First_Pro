package main

import (
	"io/ioutil"
	"fmt"
)

func main() {

	bs,err:=ioutil.ReadFile("Test.txt")
	if err!=nil{
		return
	}
	fmt.Println(string(bs))
}
