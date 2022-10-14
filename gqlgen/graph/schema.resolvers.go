package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/jbutcher93/quotes-starter/gqlgen/graph/generated"
	"github.com/jbutcher93/quotes-starter/gqlgen/graph/model"
)

func makeRequest(url string, requestType string) *http.Response {
	client := &http.Client{}
	req, _ := http.NewRequest(requestType, url, nil)
	req.Header.Set("X-Api_Key", "COCKTAILSAUCE")
	response, _ := client.Do(req)
	return response
}

// RandomQuote is the resolver for the randomQuote field.
func (r *queryResolver) RandomQuote(ctx context.Context) (*model.Quote, error) {
	response := makeRequest("http://0.0.0.0:8082/quotes", "GET")
	responseData, _ := io.ReadAll(response.Body)
	var randomQuote *model.Quote
	json.Unmarshal(responseData, &randomQuote)
	return randomQuote, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
