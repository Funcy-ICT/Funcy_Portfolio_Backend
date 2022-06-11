package middleware

import (
	"backend/pkg/model/dao"
	"backend/pkg/view"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	AuthenticationToken = "SELECT id FROM `users` WHERE `token`=? ;"
)

// Authenticate ユーザ認証を行ってContextへユーザID情報を保存する
func Authenticate(ginNextMethod gin.HandlerFunc) gin.HandlerFunc {
	var userID string

	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			log.Println("[ERROR] token is empty")
			view.ReturnErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"token is empty",
			)
		}
		//dbにtokenが存在するか
		row := dao.Conn.QueryRow(AuthenticationToken, token)
		if err := row.Scan(&userID); err != nil {
			if err == sql.ErrNoRows {
				view.ReturnErrorResponse(
					c,
					http.StatusBadRequest,
					"Bad Request",
					"token is inactive",
				)
				return
			}
			log.Println(err)
		}

		log.Println(userID)

		// ユーザIDをContextへ保存して以降の処理に利用する
		c.Set(userID, token)
		ginNextMethod(c)
	}
}
