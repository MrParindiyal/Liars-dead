package game;

type Card struct{
	Value string
	IsSpecial bool
}

type Deck struct{
	Cards []Card
}


func CreateDeck(mode string) Deck {
	cardTypes := []string{"Ace", "King", "Queen", "Joker", "Devil", "Master", "Chaos"}

	var StartingDeck Deck

	if mode == "basic" {
		StartingDeck.Cards = make([]Card, 0, 20) // TODO : make this dynamic 
		for _, val := range cardTypes[:3]{
			for i := 0; i < 6; i++ {
				card := Card{Value: val, IsSpecial: false}
				StartingDeck.Cards = append(StartingDeck.Cards, card)
			}
		}

		for i := 0; i < 2; i++ {
			StartingDeck.Cards = append(StartingDeck.Cards, Card{Value: "Joker", IsSpecial: false})
		}
	}
	return StartingDeck	

}