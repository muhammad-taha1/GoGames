// This package is used as a generic web server for games
// all apis will be provided by other game related packages
package main

import (
	"fmt"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func runGame(w http.ResponseWriter, r *http.Request) {
	game, ok := r.URL.Query()["g"]
	selectedGame := game[0]
	if ok {
		if len(selectedGame) == 0 {
			panic("Please enter a game")
		}

		switch selectedGame {
		case "ticTacToe":
			w.Write([]byte("TIC TAC TOE SEL"))
		}
		w.Write([]byte(fmt.Sprint("selected game %s\n", game)))
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/game", runGame)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
