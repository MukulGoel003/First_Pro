package main

import (
	"strings"
	"html/template"
	"os"
	"fmt"
)
var tpl *template.Template
type st struct{
	Name string
	Age int
}

func FirstThree(s string) string{
	s=strings.TrimSpace(s)
	s=s[:3]
	return s
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": FirstThree,

}

func init(){
tpl=template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
}

func main(){


	s1:=st{"mukul",22}
	s2:=st{"aman",21}
	sli:=[]st{s1,s2}

	err:=tpl.ExecuteTemplate(os.Stdout,"index.gohtml",sli)
	if err!=nil{
		fmt.Println("error")
	}
}
