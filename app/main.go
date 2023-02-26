package main

import (
	"lineChess/conf"
	"lineChess/handler"
	gameService "lineChess/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func main() {
	config := conf.Conf()
	engine := gin.Default()

	redisDB := redis.NewClient(&redis.Options{
		Addr:     config.DB_HOST_PORT,
		Password: config.DB_PASSWORD,
		DB:       0,
	})
	gs := gameService.NewGameService()
	handler.ApiHandler(engine, gs)
	handler.WsHandler(engine, redisDB)
	log.Fatal(engine.Run(config.URL))
}
