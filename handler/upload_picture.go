package handler

import (
	"log"
	"net/http"

	"github.com/Lim79Plus/go-aws/service"
	"github.com/labstack/echo/v4"
)

// UploadPicture 画像のアップロード
func UploadPicture() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Source
		file, err := c.FormFile("uploadfile")
		if err != nil {
			return err
		}
		log.Printf("Uploaded file name %v", file.Filename)

		// アップロードされたファイルを開ける
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		service.UplocadFileToS3Bucket(src, file.Filename)

		return c.String(http.StatusOK, "UploadPicture successed:"+file.Filename)
	}
}
