package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("enter the characters")
	var f string
	var words []string
	fmt.Scanln(&f)

		fe, err := http.Get("http://www-personal.umich.edu/~jlawler/wordlist")
		if err != nil {
			fmt.Println(err)
		}
		defer fe.Body.Close()

		a, err1 := ioutil.ReadAll(fe.Body)
		if err != nil {
			fmt.Println(err1)
		}

		str := string(a)
		words = strings.Split(str, "\n")


	for{
		var ans string
		for i,val:=range words{
			flag:=strings.HasPrefix(words[i],f)
			if flag==true{
				fmt.Println(i,val)
			}
		}
		fmt.Println("Have you found your word? (y/n)")
		fmt.Scanln(&ans)
		if ans=="y"{
			break
		}else if ans=="n"{
			fmt.Println("Enter the next characters of your word")
			var f1 string
			fmt.Printf("%s",f)
			fmt.Scanln(&f1)
			f+=f1
		}
	}
	var index int
	fmt.Println("enter the index of the word selected")
	fmt.Scanln(&index)
	fmt.Println(words[index])
}


