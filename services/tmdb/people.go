package tmdb

import (
	"errors"
	"fmt"

	"github.com/ctuanle/sorapql/graph/model"
)

func PeoplePopular(language *string, page *int) (*model.PopularPeople, error) {
	query := ""
	if language != nil && len(*language) != 0 {
		query = fmt.Sprintf("language=%v&", *language)
	}
	if page != nil && *page != 0 {
		query = fmt.Sprintf("%vpage=%v", query, *page)
	}

	data, err := fetcher[*model.PopularPeople]("/person/popular", query)

	if err != nil {
		return nil, errors.New("Something went wrong!")
	}

	return data, nil
}

func PersonDetail(id int, fields []string, language *string) (*model.PersonDetail, error) {
	query := ""
	if language != nil && len(*language) != 0 {
		query = "language=" + *language + "&"
	}
	var hasImage, hasExternalId, hasCredit bool
	for _, v := range fields {
		if v == "images" {
			hasImage = true
		} else if v == "external_ids" {
			hasExternalId = true
		} else if v == "combined_credits" {
			hasCredit = true
		}
	}

	if hasImage || hasCredit || hasExternalId {
		query += "append_to_response="
		if hasImage {
			query += "images,"
		}
		if hasCredit {
			query += "combined_credits,"
		}
		if hasExternalId {
			query += "external_ids"
		}
	}

	data, err := fetcher[*model.PersonDetail](fmt.Sprintf("/person/%v", id), query)

	if err != nil {
		return nil, errors.New("Something went wrong")
	}

	return data, nil
}