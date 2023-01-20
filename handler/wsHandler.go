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
	pool := websocket.NewPool()
	go pool.Start()

	engine.GET("/socket", func(ctx *gin.Context) {
		handler.wsHandler(pool, ctx)
	})
}

func (wsHandler *wsHandler) wsHandler(pool *websocket.Pool, ctx *gin.Context) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(ctx.Writer, ctx.Request)
	if err != nil {
		fmt.Fprintf(ctx.Writer, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}
