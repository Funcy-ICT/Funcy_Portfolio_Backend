package pkg

import (
	"backend/pkg/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	Server *gin.Engine
)

func init() {

	Server = gin.Default()
	Server.Use(cors.Default())
	//認証関連
	Server.POST("/sign/up", controller.SignUp())
	Server.POST("/sign/in", controller.SignIn())

	//ユーザ関連

	//作品関連
	//work := Server.Group("/work"){
	////	work.POST("",controller.)
	//}

	//グループ関連

	//検索関連

}
