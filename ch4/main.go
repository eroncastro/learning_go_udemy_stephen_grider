package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func main() {
	p := Person{FirstName: "John", LastName: "Doe"}

	fmt.Printf("%+v\n", p)

	p.FirstName = "Jane"

	fmt.Printf("%+v\n", p)

	var p2 Person
	fmt.Printf("%+v\n", p2)
}
