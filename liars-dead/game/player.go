package game

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
