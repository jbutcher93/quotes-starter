package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	// "fmt"

	"github.com/jbutcher93/quotes-starter/gqlgen/graph/generated"
	"github.com/jbutcher93/quotes-starter/gqlgen/graph/model"
)

// Quotes is the resolver for the quotes field.
func (r *queryResolver) Quotes(ctx context.Context) ([]*model.Quote, error) {
	return r.quotes, nil
	// panic(fmt.Errorf("not implemented: Quotes - quotes"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
