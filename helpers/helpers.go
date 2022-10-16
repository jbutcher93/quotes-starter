package helpers

import "net/http"

func MakeRequest(url string, requestType string) *http.Response {
	client := &http.Client{}
	req, _ := http.NewRequest(requestType, url, nil)
	req.Header.Set("X-Api-Key", "COCKTAILSAUCE")
	response, _ := client.Do(req)
	return response
}
