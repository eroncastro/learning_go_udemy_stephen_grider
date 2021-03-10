package main

import (
	"cards/cards"
	"fmt"
)

func main() {
	cardsList := cards.Deck{"Ace of Diamonds", newCard()}

	cardsList = append(cardsList, "Six of Spades")
	cardsList.Print()

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
}

func newCard() string {
	return "Five of Diamonds"
}
