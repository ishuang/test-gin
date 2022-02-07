package app

import (
	"github.com/gin-gonic/gin"
	"test-gin/controller"
)

func Routers(e *gin.Engine) {
	e.GET("/abc", controller.Abc)
}
