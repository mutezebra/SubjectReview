package client

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/mutezebra/subject-review/config"
	"github.com/mutezebra/subject-review/pkg/logger"
)

var _db *gorm.DB

func initMysql() {
	conf := config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=true&loc=Local", conf.UserName, conf.Password, conf.Address, conf.Database, conf.Charset)

	db, err := gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
	if err != nil {
		logger.Fatal(err)
	}

	DB, _ := db.DB()
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(20)
	DB.SetConnMaxLifetime(5 * time.Minute)

	_db = db
}

func GetMysqlDB(table ...string) *gorm.DB {
	if _db == nil {
		initMysql()
	}

	if table != nil {
		return _db.Table(table[0]).WithContext(context.Background())
	}
	return _db.WithContext(context.Background())
}
