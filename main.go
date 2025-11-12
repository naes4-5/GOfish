package main

include (
	"fmt"
	"errors"
	"math/rand/v2"
)

suits := {"spades", "clubs", "hearts", "diamonds"}

type struct card {
	suit string;
	int  rank;
}

type struct deck {
	cards      map[string][]card;
	next_card  card
	cards_left int;
}

type struct player {
	hand  []card;
	books int;
}

func (d deck)draw_random_card() (card, err) {
	if d.cards_left == 0 {
		return nil, errors.New("Deck empty: no more cards to draw")
	}
	rsuit := rand.intN(3)
	for len(d.cards[rsuit]) < 0 {
		rsuit = rand.intN(3)
	}
	rrank := rand.intN(len(d.cards[rsuit]) - 1)
	ret := d.cards[rsuit][rrank]
	d.cardsLeft--
	copy(cards[rsuit][rrank:], cards[rsuit][rrank+1:])
	cards[rsuit][len(cards[rsuit])-1] = nil
	cards[rsuit] = cards[rsuit][:len(cards[rsuit])-1]
	return ret, nil
}

func get_deck() deck {
	var ret deck
	for _, suit := range suits {
		next := make([]card)
		for i := 0; i < 13; i++ {
			next = append(next, card {suit: suit, rank: i})
		}
		ret[suit] = next
	}
	return ret
}

func main() {
}
