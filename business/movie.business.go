package business

import (
	"go.mongodb.org/mongo-driver/mongo"
	"main/access"
	"main/core"
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

func (m *Movie) MovieList(page int, noRecord int, genre string) (movies []*models.Movie, err error) {
	limit, offset := m.Util.GetLimitOffset(page, noRecord)
	movies, err = m.MovieAccess.GetListMovies(limit, offset, genre)
	if err != nil {
		return nil, err
	}
	return movies, nil
}
