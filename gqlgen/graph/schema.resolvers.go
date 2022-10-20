package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/jbutcher93/quotes-starter/gqlgen/graph/generated"
	"github.com/jbutcher93/quotes-starter/gqlgen/graph/model"
	"github.com/jbutcher93/quotes-starter/helpers"
)

// InsertQuote is the resolver for the insertQuote field.
func (r *mutationResolver) InsertQuote(ctx context.Context, input *model.QuoteInput) (*model.Quote, error) {
	postedQuote := &model.Quote{
		Author: input.Author,
		Quote:  input.Quote,
	}
	postBody, err := json.Marshal(&postedQuote)
	if err != nil {
		return nil, err
	}
	responseBody := bytes.NewBuffer(postBody)
	response, err := helpers.MakeRequest(ctx, "http://34.160.62.133:80/quotes", "POST", responseBody)
	if err != nil {
		return nil, err
	}
	responseData, err := helpers.HandleResponse(response)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(responseData, &postedQuote)
	return postedQuote, nil
}

// DeleteQuote is the resolver for the deleteQuote field.
func (r *mutationResolver) DeleteQuote(ctx context.Context, id *string) (*model.DeleteQuoteResponse, error) {
	response, err := helpers.MakeRequest(ctx, fmt.Sprintf("http://34.160.62.133:80/quotes/%s", *id), "DELETE", nil)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case 204:
		return &model.DeleteQuoteResponse{Code: 204, Message: "Delete successful"}, nil
	case 400:
		return &model.DeleteQuoteResponse{Code: 400, Message: "Delete unsuccessful"}, nil
	case 401:
		return &model.DeleteQuoteResponse{Code: 401, Message: "Unauthorized"}, nil
	default:
		return &model.DeleteQuoteResponse{Code: response.StatusCode, Message: "Error"}, nil
	}
}

// RandomQuote is the resolver for the randomQuote field.
func (r *queryResolver) RandomQuote(ctx context.Context) (*model.Quote, error) {
	response, err := helpers.MakeRequest(ctx, "http://34.160.62.133:80/quotes", "GET", nil)
	if err != nil {
		return nil, err
	}
	responseData, err := helpers.HandleResponse(response)
	if err != nil {
		return nil, err
	}
	var randomQuote *model.Quote
	json.Unmarshal(responseData, &randomQuote)
	return randomQuote, nil
}

// QuoteByID is the resolver for the quoteById field.
func (r *queryResolver) QuoteByID(ctx context.Context, id *string) (*model.Quote, error) {
	response, err := helpers.MakeRequest(ctx, fmt.Sprintf("http://34.160.62.133:80/quotes/%s", *id), "GET", nil)
	if err != nil {
		return nil, err
	}
	responseData, err := helpers.HandleResponse(response)
	if err != nil {
		return nil, err
	}
	var quoteByID *model.Quote
	json.Unmarshal(responseData, &quoteByID)
	return quoteByID, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
