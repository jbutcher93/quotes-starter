package helpers

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func MakeRequest(ctx context.Context, url string, requestType string, body io.Reader) *http.Response {
	auth := fmt.Sprint(ctx.Value("X-Api-Key"))
	client := &http.Client{}
	req, _ := http.NewRequest(requestType, url, body)
	req.Header.Set("X-Api-Key", auth)
	req.Header.Set("Content-Type", "application/json")
	response, _ := client.Do(req)
	return response
}

func HandleResponse(r *http.Response) ([]byte, error) {
	if r.StatusCode == 401 {
		return nil, errors.New("error: " + r.Status)
	} else {
		responseData, _ := io.ReadAll(r.Body)
		return responseData, nil
	}
}
