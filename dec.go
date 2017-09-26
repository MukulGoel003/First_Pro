package main

import "fmt"

func main() {

	s:="I am sorry"
	fmt.Println(s)
	fmt.Println(string([]byte(s)))
	fmt.Println((s[:6]))
	fmt.Printf("%T\n",s)
	for _, v := range []byte(s) {
		fmt.Println(string(v))
	}

}
