package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// UploadPicture 画像のアップロード
func UploadPicture() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "UploadPicture hello!!")
	}
}
