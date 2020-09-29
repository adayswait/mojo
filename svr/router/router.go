package router

import (
	"github.com/adayswait/mojo/handler"
	"github.com/adayswait/mojo/utils"
	"github.com/adayswait/mojo/ws"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/websocket/v2"
	"net/http"
)

func Route(app *fiber.App) {
	app.Static("/", utils.Get3wDir())
	app.Use("/ws", ws.New())
	app.Get("/ws", websocket.New(handler.Websocket))
	app.Use("/fs", filesystem.New(filesystem.Config{
		Root:   http.Dir(utils.GetFsDir()),
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
		AllowOrigins:     utils.GetWebDomain(),
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
	depWeb.Put("/create", handler.CreateDep)
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

	splanWeb := web.Group("/splan")
	splanWeb.Post("/mail", handler.SplanMail)
	splanWeb.Put("/changetime", handler.SplanChangeTime)
	splanWeb.Get("/timelocker", handler.SplanLockTime)
	splanWeb.Put("/timelocker", handler.SplanLockTime)
	splanWeb.Delete("/timelocker", handler.SplanLockTime)
	splanWeb.Put("/config/:module", handler.SplanUpdateConfig)
}
