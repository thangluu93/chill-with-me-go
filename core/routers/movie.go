package routers

import (
	"github.com/labstack/echo/v4"
	"main/core"
	"main/data"
)

type Movie struct {
	Router *echo.Group
}

func NewMovie(server *core.Server, route string) (err error, router *Movie) {
	router = &Movie{
		Router: server.Echo.Group(route),
	}

	router.Router.GET(data.MOVIE_LIST_PATH, func(c echo.Context) error {
		ctx := core.ToContextV2(&c)
		page := ctx.QueryParam("page")
		noRecord := ctx.QueryParam("noRecord")
		genre := ctx.QueryParam("genre")
		return nil
	})

	return nil, router
}
