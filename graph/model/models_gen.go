// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Link struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Address string `json:"address"`
	User    *User  `json:"user"`
}

type MoviePopular struct {
	Page         int                 `json:"page"`
	Results      []*MoviePopularItem `json:"results"`
	TotalResults int                 `json:"total_results"`
	TotalPages   int                 `json:"total_pages"`
}

type MoviePopularItem struct {
	PosterPath       *string  `json:"poster_path"`
	Adult            *bool    `json:"adult"`
	Overview         *string  `json:"overview"`
	ReleaseDate      *string  `json:"release_date"`
	GenreIds         []*int   `json:"genre_ids"`
	ID               int      `json:"id"`
	OriginalTitle    *string  `json:"original_title"`
	OriginalLanguage *string  `json:"original_language"`
	Title            *string  `json:"title"`
	BackdropPath     *string  `json:"backdrop_path"`
	Popularity       *float64 `json:"popularity"`
	VoteCount        *int     `json:"vote_count"`
	Video            *bool    `json:"video"`
	VoteAverage      *float64 `json:"vote_average"`
}

type NewLink struct {
	Title   string `json:"title"`
	Address string `json:"address"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
