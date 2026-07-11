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

	hands := game.DealNHands(deck, 6)
	fmt.Println(hands)
	room := game.NewGameRoom("chaos", 4)
	for x := range 6 {
		playerId := rand.Text()[:7]
		player := game.NewPlayer(fmt.Sprintf("player%d", x), playerId)
		player.Hand = hands[x].Cards

		fmt.Println(player.PlayerId)
		fmt.Println(player.SeatId)
		fmt.Println(player.Name)
		fmt.Println(player.Hand)
		fmt.Println(player.Lives)
		fmt.Println(player.State)
		fmt.Println("===========")
		room.AddPlayerToTable(player)
	}

	fmt.Println(room.Players[0])
	fmt.Println(room.Players[1])
	fmt.Println(room.Players[2])
	fmt.Println(room.Players[3])
	fmt.Println(room.Players[4])
	fmt.Println(room.Players[5])
	fmt.Println(room.GetRoomstats())

	// TODO
	action := game.Action{
		ActionType: "play",
		Payload: map[string]any{
			"indices": []int{0, 2},
		},
	}
	err := game.ApplyAction(room, room.Players[2], action)

	if err != nil {
	}
}
