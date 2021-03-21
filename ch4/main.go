package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	ContactInfo
}

type ContactInfo struct {
	Email   string
	ZipCode int
}

func main() {
	p := Person{FirstName: "John", LastName: "Doe"}

	fmt.Printf("%+v\n", p)

	p.FirstName = "Jane"

	fmt.Printf("%+v\n", p)

	var p2 Person
	p2.ContactInfo.Email = "lala@lala.com"
	p2.ContactInfo.ZipCode = 12345
	fmt.Printf("%+v\n", p2)

	p3 := Person{
		FirstName: "Jim",
		LastName:  "Anderson",
		ContactInfo: ContactInfo{
			Email:   "jimanderson@email.com",
			ZipCode: 12345,
		},
	}

	fmt.Printf("%+v\n", p3)
}
