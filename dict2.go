package main

import (
	"net/http"
	"fmt"
)

func main(){
	fmt.Println(http.Get("http://www-personal.umich.edu/~jlawler/wordlist"))
}
