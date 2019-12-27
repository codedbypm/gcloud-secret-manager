package main

import (
	"log"

	"github.com/codedbypm/gcloud-secret-manager/secretmanager"
)

func main() {
	// Read secret
	secretData, err := secretmanager.GetSecretData(
		"agora-secret-mongo-pass",
		"agora-polis",
	)

	if err != nil {
		log.Print(err)
		return
	}

	log.Print(string(secretData))
}
