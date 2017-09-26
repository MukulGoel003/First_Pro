package main

import "fmt"

type vehicle struct {
	color string
	doors int
}

type sedan struct {
	vehicle
	luxury bool
}

type truck struct{
	vehicle
	fourwheels bool
}

func main() {

	s:=sedan{vehicle{"blue",4},true}
	t:=truck{vehicle{"red",2},false}
	fmt.Println(s,t)
	fmt.Println(s.color,s.doors,s.luxury,t.color,t.doors,t.fourwheels)
fmt.Println(s.tran())
	fmt.Println(t.tran())
	final(t,s)
}

func(s sedan)tran()string{
	return("this is sedan")
}

func (t truck)tran()string{
	return("this is truck")
}

type trans interface {
	tran()string
}

func final(i ...trans){
	for j,v:=range i{
		fmt.Println(j,v.tran())
	}
}