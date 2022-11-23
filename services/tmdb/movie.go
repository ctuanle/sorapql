package tmdb

import (
	"github.com/ctuanle/sorapql/graph/model"
)

func MoviePopular() (*model.MoviePopular, error) {
	return fetcher[*model.MoviePopular]("/movie/popular")
}