package main

import (
	"fmt"
	"math/rand"
	// "github.com/gin-gonic/gin"
	"github.com/notnil/chess"
)
func main() {

	game := chess.NewGame()
	// generate moves until game is over
	for game.Outcome() == chess.NoOutcome {
		// select a random move
		moves := game.ValidMoves()
		move := moves[rand.Intn(len(moves))]
		game.Move(move)
	}
	// print outcome and game PGN
	fmt.Println(game.Position().Board().Draw())
	fmt.Printf("Game completed. %s by %s.\n", game.Outcome(), game.Method())
	fmt.Println(game.String())

	// app := gin.Default()
	// app.GET("/hello/:name", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	c.JSON(200, gin.H{
	// 		"message": "hello " + name,
	// 	})
	// })
	// err := app.Run(":8080")
	// if err != nil {
	// 	panic(err)
	// }
}