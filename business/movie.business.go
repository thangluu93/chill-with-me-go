package business

import (
	"main/access"
	"main/core"
)

type Movie struct {
	MovieAccess *access.Movie
}

func MovieList(page int, noRecord int, genre string) {
	util := core.UseUtil()
	util.getLimitOffset(page, noRecord)
}
