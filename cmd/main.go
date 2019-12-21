package main

import (
	"fmt"
    "source.cloud.google.com/agora-262523/gcloud-secret-manager"
)

func main() {
	const projectId = "agora-262523"
	const secretName = "agora-secret-mongo-user"
	secret, err := retrieveSecret.RetrieveSecret(projectId, secretName)
	if err != nil {
		fmt.Println(err)
		return	
	}
	fmt.Println(secret)
}    