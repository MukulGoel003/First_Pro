package main

import "fmt"

func main() {

	2d:=make([][]string,2)
	st1:=make([]string,2)
	st1[0]="abc"
	st1[1]="xyz"
	st2:=make([]string,2)
	st2[0]="ram"
	st2[1]="jio"
	2d=append(2d,st1)
	2d=append(2d,st2)
	fmt.Println(2d)



}
