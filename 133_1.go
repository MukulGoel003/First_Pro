package main

import (
	"fmt"
	//"time"
	//"time"
)

func main() {

	var c chan int = make(chan int)

	go func() {

		for i := 0; i < 10; i++ {
			c <- i

		}
		//close(c)

	}()
	//time.Sleep(time.Second*2)

	//go func() {
		for r:=range c{
			fmt.Println(r)
		}



	//}()
	var s string
	fmt.Scanln(&s)
}
