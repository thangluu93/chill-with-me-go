package routers

import (
	"github.com/labstack/echo/v4"
	"main/core"
	"main/data"
)

type User struct {
	Router *echo.Group
}

func NewUser(server *core.Server, route string) (err error, router *User) {
	router = &User{
		Router: server.Echo.Group(route),
	}

	router.Router.POST(data.LoginPath, func(c echo.Context) error {
		//ctx := core.ToContextV2(&c)
		return nil
	})

	return nil, router
}
