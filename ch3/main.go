package main

import (
	"cards/cards"
	"fmt"
)

func main() {
	var myList [10]string

	myList[0] = "Lala"

	for _, text := range myList {
		fmt.Println(text)
	}

	myList2 := []string{"lala"}

	myList3 := append(myList2, "popo")

	fmt.Println(myList2)
	fmt.Println(myList3)

	myDeck := cards.NewDeck()
	myDeck.Print()

	mySlice := []string{"a", "b", "c", "d", "e"}
	fmt.Println(len(mySlice))
	fmt.Println(mySlice[len(mySlice):])

	draw, newDeck := cards.Deal(myDeck, 6)
	fmt.Println(len(draw))
	fmt.Println(len(newDeck))
}
