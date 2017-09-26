package main

import (
	"fmt"
	"sort"
)

var arr []int
var n int
var num int

func main() {
	fmt.Println("enter the length of array")
	fmt.Scanln(&n)
	fmt.Println("enter the elements of the array")
	for i:=0;i<n;i++{
		fmt.Scanf("%d",&arr[i])
	}
	fmt.Println("enter the sum")
	fmt.Scanln(&num)
	sort.Ints(arr)
	//i:=0
	//pair(i,n,arr,num)

}

//func pair(i int,n int,arr []int, num int){
//
//}