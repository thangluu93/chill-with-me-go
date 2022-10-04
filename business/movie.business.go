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

func (m *Movie) CreateMovie(movie *models.Movie) (newMovie *models.Movie, err error) {
	newMovie, err = m.MovieAccess.CreateMovie(movie)
	if err != nil {
		return nil, err
	}
	return newMovie, nil

}

func (m *Movie) UpdateMovie(movie *models.Movie) (updatedMovie *models.Movie, err error) {
	updatedMovie, err = m.MovieAccess.UpdateMovie(movie)
	if err != nil {
		return nil, err
	}
	return updatedMovie, nil
}

func (m *Movie) DeleteMovie(movie *models.Movie) (success bool, err error) {
	_, err = m.MovieAccess.UpdateMovie(movie)
	if err != nil {
		return false, err
	}
	return true, nil
}
