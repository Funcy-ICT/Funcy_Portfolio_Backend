package pkg

import (
	"backend/pkg/controller"
	"backend/pkg/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	Server *gin.Engine
)

func init() {

	Server = gin.Default()
	Server.Use(cors.Default())
	//認証関連
	Server.GET("helth", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "server ok",
		})
	})

	Server.POST("/sign/up", controller.SignUp())
	Server.POST("/sign/in", controller.SignIn())

	//ユーザ関連

	//作品関連
	work := Server.Group("/work")
	{
		work.POST("", middleware.Authenticate(controller.CreateWork()))
		work.GET("/:id", middleware.Authenticate(controller.ReadWork()))
	}
	works := Server.Group("/works")
	{
		works.GET("/:number", middleware.Authenticate(controller.ReadWorksList()))
	}

	//グループ関連

	//検索関連

}
