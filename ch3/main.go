package main

import (
	"cards/cards"
	"fmt"
	"log"
)

func main() {
	myDeck := cards.NewDeck()

	err := myDeck.SaveToFile("popo.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Reading data")

	myNewDeck, err := cards.NewDeckFromFile("popo.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(myNewDeck.ToString())

	fmt.Println("After shuffling")

	err = myNewDeck.Shuffle()
	if err != nil {
		panic(err)
	}

	fmt.Println(myNewDeck.ToString())
}
