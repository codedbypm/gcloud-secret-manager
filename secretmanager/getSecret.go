package secretmanager

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Payload is a struct containing raw data
type Payload struct {
	Data []byte `json:"data"`
}

// Secret models any sensitive data stored in SecretManager
type Secret struct {
	Name    string  `json:"name"`
	Payload Payload `json:"payload"`
}

// getSecret retrieves a given secret from the specified Google Cloud project
func getSecret(secretName string, projectID string) (*Secret, error) {

	httpClient, err := google.DefaultClient(oauth2.NoContext)

	if err != nil {
		return nil, fmt.Errorf("Error: could not create DefaultClient (%s)", err)
	}

	// Make a request, that will call the google homepage
	url := fmt.Sprintf("https://secretmanager.googleapis.com/v1beta1/projects/%s/secrets/%s/versions/latest:access", projectID, secretName)
	req, _ := http.NewRequest("GET", url, nil)

	// // Add HTTP Headers
	req.Header.Add("X-Goog-User-Project", projectID)

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
