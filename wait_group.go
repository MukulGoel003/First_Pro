package main

import (
	"sync"
	"fmt"
)

func main() {

	c:=make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		//wg.Add(1)
		for i:=0;i<10;i++{
			c<-i
		}

		wg.Done()
	}()


	go func(){
		//wg.Add(1)
		for i:=0;i<10;i++{
			c<-i
		}

		wg.Done()
	}()

	go func(){

		wg.Wait()
		close(c)
	}()

	for i:=range c{
		fmt.Println(i)
	}
}

