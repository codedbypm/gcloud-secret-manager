package secretmanager

import (
	"context"
	"fmt"

	secrets "google.golang.org/genproto/googleapis/cloud/secrets/v1beta1"

	sm "cloud.google.com/go/secretmanager/apiv1beta1"
)

func GetSecretData(secretName string, projectID string) ([]byte, error) {

	ctx := context.Background()
	smClient, err := sm.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("sm.NewClient: %v", err)
	}

	// Build the request.
	req := &secrets.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/latest", projectID, secretName),
	}
	res, err := smClient.AccessSecretVersion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("smClient.AccessSecretVersion: %v", err)
	}

	return res.Payload.GetData(), nil
}
