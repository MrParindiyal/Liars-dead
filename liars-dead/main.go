package main

import (
	"fmt"
	"liars-dead/game"
)

func main(){
	deck := game.CreateDeck("basic")
	for i := 0; i < len(deck.Cards); i++ {
		fmt.Println(i+1, deck.Cards[i].Value)
	}
	fmt.Println("---------------------")
	deck = game.ShuffleDeck(deck)
	for i := 0; i < len(deck.Cards); i++ {
		fmt.Println(i+1, deck.Cards[i].Value)
	}
}