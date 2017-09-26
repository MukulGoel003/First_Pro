package main

import "fmt"

func main() {

	arr:=[]int{2,3,4}
	fun(arr)

}
 func fun(arr []int){
	 for i,v:=range arr{
		fmt.Println(i,v)
	 }

 }