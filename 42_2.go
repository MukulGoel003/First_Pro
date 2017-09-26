package main

import (
	"net/http"
	"io"
	//"os"
	"os"
)

func main() {
	http.HandleFunc("/",dog)
	//http.HandleFunc("/gm.jpg",dogpic)
	http.ListenAndServe(":8080",nil)
}

func dog(w http.ResponseWriter,req *http.Request){

	w.Header().Set("Content-Type","text/html;charset=utf-8")
	f,err:=os.Open("pro.txt")
	if err!=nil{
				http.Error(w,"file not found",404)
				return
			}
			defer f.Close()
	io.WriteString(w,`
	<img src="/gm.jpg">`)
}

//func dogpic(w http.ResponseWriter,req *http.Request){
//	f,err:=os.Open("pro.txt")
//	if err!=nil{
//		http.Error(w,"file not found",404)
//		return
//	}
//	defer f.Close()
//
//	io.Copy(w,f)
//
//}