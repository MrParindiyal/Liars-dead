package main

import (
	"crypto/rand"
	"fmt"
	"liars-dead/game"
)

func main() {
	deck := game.CreateDeck("chaos")
	deck = game.ShuffleDeck(deck)
	fmt.Println("---------------------")

	hands := game.DealNHands(deck, 3)
	for x := range 3 {
		playerId := rand.Text()[:7]
		player := game.NewPlayer(fmt.Sprintf("player%d", x), playerId)
		player.Hand = hands[x].Cards

		fmt.Println(player.PlayerId)
		fmt.Println(player.SeatId)
		fmt.Println(player.Name)
		fmt.Println(player.Hand)
		fmt.Println(player.Lives)
		fmt.Println(player.IsSpectator)
		fmt.Println("===========")
	}
}
