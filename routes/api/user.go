package api

import (
	"IrisFramework/app/Http/Controllers/Api"
	"github.com/kataras/iris/v12/core/router"
)

func RouteUser(api router.Party) {

	api.PartyFunc("/v1", func(v1 router.Party) {
		v1.Post("/login", Api.Login)
	})
	
}
