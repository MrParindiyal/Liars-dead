package game

import "slices"

// import (
// 	"crypto/rand"
// )

type Player struct {
	PlayerId    string
	SeatId      int
	Name        string
	Hand        []Card
	Lives       int
	IsSpectator bool
}

func NewPlayer(name string, playerId string) *Player {
	return &Player{
		PlayerId:    playerId,
		Name:        name,
		Lives:       6,
		IsSpectator: false,
	}
}

func (p *Player) DeductLife() {
	p.Lives--
}

func (p *Player) MarkSpectator() {
	p.IsSpectator = true
}

func (p *Player) MarkActive() {
	p.IsSpectator = false
}

func (p *Player) RemovePlayedCards(indices []int) []Card {
	var updatedHand []Card
	var cardsCollected []Card
	for i, card := range p.Hand {
		if !slices.Contains(indices, i) {
			updatedHand = append(updatedHand, card)
		} else {
			cardsCollected = append(cardsCollected, card)
		}
	}

	p.Hand = updatedHand
	return cardsCollected
}
