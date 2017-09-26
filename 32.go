package main

import (
	"net/http"
	"log"
	"html/template"
)

type muk int

func (m muk) ServeHTTP(w http.ResponseWriter, req *http.Request){
	err:=req.ParseForm()
	if err!=nil{
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w,"index.gohtml",req.Form)
}

var tpl *template.Template

func init(){
	tpl=template.Must(template.ParseFiles("index.gohtml"))
}

func main(){
	var m muk
	http.ListenAndServe(":8080",m)
}
