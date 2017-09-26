package main

import (
	"sync"
	"fmt"
	//"sync/atomic"
	"sync/atomic"
)

var count int32
var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go incrementor("1")
	go incrementor("2")

	wg.Wait()
	fmt.Print("final counter",count)

}

func incrementor (s string) {
	for i:=0;i<20;i++{
		atomic.AddInt32(&count,1)
		fmt.Println("process",s,"printing",i)
	}
	wg.Done()

}
