package holdem

import "fmt"
import "math/rand"
import "pokersml/cards"

type actionType struct {
    id int
    desc string
}

var CALL = &actionType{0, "call"}
var RAISE = &actionType{1, "raise"}
var FOLD = &actionType{0, "fold"}

type action struct {
    actionType *actionType
    value int
}

type bot interface {
    do() action
}

type exampleBot struct {}
func (b exampleBot) do() action {
    action := action{FOLD, 0}
    return action
}
var EXAMPLEBOT = &exampleBot{}

type Player struct {
    Name string
    Bot bot
}

type GameSettings struct {
    Speed int
}

type game struct {
    players *[]Player
    settings *GameSettings
    deck *cards.Deck
}

func StartGame(players *[]Player, settings *GameSettings) *game {
    game := game{players, settings, cards.NewDeck()}
    rand.Seed(42)
    fmt.Println("Game started!")
    return &game
}
