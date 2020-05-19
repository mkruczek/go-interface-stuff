package main

import (
	"io"
	"os"
)

func main() {

	filePath := getFilePath(os.Args)
	f, err := os.Open(filePath)
	if err != nil {
		// panic(err)
	}

	io.Copy(os.Stdout, f)
}

func getFilePath(args []string) string {
	if len(args) <= 1 {
		return "file"
	}
	return args[1]
}
