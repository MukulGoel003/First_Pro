package main

import (
	"net/http"
	//"fmt"
	"io"
	"html/template"
)

var tpl *template.Template

func init(){
	tpl=template.Must(template.ParseFiles("index.gohtml"))
}

func a (res http.ResponseWriter,req *http.Request){
	n:="hello !"
	res.Header().Set("Content-type","text/html; charset=utf-8")
	//fmt.Fprintln(res,n)
	io.WriteString(res,n)
	tpl.ExecuteTemplate(res,"index.gohtml","value passed into the template")
}

func dog (res http.ResponseWriter,req *http.Request){
	io.WriteString(res,"in func dog")
}

func me (res http.ResponseWriter, req *http.Request){
	io.WriteString(res,"hello mukul")
}

func main() {

http.HandleFunc("/",a)
http.HandleFunc("/dog/",dog)
	http.HandleFunc("/me/",me)

	http.ListenAndServe(":8080",nil)
}
