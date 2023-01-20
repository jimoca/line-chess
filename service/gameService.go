package service

import (
	"lineChess/domain"

	"github.com/notnil/chess"
)

type gameService struct {
	game *chess.Game
}

func NewGameService(game *chess.Game) domain.GameService {
	return &gameService{
		game: game,
	}
}

func (g *gameService) Hello() (res map[string]string, err error) {
	res = map[string]string{"message": "hello"}
	return
}

func (g *gameService) NewGame() {
	g.game = chess.NewGame()
	return
}

func (g *gameService) MoveStr(color chess.Color, str string) {

}
func (g *gameService) Resign(color chess.Color) {

}
func (g *gameService) Draw(offer chess.Method) {

}
