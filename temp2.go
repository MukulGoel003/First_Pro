package main

import (
	"fmt"
	"os"
	"log"
	"io"
	"strings"
)

func main(){
	name:=os.Args[1]
	//fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	tpl:=fmt.Sprint(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>HELLO WORLD!</title>
	</head>
	<body>
	<h1>`+name+`</h1>
	</body>
	</html>`)

	nf,err:=os.Create("index2.html")
	if err!=nil{
		log.Fatal("error creating file",err)
	}
	defer nf.Close()

	io.Copy(nf,strings.NewReader(tpl))
	//fmt.Println(tpl)
}
