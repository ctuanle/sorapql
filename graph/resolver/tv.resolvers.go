package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/ctuanle/sorapql/graph/generated"
	"github.com/ctuanle/sorapql/graph/model"
	"github.com/ctuanle/sorapql/services/tmdb"
)

// Tv is the resolver for the tv field.
func (r *queryResolver) Tv(ctx context.Context) (*model.Tv, error) {
	return &model.Tv{}, nil
}

// Detail is the resolver for the detail field.
func (r *tVResolver) Detail(ctx context.Context, obj *model.Tv, id int, language *string) (*model.TVDetail, error) {
	fields := graphql.CollectAllFields(ctx)
	return tmdb.TVDetail(id, fields, language)
}

// Popular is the resolver for the popular field.
func (r *tVResolver) Popular(ctx context.Context, obj *model.Tv, page *int, language *string) (*model.TVList, error) {
	return tmdb.TVPopular(page, language)
}

// TopRated is the resolver for the top_rated field.
func (r *tVResolver) TopRated(ctx context.Context, obj *model.Tv, page *int, language *string) (*model.TVList, error) {
	return tmdb.TVTopRated(page, language)
}

// OnTheAir is the resolver for the on_the_air field.
func (r *tVResolver) OnTheAir(ctx context.Context, obj *model.Tv, page *int, language *string) (*model.TVList, error) {
	return tmdb.TVOnTheAir(page, language)
}

// TV returns generated.TVResolver implementation.
func (r *Resolver) TV() generated.TVResolver { return &tVResolver{r} }

type tVResolver struct{ *Resolver }
