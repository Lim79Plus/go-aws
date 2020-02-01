package main

import (
	"log"

	"github.com/Lim79Plus/go-aws/config"
	"github.com/Lim79Plus/go-aws/router"
)

func main() {
	v := "World!"
	log.Printf("Hello %v", v)
	config.Initialize()
	router.Initialize()
}
