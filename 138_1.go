package main

import "fmt"

func main() {
	cha:=one("faa")
	cha11:=one("boo")
	cha2:=two(cha)
	cha22:=two(cha11)
	fmt.Println("final counter",<-cha2+<-cha22)

}

func one (s string) chan int{
	ch1:=make(chan int)
	go func(){
		for i:=0;i<10;i++{
			ch1<-i
			fmt.Println(s,i)
		}
		close(ch1)

	}()

	return ch1
}

func two(c2 chan int) chan int{
	c3:=make(chan int)
	go func(){
		var sum int
		for i:=range c2{

			sum+=i
		}
		c3<-sum
		close(c3)

	}()
	return c3
}