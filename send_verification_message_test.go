package tggateway_test

import (
	"errors"
	"testing"
	"time"

	tggateway "github.com/skewb1k/tg-gateway-go/v2"
)

func TestSendVerificationMessage(t *testing.T) {
	tests := []struct {
		name      string
		params    *tggateway.SendVerificationMessageParams
		expectErr error
	}{
		{
			name:      "Phone not provided",
			params:    &tggateway.SendVerificationMessageParams{},
			expectErr: tggateway.ErrPhoneNumberInvalid,
		},
		{
			name: "Phone invalid",
			params: &tggateway.SendVerificationMessageParams{
				PhoneNumber: "a",
			},
			expectErr: tggateway.ErrPhoneNumberInvalid,
		},
		{
			name: "Code length not provided",
			params: &tggateway.SendVerificationMessageParams{
				PhoneNumber: phone,
			},
			expectErr: tggateway.ErrCodeLengthRequired,
		},
		{
			name: "Code length invalid",
			params: &tggateway.SendVerificationMessageParams{
				PhoneNumber: phone,
				CodeLength:  1,
			},
			expectErr: tggateway.ErrCodeLengthInvalid,
		},
		{
			name: "Code invalid",
			params: &tggateway.SendVerificationMessageParams{
				PhoneNumber: phone,
				Code:        "1",
			},
			expectErr: tggateway.ErrCodeInvalid,
		},
		{
			name: "Request id invalid",
			params: &tggateway.SendVerificationMessageParams{
				PhoneNumber: phone,
				CodeLength:  4,
				RequestID:   "asdasd",
			},
			expectErr: tggateway.ErrRequestIDInvalid,
		},
		{
			name: "TTL invalid",
			params: &tggateway.SendVerificationMessageParams{
				PhoneNumber: phone,
				CodeLength:  4,
				TTL:         1,
			},
			expectErr: tggateway.ErrTTLInvalid,
		},
		{
			name: "Callback URL invalid",
			params: &tggateway.SendVerificationMessageParams{
				PhoneNumber: phone,
				CodeLength:  4,
				CallbackURL: "asdasd",
			},
			expectErr: tggateway.ErrCallbackURLInvalid,
		},
		{
			name: "Success",
			params: &tggateway.SendVerificationMessageParams{
				PhoneNumber: phone,
				CodeLength:  4,
			},
			expectErr: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := client.SendVerificationMessage(t.Context(), tc.params)
			if !errors.Is(err, tc.expectErr) {
				t.Errorf("unexpected error: wanted %v but got %v", tc.expectErr, err)
			}

			time.Sleep(requestInterval) // avoid flood error
		})
	}
}
