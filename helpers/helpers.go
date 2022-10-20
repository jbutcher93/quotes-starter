package helpers

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func MakeRequest(ctx context.Context, url string, requestType string, body io.Reader) (*http.Response, error) {
	headerValue := fmt.Sprint(ctx.Value("X-Api-Key"))
	client := &http.Client{}
	req, err := http.NewRequest(requestType, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Api-Key", headerValue)
	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func HandleResponse(r *http.Response) ([]byte, error) {
	if r.StatusCode > 299 {
		return nil, errors.New("error: " + r.Status)
	}
	{
		responseData, err := io.ReadAll(r.Body)
		return responseData, err
	}
}
