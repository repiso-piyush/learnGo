package main

import (
	"fmt"
)
var deckSize int = 20

func main() {
   card := newCard()
   fmt.Println(card)
}

func newCard () string{
	return "Five of Diamonds"
}

// v declared and not used.
// no new variables on the left side of :=".
// myPhone := "Iqoo" should used inside funtion only not in package level