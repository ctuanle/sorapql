package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ctuanle/sorapql/graph/generated"
	"github.com/ctuanle/sorapql/graph/model"
)

// Introduction is the resolver for the introduction field.
func (r *queryResolver) Introduction(ctx context.Context) (*model.Introduction, error) {
	return &model.Introduction{
		Name:    "ctuanle",
		Message: "Hello World!",
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
