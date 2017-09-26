package main

import (
	"path/filepath"
	"os"
	"fmt"
)

func main() {

	filepath.Walk(".", func(path string,info os.Fileinfo, err error) error){
		fmt.Println(path)
		return nil
	}

}
