package config

import (
	"IrisFramework/app/Models"
	"errors"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
	"sync"
	"xorm.io/core"
	"xorm.io/xorm"
)

var once sync.Once
var DB *xorm.Engine

func init() {
	once.Do(func() {
		dbDriver := os.Getenv("DB_CONNECTION")
		switch dbDriver {
		case "mysql":
			connectMysql()
		default:
			panic(errors.New("only supper mysql"))
		}
	})

	initTable()
	defaultConfig()
}

func connectMysql() {
	dbDriver := os.Getenv("DB_CONNECTION")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	dbParams := os.Getenv("DB_PARAMS")
	dbUser := os.Getenv("DB_USERNAME")
	dbPwd := os.Getenv("DB_PASSWORD")
	dbURL := fmt.Sprintf("%s:%s@(%s:%s)/%s?%s", dbUser, dbPwd, dbHost, dbPort, dbName, dbParams)
	var err error

	DB, err = xorm.NewEngine(dbDriver, dbURL)

	if err != nil {
		fmt.Printf("connection mysql failed, err:%v\n", err)
		log.Printf("connection mysql failed, err:%v\n", err)
		panic(err)
	}

}

func initTable()  {
	// 自动创建表
	err := DB.Sync2(new(Models.User))
	if err != nil {
		log.Printf("迁移数据结构失败:%v\n", err)
		panic(err)
	}
}

func defaultConfig()  {

	// 设置日志等级，设置显示sql，设置显示执行时间
	if os.Getenv("APP_ENV") != "production" {
		DB.SetLogger(xorm.NewSimpleLogger(logFile()))
		DB.SetLogLevel(xorm.DEFAULT_LOG_LEVEL)
		DB.ShowSQL(true)
		DB.ShowExecTime(true)
	}

	// 转换大小写
	DB.SetMapper(core.GonicMapper{})
}
