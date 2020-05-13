package main

import (
	"fmt"
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

func Shuffle(gs GameState) GameState {
	ret := clone(gs)
	ret.Deck = deck.New(deck.Deck(3), deck.Shuffle)
	return ret
}

func Deal(gs GameState) GameState {
	ret := clone(gs)
	ret.PlayerHand = make(Hand, 0, 5)
	ret.DealerHand = make(Hand, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, ret.Deck = ret.Deck[0], ret.Deck[1:]
		ret.PlayerHand = append(ret.PlayerHand, card)
		card, ret.Deck = ret.Deck[0], ret.Deck[1:]
		ret.DealerHand = append(ret.DealerHand, card)
	}

	ret.Stage = StatePlayerTurn
	return ret
}

func Hit(gs GameState) GameState {
	ret := clone(gs)
	hand := ret.CurrentPlayer()
	var card deck.Card
	card, ret.Deck = ret.Deck[0], ret.Deck[1:]
	*hand = append(*hand, card)
	if hand.Score() > 21 {
		return Stand(ret)
	}
	return ret
}

func Stand(gs GameState) GameState {
	ret := clone(gs)
	ret.Stage++
	return ret
}

func EndHand(gs GameState) GameState {
	ret := clone(gs)
	pScore, dScore := ret.PlayerHand.Score(), ret.DealerHand.Score()
	fmt.Println("=========FINAL HANDS=========")
	fmt.Println("Player: ", ret.PlayerHand, "\nScore: ", pScore)
	fmt.Println("Dealer: ", ret.DealerHand, "\nScore: ", dScore)

	switch {
	case pScore > 21:
		fmt.Println("You busted!")
		fmt.Println("You lose!")
	case dScore > 21:
		fmt.Println("Dealer busted!")
		fmt.Println("You win!")
	case pScore > dScore:
		fmt.Println("You win!")
	case dScore > pScore:
		fmt.Println("Dealer wins!")
		fmt.Println("You lose!")
	case dScore == pScore:
		fmt.Println("Draw!")
		fmt.Println("You lose!")
	}
	fmt.Println()

	ret.PlayerHand = nil
	ret.DealerHand = nil
	return ret
}

func main() {
	var gs GameState
	gs = Shuffle(gs)
	gs = Deal(gs)

	var input string
	for gs.Stage == StatePlayerTurn {
		fmt.Println("Player: ", gs.PlayerHand)
		fmt.Println("Dealer: ", gs.DealerHand.DealerString())
		fmt.Println("What will you do? (h)it or (s)tand")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			gs = Hit(gs)
		case "s":
			gs = Stand(gs)
		default:
			fmt.Println("Invalid Option: ", input)
		}
	}

	for gs.Stage == StateDealerTurn {
		// If dealer score <= 16, hit
		// If dealer has a soft 17, hit
		if gs.DealerHand.Score() <= 16 || (gs.DealerHand.Score() == 17 && gs.DealerHand.MinScore() != 17) {
			gs = Hit(gs)
		} else {
			gs = Stand(gs)
		}
	}

	gs = EndHand(gs)
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
	copy(ret.DealerHand, gs.DealerHand)

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
