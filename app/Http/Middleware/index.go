package Middleware

import (
	"sync"
)

var once sync.Once

func init() {
	once.Do(func() {
		initJWT()
		initCors()
		initUserInfo()
	})
}