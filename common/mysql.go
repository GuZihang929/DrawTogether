package common

import (
	"database/sql"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	Orm *gorm.DB
)

func GormMysql() {
	sqlDb, err := sql.Open("mysql", Config.Server.Mysql.Url)
	if err != nil {
		zap.L().Fatal("数据库连接出错：", zap.Error(err))
	}
	db, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDb}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		zap.L().Fatal("数据库连接出错：", zap.Error(err))
	} else {
		zap.L().Info("数据库连接成功")
	}
	Orm = db
}
