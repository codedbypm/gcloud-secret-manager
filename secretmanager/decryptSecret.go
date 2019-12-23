package secretmanager

import (
	"context"
	"fmt"

	cloudkms "cloud.google.com/go/kms/apiv1"
	secretManager "github.com/codedbypm/gcloud-secret-manager"
	kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"
)

// DecryptSecretSymmetric will decrypt the given secret using the specified symmetric key.
func DecryptSecretSymmetric(secretName string, projectID string, keyName string) ([]byte, error) {

	secret, err := secretManager.GetSecret(secretName, projectID)
	if err != nil {
		return nil, fmt.Errorf("secretManager: %v", err)
	}

	return decryptSymmetric(keyName, secret.Payload.Data)
}

// decryptSymmetric will decrypt the input ciphertext bytes using the specified symmetric key.
func decryptSymmetric(keyName string, ciphertext []byte) ([]byte, error) {
	ctx := context.Background()
	client, err := cloudkms.NewKeyManagementClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("cloudkms.NewKeyManagementClient: %v", err)
	}

	// Build the request.
	req := &kmspb.DecryptRequest{
		Name:       keyName,
		Ciphertext: ciphertext,
	}
	// Call the API.
	resp, err := client.Decrypt(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("Decrypt: %v", err)
	}
	return resp.Plaintext, nil
}
