package main

import (
	"time"
	"fmt"
)

func main() {

	c1:=make(chan string)
	c2:=make(chan string)

	go func(){
		for{
			c1<-"from1"
			time.Sleep(time.Second*2)
		}
	}()

	go func(){
		for{
			c2<-"from2"
			time.Sleep(time.Second*3)
		}
	}()

	go func(){
		for{
			select{
			case msg1:=<-c1:
				fmt.Println(msg1)
			case msg2:=<-c2:
				fmt.Println(msg2)
			}
		}
	}()

	var inp string
	fmt.Scanln(&inp)

}
