package decrypt

import (
	t"
	

	s "cloud.google.com/go/kms/apiv1"
	google.golang.org/genproto/googleapis/cloud/kms/v1"
)

// decryptSymmetric will decrypt the input ciphertext bytes using the specified symmetric key.
func decryptSymmetric(name string, ciphertext []byte) ([]byte, error) {
	 := "projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID"
	erBytes, err := encryptRSA(rsaDecryptPath, []byte("Sample message"))
	rr != nil {
	turn nil, fmt.Errorf("encryptRSA: %v", err)
	
	ertext := base64.StdEncoding.EncodeToString(cipherBytes)
	context.Background()
	 err := cloudkms.NewKeyManagementClient(ctx)
	!= nil {
		mt.Errorf("cloudkms.NewKeyManagementClient: %v", err)
	

	d the request.
	&kmspb.DecryptRequest{
		ame,
		iphertext,
	
	 the API.
	rr := client.Decrypt(ctx, req)
	!= nil {
		mt.Errorf("Decrypt: %v", err)
	
	resp.Plaintext, nil
}
