package business

import (
	"go.mongodb.org/mongo-driver/mongo"
	"main/access"
	"main/core"
	"main/data"
	"main/models"
)

type Movie struct {
	MovieAccess *access.Movie
	Util        core.Utility
}

func NewMovie(db *mongo.Database) *Movie {
	var movie = access.NewMovie(db, "movies")
	return &Movie{
		MovieAccess: (*access.Movie)(movie),
		Util:        *core.UseUtil(),
	}
}

func (m *Movie) MovieList(page string, noRecord string, genre string) (movies []*models.Movie, err error) {

	pageInt, errorParse := m.Util.StringToInt(page)
	if errorParse != nil {
		pageInt = 1
	}

	noRecordInt, errorParse := m.Util.StringToInt(noRecord)
	if errorParse != nil {
		noRecordInt = data.DEFAULT_PAGE_SIZE
	}

	limit, offset := m.Util.GetLimitOffset(pageInt, noRecordInt)
	movies, err = m.MovieAccess.GetListMovies(limit, offset, genre)
	
	if err != nil {
		return nil, err
	}
	return movies, nil
}
