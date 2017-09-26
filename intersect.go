package main

import (
	"fmt"
	//"sort"
	"sort"
)

func main() {


	fmt.Println("enter length of 1st array")
	var n int
	var n1 int
	fmt.Scanln(&n)
	fmt.Println("enter length of second array")
	fmt.Scanln(&n1)
	arr:=make([]int,n)
	arr2:=make([]int,n1)
	fmt.Println("enter 1st array:")
	for i:=0;i<n;i++{
		fmt.Scanln(&arr[i])
	}
	fmt.Println("enter 2nd array:")
	for i:=0;i<n1;i++{
		fmt.Scanln(&arr2[i])
	}

	sort.Ints(arr)
	sort.Ints(arr2)



	i:=0
	j:=0

	for i<n && j<n1{
		if arr[i]==arr2[j]{
			fmt.Println(arr[i])
			i++
			j++
		}else if arr[i]>arr2[j] {
			j++
		}else {
			i++
		}

	}

}
