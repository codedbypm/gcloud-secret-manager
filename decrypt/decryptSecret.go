package decrypt

import (
	"context"
	"fmt"

	cloudkms "cloud.google.com/go/kms/apiv1"
	kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"
)

// DecryptSymmetric will decrypt the input ciphertext bytes using the specified symmetric key.
func DecryptSymmetric(name string, ciphertext []byte) ([]byte, error) {
	// name := "projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID"
	// cipherBytes, err := encryptRSA(rsaDecryptPath, []byte("Sample message"))
	// if err != nil {
	//   return nil, fmt.Errorf("encryptRSA: %v", err)
	// }
	// ciphertext := base64.StdEncoding.EncodeToString(cipherBytes)
	ctx := context.Background()
	client, err := cloudkms.NewKeyManagementClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("cloudkms.NewKeyManagementClient: %v", err)
	}

	// Build the request.
	req := &kmspb.DecryptRequest{
		Name:       name,
		Ciphertext: ciphertext,
	}
	// Call the API.
	resp, err := client.Decrypt(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("Decrypt: %v", err)
	}
	return resp.Plaintext, nil
}
