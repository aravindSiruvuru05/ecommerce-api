package routers

import (
	"fmt"
	"haste/adapters/controllers"
	"haste/adapters/controllers/user"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func InitRoutes() {
	web.ErrorController(&controllers.ErrorController{})

	ns := web.NewNamespace("/api/v1",
		web.NSNamespace("/users",
			web.NSRouter(
				"/", &user.UserController{}, "get:GetAllUsers",
			),
			web.NSRouter(
				"/:id", &user.UserController{}, "get:GetById",
			),
		),
		web.NSGet("/healthcheck", func(ctx *context.Context) {
			_ = ctx.Output.Body([]byte("i am alive"))
		}),
	)
	web.AddNamespace(ns)
}

func init() {
	fmt.Println("router.go__________ router")
}
