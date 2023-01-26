package main

import (
	"lineChess/conf"
	"lineChess/handler"
	gameService "lineChess/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	config := conf.Conf()
	gs := gameService.NewGameService()
	handler.ApiHandler(engine, gs)
	handler.WsHandler(engine)
	log.Fatal(engine.Run(config.URL))
}
