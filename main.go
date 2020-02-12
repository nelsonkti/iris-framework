package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "IrisFramework/config"
	"IrisFramework/app/Http/Middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kataras/iris/v12"
	"IrisFramework/routes"
	"os"
)

func main() {
	newApp()
}

func newApp() {
	app := iris.New()

	app.Use(Middleware.CORS)
	routes.Register(app)

	app.HandleDir("/upload", "public/upload") // 设置静态资源

	app.Run(
		iris.Addr(":" + os.Getenv("APP_PORT")),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)

}
