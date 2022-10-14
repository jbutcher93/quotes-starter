package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/jbutcher93/quotes-starter/gqlgen/graph/generated"
	"github.com/jbutcher93/quotes-starter/gqlgen/graph/model"
	"github.com/jbutcher93/quotes-starter/helpers"
)

// RandomQuote is the resolver for the randomQuote field.
func (r *queryResolver) RandomQuote(ctx context.Context) (*model.Quote, error) {
	response := helpers.MakeRequest("http://0.0.0.0:8082/quotes", "GET")
	responseData, _ := io.ReadAll(response.Body)
	var randomQuote *model.Quote
	json.Unmarshal(responseData, &randomQuote)
	return randomQuote, nil
}

// QuoteByID is the resolver for the quoteById field.
func (r *queryResolver) QuoteByID(ctx context.Context, id *string) (*model.Quote, error) {
	response := helpers.MakeRequest(fmt.Sprintf("http://0.0.0.0:8082/quotes/%s", *id), "GET")
	responseData, _ := io.ReadAll(response.Body)
	var randomQuote *model.Quote
	json.Unmarshal(responseData, &randomQuote)
	return randomQuote, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
