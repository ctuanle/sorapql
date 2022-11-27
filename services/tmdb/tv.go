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
	var hasCredit, hasVideo, hasImage, hasRecommendation, hasSimilar bool
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
			query += "videos"
		}
	}

	data, err := fetcher[*model.TVDetail](fmt.Sprintf("/tv/%v", id), query)
	if err != nil {
		return nil, errors.New("Something went wrong!")
	}

	return data, nil
}