package main

import (
	"fmt"
	"sync"
)

func main() {
var wg sync.WaitGroup
	num := 5
	wg.Add(num)

	c := make(chan int)
	for i := 1; i <= num; i++ {
		go func() {
			c <- i
			wg.Done()
		}()

	}
	go func() {
		wg.Wait()
		close(c)
	}()

	pro := 1
	for j := range c {
		pro += j
		fmt.Println(j)
	}
	fmt.Println(pro)
}
