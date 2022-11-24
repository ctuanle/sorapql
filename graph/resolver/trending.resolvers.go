package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ctuanle/sorapql/graph/model"
	"github.com/ctuanle/sorapql/services/tmdb"
)

// Trending is the resolver for the trending field.
func (r *queryResolver) Trending(ctx context.Context, mediaType string, timeWindow string, page *int) (*model.Trending, error) {
	return tmdb.Trending(mediaType, timeWindow, page)
}
