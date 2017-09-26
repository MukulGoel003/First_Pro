package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(

		strings.Contains(strings.ToLower("fgx"),"op"),
	    strings.Count("hopefully","l"),
	strings.Join([]string{"a","b","c","d"},"-"),
		strings.Replace("aaaaaammmmaaaammm","m","n",6),
	)

	arr:=[]byte("test")
	fmt.Println(arr)

}
