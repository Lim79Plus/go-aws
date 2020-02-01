package handler

import (
	"io"
	"log"
	"net/http"
	"os"

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

		// Destination os上にアップロードされたファイルと同名の空ファイルを作成する
		dst, err := os.Create(file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		return c.String(http.StatusOK, "UploadPicture successed:"+file.Filename)
	}
}
