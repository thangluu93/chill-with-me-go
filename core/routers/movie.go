package routers

import (
	"github.com/labstack/echo/v4"
	"main/business"
	"main/core"
	"main/data"
	"net/http"
)

type Movie struct {
	Router *echo.Group
}

func NewMovie(server *core.Server, route string) (err error, router *Movie) {
	router = &Movie{
		Router: server.Echo.Group(route),
	}

	db := server.MongoClient.Database(server.DBName)
	movieBusiness := business.NewMovie(db)

	router.Router.GET(data.MOVIE_LIST_PATH, func(c echo.Context) error {
		ctx := core.ToContextV2(&c)
		page := ctx.QueryParam("page")
		noRecord := ctx.QueryParam("noRecord")
		genre := ctx.QueryParam("genre")

		list, err := movieBusiness.MovieList(page, noRecord, genre)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, list)
	})

	return nil, router
}
