package controller

import (
	"backend/pkg/auth"
	"backend/pkg/model/dao"
	"backend/pkg/model/dto"
	"backend/pkg/view"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {

		var sur dto.SignUpRequest
		if err := c.BindJSON(&sur); err != nil {
			view.ReturnErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"RequestBody is empty",
			)
			return
		}

		client := dao.MakeSignUpClient()
		token, err := client.Request(sur)
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

		c.JSON(http.StatusOK, view.ReturnSignResponse(token))
	}
}

func SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var sir dto.SignInRequest
		if err := c.BindJSON(&sir); err != nil {
			view.ReturnErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"RequestBody is empty",
			)
			return
		}
		client := dao.MakeSignInClient()
		token, err := client.Request(sir)
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
		c.JSON(http.StatusOK, view.ReturnSignResponse(token))
	}
}

type loginBody struct {
	Id int `json:"id"`
}

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {

		var sur loginBody
		if err := c.BindJSON(&sur); err != nil {
			view.ReturnErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"RequestBody is empty",
			)
			return
		}

		jws, err := auth.IssueUserToken(int64(sur.Id))
		if err != nil {
			log.Print(err)
			view.ReturnErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"RequestBody is empty",
			)
			return
		}

		log.Println(jws)

		c.JSON(http.StatusOK, "jwt")
	}
}
