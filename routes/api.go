package routes

import (
	"lineChess/controllers"
)

func init() {
	unAuthApi := engine.Group("api")
	{
		unAuthApi.GET("/hello", controllers.Hello)
	}
}