package main

import "fmt"
import "math/rand"
import "bytes"

type Suit struct {
    Short_desc rune
    Long_desc string
    id int
}

var CLUBS = Suit{'♣', "clubs", 0}
var HEARTS = Suit{'♥', "hearts", 1}
var DIAMONDS = Suit{'♦', "diamonds", 2}
var SPADES = Suit{'♠', "spades", 3}

var SUITES = [4]*Suit{&CLUBS, &HEARTS, &DIAMONDS, &SPADES}

type CardValue struct {
    Short_desc rune
    Long_desc string
    id int
}

var ACE = CardValue{'A', "ace", 14}
var TWO = CardValue{'2', "two", 2}
var THREE = CardValue{'3', "three", 3}
var FOUR = CardValue{'4', "four", 4}
var FIVE = CardValue{'5', "five", 5}
var SIX = CardValue{'6', "six", 6}
var SEVEN = CardValue{'7', "seven", 7}
var EIGHT = CardValue{'8', "eight", 8}
var NINE = CardValue{'9', "nine", 9}
var TEN = CardValue{'T', "ten", 10}
var JACK = CardValue{'J', "jack", 11}
var QUEEN = CardValue{'Q', "queen", 12}
var KING = CardValue{'K', "king", 13}

var CARD_VALUES = [13]*CardValue{
    &ACE, &TWO, &THREE, &FOUR, &FIVE, &SIX, &SEVEN, &EIGHT, &NINE, &TEN,
    &JACK, &QUEEN, &KING,
}

type Card struct {
    Value *CardValue
    Suit *Suit
}

func (c Card) Long_desc() string {
    return fmt.Sprintf("%v of %v", c.Value.Long_desc, c.Suit.Long_desc)
}

func (c Card) Short_desc() string {
    return fmt.Sprintf(
        "%v%v", string(c.Value.Short_desc), string(c.Suit.Short_desc),
    )
}

type deck struct {
    Cards [52]*Card
    size int
}

func NewDeck() *deck {
    deck := new(deck)
    deck.size = len(deck.Cards)
    for i, suit := range SUITES {
        for j, card_value := range CARD_VALUES {
            deck.Cards[i*len(CARD_VALUES) + j] = &Card{card_value, suit}
        }
    }
    return deck
}

func (d *deck) Pop() *Card {
    index_chosen := 0
    if d.size > 0 {
        index_chosen = rand.Intn(d.size)
    }
    card := d.Cards[index_chosen]
    index_last := d.size - 1
    d.Cards[index_chosen] = d.Cards[index_last]
    d.Cards[index_last] = nil
    d.size -= 1
    return card
}

func (d deck) ToString() string {
    var buffer bytes.Buffer

    buffer.WriteString("[")
    for _, card := range d.Cards {
        var to_write = "nil"
        if card != nil {
            to_write = card.Short_desc()
        }
        buffer.WriteString(fmt.Sprintf("%v, ", to_write))
    }
    buffer.WriteString("]")
    return buffer.String()
}

func main() {
    rand.Seed(42)
    // Build a deck
    var a_deck = NewDeck()
    for i := 1; i <= 52; i++ {
        card := a_deck.Pop()
        fmt.Printf("Popped %v\n", card.Short_desc())
        fmt.Println(a_deck.ToString())
    }

}
