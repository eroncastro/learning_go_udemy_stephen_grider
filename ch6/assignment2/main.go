package main

import "fmt"

type Triangle struct {
	Width  float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Height * t.Width
}

type Square struct {
	Length float64
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

type Shape interface {
	Area() float64
}

func printArea(s Shape) {
	fmt.Println(s.Area())
}

func main() {
	t := Triangle{Width: 10.0, Height: 10.0}
	s := Square{Length: 10.0}

	printArea(t)
	printArea(s)
}
