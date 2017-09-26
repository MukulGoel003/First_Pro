package main

import (
	"fmt"

	"os"
)

func main() {

	var t, n int
	fmt.Println("enter value of t")
	fmt.Scanln(&t)
	var i = 0
	for i < t && (t<=100 || t>=1){
		fmt.Scanln(&n)
		if n<7 || n>100{
			os.Exit(0)
		}
		var arr= make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &arr[j])
			if (arr[j]<1 || arr[j]>10) /*&& arr[n/2]!=7*/{
				os.Exit(0)
			}
		}

		if arr[0]!=1{
			i++
			fmt.Println("no")
			continue
		}
		//var size [10]int
		var j = 0
		var k = n - 1
		flag := true
		for j < n/2 {

			if arr[k] != arr[j] || arr[j]+1!=arr[j+1] {
				//fmt.Println("no")
				flag = false
				break
			}
			k--
			j++

		}
		if flag == true {
			fmt.Println("yes")

		} else {
			fmt.Println("no")
		}
		i++
	}
}






