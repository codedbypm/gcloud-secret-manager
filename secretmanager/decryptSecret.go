package secretmanager

import (
	"fmt"
)

// DecryptSecretSymmetric will decrypt the given secret using the specified symmetric key.
func DecryptSecretSymmetric(secretName string, projectID string, keyName string) ([]byte, error) {

	secret, err := getSecret(secretName, projectID)
	if err != nil {
		return nil, fmt.Errorf("secretManager: %v", err)
	}

	return decryptSymmetric(keyName, secret.Payload.Data)
}
