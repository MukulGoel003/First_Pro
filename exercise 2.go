package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {

	fmt.Println("Let's go for a walk!")
    var amt1,amt2,amt11,amt21 time.Duration

	fmt.Println("Jon started getting ready")
	amt1= time.Duration(rand.Intn(30)+61)

	fmt.Println("Jack started getting ready")
	amt2=time.Duration(rand.Intn(30)+61)

	if amt1>=amt2 {

		 go func() {

			time.Sleep(time.Second * amt2)
			fmt.Println("Jack spent ", int(amt2), " seconds getting ready")
		}()


		 func() {

			time.Sleep(time.Second * amt1)
			fmt.Println("Jon spent ", int(amt1), " seconds getting ready")
		}()

	} else{


		 go func() {

			time.Sleep(time.Second * amt1)
			fmt.Println("Jon spent ", int(amt1), " seconds getting ready")
		}()

		func() {

			time.Sleep(time.Second * amt2)
			fmt.Println("Jack spent ", int(amt2), " seconds getting ready")
		}()

	}

	amt11= time.Duration(rand.Intn(35)+11)
	amt21= time.Duration(rand.Intn(35)+11)
	fmt.Println("Arming the alarm")

	go func() {

		time.Sleep(time.Second*8)
		fmt.Println("Alarm is armed")

	}()


	if amt11>=amt21{

		go func(){

			fmt.Println("Jack started putting on shoes")

			time.Sleep(time.Second * amt21)
			fmt.Println("Jack spent ", int(amt21), " seconds putting on shoes")


		}()

		func(){

			fmt.Println("Jon started putting on shoes")
			time.Sleep(time.Second * amt11)
			fmt.Println("Jon spent ", int(amt11), " seconds putting on shoes")


		}()




	} else{

		go func(){

			fmt.Println("Jon started putting on shoes")
			time.Sleep(time.Second * amt11)
			fmt.Println("Jon spent ", int(amt11), " seconds putting on shoes")

		}()

		func(){

			fmt.Println("Jack started putting on shoes")
			time.Sleep(time.Second * amt21)
			fmt.Println("Jack spent ", int(amt21), " seconds putting on shoes")

		}()


	}

	fmt.Println("Exiting and locking the door.")

	var st string
	fmt.Scanln(&st)

}
