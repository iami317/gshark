package router

import (
	"github.com/gin-gonic/gin"
	"github.com/madneal/gshark/api"
)

func InitInitRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("init")
	{
		ApiRouter.GET("initdb", api.InitDB)    // 创建Api
		ApiRouter.POST("checkdb", api.CheckDB) // 创建Api
	}
}
