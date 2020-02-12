package Api

import (
	"IrisFramework/app/Exceptions"
	"IrisFramework/app/Http/Controllers"
	"IrisFramework/app/Http/Middleware"
	"IrisFramework/app/Http/Requests"
	"IrisFramework/app/Http/Transformers"
	"errors"
	"fmt"
	"github.com/kataras/iris/v12"
	"log"
)

func Login(ctx iris.Context)  {
	req := Requests.PostLogin{}
	ctx.ReadJSON(&req)
	fmt.Print(req.Username)


	user, err := Controllers.UserRepository.QueryByUsername(req.Username)

	if err != nil {
		ctx.JSON(Exceptions.ErrorQueryDatabase(err))
		return
	}

	log.Println(user, req)
	// If passwd are inconsistent
	if user.Passwd != req.Passwd {
		ctx.JSON(Exceptions.ErrorVerification(errors.New("用户名或密码错误")))
		return
	}

	// Login Ok
	// Get token
	token, err := Middleware.GetJWTString(user.Username, user.ID)
	if err != nil {
		ctx.JSON(Exceptions.ErrorBuildJWT(err))
	}

	res := Transformers.PostLogin{
		Username: user.Username,
		ID:       user.ID,
		Token:    token,
	}
	ctx.JSON(res)

}
