package game

func LegalActions(r *GameRoom, p *Player) []Action {

	switch r.State {
	case StateLobby:
		return []Action{{ActionType: "ready"}, {ActionType: "leave"}}
	case StateCountdown:
		return []Action{}
	case StateInRound:
		if p.IsSpectator {
			return []Action{{ActionType: "leave"}}
		}
		if r.CurrentTurnSeat != p.SeatId {
			return []Action{{ActionType: "leave"}}
		}
		var legal []Action

		// TODO : implement legal actions

		return legal

	case StateRoundEnd:
		return []Action{{ActionType: "leave"}}
	case StateGameOver:
		return []Action{{ActionType: "leave"}}
	default:
		return []Action{}
	}

}

func ApplyAction(r *GameRoom, p *Player, a Action) error {
	return nil // TODO : apply rules to player actions
}
