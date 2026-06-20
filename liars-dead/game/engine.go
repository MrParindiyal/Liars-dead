package game

func RunGameLoop(r *GameRoom){
	for r.State != StateGameOver{
		currentPlayer := r.Players[r.CurrentTurnSeat]
		
		// TODO : fix logic, add player count check
		legal := LegalActions(r, currentPlayer)
		ApplyAction(r, currentPlayer, legal[0])
		r.AdvanceTurn()
	}
}