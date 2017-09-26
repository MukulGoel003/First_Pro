package main

import "fmt"

func main() {
	sl:=make([][]int,0)

	for i:=0;i<3 ;i++  {
		sl1:=make([]int,0)
		for j:=0;j<4 ;j++  {
			sl1=append(sl1,j)
		}
		sl=append(sl,sl1)
	}
	fmt.Println(sl)
}
