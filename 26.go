package main

import (
	"net"
	"log"
	"io"
	"fmt"
	"bufio"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn){
	defer conn.Close()
	request(conn)
	respond(conn)
}

func request(conn net.Conn){
	i:=0
	scanner:=bufio.NewScanner(conn)
	for scanner.Scan(){
		ln:=scanner.Text()
		fmt.Println(ln)
		if i==0{
			m:=strings.Fields(ln)[0]
			fmt.Println("***METHOD",m)
		}
		if ln==""{
			break
		}
		i++
	}
}

func respond(conn net.Conn){
	body:=`<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><TITLE></TITLE><body>HELLO WORLD</body></html>`

}
