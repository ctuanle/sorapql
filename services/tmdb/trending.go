package tmdb

import (
	"errors"
	"fmt"

	"github.com/ctuanle/sorapql/graph/model"
)

var MediaType = map[string]string{ "all": "", "movie": "", "tv": "" }
var TimeWindow = map[string]string{ "day": "", "week": "" }

func Trending(mediaType string, timeWindow string, page *int) (*model.Trending, error) {
	if _, ok := MediaType[mediaType]; !ok {
		return nil, errors.New("Media type must be 'all', 'movie' or 'tv'")
	}
	if _, ok := TimeWindow[timeWindow]; !ok {
		return nil, errors.New("Time window must be 'day' or 'week'")
	}
	
	query := ""
	if page != nil && *page != 0 {
		query = fmt.Sprintf("page=%v", *page)
	}

	data, err := fetcher[*model.Trending](fmt.Sprintf("/trending/%v/%v", mediaType, timeWindow), query)
	if err != nil {
		return nil, errors.New("Something went wrong!")
	}

	return data, nil
}