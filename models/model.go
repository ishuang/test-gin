// 数据库初始化，建立连接
package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"test-gin/pkg/logging"
	"test-gin/pkg/setting"
	"time"
)
type Model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
var Db *gorm.DB

// Setup initializes the database instance
func Setup() {
	db, err:= gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name)))

	if err != nil {
		logging.Fatal("models.Setup err: %v", err)
		panic(err)
	}
	////设置连接池
	////空闲
	//sqlDB, err := db.DB()
	//if err != nil {
	//	logging.Fatal("get sqlDB err:%v", err)
	//}
	//
	//sqlDB.SetMaxIdleConns(mysqlConfig.MaxIdleConn)
	////打开
	//sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpenConn)
	////超时
	//sqlDB.SetConnMaxLifetime(time.Second * time.Duration(mysqlConfig.ConnMaxLifeTime))

	Db = db
}
