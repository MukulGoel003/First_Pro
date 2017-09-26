package main

import (
	"net/http"
	"fmt"
)

type muk int

func (m muk) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"in the function")
}

func main(){
	var m muk
http.ListenAndServe(":8080",m) //m is of type handler as muk is implementing the method of the same signature as by the handler interface
//or any value of muk type is implicitly implementing handler interface
}
