package main

import (
	"html/template"
	"os"
	"log"
)

func main(){
	tpl:=template.Must(template.ParseGlob("First_Pro/index.gohtml"))
singer:=[]string{"lata","asha","manna"}
	err:=tpl.ExecuteTemplate(os.Stdout,"index.gohtml",singer)
	if err!=nil{
		log.Fatalln(err)
	}

	//err=tpl.ExecuteTemplate(os.Stdout,"index2.html",nil)
	//if err!=nil{
	//	log.Fatalln(err)
	//}
}