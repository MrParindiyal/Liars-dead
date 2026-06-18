package game

// import "crypto/rand"

type GameRoom struct {
	Id                string
	Players           []*Player
	GameMode          string
	MaxAllowedPlayers int
	PlayedStack       []Card
	nextSeatId        int
}

func NewGameRoom(mode string, maxPlayers int) *GameRoom {
	return &GameRoom{
		// Id: rand.Text()[2:8],  // TODO : how to assign room id?
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
		if !p.IsSpectator {
			active = append(active, p)
		}
	}
	return active
}
