package main

import (
	"sync"
	"fmt"
	"time"
)

func main() {
	//arr:=[]int{15,2,3,45,5,6,7,8,9,10}
	//fmt.Println("array is : ",arr)
	j:=one(1,2,3,4,5,6,7,8,9,10)
c1:=twosq(j)
	c2:=twosq(j)
	out:=threemer(c1,c2)
	for v:=range out{
		fmt.Println(v)
	}


}

func one(i ...int)chan int{
	out:=make(chan int)
	go func() {
		for _,r:=range i{
			out<-r
			//fmt.Println("1st fn : ",r)
			time.Sleep(time.Millisecond)
		}
		close(out)
	}()
	return out
}

func twosq(in chan int)chan int{
	out:=make(chan int)
	go func() {
		for k:=1;k<=10;k++{
			for v:=range in{
				out<-v*v
				//fmt.Println("gugu",v)
			}
		}

		close(out)
	}()
	return out
}

func threemer(in ... chan int)chan int{
	out:=make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(in))

	for _,v:=range in{
		go func() {
			for k:=range v{
				out<-k
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out

}