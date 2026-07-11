package game

import "math/rand/v2"

type Card struct {
	Value     string
	IsSpecial bool
}

type Deck struct {
	Cards []Card
	Mode  string
}

func CreateDeck(mode string) Deck {
	cardTypes := []string{"King", "Queen", "Ace", "Joker", "Devil", "Master", "Chaos"} // for reference

	var StartingDeck Deck
	StartingDeck.Mode = mode

	if mode == "bluff" {
		StartingDeck.Cards = make([]Card, 0, 52)
		standardCardTypes := []string{"King", "Queen", "Ace", "Joker", "10", "9", "8", "7", "6", "5", "4", "3", "2"}
		for _, val := range standardCardTypes {
			for range 4 {
				StartingDeck.Cards = append(StartingDeck.Cards, Card{Value: val, IsSpecial: false})
			}
		}
		return StartingDeck
	}

	StartingDeck.Cards = make([]Card, 0, 21) // TODO : make this dynamic
	for _, val := range cardTypes[:2] {
		for range 5 { // alternate for i := 0; i < 5; i++ since i was never used.
			StartingDeck.Cards = append(StartingDeck.Cards, Card{Value: val, IsSpecial: false})
		}
	}

	if mode == "basic" {
		StartingDeck.Cards = append(StartingDeck.Cards, Card{Value: "King", IsSpecial: false})
		StartingDeck.Cards = append(StartingDeck.Cards, Card{Value: "Queen", IsSpecial: false})

		for range 6 {
			StartingDeck.Cards = append(StartingDeck.Cards, Card{Value: "Ace", IsSpecial: false})
		}

		for range 2 {
			StartingDeck.Cards = append(StartingDeck.Cards, Card{Value: "Joker", IsSpecial: false})
		}
	}

	if mode == "devil" {
		baseDeck := CreateDeck("basic")
		baseDeck.Cards = append(baseDeck.Cards, Card{Value: "Devil", IsSpecial: true})
		StartingDeck = baseDeck
	}

	if mode == "chaos" {
		StartingDeck.Cards = append(StartingDeck.Cards, Card{Value: "Master", IsSpecial: true})
		StartingDeck.Cards = append(StartingDeck.Cards, Card{Value: "Chaos", IsSpecial: true})
	}

	return StartingDeck

}

func ShuffleDeck(rawDeck Deck) Deck {
	// Knuth shuffle implementation for shuffling decks

	cardLen := len(rawDeck.Cards)
	for i := range cardLen {
		randindex := rand.IntN(cardLen-i) + i
		rawDeck.Cards[randindex], rawDeck.Cards[i] = rawDeck.Cards[i], rawDeck.Cards[randindex]
	}

	return rawDeck
}

func DealNHands(deck Deck, n int) []Deck {
	// TODO : add checks for n in main

	sizeOfHand := 5
	if deck.Mode == "chaos" {
		sizeOfHand = 3
	}
	out := []Deck{}
	lb := 0

	for range n {
		ub := lb + sizeOfHand
		hand := Deck{}
		hand.Cards = make([]Card, sizeOfHand)
		copy(hand.Cards, deck.Cards[lb:ub])
		out = append(out, hand)
		lb = ub
	}

	return out
}
