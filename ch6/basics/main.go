package main

import "fmt"

// Interface type -> abstract
// Interfaces are implicit in Go
type GreetingBot interface {
	Greet(string) string
}

// Concrete type
type EnglishBot struct{}

func (EnglishBot) Greet(name string) string {
	return "Hello, " + name
}

// Concrete type
type SpanishBot struct{}

func (SpanishBot) Greet(name string) string {
	return "Hola, " + name
}

func printGreeting(g GreetingBot, name string) {
	fmt.Println(g.Greet(name))
}

func main() {
	englishBot := EnglishBot{}
	spanishBot := SpanishBot{}

	printGreeting(englishBot, "John")
	printGreeting(spanishBot, "John")
}
