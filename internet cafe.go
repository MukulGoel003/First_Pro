package main

import (
	"fmt"
	"time"
	"math/rand"

)

func main() {

	var arr [25]int
	for i:=0;i<=24;i++{
		arr[i]=i+1
	}

	fmt.Println()

    var c2t chan int=make(chan int)
	var t2c chan int=make(chan int)
	for i:=1;i<=8;i++{
		go computers(i,c2t,t2c)
	}
	go tourists(arr,t2c,c2t)

	var inp string
	fmt.Scanln(&inp)

}

func tourists(arr [25]int,t2c,c2t chan int) {
	t2c <- arr[0]
	for i := 1; i <= 25; i++ {
		dup_c2t := <-c2t
		if dup_c2t != 0 {
			t2c <- arr[i]
		}
	}
}
func computers(i int,c2t,t2c chan int) {
	dup := <-t2c
	for dup != 0 {
			var amt time.Duration
			amt = time.Duration(rand.Intn(5))
			fmt.Println("Tourist ", dup, " is online on comp ", i)
			time.Sleep(time.Second * amt)
			fmt.Println("Tourist ", dup, " is done, having spent ", int(amt), " seconds online")
			c2t <- i
		/* else if dup == 0 {
			fmt.Println("The place is empty, let's close up and go to the beach!")

		} else if dup != 0 {
			fmt.Println("Tourist ", dup, " is waiting for the turn")
		}*/
	}
	fmt.Println("The place is empty, let's close up and go to the beach!")
}