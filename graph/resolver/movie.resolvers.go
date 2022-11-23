package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ctuanle/sorapql/graph/generated"
	"github.com/ctuanle/sorapql/graph/model"
	"github.com/ctuanle/sorapql/services/tmdb"
)

// Popular is the resolver for the popular field.
func (r *movieResolver) Popular(ctx context.Context, obj *model.Movie, page *int, language *string, region *string) (*model.MovieList, error) {
	return tmdb.MoviePopular(page, language, region)
}

// TopRated is the resolver for the top_rated field.
func (r *movieResolver) TopRated(ctx context.Context, obj *model.Movie, page *int, language *string, region *string) (*model.MovieList, error) {
	return tmdb.MovieTopRated(page, language, region)
}

// Upcoming is the resolver for the upcoming field.
func (r *movieResolver) Upcoming(ctx context.Context, obj *model.Movie, page *int, language *string, region *string) (*model.MovieList, error) {
	return tmdb.MovieUpcoming(page, language, region)
}

// Movie is the resolver for the movie field.
func (r *queryResolver) Movie(ctx context.Context) (*model.Movie, error) {
	return &model.Movie{}, nil
}

// Movie returns generated.MovieResolver implementation.
func (r *Resolver) Movie() generated.MovieResolver { return &movieResolver{r} }

type movieResolver struct{ *Resolver }
