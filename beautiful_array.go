package main

import (
	"fmt"
	"os"
)

func main() {

	var flag bool
	var t int
	fmt.Scanln(&t)

	for t>0{
		var n int
		fmt.Scanln(&n)
		var arr=make([]int,n)
		for i:=0;i<n;i++{
			fmt.Scanf("%d",&arr[i])
		}
		for j:=0;j<n-1;j++{
			for k:=j+1;k<n;k++{

				flag=check(arr,n,arr[j]*arr[k])
				if flag == false{
					fmt.Println("no")
					os.Exit(0)
				}
			}
		}
		fmt.Println("yes")
		t--
	}
}

func check(arr []int,n,num int ) bool{
	for i:=0;i<n;i++{
		if num==arr[i]{
			return true
		}
	}
	return false
}
