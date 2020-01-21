package main

import (
	"log"
	"net/http"

	"github.com/Lim79Plus/go-aws/config"
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
	// e.POST("/upload", handler.UploadPicture)
	e.Static("/", "view/form.html")
	e.POST("/upload", UploadPicture)

	e.Logger.Fatal(e.Start(":1234"))
}

// UploadPicture tt
func UploadPicture(c echo.Context) error {
	return c.String(http.StatusOK, "UploadPicture hello!!")
}
