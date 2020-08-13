package config

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"os"
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
	db, err := gorm.Open("mysql", "homestead:secret@(192.168.10.10:3306)/homestead?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.SetLogger(&GormLogger{})
	db.LogMode(true)
	db.SingularTable(true)

	// test
	Gorm = db

	//var userModel Models.User
	//Gorm.Where("username like ?","%风%").Find(&userModel)
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
