package main

import (
	"fmt"
	"liars-dead/game"
)

func main() {
	deck := game.CreateDeck("chaos")
	// for i := 0; i < len(deck.Cards); i++ {
	// 	fmt.Println(i+1, deck.Cards[i].Value)
	// }
	// fmt.Println("---------------------")

	deck = game.ShuffleDeck(deck)
	// for i := 0; i < len(deck.Cards); i++ {
	// 	fmt.Println(i+1, deck.Cards[i].Value)
	// }
	fmt.Println("---------------------")

	hands := game.DealNHands(deck, 3)
	for x := range 3 {
		player := game.NewPlayer(fmt.Sprintf("player%d", x), x)
		player.Hand = hands[x].Cards

		fmt.Println(player.Uid)
		fmt.Println(player.PlayerId)
		fmt.Println(player.Name)
		fmt.Println(player.Hand)
		fmt.Println(player.Lives)
		fmt.Println(player.IsSpectator)
		fmt.Println("===========")
	}
}
