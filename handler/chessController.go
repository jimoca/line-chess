package handler

import (
	"bytes"
	"fmt"
	"lineChess/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/notnil/chess"
)

type ResponseError struct {
	Message string `json:"message"`
}

type chessHandler struct {
	gameService domain.GameService
}

func NewChessHandler(engine *gin.Engine, gs domain.GameService) {
	controller := &chessHandler{
		gameService: gs,
	}

	socket := engine.Group("/socket")
	socket.GET("", controller.wsHandler)
	api := engine.Group("/api")
	api.GET("/hello", controller.Hello)
}

func (c *chessHandler) Hello(ctx *gin.Context) {
	res, err := c.gameService.Hello()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Internal Server Error",
		})
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *chessHandler) wsHandler(ctx *gin.Context) {
	upGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	ws, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
	}

	defer func() {
		closeSocketErr := ws.Close()
		if closeSocketErr != nil {
			log.Println(err)
		}
	}()
	for {
		msgType, msg, err := ws.ReadMessage()
		fmt.Printf("Message Type: %d, Message: %s\n", msgType, string(msg))
		if err != nil {
			log.Println(err)
			break
		}
		if string(msg) == "newGame" {
			game := chess.NewGame()
			for game.Outcome() == chess.NoOutcome {
				moves := game.ValidMoves()
				response(ws, &err, &map[string]string{"valid moves": returnValidMove(&moves)})
				msgType, msg, err := ws.ReadMessage()
				messsge := string(msg)
				fmt.Printf("Message Type: %d, Message: %s\n", msgType, messsge)
				if err != nil {
					log.Println(err)
					break
				}
				if checkValidMove(&moves, &messsge) {
					game.MoveStr(messsge)
				} else {
					response(ws, &err, &map[string]string{"waring": "invalid move !"})
				}
			}
		}
	}

}

func response(ws *websocket.Conn, err *error, msg *map[string]string) {
	*err = ws.WriteJSON(*msg)
	if err != nil {
		log.Println(err)
	}
}

func returnValidMove(moves *[]*chess.Move) string {
	var buffer bytes.Buffer
	for _, move := range *moves {
		buffer.WriteString(move.String() + " ")
	}
	return string(buffer.String())
}

func checkValidMove(moves *[]*chess.Move, receiveMove *string) bool {
	for _, move := range *moves {
		if move.String() == *receiveMove {
			return true
		}
	}

	return false
}
