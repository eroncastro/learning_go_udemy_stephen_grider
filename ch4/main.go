package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	ContactInfo
}

func (p Person) Print() {
	fmt.Printf("%+v\n", p)
}

type ContactInfo struct {
	Email   string
	ZipCode int
}

func main() {
	p := Person{FirstName: "John", LastName: "Doe"}

	p.Print()

	p.FirstName = "Jane"

	p.Print()

	var p2 Person
	p2.ContactInfo.Email = "lala@lala.com"
	p2.ContactInfo.ZipCode = 12345
	p2.Print()

	p3 := Person{
		FirstName: "Jim",
		LastName:  "Anderson",
		ContactInfo: ContactInfo{
			Email:   "jimanderson@email.com",
			ZipCode: 12345,
		},
	}

	p3.Print()
}
