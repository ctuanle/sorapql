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

func TVDetail(id int, fields []string, language *string) (*model.TVDetail, error) {
	query := ""
	if language != nil && len(*language) != 0 {
		query += "language=" + *language + "&"
	}
	// TODO: if fields contains only one of these, should call specific endpoint
	var hasCredit, hasVideo, hasImage, hasRecommendation, hasSimilar, hasEx, hasImdb bool
	for _, field := range fields {
		if field == "credits" {
			hasCredit = true
		} else if field == "videos" {
			hasVideo = true
		} else if field == "images" {
			hasImage = true
		} else if field == "recommendations" {
			hasRecommendation = true
		} else if field == "similar" {
			hasSimilar = true
		} else if field == "external_ids" {
			hasEx = true
		} else if field == "imdb_rating" {
			hasImdb = true
		}
	}

	if hasCredit || hasImage || hasRecommendation || hasSimilar || hasVideo {
		query += "append_to_response="
		if hasCredit {
			query += "credits,"
		}
		if hasImage {
			query += "images,"
		}
		if hasRecommendation {
			query += "recommendations,"
		}
		if hasSimilar {
			query += "similar,"
		}
		if hasVideo {
			query += "videos,"
		}
		if hasEx {
			query += "external_ids"
		}
	}

	data, err := fetcher[*model.TVDetail](fmt.Sprintf("/tv/%v", id), query)
	if err != nil {
		return nil, errors.New("Something went wrong!")
	}

	if hasImdb {
		var imdb_id string
		if data.ExternalIds != nil && len(*data.ExternalIds.ImdbID) != 0 {
			imdb_id = *data.ExternalIds.ImdbID
		} else {
			res, _ := fetcher[*model.MovieExternalIds](fmt.Sprintf("/tv/%v/external_ids", id), "")
			if res.ImdbID != nil && len(*res.ImdbID) != 0 {
				imdb_id = *res.ImdbID
			}
		}

		if len(imdb_id) != 0 {
			rating, _ := imdb_rate(imdb_id)
			data.ImdbRating = rating
		}
	}

	return data, nil
}