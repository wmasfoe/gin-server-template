package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"go-blog-server/global"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.CONFIG.System.Env)
	r := gin.Default()
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	publicGroup := r.Group("api")
	publicGroup.GET("/", func(c *gin.Context) {
		c.String(200, "NB")
	})

	return r
}
