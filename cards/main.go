package main

import "fmt"

// var card string
// You can initalize a var outside func main but  you must assign its value inside the main function

func main() {
	//var card string = "Ace of Spades"
	//card := "Ace of Spades",
	card := newCard()
	//card = "Five of Diamonds"

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
