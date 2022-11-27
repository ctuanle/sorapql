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

func MovieTopRated(page *int, language *string, region *string) (*model.MovieList, error) {
	query := queryListBuilder(page, language, region)

	data, err := fetcher[*model.MovieList]("/movie/top_rated", query)
	if err != nil {
		return nil, errors.New("Something went wrong!")
	}

	return data, nil
}

func MovieUpcoming(page *int, language *string, region *string) (*model.MovieList, error) {
	query := queryListBuilder(page, language, region)

	data, err := fetcher[*model.MovieList]("/movie/upcoming", query)
	if err != nil {
		return nil, errors.New("Something went wrong!")
	}
	
	return data, nil
}

func MovieDetail(id int, fields []string, language *string) (*model.MovieDetail, error) {
	query := ""
	if language != nil && len(*language) != 0 {
		query = "language=" + *language
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
		query += "&append_to_response="
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

	data, err := fetcher[*model.MovieDetail](fmt.Sprintf("/movie/%v", id), query)
	if err != nil {
		return nil, errors.New("Something went wrong!")
	}

	return data, nil
}