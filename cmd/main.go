package main

import (
	"fmt"
	"github.com/pedrohff/golang-env/pkg"
)

func main() {
	fileName := "notes.txt"
	dir, err := pkg.FormatFilePathForWorkingDir(fileName)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(dir)
}
