package game

import "slices"

type PlayerState string

type Player struct {
	PlayerId string
	SeatId   int
	Name     string
	Hand     []Card
	Lives    int
	State    PlayerState
}

const (
	Spectator    PlayerState = "spectator"
	LobbyWaiting PlayerState = "waiting"
	LobbyReady   PlayerState = "ready"
	InGame       PlayerState = "playing"
)

func NewPlayer(name string, playerId string) *Player {
	return &Player{
		PlayerId: playerId,
		Name:     name,
		Lives:    6,
		State:    LobbyWaiting,
	}
}

func (p *Player) DeductLife() {
	p.Lives--
}

func (p *Player) MarkSpectator() {
	p.State = Spectator
}

func (p *Player) MarkActive() {
	p.State = LobbyWaiting
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
