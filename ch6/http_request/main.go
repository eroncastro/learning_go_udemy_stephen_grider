package main

import (
	"fmt"
	"io"
	"net/http"
)

type LogWriter struct{}

func (LogWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}

func main() {
	resp, err := http.Get("https://google.com/")
	if err != nil {
		panic(err)
	}

	lw := LogWriter{}

	io.Copy(lw, resp.Body)
}
