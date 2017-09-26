package main

import (
	"bufio"
	"strings"
	"fmt"
)

func main(){
	s:="Hey there\nall ok\ngood bye"
	scanner:=bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan(){
		line:=scanner.Text()
		fmt.Println(line)
	}
}
