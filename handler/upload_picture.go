package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// UploadPicture 画像のアップロード
func UploadPicture(c echo.Context) error {
	return c.String(http.StatusOK, "UploadPicture hello!!")
}
