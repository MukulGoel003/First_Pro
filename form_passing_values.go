package main

import (
	"net/http"
	"io"
)

func main(){
	http.HandleFunc("/",foo)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func foo(w http.Response, req*http.Request){
	v:=req.FormValue("q")
	w.Header().Set("Content_type","text/html: charset=utf-8")
	io.WriteString(w,`
	<form method="post"
	<input type="text" name='q'>
	<input type="submit">
	</form>
	<br>`+v)

}