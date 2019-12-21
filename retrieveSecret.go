package retrieveSecret

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Payload struct {
	Data string `json:"data"`
}

type Secret struct {
    Name string `json:"name"`
	Payload Payload `json:"payload"`
}

func RetrieveSecret(projectId string, secretName string) (*Secret, error) {
	
	httpClient, err := google.DefaultClient(oauth2.NoContext)

	if err != nil {
		return nil, fmt.Errorf("Error: could not create DefaultClient (%s)", err)
	}

	// Make a request, that will call the google homepage
	url := fmt.Sprintf("https://secretmanager.googleapis.com/v1beta1/projects/%s/secrets/%s/versions/latest:access", projectId, secretName)
	req, _ := http.NewRequest("GET", url, nil)

	// // Add HTTP Headers
	req.Header.Add("X-Goog-User-Project", projectId)

	// Execute the request
	res, err := httpClient.Do(req)

	// If the request failed, log the error
	if err != nil {
		return nil, fmt.Errorf("Error: could not retrieve secret %s (%s)", secretName, err)
	}
	
	// The client must close the response when done	
	defer res.Body.Close()

	// Decode the JSON
	var secret Secret
	decodeError := json.NewDecoder(res.Body).Decode(&secret)
	if decodeError != nil {
		return nil, fmt.Errorf("Error: could not decode secret %s (%s)", secretName, decodeError)
	}

	return &secret, nil
}