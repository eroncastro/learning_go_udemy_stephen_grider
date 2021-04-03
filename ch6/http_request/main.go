package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://google.com/")
	if err != nil {
		panic(err)
	}

	bs := make([]byte, 100000)

	resp.Body.Read(bs)

	fmt.Println(string(bs))
}
