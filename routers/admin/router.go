package admin

import (
	"github.com/gin-gonic/gin"
	"test-gin/controller"
	"test-gin/middlewares"
)

func Routers(e *gin.Engine) {
	v := e.Group("/admin").Use(middlewares.Test())
	{
		v.GET("/list", controller.List)
	}
}
