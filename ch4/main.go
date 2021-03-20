package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func main() {
	p := Person{FirstName: "John", LastName: "Doe"}

	fmt.Println(p)

	p.FirstName = "Jane"

	fmt.Println(p)
}
