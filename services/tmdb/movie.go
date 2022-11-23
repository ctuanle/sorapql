package tmdb

import (
	"errors"
	"fmt"

	"github.com/ctuanle/sorapql/graph/model"
)

func queryListBuilder(page *int, language *string, region *string) string {
	query := ""
	if page != nil && *page != 0 && *page != 1 {
		query = fmt.Sprintf("%vpage=%v&", query, *page)
	}
	if language != nil && len(*language) != 0 {
		query = fmt.Sprintf("%vlanguage=%v&", query, *language)
	}
	if region != nil && len(*region) != 0 {
		query = fmt.Sprintf("%vregion=%v&", query, *language)
	}
	return query
}

func MoviePopular(page *int, language *string, region *string) (*model.MovieList, error) {
	query := queryListBuilder(page, language, region)

	data, err := fetcher[*model.MovieList]("/movie/popular", query)
	if err != nil {
		return nil, errors.New("Something went wrong!")
	}

	return data, nil
}

func TopRated(page *int, language *string, region *string) (*model.MovieList, error) {
	query := queryListBuilder(page, language, region)

	data, err := fetcher[*model.MovieList]("/movie/top_rated", query)
	if err != nil {
		return nil, errors.New("Something went wrong!")
	}

	return data, nil
}

func Upcoming(page *int, language *string, region *string) (*model.MovieList, error) {
	query := queryListBuilder(page, language, region)

	data, err := fetcher[*model.MovieList]("/movie/upcoming", query)
	if err != nil {
		return nil, errors.New("Something went wrong!")
	}
	
	return data, nil
}