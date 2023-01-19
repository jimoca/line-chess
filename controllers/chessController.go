package controllers

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"log"
	"github.com/notnil/chess"
	"bytes"
)


func Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello!",
	})
}

func NewGame(ctx *gin.Context) {
	
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
	var game *chess.Game;
	for {		
		msgType, msg, err := ws.ReadMessage()
		fmt.Printf("Message Type: %d, Message: %s\n", msgType, string(msg))
		if err != nil {
			log.Println(err)
			break;
		}
		if(string(msg) == "newGame") {
			game = chess.NewGame()
			for game.Outcome() == chess.NoOutcome {
				moves := game.ValidMoves()
				response(ws, &err, &map[string]string{"valid moves": returnValidMove(&moves)})
				msgType, msg, err := ws.ReadMessage()
				messsge := string(msg)
				fmt.Printf("Message Type: %d, Message: %s\n", msgType, messsge)
				if err != nil {
					log.Println(err)
					break;
				}
				if(checkValidMove(&moves, &messsge)) {
					game.MoveStr(messsge)
				}
			}
		}
	}
	
}

func response(ws *websocket.Conn,err *error, msg *map[string]string) {
	*err = ws.WriteJSON(*msg)
	if err != nil {
		log.Println(err)
	}
}

func returnValidMove(moves *[]*chess.Move) string{
	var buffer bytes.Buffer
	for _, move := range *moves {
		buffer.WriteString(move.String() + " ")
	}
	return string(buffer.String())
}


func checkValidMove(moves *[]*chess.Move, receiveMove *string) bool{
		for _, move := range *moves {
			if move.String() == *receiveMove {
				return true
			}
		}
	
		return false
}


