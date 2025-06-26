package controller

import (
	"backend/file-server/pkg/view"
	"fmt"
	"log"
	"net/http"
	"os"

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

func DeleteImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		fileName := c.Param("filename")

		if fileName == "" {
			view.ReturnErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"filename is required",
			)
			return
		}

		// uploadimagesディレクトリから削除
		filePath := fmt.Sprintf("uploadimages/%s", fileName)

		err := os.Remove(filePath)
		if err != nil {
			log.Printf("[ERROR] Failed to delete file: %s, error: %v\n", filePath, err)
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				"Failed to delete file",
			)
			return
		}

		log.Printf("[INFO] Successfully deleted file: %s\n", filePath)
		c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
	}
}
