package pkg

import (
	"backend/file-server/pkg/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var (
	//Server gin flameworkのserver
	Server *gin.Engine
)

func init() {

	Server = gin.Default()
	//allows all origins
	Server.Use(cors.Default())

	Server.Use(static.Serve("/", static.LocalFile("./uploadimages", true)))
	Server.POST("/upload/file", controller.UploadImage())
}
