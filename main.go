package main

import (
	"github.com/adeven/go_demo/gogo"
	"github.com/adeven/goenv"
	"log"
)

func main() {
	log.Println("starting demo app")
	gogo.Gogo(goenv.DefaultGoenv())
	select {}
}
