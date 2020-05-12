package main

import (
	"fmt"
	"strings"

	"github.com/Tak1za/go-deck"
)

//Hand exported
type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

//DealerString exported
func (h Hand) DealerString() string {
	return h[0].String() + ", **HIDDEN**"
}

func main() {
	cards := deck.New(deck.Deck(3), deck.Shuffle)
	var card deck.Card
	var player, dealer Hand
	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&player, &dealer} {
			card, cards = cards[0], cards[1:]
			*hand = append(*hand, card)
		}
	}

	var input string
	for input != "s" {
		fmt.Println("Player: ", player)
		fmt.Println("Dealer: ", dealer.DealerString())
		fmt.Println("What will you do? (h)it or (s)tand")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			card, cards = cards[0], cards[1:]
			player = append(player, card)

		}
	}

	fmt.Println("=========FINAL HANDS=========")
	fmt.Println("Player: ", player)
	fmt.Println("Dealer: ", dealer)
}
