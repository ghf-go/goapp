package utils

import (
	"fmt"
	"time"

	"github.com/ghf-go/goapp/apppackage/app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 全局数据库连接
var DB *gorm.DB

// InitDB 按配置初始化 gorm 连接并自动迁移表结构
func InitDB() {
	c := Conf.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		c.User, c.Password, c.Host, c.Port, c.Database, c.Charset, c.ParseTime, c.Loc)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败: " + err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("获取数据库实例失败: " + err.Error())
	}
	setPool(sqlDB)
	if err := db.AutoMigrate(&models.ProjectModel{}, &models.BuildRecordModel{}); err != nil {
		panic("自动迁移表结构失败: " + err.Error())
	}
	DB = db
	Log("数据库初始化完成 %s:%d/%s", c.Host, c.Port, c.Database)
}

func setPool(sqlDB interface {
	SetMaxOpenConns(int)
	SetMaxIdleConns(int)
	SetConnMaxLifetime(time.Duration)
	SetConnMaxIdleTime(time.Duration)
}) {
	c := Conf.DB
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	if d, err := time.ParseDuration(c.ConnMaxLifetime); err == nil {
		sqlDB.SetConnMaxLifetime(d)
	}
	if d, err := time.ParseDuration(c.ConnMaxIdleTime); err == nil {
		sqlDB.SetConnMaxIdleTime(d)
	}
}
