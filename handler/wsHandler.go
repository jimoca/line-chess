package handler

import (
	"fmt"
	"lineChess/domain"
	"lineChess/pkg/websocket"

	"github.com/gin-gonic/gin"
)

type wsHandler struct {
	ws domain.Websocket
}

func NewWsHandler(engine *gin.Engine, ws domain.Websocket) {
	handler := &wsHandler{
		ws: ws,
	}

	socket := engine.Group("/socket")
	socket.GET("", handler.wsHandler)
}

func (wsHandler *wsHandler) wsHandler(ctx *gin.Context) {
	ws, err := websocket.Upgrade(ctx.Writer, ctx.Request)
	if err != nil {
		fmt.Fprintf(ctx.Writer, "%+V\n", err)
	}
	go websocket.Writer(ws)
	websocket.Reader(ws)
}
