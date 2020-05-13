package main

import (
	"strings"

	"github.com/Tak1za/go-deck"
)

type Stage int8

const (
	StatePlayerTurn Stage = iota
	StateDealerTurn
	StateHandOver
)

type GameState struct {
	Deck       []deck.Card
	Stage      Stage
	PlayerHand Hand
	DealerHand Hand
}

//Hand type to handle a player's or dealer's hands
type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}
	for _, c := range h {
		if c.Rank == deck.Ace {
			return minScore + 10
		}
	}
	return minScore
}

func (h Hand) MinScore() int {
	score := 0
	for _, c := range h {
		score += min(int(c.Rank), 10)
	}

	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//DealerString method to print dealer's hand in hidden format
func (h Hand) DealerString() string {
	return h[0].String() + ", **HIDDEN**"
}

func main() {
	// cards := deck.New(deck.Deck(3), deck.Shuffle)
	// var card deck.Card
	// var player, dealer Hand
	// for i := 0; i < 2; i++ {
	// 	for _, hand := range []*Hand{&player, &dealer} {
	// 		card, cards = cards[0], cards[1:]
	// 		*hand = append(*hand, card)
	// 	}
	// }

	// var input string
	// for input != "s" {
	// 	fmt.Println("Player: ", player)
	// 	fmt.Println("Dealer: ", dealer.DealerString())
	// 	fmt.Println("What will you do? (h)it or (s)tand")
	// 	fmt.Scanf("%s\n", &input)
	// 	switch input {
	// 	case "h":
	// 		card, cards = cards[0], cards[1:]
	// 		player = append(player, card)
	// 	}
	// }

	// //If dealer score <= 16, hit
	// //If dealer has a soft 17, hit
	// for dealer.Score() <= 16 || (dealer.Score() == 17 && dealer.MinScore() != 17) {
	// 	card, cards = cards[0], cards[1:]
	// 	dealer = append(dealer, card)
	// }

	// pScore, dScore := player.Score(), dealer.Score()
	// fmt.Println("=========FINAL HANDS=========")
	// fmt.Println("Player: ", player, "\nScore: ", pScore)
	// fmt.Println("Dealer: ", dealer, "\nScore: ", dScore)

	// switch {
	// case pScore > 21:
	// 	fmt.Println("You busted!")
	// case dScore > 21:
	// 	fmt.Println("Dealer busted!")
	// case pScore > dScore:
	// 	fmt.Println("You win!")
	// case dScore > pScore:
	// 	fmt.Println("Dealer wins!")
	// case dScore == pScore:
	// 	fmt.Println("Draw!")
	// }
}

func clone(gs GameState) GameState {
	ret := GameState{
		Deck:       make([]deck.Card, len(gs.Deck)),
		Stage:      gs.Stage,
		PlayerHand: make(Hand, len(gs.PlayerHand)),
		DealerHand: make(Hand, len(gs.DealerHand)),
	}

	copy(ret.Deck, gs.Deck)
	copy(ret.PlayerHand, gs.PlayerHand)
	copy(ret.DealerHand, gs.PlayerHand)

	return ret
}

func (gs *GameState) CurrentPlayer() *Hand {
	switch gs.Stage {
	case StatePlayerTurn:
		return &gs.PlayerHand
	case StateDealerTurn:
		return &gs.DealerHand
	default:
		panic("It isn't currently any player's turn")
	}
}
