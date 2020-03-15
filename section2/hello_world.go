package main // executable package - Go has two types: executable and reusable

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println("Hello, World!")
}
