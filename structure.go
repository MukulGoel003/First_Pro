package main

import "fmt"

type person struct {
	fname,iname string
	favfood []string
}

func main() {

	p :=person{"mukul","goel",[]string{"pizza","gulab jamun","thandai",}}
	fmt.Println(p.favfood)
	for i,v :=range p.favfood{
		fmt.Println(v,i)
	}
	fmt.Println(p.walk())

}

func (p person) walk () string{

	s:=fmt.Sprintln(p.fname," is walking.")
return s
}
