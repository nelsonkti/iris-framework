package Middleware

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12/context"
)

var Logined context.Handler

type UserInfo struct {
	ID       int64
	Username string
}

func initUserInfo() {
	Logined = func(ctx context.Context) {
		jwtInfo := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
		id := int64(jwtInfo["userId"].(float64))
		username := jwtInfo["userName"].(string)
		UserInfo := UserInfo{
			ID:       id,
			Username: username,
		}
		ctx.Values().Set("UserInfo", UserInfo)
		ctx.Next()
	}
}
