package tggateway_test

import (
	"errors"
	"testing"
	"time"

	tggateway "github.com/skewb1k/tg-gateway-go/v2"
)

func TestRevokeVerificationMessage(t *testing.T) {
	tests := []struct {
		name      string
		params    *tggateway.RevokeVerificationMessageParams
		expectErr error
	}{
		{
			name:      "Request id not provided",
			params:    &tggateway.RevokeVerificationMessageParams{},
			expectErr: tggateway.ErrRequestIDRequired,
		},
		{
			name: "Request id invalid",
			params: &tggateway.RevokeVerificationMessageParams{
				RequestID: "a",
			},
			expectErr: tggateway.ErrRequestIDInvalid,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := client.RevokeVerificationMessage(t.Context(), tc.params)
			if !errors.Is(err, tc.expectErr) {
				t.Errorf("unexpected error: wanted %v but got %v", tc.expectErr, err)
			}

			time.Sleep(requestInterval) // avoid flood error
		})
	}
}
