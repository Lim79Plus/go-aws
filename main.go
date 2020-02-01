package main

import (
	"github.com/Lim79Plus/go-aws/config"
	"github.com/Lim79Plus/go-aws/router"
)

func main() {
	config.Initialize()
	router.Initialize()
}
