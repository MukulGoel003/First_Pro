package main

import "fmt"

func main() {

	fan:=make([][]string,2)
	st1:=make([]string,2)
	st1[0]="abc"
	st1[1]="xyz"
	st2:=make([]string,2)
	st2[0]="ram"
	st2[1]="jio"
	fan=append(fan,st1)
	fan=append(fan,st2)
	fmt.Println(fan)



}
