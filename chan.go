package main

import (
	"fmt"
	"time"
)


func main() {

	var c chan string=make(chan string)

	go pinger(c<-)
	go ponger(c<-)
	//go pinger(c)
	go printer(<-c)

	var inp string
	fmt.Scanln(&inp)

}

func ponger(c  chan<- string){
	for i:=0;i<3;i++{
		c<-"pong"
	}

}

func pinger(c chan<-string){
	for i:=0;i<2;i++{
	c<-"ping"
	}
}

func printer(c<-chan string){
	for{
		msg:=<-c
		fmt.Println(msg)
		time.Sleep(time.Second*1)
	}
}