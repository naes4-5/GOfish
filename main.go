package main

import (
	"fmt"
	"errors"
	"math/rand/v2"
	"log"
)

var suits [4]string = [4]string{"spades", "clubs", "hearts", "diamonds"}

type card struct {
	suit string;
	rank int;
}

type deck struct {
	cards      map[string][]card;
	// next_card  card
	cards_left int;
}

type player struct {
	hand  []card;
	books int;
}

func (d deck)draw_random_card() (card, error) {
	if d.cards_left == 0 {
		return card {}, errors.New("Deck empty: no more cards to draw")
	}
	rsuit := suits[rand.IntN(3)]
	for len(d.cards[rsuit]) < 0 {
		rsuit = suits[rand.IntN(3)]
	}
	rrank := rand.IntN(len(d.cards[rsuit]) - 1)
	ret := d.cards[rsuit][rrank]
	d.cards_left--
	copy(d.cards[rsuit][rrank:], d.cards[rsuit][rrank+1:])
	d.cards[rsuit] = d.cards[rsuit][:len(d.cards[rsuit])-1]
	return ret, nil
}

func get_deck() deck {
	var ret deck
	ret.cards_left = 52
	for _, suit := range suits {
		next := []card{}
		for i := 0; i < 13; i++ {
			next = append(next, card {suit: suit, rank: i})
		}
		ret.cards[suit] = next
	}
	return ret
}

func main() {
	deck1 := get_deck()
	first_card, err := deck1.draw_random_card()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d of %s", first_card.rank, first_card.suit)
}
