package tmdb

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func fetcher[T any](path string) (T, error) {
	key := os.Getenv("TMDB_KEY")
	base := os.Getenv("TMDB_URL")
	url := fmt.Sprintf("%v%v?api_key=%v", base, path, key)

	var data T
	tmdbClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 5 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("tmdb : ", err)
		return data, nil
		// log.Fatal(err)
	}

	// req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := tmdbClient.Do(req)
	if getErr != nil {
		log.Println("tmdb : ", getErr) 
		return data, nil
		// log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		// log.Fatal(readErr)
		log.Println("tmdb : ", readErr)
		return data, nil
	}

	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		log.Println("tmdb : ", jsonErr)
		return data, nil
		// log.Fatal(jsonErr)
	}

	return data, nil
}
