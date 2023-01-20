package domain

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/notnil/chess"
)

type Game struct {
	ID         int64         `json: "id"`
	ValidMoves []string      `json: "validMoves"`
	OutCome    chess.Outcome `json: "outCome"`
}

type GameService interface {
	Hello() (res map[string]string, err error)
	NewGame()
	MoveStr(color chess.Color, str string)
	Resign(color chess.Color)
	Draw(offer chess.Method)
}

type Websocket interface {
	Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error)
	Reader(conn *websocket.Conn)
	Writer(conn *websocket.Conn)
}
