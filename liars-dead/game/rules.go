package game

import "strconv"

func LegalActions(r *GameRoom, p *Player) []Action {

	switch r.State {
	case StateLobby:
		return []Action{{ActionType: "ready"}, {ActionType: "leave"}, {ActionType: "unready"}}
	case StateCountdown:
		return []Action{{ActionType: "unready"}}
	case StateInRound:
		if p.State == Spectator {
			return []Action{{ActionType: "leave"}}
		}
		if r.CurrentTurnSeat != p.SeatId {
			if r.LastPlayedBy == nil || r.LastPlayedBy.SeatId != p.SeatId {
				return []Action{{ActionType: "leave"}}
			}
		}
		var legal []Action
		if r.LastPlayedSize > 0 {
			legal = append(legal, Action{ActionType: "check"})
		}
		legal = append(legal, Action{ActionType: "play"})
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

	switch a.ActionType {
	case "leave":
		if r.State == StateLobby {
			var remainingPlayers []*Player
			for _, player := range r.Players {
				if player.PlayerId != p.PlayerId {
					remainingPlayers = append(remainingPlayers, player)
				}
			}
			r.Players = remainingPlayers
		}
		p.State = Spectator

	case "ready":
		p.State = LobbyReady

	case "unready":
		p.State = LobbyWaiting

	case "check":
		victim := r.LastPlayedBy
		victimHandSize := r.LastPlayedSize

		if r.GameMode == "bluff" {
			for _, card := range r.PlayedStack[len(r.PlayedStack)-victimHandSize : len(r.PlayedStack)] {
				if card.Value != "Joker" && card.Value != r.TableSubject {
					victim.Hand = append(victim.Hand, r.PlayedStack...) // let the last player collect entire stack
					r.PlayedStack = nil
					return nil
				}
			}
			// last hand was valid, player collects all
			p.Hand = append(p.Hand, r.PlayedStack...)
			r.PlayedStack = nil

		} else {
			for _, card := range r.PlayedStack[len(r.PlayedStack)-victimHandSize : len(r.PlayedStack)] {
				{
					if card.Value != r.TableSubject && !card.IsSpecial {
						victim.DeductLife() //TODO : modify with probability later
					} else if card.Value == "Devil" {
						for _, target := range r.Players {
							if target.PlayerId != victim.PlayerId {
								target.DeductLife()
							}
						}
					} else if card.Value == "Master" {
						targetSeatId := a.Payload["targetSeatId"].(int)
						if victim.SeatId != targetSeatId {
							for _, player := range r.Players {
								if player.SeatId == targetSeatId {
									player.DeductLife()
								}
							}
						}
					} else if card.Value == "Chaos" {
						for _, player := range r.Players {
							playerTarget := a.Payload[strconv.Itoa(player.SeatId)].(int)
							for _, target := range r.Players {
								if target.SeatId == playerTarget {
									target.DeductLife()
									break
								}
							}
						}
					}
				}
			}
		}
	case "play":
		indices := a.Payload["indices"].([]int)
		played := p.RemovePlayedCards(indices)
		r.PlayedStack = append(r.PlayedStack, played...) //unpacks 'played' and appends one at a time
		r.LastPlayedSize = len(played)
		r.LastPlayedBy = p
	}

	return nil
}
