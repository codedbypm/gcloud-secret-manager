package retrieveSecret

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Payload struct {
	data string
}

type Secret struct {
    name string
	payload Payload
}

func RetrieveSecret(projectId string, secretName string, accessToken string) Secret {
	// Create a new context
	captchaContext := context.Background()
	// With a deadline of 10 seconds
	captchaContext, _ = context.WithTimeout(captchaContext, 10*time.Second)

	// Make a request, that will call the google homepage
	url := fmt.Sprintf("https://secretmanager.googleapis.com/v1beta1/projects/%s/secrets/%s/versions/latest:access", projectId, secretName)
	req, _ := http.NewRequest("GET", url, nil)

	// Associate the cancellable context we just created to the request
	req = req.WithContext(captchaContext)

	// Add HTTP Headers
	authorization := fmt.Sprintf("Bearer %s", accessToken)
	req.Header.Add("Authorization", authorization)
	req.Header.Add("X-Goog-User-Project" projectId)

	// Create a new HTTP client and execute the request
	httpClient := &http.Client{}	
	
	// Execute the request
	res, err := httpClient.Do(req)

	// If the request failed, log the error
	if err != nil {
		log.Fatalln(err)		
		return
	}
	
	// The client must close the response when done	
	defer res.Body.Close()

	// Decode the JSON
	var secret Secret
	decodeError := json.NewDecoder(res.Body).Decode(&secret)
	if decodeError != nil {
		log.Fatalln(decodeError)	
		return	
	}

	return secret
}