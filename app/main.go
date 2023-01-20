package main

import (
	"lineChess/conf"
	chessController "lineChess/handler"
	gameService "lineChess/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/notnil/chess"
)

func main() {
	engine := gin.Default()
	config := conf.Conf()
	var game *chess.Game
	gs := gameService.NewGameService(game)
	chessController.NewChessController(engine, gs)
	log.Fatal(engine.Run(config.URL))
}
