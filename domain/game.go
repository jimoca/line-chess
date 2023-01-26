package domain

import (
	"github.com/notnil/chess"
)

type Game struct {
	ID         int64         `json: "id"`
	ValidMoves []string      `json: "validMoves"`
	OutCome    chess.Outcome `json: "outCome"`
}
type GameService interface {
	Hello() (res map[string]string, err error)
}
