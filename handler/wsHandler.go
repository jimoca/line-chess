package handler

import (
	"fmt"
	"lineChess/domain"
	"lineChess/pkg/ws"

	"github.com/gin-gonic/gin"
)

type wsHandler struct {
	ws domain.Websocket
}

func NewWsHandler(engine *gin.Engine, dws domain.Websocket) {
	handler := &wsHandler{
		ws: dws,
	}
	pool := ws.NewPool()
	go pool.Start()

	engine.GET("/socket", func(ctx *gin.Context) {
		handler.wsHandler(pool, ctx)
	})
}

func (wsHandler *wsHandler) wsHandler(pool *ws.Pool, ctx *gin.Context) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := ws.Upgrade(ctx.Writer, ctx.Request)
	if err != nil {
		fmt.Fprintf(ctx.Writer, "%+v\n", err)
	}

	client := &ws.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}
