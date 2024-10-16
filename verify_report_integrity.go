package tggateway

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
)

// Use this method to verify the integrity of the delivery report from callback(callback_url).
// The function retrieves the 'X-Request-Timestamp' and 'X-Request-Signature' headers from the request,
// reads the body of the HTTP POST request, and forms a data-check-string by concatenating the timestamp
// and body with a newline character. It then hashes the API token with SHA256, creates an HMAC using
// the token hash, and compares the computed HMAC signature with the received signature to verify authenticity.
func (c Client) VerifyReportIntegrity(r *http.Request) (bool, error) {
	timestamp := r.Header.Get("X-Request-Timestamp")
	requestSignature := r.Header.Get("X-Request-Signature")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return false, fmt.Errorf("could not read post body: %v", err)
	}
	defer r.Body.Close()

	dataCheckString := fmt.Sprintf("%s\n%s", timestamp, string(body))

	apiTokenHash := sha256.Sum256([]byte(c.token))

	mac := hmac.New(sha256.New, apiTokenHash[:])
	mac.Write([]byte(dataCheckString))
	expectedSignature := hex.EncodeToString(mac.Sum(nil))

	if hmac.Equal([]byte(expectedSignature), []byte(requestSignature)) {
		return true, nil
	}

	return false, fmt.Errorf("signature mismatch")
}
