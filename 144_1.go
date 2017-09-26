package main

import "fmt"

func main() {

for n:=range two(two(one(2,3))){
	fmt.Println(n)
}

}

func one(n ... int) chan int{

	out:=make(chan int)
	go func() {
		for _,num:=range n{
			out<-num
		}
		close(out)
	}()
	return out

}

func two(c chan int)chan int{
	out:=make(chan int)
	go func() {
		for gh:=range c{
			out<-gh*gh
		}
		close(out)
	}()
	return out
}