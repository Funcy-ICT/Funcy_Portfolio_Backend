package controller

import (
	"backend/pkg/model/dao"
	"backend/pkg/model/dto"
	"backend/pkg/view"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateWork() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("userID")
		if userID == "" {
			log.Println("[ERROR] userID is empty")
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"InternalServerError",
				"userID is empty",
			)
			return
		}
		var cwr dto.CreateWorkRequest
		if err := c.BindJSON(&cwr); err != nil {
			view.ReturnErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"RequestBody is empty",
			)
			return
		}

		client := dao.MakeCreateWorkClient()
		_, err := client.Request(userID, cwr)
		if err != nil {
			log.Println(err)
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				err.Error(),
			)
			return
		}

		c.JSON(http.StatusOK, view.ReturnCreateWork(cwr))
	}
}

func ReadWork() gin.HandlerFunc {
	return func(c *gin.Context) {
		workID := c.Param("id")
		if workID == "" {
			log.Println("[ERROR] workID is empty")
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"InternalServerError",
				"workID is empty",
			)
			return
		}
		client := dao.MakeReadWorkClient()
		workInfo, err := client.Request(workID)
		if err != nil {
			log.Println(err)
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				err.Error(),
			)
			return
		}

		c.JSON(http.StatusOK, workInfo)
	}
}

func ReadWorksList() gin.HandlerFunc {
	return func(c *gin.Context) {
		number := c.Param("number")

		if number == "" {
			log.Println("[ERROR] number is empty")
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"InternalServerError",
				"number is empty",
			)
			return
		}
		client := dao.MakeReadWorksListClient()
		workInfo, err := client.Request(number)
		if err != nil {
			log.Println(err)
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				err.Error(),
			)
			return
		}

		c.JSON(http.StatusOK, workInfo)
	}
}
