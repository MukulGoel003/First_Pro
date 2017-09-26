package main

import (
	"net"
	"fmt"
)

func main() {

	li,err:=net.Listen("tcp",":8080")
	if err!=nil{
		fmt.Println("error1",err)
	}
	defer li.Close()

	for{
		conn,err:=li.Accept()
		if err!=nil{
			fmt.Println("error 2",err)
		}
		fmt.Fprintln(conn,"i see u connected")
	}
}
