package main

import (
	"bufio"
	"strings"
	"fmt"
)

func main(){
	s:="Hey there\nI am happy"
	scanner:=bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan(){
		line:=scanner.Text()
		fmt.Println(line)
	}

}