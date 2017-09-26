package main

import (
	"os"
	"fmt"
)

func main() {

	dir,err:=os.Open(".")
	if err!=nil{
		return
	}

	defer dir.Close()

	fd,err:=dir.Readdir(0)
	if err!=nil{
		return
	}

	for _,fi:=range fd{
		fmt.Println(fi.Name())
	}
}
