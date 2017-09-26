package main

import (
	"os"
	"fmt"
)

func main() {


	file,err:=os.Create("test.txt")
	if err!=nil{
		return
	}

	file.WriteString("mukul goel")
file.Close()
	file,err= os.Open("test.txt")
	if err !=nil{
		return
	}

	defer file.Close()

	stat,err:=file.Stat()
	if err!=nil {
		return
	}

	bs:=make([]byte,stat.Size())
	f,err:=file.Read(bs)

	if err!=nil{
		return
	}
    fmt.Printf("%T",f)
	str:=string(bs)
	fmt.Println(str)

}
