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

// Detail is the resolver for the detail field.
func (r *peopleResolver) Detail(ctx context.Context, obj *model.People, id int, language *string) (*model.PersonDetail, error) {
	fields := graphql.CollectAllFields(ctx)
	return tmdb.PersonDetail(id, fields, language)
}

// Popular is the resolver for the popular field.
func (r *peopleResolver) Popular(ctx context.Context, obj *model.People, language *string, page *int) (*model.PopularPeople, error) {
	return tmdb.PeoplePopular(language, page)
}

// People is the resolver for the people field.
func (r *queryResolver) People(ctx context.Context) (*model.People, error) {
	return &model.People{}, nil
}

// People returns generated.PeopleResolver implementation.
func (r *Resolver) People() generated.PeopleResolver { return &peopleResolver{r} }

type peopleResolver struct{ *Resolver }
