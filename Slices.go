package main

import "fmt"

type mj struct{
	x int32
	y float32
	z string
}

func main() {
var mjs mj
	c:=&mjs

	fmt.Println(c.x)
}
