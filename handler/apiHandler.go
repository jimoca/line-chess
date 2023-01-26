package handler

import (
	"lineChess/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Message string `json:"message"`
}

type apiHandler struct {
	gameService domain.GameService
}

func ApiHandler(engine *gin.Engine, gs domain.GameService) {
	controller := &apiHandler{
		gameService: gs,
	}
	api := engine.Group("/api")
	api.GET("/hello", controller.Hello)
}

func (c *apiHandler) Hello(ctx *gin.Context) {
	res, err := c.gameService.Hello()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Internal Server Error",
		})
	}
	ctx.JSON(http.StatusOK, res)
}
