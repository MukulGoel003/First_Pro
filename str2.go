package main

import "fmt"

type circle struct {
	length int32
	br int32

}

func main() {
	c:=circle{6,9}
	fmt.Println(c.area())

}

func (c1 *circle) area() int32{
	return c1.length*c1.br
}