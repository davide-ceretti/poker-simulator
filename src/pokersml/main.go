package main

import "fmt"
import "pokersml/holdem"


func main() {
    fmt.Println("Running main.go")
    davide := holdem.Player{"Davide", holdem.EXAMPLEBOT}
    luke := holdem.Player{"Luke", holdem.EXAMPLEBOT}

    var players = &[]holdem.Player {davide, luke}
    settings := &holdem.GameSettings{1}
    holdem.StartGame(players, settings)
}
