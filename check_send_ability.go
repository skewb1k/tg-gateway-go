package tggateway

import (
	"context"
	"encoding/json"
	"fmt"
)

type CheckSendAbilityParams struct {
	// The phone number for which you want to check our ability to send a verification message, in the E.164 format.
	PhoneNumber string `json:"phone_number"`
}

// Use this method to optionally check the ability to send a verification message to the specified phone number. If the
// ability to send is confirmed, a fee will apply according to the pricing plan. After checking, you can send a
// verification message using the sendVerificationMessage method, providing the request_id from this response.

// Within the scope of a request_id, only one fee can be charged. Calling sendVerificationMessage once with the returned
// request_id will be free of charge, while repeated calls will result in an error. Conversely, calls that don't include
// a request_id will spawn new requests and incur the respective fees accordingly. Note that this method is always free
// of charge when used to send codes to your own phone number.
func (c Client) CheckSendAbility(ctx context.Context, params *CheckSendAbilityParams) (RequestStatus, error) {
	if params == nil {
		panic("params cannot be nil")
	}

	var result RequestStatus
	resultBytes, err := c.makeAPIRequest(ctx, "checkSendAbility", params, &result)
	if err != nil {
		return RequestStatus{}, err
	}

	if err := json.Unmarshal(resultBytes, &result); err != nil {
		return RequestStatus{}, fmt.Errorf("failed to unmarshal result: %w", err)
	}

	return result, nil
}
