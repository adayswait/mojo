package router

import "net/http"
import "github.com/gofiber/cors"
import "github.com/gofiber/fiber"
import "github.com/gofiber/websocket"
import "github.com/gofiber/fiber/middleware"
import "github.com/adayswait/mojo/handler"

func Route(app *fiber.App) {
	// app.Static("/", "../cli/")
	app.Get("/ws", websocket.New(handler.Websocket))
	app.Use("/fs", middleware.FileSystem(middleware.FileSystemConfig{
		Root: http.Dir("/opt/jesse/git"),
		// Index:  "index.html",
		Browse: true,
	}))

	//服务间调用, 不使用cookie
	api := app.Group("/api")
	api.Use(cors.New())
	authApi := api.Group("/auth")

	authApi.Get("/login", handler.Login)
	authApi.Get("/logout", handler.Logout)
	authApi.Get("/register", handler.Register)

	//用于web页面展示的请求, 允许cookie等
	web := app.Group("/web")
	web.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://10.1.1.248:8080"},
		AllowCredentials: true,
	}))
	web.Get("/svnlog", handler.SvnLog)
	authWeb := web.Group("/auth")
	authWeb.Get("/login", handler.Login)
	authWeb.Get("/logout", handler.Logout)
	authWeb.Get("/register", handler.Register)

	//db工具
	dbWeb := web.Group("/db")
	dbWeb.Get("/:table?/:key?", handler.ViewDB)
	dbWeb.Put("/:table?/:key?", handler.UpdateDB)
	dbWeb.Post("/:table?/:key?", handler.NewDB)
	dbWeb.Delete("/:table?/:key?", handler.DeleteDB)

	rsyncWeb := web.Group("/rsync")
	rsyncWeb.Get("/", handler.Rsync)
}
