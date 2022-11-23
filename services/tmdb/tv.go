package tmdb

import (
	"errors"
	"fmt"

	"github.com/ctuanle/sorapql/graph/model"
)

func queryListTvBuilder(page *int, language *string) string {
	query := ""
	if page != nil && *page != 0 && *page != 1 {
		query = fmt.Sprintf("%vpage=%v&", query, *page)
	}
	if language != nil && len(*language) != 0 {
		query = fmt.Sprintf("%vlanguage=%v&", query, *language)
	}
	return query
}

func TVPopular(page *int, language *string) (*model.TVList, error) {
	query := queryListTvBuilder(page, language)

	data, err := fetcher[*model.TVList]("/tv/popular", query)
	if err != nil {
		return nil, errors.New("Something went wrong!")
	}

	return data, nil
}

func TVTopRated(page *int, language *string) (*model.TVList, error) {
	query := queryListTvBuilder(page, language)

	data, err := fetcher[*model.TVList]("/tv/top_rated", query)
	if err != nil {
		return nil, errors.New("Something went wrong!")
	}

	return data, nil
}

func TVOnTheAir(page *int, language *string) (*model.TVList, error) {
	query := queryListTvBuilder(page, language)

	data, err := fetcher[*model.TVList]("/tv/on_the_air", query)
	if err != nil {
		return nil, errors.New("Something went wrong!")
	}
	
	return data, nil
}