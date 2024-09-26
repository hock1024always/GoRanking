package dao

import (
	"Ranking/config"
	"Ranking/pkg/logger"
	"gorm.io/driver/mysql" // 引入 MySQL 驱动
	"gorm.io/gorm"         // 引入 Gorm
	"time"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	// 使用 gorm.Open 打开 MySQL 数据库连接
	Db, err = gorm.Open(mysql.Open(config.Mysqldb), &gorm.Config{})
	if err != nil {
		logger.Error(map[string]interface{}{"mysql connect error": err.Error()})
		return // 连接失败，提前返回
	}

	// 获取底层的 sql.DB
	sqlDB, err := Db.DB()
	if err != nil {
		logger.Error(map[string]interface{}{"get DB instance error": err.Error()})
		return // 获取 DB 实例失败，提前返回
	}

	// 配置连接池参数
	sqlDB.SetMaxIdleConns(10)           // 设置最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 设置最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置连接的最大可重用时长
}
