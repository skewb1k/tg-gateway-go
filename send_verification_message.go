package tggateway

import (
	"context"
	"fmt"
)

type SendVerificationMessageParams struct {
	// The phone number to which the verification code was sent, in the E.164 format.
	PhoneNumber string `json:"phone_number"`
	// The unique identifier of a previous request from checkSendAbility.
	// If provided, this request will be free of charge.
	RequestId string `json:"request_id,omitempty"`
	// Username of the Telegram channel from which the code will be sent.
	// The specified channel, if any, must be verified and owned by the same account who owns the Gateway API token.
	SenderUsername string `json:"sender_username,omitempty"`
	// The verification code. Use this parameter if you want to set the verification code yourself.
	// Only fully numeric strings between 4 and 8 characters in length are supported.
	// If this parameter is set, code_length is ignored.
	Code string `json:"code,omitempty"`
	// The length of the verification code if Telegram needs to generate it for you.
	// Supported values are from 4 to 8.
	// This is only relevant if you are not using the code parameter to set your own code.
	// Use the checkVerificationStatus method with the code parameter to verify the code entered by the user.
	CodeLength int `json:"code_length,omitempty"`
	// An HTTPS URL where you want to receive delivery reports related to the sent message, 0-256 bytes.
	CallbackUrl string `json:"callback_url,omitempty"`
	// Custom payload, 0-128 bytes. This will not be displayed to the user, use it for your internal processes.
	Payload string `json:"payload,omitempty"`
	// Time-to-live (in seconds) before the message expires and is deleted.
	// The message will not be deleted if it has already been read.
	// If not specified, the message will not be deleted.
	// Supported values are from 60 to 86400.
	Ttl int `json:"ttl,omitempty"`
}

// Use this method to send a verification message.
// Charges will apply according to the pricing plan for each successful message delivery.
// Note that this method is always free of charge when used to send codes to your own phone number.
func (c Client) SendVerificationMessage(ctx context.Context, params *SendVerificationMessageParams) (*RequestStatus, error) {
	var resp struct {
		Ok     bool           `json:"ok"`
		Error  *string        `json:"error"`
		Result *RequestStatus `json:"result"`
	}

	err := c.makeAPIRequest(ctx, "sendVerificationMessage", params, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf("send verification message failed: %w", c.mapErr(*resp.Error))
	}

	return resp.Result, nil
}
