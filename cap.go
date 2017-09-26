package main

import "fmt"

func main() {

	var sl []int
	arr:=[]int{2,3,4,3,3,3,2,3,3,3,3,3,3}
	sl=arr[:8]
	fmt.Println(cap(sl),len(sl))
}
