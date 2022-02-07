package main

import (
	"fmt"
	"log"
	"net/http"
	"test-gin/models"
	"test-gin/pkg/logging"
	"test-gin/pkg/setting"
	"test-gin/routers"
	"test-gin/routers/admin"
	"test-gin/routers/app"
)
func init() {
	setting.Setup()
	logging.Setup()
	models.Setup()
}
func main() {
	// 加载多个APP的路由配置
	routers.Include(admin.Routers,app.Routers)
	// 初始化路由
	r := routers.Init()
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20 //1MB
	server := &http.Server{
		Addr:           endPoint,
		Handler:        r,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)
	server.ListenAndServe()
}


