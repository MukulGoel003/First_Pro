package main

import "fmt"

type cus struct {
	mj int32
	ned string
	hj float32

}

//c:=new(cus)
var c cus

func main() {

	fmt.Println(c.mj,c.ned,c.hj)
	var gh int32
	fmt.Scanf("%d",&gh)
	fmt.Println(gh)

}
