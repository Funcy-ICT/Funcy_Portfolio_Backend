package controller

import (
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

type UploadResponse struct {
	URLs []string `json:"urls"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func UploadImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		var urls []string
		form, err := c.MultipartForm()
		if err != nil {
			log.Printf("[ERROR] Failed to parse multipart form: %v", err)
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Error:   "Upload Failed",
				Message: "リクエストの形式が正しくありません",
			})
			return
		}

		files := form.File["file"]
		if len(files) == 0 {
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Error:   "Upload Failed",
				Message: "アップロードするファイルが選択されていません",
			})
			return
		}

		for _, file := range files {
			uuID, err := uuid.NewRandom()
			if err != nil {
				log.Printf("[ERROR] Failed to generate UUID: %v", err)
				c.JSON(http.StatusInternalServerError, ErrorResponse{
					Error:   "Upload Failed",
					Message: "サーバー内部エラーが発生しました。しばらく時間をおいて再度お試しください",
				})
				return
			}

			fileName := fmt.Sprintf("uploadimages/%s%s", uuID, file.Filename)

			err = c.SaveUploadedFile(file, fileName)
			if err != nil {
				log.Printf("[ERROR] Failed to save uploaded file: %v", err)
				c.JSON(http.StatusInternalServerError, ErrorResponse{
					Error:   "Upload Failed",
					Message: "ファイルの保存に失敗しました。しばらく時間をおいて再度お試しください",
				})
				return
			}

			if fileName == "" {
				log.Println("[ERROR] Generated filename is empty")
				c.JSON(http.StatusInternalServerError, ErrorResponse{
					Error:   "Upload Failed",
					Message: "ファイル名の生成に失敗しました。しばらく時間をおいて再度お試しください",
				})
				return
			}

			urlName := fmt.Sprintf("http://localhost:3004/%s%s", uuID, file.Filename)
			urls = append(urls, urlName)
		}

		c.JSON(http.StatusOK, UploadResponse{URLs: urls})
	}
}
