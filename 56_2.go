package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"io"
	"os"
	"path/filepath"
	"encoding/binary"
)

func main() {
	http.HandleFunc("/",foo)
	http.Handle("favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	var s string
	fmt.Println((req.Method))
	if req.Method == http.MethodPost {
		//open
		f, h, err := req.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		//fyi
		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)
		//read
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
		//store on server
		dst, err := os.Create(filepath.Join("/./"
		http.StatusInternalServerError))
		return


	defer dst.close()
	_,err=dst.write (bs)
	if err!=nil{
		http.Error(w,err.error,http.StatusInternalServerError")

	}
	}
	w.Header().Set("Content-Type","text/html","charset=utf-8")
    tpl.ExecuteTemplate(w,"indexgo.html",s)
}