package router

import (
	"github.com/adayswait/mojo/handler"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/websocket"
	"net/http"
)

func Route(app *fiber.App) {
	app.Static("/", "./dist/")
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

	visitorWeb := web.Group("/visitor")
	visitorWeb.Get("/breakdep", handler.BreakDep)

	//身份认证
	authWeb := web.Group("/auth")
	authWeb.Get("/login", handler.Login)
	authWeb.Get("/logout", handler.Logout)
	authWeb.Get("/register", handler.Register)

	depWeb := web.Group("/dep")
	depWeb.Get("/test", handler.Test)
	depWeb.Get("/submit", handler.SubmitDep)
	depWeb.Get("/commithistory", handler.CommitHistory)
	depWeb.Get("/progress", handler.ProgressList)
	depWeb.Delete("/", handler.DeleteDep)

	//db工具
	dbWeb := web.Group("/db")
	dbWeb.Get("/:table?/:key?", handler.ViewDB)
	dbWeb.Put("/:table?/:key?", handler.UpdateDB)
	dbWeb.Post("/:table?/:key?", handler.NewDB)
	dbWeb.Delete("/:table?/:key?", handler.DeleteDB)

	chatWeb := web.Group("/chat")
	chatWeb.Post("/:to", handler.Chat)
}
