package main

import (
	"io"
	"os"
)

func main() {
	filePath := os.Args[1] //check if exist

	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, f)
}
