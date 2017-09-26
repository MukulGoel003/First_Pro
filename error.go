package main

import (
	"os"
	"fmt"
	"log"
)

func main() {

	_,err:=os.Open("no-file.txt")
	if err!=nil{
		fmt.Println("error happened")
		//panic(err)
		log.Fatalln(err)
	}

}
