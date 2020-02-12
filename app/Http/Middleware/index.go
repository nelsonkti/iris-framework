package Middleware

import (
	"log"
	"sync"
)

var once sync.Once

func init() {
	log.Print("【中间件】 初始化...\n")
	once.Do(func() {
		initJWT()
		initCors()
		initUserInfo()
	})
}