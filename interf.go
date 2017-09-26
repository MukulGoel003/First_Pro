package main

import "fmt"

func main() {

	var val interface{}=9.9
	fmt.Println(val)
	//fmt.Println(int(val))
	fmt.Println(val.(float64))

}
