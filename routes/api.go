package routes

import (
	"lineChess/controllers"
)

func init() {
	socket := engine.Group("/socket")
	{
		socket.GET("", controllers.NewGame)
	}
	unAuthApi := engine.Group("api")
	{
		unAuthApi.GET("/hello", controllers.Hello)
	}
}