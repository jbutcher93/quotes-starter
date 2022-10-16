package helpers

import (
	"io"
	"net/http"
)

func MakeRequest(url string, requestType string, body io.Reader) *http.Response {
	client := &http.Client{}
	req, _ := http.NewRequest(requestType, url, body)
	req.Header.Set("X-Api-Key", "COCKTAILSAUCE")
	req.Header.Set("Content-Type", "application/json")
	response, _ := client.Do(req)
	return response
}
