package controller

import (
	"backend/file-server/pkg/view"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Image struct {
	Image string `json:"image"`
}

type UploadRequest struct {
	Images []Image `json:"images"`
}

func UploadImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		var urls []string
		form, _ := c.MultipartForm()
		files := form.File["file"]

		if len(files) == 0 {
			view.ReturnErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"not exist file",
			)
		}

		for _, file := range files {

			uuID, err := uuid.NewRandom()
			if err != nil {
				log.Println("uuid generate is failed")
			}

			fileName := fmt.Sprintf("uploadimages/%s%s", uuID, file.Filename)

			err = c.SaveUploadedFile(file, fileName)
			if err != nil {
				log.Println("[ERROR] Faild Bind JSON　\n ", err)
				c.JSON(http.StatusBadRequest, "Request is error")
				view.ReturnErrorResponse(
					c,
					http.StatusBadRequest,
					"Bad Request",
					"Request is error",
				)
				return
			}
			if fileName == "" {
				view.ReturnErrorResponse(
					c,
					http.StatusBadRequest,
					"Bad Request",
					"file name is null",
				)
			}
			urlName := fmt.Sprintf("http://localhost:3004/%s%s", uuID, file.Filename)
			urls = append(urls, urlName)
		}

		c.JSON(http.StatusOK, view.UploadResponse(urls))
	}

}
