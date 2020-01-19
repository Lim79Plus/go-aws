package main

import (
	"github.com/Lim79Plus/go-aws/config"
	"github.com/Lim79Plus/go-aws/service"
	"log"
)

func main() {
	v := "World!"
	log.Printf("Hello %v", v)
	config.Init()
	service.S3Access()
}
