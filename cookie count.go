package main

import (
	"net/http"
	"strconv"
	"log"
	"io"
)

func main() {
	http.HandleFunc("/",foo)
	http.Handle("favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func foo(res http.ResponseWriter,req *http.Request){
	cookie,err:=req.Cookie("my-cookie")

	if err!=nil{
		cookie=&http.Cookie{
			Name:"my-cookie",
			Value:"0",
		}
	}

	count,err:=strconv.Atoi(cookie.Value)
	if err!=nil{
		log.Fatalln(err)
	}
	count++
	cookie.Value=strconv.Itoa(count)

	http.SetCookie(res,cookie)

	io.WriteString(res,cookie.Value)
}