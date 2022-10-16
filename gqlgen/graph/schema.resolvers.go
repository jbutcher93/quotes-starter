package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/jbutcher93/quotes-starter/gqlgen/graph/generated"
	"github.com/jbutcher93/quotes-starter/gqlgen/graph/model"
	"github.com/jbutcher93/quotes-starter/helpers"
)

// InsertQuote is the resolver for the insertQuote field.
func (r *mutationResolver) InsertQuote(ctx context.Context, input *model.QuoteInput) (*model.Quote, error) {
	Quote := &model.Quote{
		Author: input.Author,
		Quote:  input.Quote,
	}
	postBody, _ := json.Marshal(&Quote)
	responseBody := bytes.NewBuffer(postBody)
	response := helpers.MakeRequest("http://0.0.0.0:8082/quotes", "POST", responseBody)

	/*
		Getting back our newly created UUID and unmarshalling into our Quote instance
		to share with user
	*/
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(responseData, &Quote)
	return Quote, nil
}

// RandomQuote is the resolver for the randomQuote field.
func (r *queryResolver) RandomQuote(ctx context.Context) (*model.Quote, error) {
	response := helpers.MakeRequest("http://0.0.0.0:8082/quotes", "GET", nil)
	responseData, _ := io.ReadAll(response.Body)
	var randomQuote *model.Quote
	json.Unmarshal(responseData, &randomQuote)
	return randomQuote, nil
}

// QuoteByID is the resolver for the quoteById field.
func (r *queryResolver) QuoteByID(ctx context.Context, id *string) (*model.Quote, error) {
	response := helpers.MakeRequest(fmt.Sprintf("http://0.0.0.0:8082/quotes/%s", *id), "GET", nil)
	responseData, _ := io.ReadAll(response.Body)
	var randomQuote *model.Quote
	json.Unmarshal(responseData, &randomQuote)
	return randomQuote, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
