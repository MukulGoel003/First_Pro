package main

import (
	"fmt"
	//"time"
)

func main() {
	c:=one(3,6,7,1,0)

	xyz:=two(c)
	abc:=two(c)
	/*for i:=range abc{
		fmt.Println("abc",i)
	}

	for j:=range xyz{
	fmt.Println("xyz",j)
	}*/

	cha:=three(abc,xyz)
//time.Sleep(time.Second*1)
	//close(cha)
	for r:=0;r<5;r++{
		fmt.Println(<-cha)  //drawback of this approach is that ,it can be used for only limited number and known number of channels
		                    //limited because in the last function, goroutine is to be specified for each channel()not possible for large no. of channels
	}                       //known because for printing out the channel values, range can't be used because the channel was not closed

	//var str string
	//fmt.Scanln(&str)
}

func one(num ... int)chan int{
	c:=make(chan int)
	go func() {
		for i:=range num{
			c<-i
		}
		close(c)
	}()
	return c
}

func two(c chan int)chan int{

	c1:=make(chan int)
	go func() {

		for l:=range c{
			c1<-l*l
		}
		close(c1)
	}()
return c1
}

func three(c1,c2 chan int) chan int{
	out:=make(chan int)
	go func(){
		for{
			out<-<-c1
		}
	}()

	go func(){
		for{
			out<-<-c2
		}
	}()
	return out
}