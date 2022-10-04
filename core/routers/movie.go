package routers

import (
	"github.com/labstack/echo/v4"
	"main/business"
	"main/core"
	"main/data"
	"main/models"
	"net/http"
)

type Movie struct {
	Router        *echo.Group
	MovieBusiness *business.Movie
}

func (m *Movie) movieListCtrl(c echo.Context) error {
	ctx := core.ToContextV2(&c)
	page := ctx.QueryParam("page")
	noRecord := ctx.QueryParam("noRecord")
	genre := ctx.QueryParam("genre")

	list, err := m.MovieBusiness.MovieList(page, noRecord, genre)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, list)
}

func (m *Movie) movieCreateCtrl(c echo.Context) error {
	ctx := core.ToContextV2(&c)
	movie := new(models.Movie)
	if err := ctx.Bind(movie); err != nil {
		return err
	}
	movie.IsAchieve = false
	newMovie, err := m.MovieBusiness.CreateMovie(movie)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, newMovie)
}

func (m *Movie) movieUpdateCtrl(c echo.Context) error {
	ctx := core.ToContextV2(&c)
	movie := new(models.Movie)
	if err := ctx.Bind(movie); err != nil {
		return err
	}
	movie.IsAchieve = false
	newMovie, err := m.MovieBusiness.UpdateMovie(movie)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, newMovie)
}

func (m *Movie) movieDeleteCtrl(c echo.Context) error {
	ctx := core.ToContextV2(&c)
	movie := new(models.Movie)
	if err := ctx.Bind(movie); err != nil {
		return err
	}
	movie.IsAchieve = false
	success, err := m.MovieBusiness.DeleteMovie(movie)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, success)
}

func NewMovie(server *core.Server, route string) (err error, router *Movie) {
	db := server.MongoClient.Database(server.DBName)
	router = &Movie{
		Router:        server.Echo.Group(route),
		MovieBusiness: business.NewMovie(db),
	}

	router.Router.GET(data.MovieListPath, router.movieListCtrl)
	router.Router.POST(data.MovieCreatePath, router.movieCreateCtrl)
	router.Router.PUT(data.MovieUpdatePath, router.movieUpdateCtrl)
	router.Router.DELETE(data.MovieDeletePath, router.movieDeleteCtrl)

	return nil, router
}
