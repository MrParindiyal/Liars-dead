package game

import (
	"crypto/rand"
	"fmt"
)

type GameState string

type GameRoom struct {
	Id                string
	Players           []*Player
	GameMode          string
	MaxAllowedPlayers int
	PlayedStack       []Card
	CurrentTurnSeat   int
	State             GameState
	nextSeatId        int
	leader            string
	TableSubject      string
	LastPlayedSize    int
	LastPlayedBy      *Player
}

const (
	StateLobby     GameState = "lobby"
	StateCountdown GameState = "countdown"
	StateInRound   GameState = "in_round"
	StateRoundEnd  GameState = "round_end"
	StateGameOver  GameState = "game_over"
)

func NewGameRoom(mode string, maxPlayers int) *GameRoom {
	return &GameRoom{
		Id:                rand.Text()[2:8],
		GameMode:          mode,
		MaxAllowedPlayers: maxPlayers,
	}
}

func (r *GameRoom) AddPlayerToTable(p *Player) {
	if len(r.ActivePlayers()) >= r.MaxAllowedPlayers {
		p.MarkSpectator()
	}
	seatId := r.nextSeatId
	p.SeatId = seatId
	r.Players = append(r.Players, p)
	r.nextSeatId++
}

func (r *GameRoom) ActivePlayers() []*Player {
	var active []*Player
	for _, p := range r.Players {
		if p.State != Spectator {
			active = append(active, p)
		}
	}
	return active
}

func (r *GameRoom) AdvanceTurn() {
	for _, p := range r.ActivePlayers() {
		if p.SeatId > r.CurrentTurnSeat {
			r.CurrentTurnSeat = p.SeatId
			return
		}
	}

	r.CurrentTurnSeat = r.ActivePlayers()[0].SeatId
}

func (r *GameRoom) GetRoomstats() string {
	stats := "ID: %s\nMode: %s\nCapacity: %d/%d\nState: %s\nSubject: %s\nLastPlayedSize: %d"
	return fmt.Sprintf(stats, r.Id, r.GameMode, len(r.Players), r.MaxAllowedPlayers, r.State, r.TableSubject, r.LastPlayedSize)
}
