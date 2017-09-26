package main

import (
	"net"
	"fmt"
	"log"
	"bufio"
)

func main(){
	li,err:=net.Listen("tcp",":8080")
	if err!=nil{
		fmt.Println("error1")
	}
	defer li.Close()

	for{
		conn,err:=li.Accept()
		if err!=nil{
			log.Println("error")
		}
		go handle(conn)
	}
}

func handle(conn net.Conn){
	scanner:=bufio.NewScanner(conn)
	for scanner.Scan(){
		ln:=scanner.Text()
		fmt.Println(ln)
		fmt.Fprintln(conn,ln)
	}
	defer conn.Close()
}
