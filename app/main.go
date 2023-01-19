package main

import (
	"lineChess/routes"
	"log"
)
func main() {
	engine := routes.New()
	log.Fatal(engine.Run())
}
