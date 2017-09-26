package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/",set)
	http.HandleFunc("/read",read)
	http.Handle("/favicon.ico",http.NotFoundHandler())
http.ListenAndServe(":8080",nil)
}

func set(w http.ResponseWriter,req *http.Request){
	http.SetCookie(w,&http.Cookie{
		Name:"mycookie",
		Value:"v",
	})
	fmt.Fprintln(w,"cookie written")

}

func read(w http.ResponseWriter,req *http.Request){
	c,err:=req.Cookie("mycookie")
	if err!=nil{
		http.Error(w,err.Error(),http.StatusNoContent)
		return
	}
	fmt.Fprintln(w,"YOUR COOKIE:",c)
}
