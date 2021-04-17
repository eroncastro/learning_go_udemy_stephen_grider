package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("File argument is missing.")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
