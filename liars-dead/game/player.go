package game

import "crypto/rand"

type Player struct {
	Uid         string
	PlayerId    int
	Name        string
	Hand        []Card
	Lives       int
	IsSpectator bool
}

func NewPlayer(name string, playerNum int) Player {
	return Player{
		Uid:         (rand.Text())[:6],
		PlayerId:    playerNum,
		Name:        name,
		Lives:       6,
		IsSpectator: false,
	}
}
