package router

import (
	"github.com/Lim79Plus/go-aws/handler"
	"github.com/labstack/echo/v4"
)

// Initialize router
func Initialize() {
	e := echo.New()
	e.Static("/", "view/form.html")
	e.POST("/upload", handler.UploadPicture(nil))

	e.Logger.Fatal(e.Start(":1234"))
}
