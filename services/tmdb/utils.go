package tmdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ctuanle/sorapql/graph/model"
)

// func queryBuilder(params map[string]any) string {
// 	query := ""
// 	for k, v := range params {
// 		query = fmt.Sprintf("%v%v=%v&", query, k, v)
// 	}
// 	return query
// }

func getJson[T any](url string) (T, error) {
	fmt.Println(url)
	var data T
	tmdbClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 5 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("getJson(NewRequest)", url, ":", err)
		return data, err
	}

	res, err := tmdbClient.Do(req)
	if err != nil {
		log.Println("getJson(DoGet)", url, " : ", err)
		return data, err
	}
	if res.StatusCode != 200 {
		log.Println("get(DoGet)", res.Status)
		return data, errors.New("!=200")
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		// log.Fatal(readErr)
		log.Println("getJson(ReadBody)", url, " : ", err)
		return data, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("getJson(Unmarshal)", url, " : ", err)
		return data, err
		// log.Fatal(jsonErr)
	}

	return data, nil
}

func fetcher[T any](path string, query string) (T, error) {
	key := os.Getenv("TMDB_KEY")
	base := os.Getenv("TMDB_URL")

	url := fmt.Sprintf("%v%v?api_key=%v&%v", base, path, key, query)

	return getJson[T](url)
}

type IMDBResponse struct {
	Rating *model.IMDBRating `json:"rating"`
}

func imdb_rate(imdb_id string) (*model.IMDBRating, error) {
	imdb_api_host := os.Getenv("IMDB_API_URl")

	data, err := getJson[IMDBResponse](fmt.Sprintf("%v/title/%v", imdb_api_host, imdb_id))
	if err != nil {
		log.Println("imdb", imdb_id)
		return nil, err
	}
	return data.Rating, nil
}