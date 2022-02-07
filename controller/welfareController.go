package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"test-gin/models"
	"test-gin/pkg/app"
	"test-gin/pkg/e"
	"test-gin/pkg/logging"
)

func List(c *gin.Context) {
	appG := app.Gin{C: c}
	//c.String(http.StatusOK, "上传成功！")
	appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
	return
}

func Abc(c *gin.Context) {
	logging.Debug("哈哈哈哈哈")
	var welfare models.WelfareModel
    id:=c.DefaultQuery("id","")
	models.Db.Find(&welfare,"id=?", id).Row()
	//models.NewDb().Find(&welfare).Row()
	fmt.Printf("===%+v",welfare)
    c.JSON(http.StatusOK,welfare)
	//c.String(http.StatusOK, "abc")
}
