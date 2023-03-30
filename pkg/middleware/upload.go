package middleware

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

func UploadFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		file, err := ctx.FormFile("image")
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, err)
		}

		src, err := file.Open()
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, err)
		}
		defer src.Close()

		tempFile, err := ioutil.TempFile("uploads", "image-*.png")
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, err)
		}
		defer tempFile.Close()

		if _, err = io.Copy(tempFile, src); err != nil {
			return ctx.JSON(http.StatusBadRequest, err)
		}

		data := tempFile.Name()

		ctx.Set("dataFile", data)
		return next(ctx)
	}
}
