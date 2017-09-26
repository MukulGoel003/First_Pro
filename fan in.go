package main

import (
	"fmt"

)

func main() {

	c:=two(one("joey"),one("chandler"))
	for i:=0;i<10;i++{
		fmt.Println(<-c)
	}

}

func one(s string) chan string{
	c:=make(chan string)
	go func() {
		for i:=0;i<10;i++{
			c<-s
		}
		close(c)
	}()
	return c
}

func two(c1,c2 chan string) chan string{
	c:=make(chan string)
	go func() {
		for{
			c<- <-c1
		}

	}()

	go func() {
		for{
			c<- <-c2
		}
	}()

	return c
}