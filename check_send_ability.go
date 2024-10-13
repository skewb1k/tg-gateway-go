package tggateway

import (
	"context"
)

type CheckSendAbilityParams struct {
	// The phone number for which you want to check our ability to send a verification message, in the E.164 format.
	PhoneNumber string `json:"phone_number"`
}

// Use this method to check the ability to send a verification message to the specified phone number.
// If the ability to send is confirmed, a fee will apply according to the pricing plan.
// After checking, you can send a verification message using the sendVerificationMessage method, providing the request_id from this response.
// Within the scope of a request_id, only one fee can be charged.
// Calling sendVerificationMessage once with the returned request_id will be free of charge, while repeated calls will result in an error.
// Conversely, calls that don't include a request_id will spawn new requests and incur the respective fees accordingly.
// Note that this method is always free of charge when used to send codes to your own phone number.
func (c Client) CheckSendAbility(ctx context.Context, params *CheckSendAbilityParams) (*RequestStatus, error) {
	var resp struct {
		Ok     bool           `json:"ok"`
		Error  *string        `json:"error"`
		Result *RequestStatus `json:"result"`
	}

	err := c.makeAPIRequest(ctx, "checkSendAbility", params, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, c.mapErr(*resp.Error)
	}

	return resp.Result, nil
}
