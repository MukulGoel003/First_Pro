package main

import (

	"fmt"
)

//type shape
type shape interface{
	area() int
}

type circl struct{
	rad int

}

type rect struct {
	len int
	bre int
}

func (c *circl) area() int{
	return int(3*c.rad*c.rad)
}

func (r *rect) area() int{
	return r.len*r.bre
}

func ran(sh ...shape) {

	for _,s:=range sh{
		fmt.Println(s.area())
	}

}

func main() {
	c:=circl{2}
	r:=rect{3,3}
	fmt.Println(c.area(),r.area())
	ran(&c,&r)

}
