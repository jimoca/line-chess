package service

import (
	"lineChess/domain"
)

type gameService struct {
}

func NewGameService() domain.GameService {
	return &gameService{}
}

func (g *gameService) Hello() (res map[string]string, err error) {
	res = map[string]string{"message": "hello"}
	return
}
