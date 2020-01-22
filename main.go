package main

import (
	"log"

	"github.com/Lim79Plus/go-aws/config"
	"github.com/Lim79Plus/go-aws/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	v := "World!"
	log.Printf("Hello %v", v)
	config.Init()
	// service.S3Access()

	router()
}

func router() {
	e := echo.New()
	e.Static("/", "view/form.html")
	e.POST("/upload", handler.UploadPicture())

	e.Logger.Fatal(e.Start(":1234"))
}
