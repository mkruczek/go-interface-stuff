package main

import (
	"fmt"
	"io"
	"net/http"
)

type storage struct {
	content string
}

func (s storage) Write(p []byte) (n int, err error) {
	s.content = string(p)
	fmt.Println(s.content)
	return len(p), nil
}

func main() {

	resp, err := http.Get("http://google.com")
	if err != nil {
		panic(err)
	}

	io.Copy(storage{}, resp.Body) 
}
