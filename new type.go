package main

import "fmt"

type gator int

type flamingo bool

func main() {

	var g1 gator
	var f flamingo
	f=true
	f.metgator()
	g1=23
	fmt.Println(g1)
	fmt.Printf("\n%T",g1)
 var x int
	fmt.Println(x)
	fmt.Printf("%T",x)
	g1.metgator()
	final(g1,f)
}


func (g1 gator) metgator()string{
	fmt.Println("hello gator")
	return ("hey")
}

func(f flamingo) metgator()string{
fmt.Println("hello flamingo")
	return ("hello")
}

type inter interface {
	metgator() string
}

func final (hj...inter){
	for _,v:=range hj{
		fmt.Println(v.metgator())
	}
}