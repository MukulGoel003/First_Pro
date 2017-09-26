package main

import (
	"fmt"
	"strings"
)

func main() {

	dict:=map[int]string{
		1:"abacus",
		2:"apple",
		3:"boy",
		4:"zebra",
		5:"android",
		6:"zen",
		7:"baby",
		8:"cage",
	}

	fmt.Println("enter the first character")
	var f string
	fmt.Scanln(&f)
	for i,val:=range dict{
		flag:=strings.HasPrefix(dict[i],f)
		if flag==true{
			fmt.Println(i,val)
		}
	}

	var index int
	fmt.Println("enter the index of the word selected")
	fmt.Scanln(&index)
	fmt.Println(dict[index])
}
