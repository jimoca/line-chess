package service

import (
	"bytes"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/notnil/chess"
)

type gameSocket struct {
	MoveStr chan string
}

func NewGameSocket() *gameSocket {
	return &gameSocket{
		MoveStr: make(chan string),
	}
}

func (g *gameSocket) NewGame() {

	game := chess.NewGame()
	for game.Outcome() == chess.NoOutcome {
		hi := <-g.MoveStr
		fmt.Printf("adsjflksjlei: %s", hi)
		// moves := game.ValidMoves()
		// response(ws, &err, &map[string]string{"valid moves": returnValidMove(&moves)})
		// msgType, msg, err := ws.ReadMessage()
		// messsge := string(msg)
		// fmt.Printf("Message Type: %d, Message: %s\n", msgType, messsge)
		// if err != nil {
		// 	log.Println(err)
		// 	break
		// }
		// if checkValidMove(&moves, &messsge) {
		// 	game.MoveStr(messsge)
		// } else {
		// 	response(ws, &err, &map[string]string{"waring": "invalid move !"})
		// }
	}

	return
}

func (g *gameSocket) Resign(color chess.Color) {

}
func (g *gameSocket) Draw(offer chess.Method) {

}

func response(ws *websocket.Conn, err *error, msg *map[string]string) {
	*err = ws.WriteJSON(*msg)
	if err != nil {
		log.Println(err)
	}
}

func returnValidMove(moves *[]*chess.Move) string {
	var buffer bytes.Buffer
	for _, move := range *moves {
		buffer.WriteString(move.String() + " ")
	}
	return string(buffer.String())
}

func checkValidMove(moves *[]*chess.Move, receiveMove *string) bool {
	for _, move := range *moves {
		if move.String() == *receiveMove {
			return true
		}
	}

	return false
}
