package main

import "fmt"

type cir struct{
	l int32
	b int32
	c string
}

type cir2 struct{
	e int32
	g string
	cir
}

func main() {

	c2:=new(cir2)
	//c:=cir{2,3,"linh"}
	fmt.Println(c2.cir.l)
	c2.cir.example()
	c2.example()

}

func (c*cir) example() {
	fmt.Println("hello there!")
}
