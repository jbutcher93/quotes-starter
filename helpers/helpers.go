package helpers

import (
	"errors"
	"io"
	"net/http"
)

func MakeRequest(auth string, url string, requestType string, body io.Reader) *http.Response {
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
