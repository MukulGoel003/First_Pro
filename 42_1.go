package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/",dogg)
	http.ListenAndServe(":8080",nil)
}
func dogg(w http.ResponseWriter,req *http.Request){
	w.Header().Set("Content-Type","text/html;charset=utf-8")

	io.WriteString(w,`
	<!--Not serving from our server-->
	<img src="https://www.w3schools.com/w3images/fjords.jpg">
	`)

}
