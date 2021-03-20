package main

import "fmt"

func main() {
	var numbers []int

	for i := 0; i <= 10; i++ {
		numbers = append(numbers, i)
	}

	for number := range numbers {
		if number%2 != 0 {
			fmt.Printf("%d is odd\n", number)
		} else {
			fmt.Printf("%d is even\n", number)
		}
	}
}
