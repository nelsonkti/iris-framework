package config

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"os"
	"time"
)

type GormLogger struct{}
var Gorm *gorm.DB

func initDB() {
	once.Do(func() {
		dbDriver := os.Getenv("DB_CONNECTION")
		switch dbDriver {
		case "mysql":
			initMysql()
			break
		default:
			panic(errors.New("only supper mysql"))
			break
		}
	})

}

/**
  初始化mysql
 */
func initMysql() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),os.Getenv("DB_DATABASE"), os.Getenv("DB_PARAMS"))

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("failed to connect database")
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	sqlDB.SetMaxIdleConns(1024)
	sqlDB.SetMaxOpenConns(1024)
	sqlDB.SetConnMaxLifetime(time.Minute * 10)

	Gorm = db
}


/**
  打印日志
 */
func (*GormLogger) Print(v ...interface{}) {

	switch v[0] {
	case "sql":
		// [gorm] [sql] SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((username like ?)) []interface {}{"%萧风%"}  [1.880301ms]
		msg := fmt.Sprintf("[%s] [%s] [%s] %s %#v ", "gorm", v[0], v[2], v[3], v[4])
		Log.Info(msg)
	case "log":
		log.WithFields(log.Fields{"module": "gorm", "type": "log"}).Print(v[2])
	}
}
