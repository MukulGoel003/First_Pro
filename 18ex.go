package main

import (
	"html/template"
	"os"
	"fmt"
)


type hotel struct {
	name string
	address string
	zip int64
}
func main(){
	hotel1:=hotel{"amaze","shakuntala palace",110085}
	hotel2:=hotel{"abc hotel","abc address",110006}
	sli:=[]hotel{hotel1,hotel2}
	tl:=template.Must(template.ParseFiles("index3.html"))
	err:=tl.ExecuteTemplate(os.Stdout,"index3.html",sli)
if err!=nil{
	fmt.Println("error")
}
	fmt.Println(sli)
}