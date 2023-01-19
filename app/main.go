package main

import (
	"lineChess/routes"
	"lineChess/conf"
	"log"
)
func main() {
	config := conf.Conf()
	engine := routes.New()
	log.Fatal(engine.Run(config.URL))
}
