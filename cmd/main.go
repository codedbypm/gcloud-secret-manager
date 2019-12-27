package main

import (
	"log"

	"github.com/codedbypm/gcloud-secret-manager/secretmanager"
)

func main() {
	// Read secret
	secret, err := secretmanager.GetSecretData(
		"agora-secret-mongo-user",
		"agora-polis",
		"agora-crypto-key",
	)

	if err != nil {
		log.Print(err)
		return
	}

	log.Print(secret)

}
