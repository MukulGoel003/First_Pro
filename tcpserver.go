package main

import (
	"net"
	"fmt"
	"encoding/gob"
)

func server(){
	ln,err:=net.Listen("tcp",":9999")
	if err!=nil{
		fmt.Println(err)
		return
	}
	for{
		c,err:=ln.Accept()
		if err!=nil{
			fmt.Println(err)
			continue
		}

		go handleserverconnection(c)
	}
}


func handleserverconnection(c net.Conn){
	var msg string
	err:=gob.NewDecoder(c).Decode(&msg)
	if err!=nil{
		fmt.Println(err)

	}else{
		fmt.Println("received",msg)
	}
	c.Close()
}

func client(){
	c,err:=net.Dial("tcp","127.0.0.1:9999")
	if err!=nil{
		fmt.Println(err)
		return
	}

	msg:="hey there!"
	fmt.Println("Sending",msg)
	err=gob.NewEncoder(c).Encode(msg)
	if err!=nil{
		fmt.Println(err)
	}
	c.Close()
}

func main() {

	go server()
	go client()

	var inp string
	fmt.Scanln(&inp)

}
