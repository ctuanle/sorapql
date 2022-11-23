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
)

func fetcher[T any](path string, query string) (T, error) {
	key := os.Getenv("TMDB_KEY")
	base := os.Getenv("TMDB_URL")
	url := fmt.Sprintf("%v%v?api_key=%v&%v", base, path, key, query)
	fmt.Println(url)

	var data T
	tmdbClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 5 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("tmdb(NewRequest)", path, ":", err)
		return data, err
	}

	res, err := tmdbClient.Do(req)
	if err != nil {
		log.Println("tmdb(DoGet)", path, " : ", err)
		return data, err
	}
	if res.StatusCode != 200 {
		log.Println("tmdb(DoGet)", res.Status)
		return data, errors.New("!=200")
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		// log.Fatal(readErr)
		log.Println("tmdb(ReadBody)", path, " : ", err)
		return data, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("tmdb(Unmarshal)", path, " : ", err)
		return data, err
		// log.Fatal(jsonErr)
	}

	return data, nil
}
