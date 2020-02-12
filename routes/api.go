package routes

import (
	"IrisFramework/app/Http/Middleware"
	api2 "IrisFramework/routes/api"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

func Register(api *iris.Application) {
	apiPrefix := api.Party("/api").AllowMethods(iris.MethodOptions)
	{
		//apiPrefix.Get("/home", func(ctx iris.Context) {
		//	ctx.WriteString("Hello from api")
		//})
		
		api2.RouteUser(apiPrefix)
		
		

		apiPrefix.PartyFunc("/v1", func(v1 router.Party) {
			v1.Get("/", func(ctx iris.Context) {
				ctx.WriteString("hello v1")
			})
			
			v1.Use(Middleware.JWT.Serve, Middleware.Logined)

			v1.Get("/", func(ctx iris.Context) {
				ctx.WriteString("hello v1")
			})

		})

	}

}
