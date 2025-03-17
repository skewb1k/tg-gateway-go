package tggateway_test

import (
	"errors"
	"testing"
	"time"

	tggateway "github.com/skewb1k/tg-gateway-go/v2"
)

func TestCheckSendAbility(t *testing.T) {
	tests := []struct {
		name      string
		params    *tggateway.CheckSendAbilityParams
		expectErr error
	}{
		{
			name:      "Phone not provided",
			params:    &tggateway.CheckSendAbilityParams{},
			expectErr: tggateway.ErrPhoneNumberInvalid,
		},
		{
			name: "Phone invalid",
			params: &tggateway.CheckSendAbilityParams{
				PhoneNumber: "a",
			},
			expectErr: tggateway.ErrPhoneNumberInvalid,
		},
		{
			name: "Success",
			params: &tggateway.CheckSendAbilityParams{
				PhoneNumber: phone,
			},
			expectErr: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := client.CheckSendAbility(t.Context(), tc.params)
			if !errors.Is(err, tc.expectErr) {
				t.Errorf("unexpected error: wanted %v but got %v", tc.expectErr, err)
			}

			time.Sleep(requestInterval) // avoid flood error
		})
	}
}
